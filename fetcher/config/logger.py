import logging

try:
    from config.settings import PROJECT_NAME
except ModuleNotFoundError:
    from fetcher.config.settings import PROJECT_NAME


def start_logger():
    logger = logging.getLogger(PROJECT_NAME)
    handler = logging.StreamHandler()
    formatter = logging.Formatter(
        "%(asctime)s - [%(levelname)s] *%(name)s* [%(module)s.%(funcName)s:%(lineno)d]: %("
        "message)s"
    )
    handler.setFormatter(formatter)
    logger.addHandler(handler)
    logger.setLevel(logging.INFO)
    return logger


logger = start_logger()
