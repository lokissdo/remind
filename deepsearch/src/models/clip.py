from io import BytesIO
import logging
from typing import Any, Callable, List
from PIL import Image
import PIL
import torch
import clip 
from src.schema.dtypes import FeatureModel
from src.models.base import BaseModel


class ClipModel(BaseModel):
    def __init__(self, model_name: str):
        self.model_name = model_name
        self.model: torch.nn.Module = None
        self.preprocessor: Callable[[PIL.Image], torch.Tensor] = None
        self.device: str = None

        super().__init__()
    
    def load_model(self):
        device = "cuda" if torch.cuda.is_available() else "cpu"

        if device != "cuda": 
            logging.warning("CLIP not running on CUDA!")

        self.model, self.preprocessor = clip.load(self.model_name, device=device)
        logging.info(f"CLIP {self.model_name} successfully loaded")
    
    def unload_model(self):
        del self.model
        del self.preprocessor
        logging.info(f"CLIP {self.model_name} successfully unloaded")

    def embed(self, data: Any) -> FeatureModel:
        if isinstance(data, str):
            return self._embed_text(self._tokenize_text(data))
        elif isinstance(data, list) and all(isinstance(item, str) for item in data):
            return self._embed_text(self._tokenize_texts(data))
        elif isinstance(data, bytes):
            return self._embed_image(self._process_image(data))
        elif isinstance(data, list) and all(isinstance(item, bytes) for item in data):
            return self._embed_image(self._process_images(data))
        else:
            raise ValueError(f"Unsupported data type {type(data)}")
    
    def _embed_text(self, text_tensor: torch.Tensor) -> FeatureModel:
        with torch.no_grad():
            text_feature = self.model.encode_text(text_tensor.to(self.device))
            text_feature = text_feature / text_feature.norm(dim=-1, keepdim=True)
        # features have type float16
        features = text_feature.cpu().numpy()
        return FeatureModel.from_numpy(features)
    
    def _embed_image(self, image_tensor: torch.Tensor) -> FeatureModel:
        with torch.no_grad():
            image_feature = self.model.encode_image(image_tensor)
            image_feature = image_feature / image_feature.norm(dim=-1, keepdim=True)
        # features have type float16
        features = image_feature.cpu().numpy()
        return FeatureModel.from_numpy(features)
    
    def _tokenize_text(self, text: str) -> torch.Tensor:
        return clip.tokenize(text).to(self.device)

    def _tokenize_texts(self, texts: List[str]) -> torch.Tensor:
        return torch.cat([self._tokenize_text(text) for text in texts])

    def is_valid_image(self, image_data: bytes) -> bool:
        try:
            Image.open(BytesIO(image_data))
            return True
        except PIL.UnidentifiedImageError:
            return False

    def _process_image(self, image_data: bytes) -> torch.Tensor:
        if not self.is_valid_image(image_data):
            raise ValueError("Invalid image data")
        image = Image.open(BytesIO(image_data)).convert('RGB')
        return self.preprocessor(image).unsqueeze(0).to(self.device)

    def _process_images(self, images_data: List[bytes]) -> torch.Tensor:
        images = []
        for im_data in images_data:
            if im_data == b'': continue
            images.append(Image.open(BytesIO(im_data)))
        return torch.stack([
            self.preprocessor(image) for image in images
        ]).to(self.device)
