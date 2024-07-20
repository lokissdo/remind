import logging
from typing import Mapping


class Logger:
    def __init__(self, name: str, level: int = logging.DEBUG) -> None:
        self.logger = logging.getLogger(name)
        self.logger.setLevel(level)
        self.logger.propagate = False
        
        for log_handler in self.logger.handlers[:]:
            self.logger.removeHandler(log_handler)
    
    def info(
        self, message: object, *args: object, tags: Mapping[str, object] | None = None
    ) -> None:
        self.logger.info(message, *args, extra=tags)
        
    def warn(
        self, message: object, *args: object, tags: Mapping[str, object] | None = None
    ) -> None:
        self.logger.warning(message, *args, extra=tags)
        
    def error(
        self, message: object, *args: object, tags: Mapping[str, object] | None = None
    ) -> None:
        self.logger.error(message, *args, extra=tags)
        
    def debug(
        self, message: object, *args: object, tags: Mapping[str, object] | None = None
    ) -> None:
        self.logger.debug(message, *args, extra=tags)
    
    def exception(
        self, message: object, *args: object, tags: Mapping[str, object] | None = None
    ) -> None:
        self.logger.exception(message, *args, extra=tags)