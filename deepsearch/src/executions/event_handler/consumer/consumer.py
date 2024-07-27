import base64
import json
import pika
import logging

from src.core.config import app_config

from src.executions.event_handler.consumer.payload import JournalCreated, ImageCreated
from src.models.base import BaseModel
from src.storage.document.handler import save_document_to_db
from src.storage.image.handler import save_image_to_db

from src.executions.event_handler.publisher.publisher import Publisher


def create_callback(clip_model: BaseModel, bert_model: BaseModel):
    def process_body(body_dict, model_fields, model_class, db_model, save_func):
        filtered_dict = {key: value for key, value in body_dict.items() if key in model_fields}
        if filtered_dict.get('type') == "image":
            filtered_dict['content'] = base64.b64decode(filtered_dict['content'])
        try:
            instance = model_class(**filtered_dict)
            result = save_func(db_model, instance.model_dump())
            if result is not None:
                print(result)
                publisher = Publisher(filtered_dict.get("type"))
                print(result)
                publisher.connect()
                print(result)
                response = {
                    "type": filtered_dict.get("type"),
                    "id": result
                }
                print(json.dumps(response))
                publisher.publish(json.dumps(response), "application/json")
                publisher.close()
        except Exception as e:
            logging.error(f"Failed to validate the body: {e}")

    def callback(ch, method, properties, body):
        try:
            body_dict = json.loads(body)
        except json.JSONDecodeError:
            logging.error("Failed to decode JSON body")
            # TODO: Handle the rollback, notify the journal service
            return
        
        if body_dict.get("type") is not None and body_dict.get("type") == "journal": 
            process_body(body_dict, JournalCreated.model_fields, JournalCreated, bert_model, save_document_to_db)
        elif body_dict.get("type") is not None and body_dict.get("type") == "image":
            process_body(body_dict, ImageCreated.model_fields, ImageCreated, clip_model, save_image_to_db)
        
    return callback


def consumer(clip_model: BaseModel, bert_model: BaseModel):
    connection = pika.BlockingConnection(pika.ConnectionParameters(app_config.rabbitmq_host))
    channel = connection.channel()
    
    # Declare the exchange
    exchange_name = app_config.consumer_exchange_name
    channel.exchange_declare(
        exchange=exchange_name, 
        exchange_type='direct',
        durable=True,
        auto_delete=False,
        internal=False,
        arguments=None
    )
    
    # Declare the queue
    result = channel.queue_declare(
        queue='',
        durable=True,
        exclusive=True,
        auto_delete=False,
        arguments=None
    )
    queue_name = result.method.queue
    
    # Bind the queue to the exchange
    routing_key = app_config.consumer_routing_key
    channel.queue_bind(
        exchange=exchange_name,
        queue=queue_name,
        routing_key=routing_key
    )

    # Set up the consumer
    channel.basic_consume(
        queue=queue_name,
        on_message_callback=create_callback(clip_model, bert_model),
        auto_ack=True
    )
    
    logging.info(" [*] Waiting for messages. To exit press Ctrl+C to exit")
    try:
        channel.start_consuming()
    except KeyboardInterrupt:
        logging.info(" [*] Stopping consumer")
        channel.stop_consuming()
        
    # Close the connection
    connection.close()    
    

