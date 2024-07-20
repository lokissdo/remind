import os
from typing import Dict, List, Optional

from src.embedding_models.config import EmbeddingModelConfig
from src.llm_models.base import BaseLLM
from src.llm_models.gemini import Gemini
from src.types.dtypes import MediaData, QueryResult
from src.types.media_type import MEDIA_TYPE
from src.types.utils import SourceUtils
from src.vector_database.base import BaseVectorDatabase
from src.vector_database.chromadb import ChromaDB


class App:
    def __init__(
            self,
            embedding_model_config: Optional[EmbeddingModelConfig] = None,
            vector_database: Optional[BaseVectorDatabase] = None,
            llm: Optional[BaseLLM] = None,
    ):
        self.embedding_model_config = (
            embedding_model_config
            if embedding_model_config
            else EmbeddingModelConfig()
        )
        self.vector_database = (
            vector_database
            if vector_database
            else ChromaDB(embedding_model_config=self.embedding_model_config)
        )

        self.llm = llm if llm else Gemini(self.vector_database)
        self.source_utils = SourceUtils()

    def add_data(self, source: str):
        self.source_utils.add_data(
            source, self.embedding_model_config, self.vector_database
        )

    def query(
            self, query: str, media_types: List[MEDIA_TYPE] = [MEDIA_TYPE.IMAGE], n_results: int = 1
    ) -> QueryResult:
        data = self.get_data(query, media_types, n_results)
        print(data)
        response = self.llm.query(query, data)
        return response

    def get_data(
            self, query: str, media_types: List[MEDIA_TYPE] = [MEDIA_TYPE.IMAGE], n_results: int = 1
    ) -> Dict[MEDIA_TYPE, List[MediaData]]:
        return self.source_utils.get_data(
            query, media_types, self.embedding_model_config, self.vector_database, n_results
        )

    def run(self):
        import subprocess
        try:
            import streamlit
            streamlit_file_path = os.path.join(os.path.dirname(__file__), 'streamlit_app.py')
            run_process = subprocess.Popen(['streamlit', 'run', streamlit_file_path, '> NUL'])
            run_process.communicate()
        except ModuleNotFoundError:
            raise ModuleNotFoundError(
                "The required dependencies for ui are not installed."
                ' Please install with `pip install --upgrade "deepsearchai[ui]"`'
            )


if __name__ == "__main__":
    app = App()
    app.add_data("/Users/nguyentienthanh/Downloads/atoz_japanese_food.jpeg")
    app.add_data("/Users/nguyentienthanh/Downloads/japan_tradition.jpeg")
    app.add_data("/Users/nguyentienthanh/Downloads/ramen_with_broth.jpeg")
    
    response = app.query("What do you know about Japan?", media_types=[MEDIA_TYPE.IMAGE], n_results=2)