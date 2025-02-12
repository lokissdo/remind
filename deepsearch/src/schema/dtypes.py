from typing import Generic, Optional, TypeVar, List

import numpy as np
from pydantic import BaseModel, validator

T = TypeVar('T')

class NumpyArrayModel(BaseModel, Generic[T]):
    data: list[T]
    shape: list[int]

    @classmethod
    def from_numpy(cls, arr: np.ndarray):
        data = arr.flatten().tolist()
        shape = arr.shape
        return cls(data=data, shape=shape)


class FeatureModel(BaseModel):
    feature: NumpyArrayModel

    @classmethod
    def from_numpy(cls, arr: np.ndarray):
        return cls(feature=NumpyArrayModel.from_numpy(arr))


class ImageModel(BaseModel):
    imageb64: str


class ImageListModel(BaseModel):
    imageb64_ls: List[str]


class ImageUrlModel(BaseModel):
    url: str


class TextModel(BaseModel):
    text: str


class TextsModel(BaseModel):
    texts: List[str]


class JournalModel(BaseModel):
    journal_id: str
    text_list: Optional[List[str]]
    imageb64_list: Optional[List[str]]
    audio_list: Optional[List[str]]


class DocumentDto(BaseModel):
    journal_id: str
    content: str


class ImageDto(BaseModel):
    journal_id: str
    content: bytes


class QueryModel(BaseModel):
    username: str
    text: str
    limit: int
    

