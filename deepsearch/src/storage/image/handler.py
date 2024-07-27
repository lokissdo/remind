from datetime import datetime, timezone

from sqlalchemy import select
from src.schema.dtypes import QueryModel
from src.storage.connection import get_session
from src.storage.image.model import ImageModel
from src.models.base import BaseModel


def get_image_database(
    model: BaseModel,
    query: QueryModel,
) -> list[ImageModel] | None:
    """Gets the image entries from database based on the query"""
    query_embedding = model.embed(query.text).feature.data
    
    limit = query.limit
    if limit is None:
        limit = 10

    with get_session() as db:
        images = db.execute(
            select(ImageModel)
            .where(ImageModel.username == query.username)
            .order_by(ImageModel.embedding.op("<=>")(query_embedding))
            .limit(limit)
        ).all()

        if not images:
            return []

        # get a list of IMAGE objects
        return [image for (image,) in images]


def save_image_to_db(
    model: BaseModel,
    image_data: dict[str, any]
) -> int:
    """Saves image entry to database"""
    embedding = model.embed(image_data.get("content", "")).feature.data
    print(f"Length of embedding {len(embedding)}")
    with get_session() as db:
        try:
            new_image = ImageModel(
                image_id = image_data.get(ImageModel.image_id.key),
                journal_id=image_data.get(ImageModel.journal_id.key),
                username=image_data.get(ImageModel.username.key, ""),
                embedding=embedding,
                created_at=datetime.now(timezone.utc),
                updated_at=datetime.now(timezone.utc),
            )
            db.add(new_image)
            db.commit()
            return image_data.get(ImageModel.image_id.key)
        except Exception:
            db.rollback()
            return None