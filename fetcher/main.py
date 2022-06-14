from config.logger import logger
from grpc_server.serve import run_server

if __name__ == "__main__":
    logger.info("Running fetcher main.")

    run_server()
