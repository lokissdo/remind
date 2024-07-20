from typing import Optional
from pydantic import BaseModel


class QueryModel(BaseModel):
    query: str
    context: Optional[str]
    user_id: Optional[str]
    timestamp: Optional[str]
    session_id: Optional[str]