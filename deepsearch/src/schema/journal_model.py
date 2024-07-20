from typing import List, Optional
from pydantic import BaseModel, field_validator

class JournalModel(BaseModel):
    journal_id: str
    text_list: Optional[List[str]]
    imageb64_list: Optional[List[str]]
    audio_list: Optional[List[str]]
    