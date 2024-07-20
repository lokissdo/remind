from typing import Any

from src.schema.dtypes import FeatureModel


class BaseModel:
    def __init__(self, **kwargs):
        for key, value in kwargs.items():
            setattr(self, key, value)
    
    def __name__(self):
        return self.__class__.__name__
    
    def load_model(self):
        raise NotImplementedError
    
    def unload_model(self):
        raise NotImplementedError
    
    def embed(self, data: Any) -> FeatureModel:
        raise NotImplementedError