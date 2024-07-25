from typing import Dict, List, Type

from pgvector.sqlalchemy import Vector
from sqlalchemy import BigInteger, Column, DateTime, Integer, String, func, text
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()


class Document(Base):
    __tablename__ = "document"
    id = Column(BigInteger, primary_key=True, index=True, autoincrement=True)
    journal_id = Column(Integer, nullable=False)
    username = Column(String, nullable=False)
    embedding = Column(Vector(768), nullable=False)
    created_at = Column(DateTime(timezone=False), server_default=func.now())
    updated_at = Column(
        DateTime(timezone=False),
        nullable=False,
        server_default=text("CURRENT_TIMESTAMP"),
        onupdate=func.now(),
    )

    @classmethod
    def to_json(
        cls: Type["Document"], documents: List["Document"]
    ) -> List[Dict[str, any]]:
        """Get a list of document entries and convert them to dictionaries.

        Returns data in a map format. Only returns 'id', 'title', 'url', 'content', 'tokens', 'created_at', and 'updated_at' fields.
        """

        return [
            {
                Document.id.key: document.id,
                Document.journal_id.key: document.journal_id,
                Document.content.key: document.content,
                Document.created_at.key: document.created_at,
                Document.updated_at.key: document.updated_at,
            }
            for document in documents
        ]
