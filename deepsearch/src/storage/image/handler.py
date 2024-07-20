from datetime import datetime, timezone

from sqlalchemy import select
from src.storage.connection import get_session
from src.storage.image.model import ImageModel
from src.models.base import BaseModel


def get_image_database(
    model: BaseModel,
    query: str,
) -> list[ImageModel] | None:
    """Gets the image entries from database based on the query"""
    query_embedding = model.embed(query).feature.data

    with get_session() as db:
        images = db.execute(
            select(ImageModel)
            .order_by(ImageModel.embedding.op("<=>")(query_embedding))
            .limit(3)
        ).all()

        if not images:
            return []

        # get a list of IMAGE objects
        return [image for (image,) in images]


def save_image_to_db(
    model: BaseModel,
    image_data: dict[str, any]
) -> None:
    """Saves image entry to database"""
    embedding = model.embed(image_data.get(ImageModel.content.key, "")).feature.data
    print(f"Length of embedding {len(embedding)}")
    with get_session() as db:
        try:
            new_image = ImageModel(
                journal_id=image_data.get(ImageModel.journal_id.key),
                content=image_data.get(ImageModel.content.key, ""),
                embedding=embedding,
                created_at=datetime.now(timezone.utc),
                updated_at=datetime.now(timezone.utc),
            )
            db.add(new_image)
            db.commit()

        except Exception:
            db.rollback()