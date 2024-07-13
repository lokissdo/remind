from typing import Any

from types.data_source import DataSource
from types.media_type import MEDIA_TYPE


class BaseEmbeddingModel:
    def __init__(self):
        pass
    
    def get_media_encoding(
        self, data: Any, dataType: MEDIA_TYPE, datasource: DataSource
    ):
        raise NotImplementedError
    
    def get_text_encoding(self, query: str):
        raise NotImplementedError
    
    def get_collection_name(self, media_type: MEDIA_TYPE):
        raise NotImplementedError