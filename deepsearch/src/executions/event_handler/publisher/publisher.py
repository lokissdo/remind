import json
import uuid
import pika
import logging
from datetime import datetime
from typing import Any, Dict, List

from src.core.config import app_config

# Constants
PUBLISH_MANDATORY = False


class Publisher:
    def __init__(self, message_type_name: str):
        self.amqp_url = app_config.rabbitmq_host
        self.exchange_name = app_config.publisher_exchange_name
        self.binding_key = app_config.publisher_routing_key
        self.message_type_name = message_type_name
        self.connection = None
        self.channel = None
        
    def connect(self):
        print("Connecting to RabbitMQ")
        self.connection = pika.BlockingConnection(pika.ConnectionParameters(self.amqp_url))
        self.channel = self.connection.channel()
        print("hello")
    
    def configure(self, exchange_name: str = None, binding_key: str = None, message_type_name: str = None):
        if exchange_name:
            self.exchange_name = exchange_name
        if binding_key:
            self.binding_key = binding_key
        if message_type_name:
            self.message_type_name = message_type_name

    def publish(self, body: str, content_type: str):
        if not self.channel:
            raise Exception("Channel is not initialized. Call connect() first.")
        logging.info(f"Publishing message to exchange {self.exchange_name} with routing key {self.binding_key}")
        
        properties = pika.BasicProperties(
            content_type=content_type,
            delivery_mode=pika.spec.PERSISTENT_DELIVERY_MODE,
            message_id=str(uuid.uuid4()),
            timestamp=int(datetime.now().timestamp()),
            type=self.message_type_name,
        )
        
        self.channel.basic_publish(
            exchange=self.exchange_name,
            routing_key=self.binding_key,
            body=body,
            properties=properties,
            mandatory=PUBLISH_MANDATORY,
        )
        
    def close(self):
        if self.channel:
            self.channel.close()
        if self.connection:
            self.connection.close()