from pathlib import Path
from typing import Any, Dict, List, Union

from src.models.base import BaseModel
from src.models.clip import ClipModel
from src.models.utils import chunk_data, read_file
from src.storage.document.handler import (
    get_document_database,
    save_document_to_db,
)
from src.storage.document.model import Document


class IngestChunksError(Exception):
    """Raised when there's an error ingesting chunks to the database"""

    def __init__(self, error: Union[BaseException, str]) -> None:
        self.message = f"Failed to save chunks to database: {error}"
        super().__init__(self.message)


class QueryDataError(Exception):
    """Raised when there's an error querying data"""

    def __init__(self, error: Union[BaseException, str]) -> None:
        self.message = f"Failed to get data with query: {error}"
        super().__init__(self.message)


def create_chunk_data(chunks: List[str]) -> List[Dict[str, Any]]:
    return [
        {
            Document.journal_id.key: 1,
            Document.content.key: chunk,
            Document.embedding.key: chunk,
        }
        for chunk in chunks
    ]


def ingest_chunks_to_db_example(model: BaseModel) -> Dict[str, str]:
    try:
        # Read file from directory
        script_dir = Path(__file__).parent
        file_content = read_file(script_dir / "document.txt")
        chunks = chunk_data(file_content)

        for data in create_chunk_data(chunks):
            print(data)
            save_document_to_db(model, data)

    except BaseException as e:
        raise IngestChunksError(str(e)) from e

    return {"message": "Save chunks to database completed"}


def get_data_with_query_example(model: BaseModel, query: str) -> Dict[str, List[Dict[str, Any]]]:
    try:
        documents = get_document_database(model, query)
        return {"documents": Document.to_json(documents)}
    except BaseException as e:
        raise QueryDataError(str(e)) from e
