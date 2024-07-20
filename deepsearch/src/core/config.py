from configparser import ConfigParser, ExtendedInterpolation
from functools import lru_cache
import logging
import os
from pathlib import Path
from pydantic_settings import BaseSettings

from src.core.log.logger import Logger


class AppConfig(BaseSettings):
    environment: str = "dev"
    #Google API
    google_api_key: str = ""
    #logging
    log_level: str = ""
    
    def set_environment_variables(self) -> None:
        """Set the environment variable"""
        os.environ["GOOGLE_API_KEY"] = self.google_api_key

@lru_cache
def get_config(
    environment: str = os.environ.get("ENVIRONMENT", "dev"),
    secret_init_path: str = os.environ.get("SECRET_INIT_PATH", "configs/secret.ini"),
) -> AppConfig:
    parser = ConfigParser(interpolation=ExtendedInterpolation())
    config_paths = parser.read(
        [Path(secret_init_path), Path("configs/" + environment + ".ini")],
        encoding="utf-8",
    )
    
    logger = Logger("config", logging.INFO)
    logger.info(f"Reading config files: {config_paths}")
    
    output_dict = {}
    
    sections = parser.sections()
    for section in sections:
        output_dict.update(dict(parser.items(section)))
    
    return AppConfig(**output_dict)


# initialize the app config
app_config = get_config()
app_config.set_environment_variables()

# this is in string format ("WARNINING", "DEBUG" etc)
log_level_str = app_config.log_level
log_level = (
    getattr(logging, log_level_str) if hasattr(logging, log_level_str) else logging.INFO
)
logger = Logger("app", log_level)