from contextlib import asynccontextmanager
import logging
import requests
import torch
import clip
from PIL import Image
import logging
from fastapi import FastAPI, File, UploadFile
from fastapi.middleware.cors import CORSMiddleware

from src.storage.image.handler import get_image_database, save_image_to_db
from src.storage.document.handler import get_document_database, save_document_to_db
from src.models.bert import BertModel
from src.models.clip import ClipModel
from src.schema.dtypes import DocumentDto, FeatureModel, ImageDto, JournalModel, TextModel, TextsModel, ImageListModel, ImageModel, ImageUrlModel
from typing import Any, Dict, List, Annotated
from io import BytesIO

from src.core.config import app_config


@asynccontextmanager
async def lifespan(app: FastAPI):
    logging.info("Loading models")
    clip_model_name = app_config.clip_model_name
    bert_model_name = app_config.bert_model_name

    app.state.clip_model = ClipModel(clip_model_name)
    app.state.clip_model.load_model()

    app.state.bert_model = BertModel(bert_model_name)
    app.state.bert_model.load_model()

    yield

    app.state.clip_model.unload_model()
    app.state.bert_model.unload_model()


app = FastAPI(lifespan=lifespan)

origins = [
    "http://localhost",
    "http://localhost:8080",
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


@app.get("/")
def read_root() -> Dict[str, str]:
    return {"API Name": app.state.model_name }


@app.get("/api/v1/query/")
def query_data(query: TextModel) -> Any:
    logging.info(f"Querying data with {query.text}")
    query = query.text
    images = get_image_database(app.state.clip_model, query)
    return {"images": [image.id for image in images]}


@app.post("/api/v1/document")
async def encode_document(data: DocumentDto) -> TextModel:
    logging.info(f"Encoding document")
    save_document_to_db(app.state.bert_model, data)
    return {"message": "Save chunk to database completed"}


@app.post("/api/v1/image")
async def encode_image(data: ImageDto) -> TextModel:
    logging.info(f"Encoding image")
    image_data = {
        "journal_id": data.journal_id,
        "content": data.content,
    }
    save_image_to_db(app.state.clip_model, image_data)
    return TextModel(text="Image saved to DB")
