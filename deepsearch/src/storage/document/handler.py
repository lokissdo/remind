from datetime import datetime, timezone
import logging

from sqlalchemy.sql import select, update

from src.models.base import BaseModel
from src.schema.dtypes import QueryModel
from src.storage.connection import get_session
from src.storage.document.model import Document


def get_document_database(
    model: BaseModel,
    query: QueryModel,
) -> list[Document] | None:
    """Gets the document entries from database based on the query"""
    query_embedding = model.embed(query.text).feature.data
    
    limit = query.limit
    if limit is None:
        limit = 10

    with get_session() as db:
        documents = db.execute(
            select(Document)
            .where(Document.username == query.username)
            .order_by(Document.embedding.op("<=>")(query_embedding))
            .limit(limit)
        ).all()

        if not documents:
            return []

        # get a list of DOCUMENT objects
        return [document for (document,) in documents]


def set_all_documents_to_is_moved() -> None:
    """Sets all document entries to is_moved"""
    with get_session() as db:
        db.execute(
            update(Document).where(Document.is_moved.is_(False)).values(is_moved=True)
        )
        db.commit()


def save_document_to_db(
    model: BaseModel,
    document_data: dict[str, any]
) -> int:
    """Saves document entry to database"""
    embedding = model.embed(document_data.get("content", "")).feature.data
    logging.info(f"Length of embedding {len(embedding)}")
    with get_session() as db:
        try:
            new_document = Document(
                journal_id=document_data.get(Document.journal_id.key),
                username=document_data.get(Document.username.key, ""),
                embedding=embedding,
                created_at=datetime.now(timezone.utc),
                updated_at=datetime.now(timezone.utc),
            )
            db.add(new_document)
            db.commit()
            
            return document_data.get(Document.journal_id.key)
        except Exception:
            db.rollback()
            return None
