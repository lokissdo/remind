from typing import Optional

from src.embedding_models.base import BaseEmbeddingModel
from src.embedding_models.blip_image_captioning import BlipImageCaptioning
from src.embedding_models.clip import Clip
from src.types.media_type import MEDIA_TYPE


class EmbeddingModelConfig:
    def __init__(
        self,
        image_embedding_model: Optional[BaseEmbeddingModel] = None,
        image_captioning_model: Optional[BaseEmbeddingModel] = BlipImageCaptioning(),
    ):
        if not image_embedding_model:
            image_embedding_model = Clip()
        
        image_related_embedding_models = [image_embedding_model]
        if image_captioning_model:
            image_related_embedding_models.append(image_captioning_model)
        
        self.llm_models = {
            MEDIA_TYPE.IMAGE: image_related_embedding_models,
        }

    def get_embedding_model(self, media_type: MEDIA_TYPE):
        return self.llm_models.get(media_type, [])