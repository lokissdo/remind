from datetime import datetime, timezone

from sqlalchemy.sql import select, update

from src.models.base import BaseModel
from src.storage.connection import get_session
from src.storage.document.model import Document


def get_document_database(
    model: BaseModel,
    query: str,
) -> list[Document] | None:
    """Gets the document entries from database based on the query"""
    query_embedding = model.embed(query).feature.data

    with get_session() as db:
        documents = db.execute(
            select(Document)
            .order_by(Document.embedding.op("<=>")(query_embedding))
            .limit(3)
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
) -> None:
    """Saves document entry to database"""
    embedding = model.embed(document_data.get(Document.content.key, "")).feature.data
    print(f"Length of embedding {len(embedding)}")
    with get_session() as db:
        try:
            new_document = Document(
                journal_id=document_data.get(Document.journal_id.key),
                content=document_data.get(Document.content.key, ""),
                embedding=embedding,
                created_at=datetime.now(timezone.utc),
                updated_at=datetime.now(timezone.utc),
            )
            db.add(new_document)
            db.commit()

        except Exception:
            db.rollback()
            raise
