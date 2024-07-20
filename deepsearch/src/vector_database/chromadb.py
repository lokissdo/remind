from typing import Any, Dict, List, Optional

import chromadb
from chromadb.errors import InvalidDimensionException
from chromadb import Collection, QueryResult

from src.core.config import logger
from src.embedding_models.base import BaseEmbeddingModel
from src.embedding_models.config import EmbeddingModelConfig
from src.types.data_source import DataSource
from src.types.dtypes import MediaData
from src.types.media_type import MEDIA_TYPE
from src.vector_database.base import BaseVectorDatabase
from src.vector_database.configs.chromadb import ChromaDBConfig


class ChromaDB(BaseVectorDatabase):
    """Vector database using ChromaDB"""
    
    BATCH_SIZE = 100
    
    def __init__(
        self,
        config: Optional[ChromaDBConfig] = None,
        embedding_model_config: EmbeddingModelConfig = EmbeddingModelConfig()
    ):
        """Initialize a new ChromaDB instance

        :param config: Configuration options for Chroma, defaults to None
        :type config: Optional[ChromaDbConfig], optional
        """
        self.config = ChromaDBConfig() if config is None else config
        self.client = chromadb.Client(self.config.settings)
        self.embedding_model_config = embedding_model_config
        self._set_all_collections()
    
    def add(
        self,
        data: Any,
        datasource: DataSource,
        file: str,
        source: str,
        media_type: MEDIA_TYPE,
        embedding_model: BaseEmbeddingModel,
    ):
        encodings_json = embedding_model.get_media_encoding(
            data, media_type, datasource
        )
        
        embeddings = encodings_json.get("embedding", None)
        
        documents = (
            [file]
            if not encodings_json.get("documents")
            else encodings_json.get("documents")
        )
        
        size = len(documents)
        if embeddings is not None and len(embeddings) != size:
            raise ValueError(
                "Cannot add documents to chromadb with inconsistent embeddings"
            )
        
        metadata = self._construct_metadata(
            encodings_json.get("metadata", None), source, file, len(documents)
        )
        
        ids = encodings_json.get("ids", [])
        
        collection = self._get_or_create_collection(
            embedding_model.get_collection_name(media_type)
        )
        
        for i in range(0, len(documents), self.BATCH_SIZE):
            logger.info(
                "Inserting batches from {} to {} in chromadb".format(
                    i, min(len(documents), i + self.BATCH_SIZE)
                )
            )
            if embeddings is not None:
                collection.add(
                    embeddings=embeddings[i : i + self.BATCH_SIZE],
                    documents=documents[i : i + self.BATCH_SIZE],
                    ids=ids[i : i + self.BATCH_SIZE],
                    metadatas=metadata[i : i + self.BATCH_SIZE],
                )
            else:
                collection.add(
                    documents=documents[i : i + self.BATCH_SIZE],
                    ids=ids[i : i + self.BATCH_SIZE],
                    metadatas=metadata[i : i + self.BATCH_SIZE],
                )
    
    def query(
        self,
        query: str,
        n_results: int,
        media_type: MEDIA_TYPE,
        distance_threshold: float,
        embedding_model: BaseEmbeddingModel
    ) -> List[MediaData]:
        response = embedding_model.get_text_encoding(query)
        input_embeddings = response.get("embedding", None)
        input_raw_query = response.get("text", None)
        if input_embeddings:
            query_params = {
                "query_embeddings": [input_embeddings],
                "n_results": n_results,
            }
        else:
            query_params = {"query_texts": [input_raw_query], "n_results": n_results}
                    
        collection = self._get_or_create_collection(
            embedding_model.get_collection_name(media_type),
        )
        try:
            results = collection.query(**query_params)
        except InvalidDimensionException as e:
            raise InvalidDimensionException(
                e.message()
                + ". This is commonly a side-effect when an embedding function, different from the one used to"
                " add the embeddings, is used to retrieve an embedding from the database."
            ) from None
        filtered_results = self.filter_query_result_by_distance(
            results, distance_threshold
        )
        if len(filtered_results) == 0:
            return []

        keys = ["documents", "metadatas", "distances"]
        data = [filtered_results.get(key)[0] for key in keys]

        key = ["document", "metadata", "distance"]
        media_data = [
            dict(zip(key, values)) for values in zip(*data)
        ]

        return media_data

    def filter_query_result_by_distance(
        self, query_result: QueryResult, distance_threshold: float
    ) -> QueryResult:
        filtered_result: QueryResult = {
            "ids": [],
            "embeddings": [],
            "documents": [],
            "metadatas": [],
            "distances": [],
        }
        
        for i, _ in enumerate(query_result["ids"]):
            if query_result["distances"][i] is None:
                continue

            filtered_subresult = {key: [] for key in filtered_result.keys()}

            for j, distance in enumerate(query_result["distances"][i]):
                if distance < distance_threshold:
                    continue

                filtered_subresult["ids"].append(query_result["ids"][i][j])
                filtered_subresult["distances"].append(distance)

                for key in ["embeddings", "documents", "metadatas"]:
                    if key in query_result and query_result[key]:
                        filtered_subresult[key].append(query_result[key][i][j])

            for key, value in filtered_subresult.items():
                if value:
                    filtered_result[key].append(value)

        return filtered_result

    def get_existing_document_ids(
        self, metadata_filters, collection_name: str
    ) -> List[str]:
        query_args = {"where": self._generate_where_clause(metadata_filters)}
        collection = self._get_or_create_collection(collection_name)

        results = []
        offset = 0
        first_iteration = True
        while offset != -1 or first_iteration:
            first_iteration = False
            query_result = collection.get(
                **query_args, offset=offset, limit=self.BATCH_SIZE
            )
            metadatas = query_result.get("metadatas", [])
            document_ids = list(
                map(lambda metadata: metadata.get("document_id", []), metadatas)
            )
            results.extend(document_ids)
            offset = offset + min(self.BATCH_SIZE, len(query_result.get("ids")))
            if len(query_result.get("ids")) == 0:
                break
        return results
    
    def count(self) -> Dict[str, int]:
        """
        Count number of documents/chunks embedded in the database.

        :return: number of documents
        """
        return self._get_collection_count()
    
    def delete(self, where, media_type: Optional[MEDIA_TYPE] = None):
        if not media_type or media_type == MEDIA_TYPE.IMAGE:
            media_collections = self.collections.get(MEDIA_TYPE.IMAGE)
            for collection in media_collections:
                collection.delete(where=where)
        
    def reset(self):
        """
        Resets the database. Deletes all embeddings irreversibly.
        """
        # Delete all data from the collection
        for media_collections in self.collections.values():
            for collection in media_collections:
                try:
                    self.client.delete_collection(collection.name)
                except ValueError:
                    raise ValueError(
                        "For safety reasons, resetting is disabled. "
                        "Please enable it by setting `allow_reset=True` in your ChromaDbConfig"
                    ) from None
        self._set_all_collections()

    def _generate_where_clause(self, where_clause: Dict[str, any]):
        # If only one filter is supplied, return it as is
        # (no need to wrap in $and based on chroma docs)
        if not where_clause:
            return {}
        if len(where_clause.keys()) == 1:
            value = list(where_clause.values())[0]
            key = list(where_clause.keys())[0]
            if isinstance(value, list):
                where_filter = {key: {"$in": value}}
            else:
                where_filter = {key: value}
            return where_filter
        where_filters = []
        for k, v in where_clause.items():
            if isinstance(v, list):
                where_filters.append({k: {"$in": v}})
            if isinstance(v, str):
                where_filters.append({k: v})
        return {"$and": where_filters}

    def _set_all_collections(self):
        collections = {}
        for item in self.embedding_model_config.llm_models.items():
            collections[item[0]] = []
            for embedding_model in item[1]:
                media_type = item[0]
                collection_name = embedding_model.get_collection_name(media_type)
                collections[media_type].append(
                    self._get_or_create_collection(collection_name)
                )
        self.collections = collections
        
    def _get_or_create_collection(
        self,
        collection_name: str,
    ) -> Collection:
        return self.client.get_or_create_collection(
            name=collection_name,
            embedding_function=self.config.embedding_function,
            metadata={"hnsw:space": "cosine"},
        )
        
    def _get_collection_count(self):
        collections_count = {}
        for media_collections in self.collections.values():
            for collection in media_collections:
                collections_count[collection.name] = collection.count()
        return collections_count
        