from typing import Any, Dict
from fastapi import FastAPI
import torch

from src.core.config import logger
from src.schema.journal_model import JournalModel
from src.schema.query_model import QueryModel


app = FastAPI()


@app.lifespan.on_event("startup")
async def on_startup():
    device = "cuda" if torch.cuda.is_available() else "cpu"

    if device != "cuda": 
        logger.warn("CLIP not running on CUDA!")

    # model_name = "ViT-L/14@336px"
    # model, preprocess = clip.load(model_name, device=device)
    # logging.info(f"CLIP {model_name} successfully loaded")
    # app.state.device = device
    # app.state.model_name = model_name
    # app.state.model = model
    # app.state.preprocesor = preprocess

@app.post("/api/journal")
def encode_journal(data: JournalModel) -> Dict[Any]:
    texts = data.text_list
    texts_tokenized = torch.cat([clip])


@app.get("/api/search_journal")
def search_journal(data: QueryModel):
    pass