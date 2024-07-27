from contextlib import asynccontextmanager
import json
import logging
import threading
from fastapi import FastAPI, File, Query, UploadFile
from fastapi.middleware.cors import CORSMiddleware

from src.executions.caching.redis import RedisManager
from src.executions.event_handler.publisher.publisher import Publisher
from src.storage.image.handler import get_image_database, save_image_to_db
from src.storage.document.handler import get_document_database, save_document_to_db
from src.models.bert import BertModel
from src.models.clip import ClipModel
from src.schema.dtypes import DocumentDto, ImageDto, QueryModel, TextModel
from typing import Any, Dict, List, Annotated
from io import BytesIO

from src.core.config import app_config

from src.executions.event_handler.consumer.consumer import consumer

logging.basicConfig(level=logging.INFO)


@asynccontextmanager
async def lifespan(app: FastAPI):
    logging.info("Loading models")
    clip_model_name = app_config.clip_model_name
    bert_model_name = app_config.bert_model_name

    app.state.clip_model = ClipModel(clip_model_name)
    app.state.clip_model.load_model()

    app.state.bert_model = BertModel(bert_model_name)
    app.state.bert_model.load_model()
    
    logging.info("Starting consumer thread")
    consumer_thread = threading.Thread(
        target=consumer, 
        args=(app.state.clip_model, app.state.bert_model,), 
        daemon=True
    )
    consumer_thread.start()

    yield

    app.state.clip_model.unload_model()
    app.state.bert_model.unload_model()


app = FastAPI(lifespan=lifespan)

redis_manager = RedisManager(app_config.redis_host)

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
    return {"API Name": "PGVector API"}


@app.get("/api/v1/query/")
def query_data(
    username: str = Query(...),
    text: str = Query(...),
    limit: int = Query(...)
) -> Any:
    logging.info(f"Querying data with {text}")
    
    query = QueryModel(username=username, text=text, limit=limit)
    
    # Generate a unique cache key based on the query parameters
    cache_key = f"{username}-{text}-{limit}"
    cached_response = redis_manager.get(cache_key)
    if cached_response:
        logging.info("Returning cached response")
        return json.loads(cached_response)
    
    images = get_image_database(app.state.clip_model, query)
    texts = get_document_database(app.state.bert_model, query)
    
    response = {
        "images": [image.image_id for image in images],
        "journals": [text.journal_id for text in texts]
    }
    
    redis_manager.set(cache_key, json.dumps(response), app_config.cache_expire)
    
    return response
    

@app.post("/api/v1/document")
async def encode_document(data: DocumentDto) -> TextModel:
    logging.info(f"Encoding document")
    save_document_to_db(
        app.state.bert_model, 
        {
            "content": data.content,
            "journal_id": data.journal_id
        }
    )
    return TextModel(text="Save chunk to database completed")


@app.post("/api/v1/image")
async def encode_image(data: ImageDto) -> TextModel:
    logging.info(f"Encoding image")
    image_data = {
        "journal_id": data.journal_id,
        "content": data.content,
    }
    save_image_to_db(app.state.clip_model, image_data)
    return TextModel(text="Image saved to DB")
