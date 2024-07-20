import logging
import os
from configparser import ConfigParser, ExtendedInterpolation
from functools import lru_cache
from pathlib import Path

import typer
from pydantic_settings import BaseSettings

# from app.core.log.logger import Logger


class AppConfig(BaseSettings):
    # Environment
    environment: str = "dev"
    # Log Level
    log_level: str = "INFO"

    # Clip Model
    clip_model_name: str = ""

    # Bert Model
    bert_model_name: str = ""

    # Postgres
    postgres_db_user: str = ""
    postgres_db_password: str = ""
    postgres_db_host: str = ""
    postgres_db_port: int = 0
    postgres_db_name: str = ""
    postgres_max_overflow: int = 0
    postgres_pool_size: int = 0
    postgres_pool_timeout: int = 0
    postgres_pool_recycle: int = 0


@lru_cache
def get_config(
    environment: str = os.environ.get("ENVIRONMENT", "dev"),
    secret_init_path: str = os.environ.get("SECRET_INI_PATH", "configs/secret.ini"),
) -> AppConfig:
    # parser will be used to contain the main overall rendered config
    parser = ConfigParser(interpolation=ExtendedInterpolation())

    # read raw secret value first
    secret_parser = ConfigParser(interpolation=ExtendedInterpolation())
    _ = secret_parser.read(
        [Path(secret_init_path)],
        encoding="utf-8",
    )

    # special character $ must be escaped for use in extended interpolation later with overall config
    secret_sections = secret_parser.sections()
    for section in secret_sections:
        section_items = dict(secret_parser.items(section, raw=True))
        parser.add_section(section)
        for key, value in section_items.items():
            safe_value = value.replace("$", "$$")
            # we set the safe secret value for later extended interpolation reference
            parser.set(section, key, safe_value)

    # read and render the main config
    config_paths = parser.read(
        [Path("configs/" + environment + ".ini")],
        encoding="utf-8",
    )

    # at this stage we haven't get the env var setup, so we cant use the global logger
    # logger = Logger("config", logging.INFO)
    # logger.info("Reading config from paths: %s", config_paths)

    output_dict = {}

    sections = parser.sections()
    for section in sections:
        output_dict.update(dict(parser.items(section)))

    return AppConfig(**output_dict)


# initialize the app config
app_config = get_config()
# app_config.set_environment_variables()

# log_level = (
#     getattr(logging, app_config.log_level)
#     if hasattr(logging, app_config.log_level)
#     else logging.INFO
# )
# logger = Logger("app", log_level)


# def main() -> None:
#     """Print config value of specified key."""
#     # typer.echo(app_config.dict().get(key))
#     print(app_config.model_dump())


# if __name__ == "__main__":
#     main()
