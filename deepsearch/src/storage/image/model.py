from typing import Dict, List, Type
from pgvector.sqlalchemy import Vector
from sqlalchemy.dialects.postgresql import BYTEA
from sqlalchemy import BigInteger, Column, DateTime, Integer, func, text, String
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()


class ImageModel(Base):
    __tablename__ = "image"
    id = Column(BigInteger, primary_key=True, index=True, autoincrement=True)
    image_id = Column(BigInteger, nullable=False)
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
        cls: Type["ImageModel"], images: List["ImageModel"]
    ) -> List[Dict[str, any]]:
        """Get a list of image entries and convert them to dictionaries.

        Returns data in a map format. 
        """

        return [
            {
                ImageModel.id.key: image.id,
                ImageModel.journal_id.key: image.journal_id,
            }
            for image in images
        ]