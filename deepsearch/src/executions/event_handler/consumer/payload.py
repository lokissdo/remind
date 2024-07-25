from typing import List
from pydantic import BaseModel, Field


class JournalCreated(BaseModel):
    type: str = Field(..., description="Journal created")
    journal_id: int = Field(..., description="Journal ID")
    content: str = Field(..., description="The content of journal")
    username: str = Field(..., description="Username")


class ImageCreated(BaseModel):
    type: str = Field(..., description="Image created")
    journal_id: int = Field(..., description="Journal ID")
    image_id: int = Field(..., description="Image ID")
    content: bytes = Field(..., description="The content of image")
    username: str = Field(..., description="Username")
