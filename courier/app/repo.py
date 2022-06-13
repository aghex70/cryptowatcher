from config import settings
from config.logger import logger
from pymongo import MongoClient
from pymongo.errors import ConnectionFailure


def get_database():
    logger.info("Connecting to database. URI: %s", settings.MONGO_URI)
    client = MongoClient(settings.MONGO_URI)
    try:
        logger.info("Checking database connection...")
        client.admin.command("ping")
    except ConnectionFailure:
        logger.error("Error connecting to database")
    logger.info("Connected to database. Database: %s", settings.MONGO_DATABASE)
    return client[settings.MONGO_DATABASE]


def get_collection(database):
    """Get database collection"""
    logger.info("Connecting to collection. Collection: %s", settings.MONGO_COLLECTION)
    collection = database[settings.MONGO_COLLECTION]
    logger.info("Connected to collection. Collection: %s", settings.MONGO_COLLECTION)
    return collection
