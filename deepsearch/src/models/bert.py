from typing import Any
from sentence_transformers import SentenceTransformer
from src.schema.dtypes import FeatureModel
from src.models.base import BaseModel


class BertModel(BaseModel):
    def __init__(self, model_name: str):
        self.model_name = model_name
        self.model = None

        super().__init__()

    def load_model(self):
        self.model = SentenceTransformer(self.model_name)
    
    def unload_model(self):
        del self.model
    
    def embed(self, data: Any) -> FeatureModel:
        if isinstance(data, str):
            return self._embed_text(data)
        else:
            raise ValueError(f"Unsupported data type {type(data)}")
    
    def _embed_text(self, text: str) -> FeatureModel:
        features = self.model.encode(text)
        return FeatureModel.from_numpy(features)