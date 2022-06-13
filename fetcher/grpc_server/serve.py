import os
from concurrent import futures

import grpc
from app.fetcher import FetcherService
from config.logger import logger
from grpc_server import fetcher_pb2_grpc


def run_server():
    logger.info("Running run_server.")
    _serve()


def _serve(max_workers=2):
    logger.info("Running _serve.")

    svc = FetcherService()
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=max_workers))
    fetcher_pb2_grpc.add_FetcherServiceServicer_to_server(svc, server)
    address = f'{os.environ["HOST"]}:{os.environ["PORT"]}'
    server.add_insecure_port(address)
    logger.info("Starting server. Address: %s", address)
    server.start()
    logger.info("Server started. Address: %s", address)
    server.wait_for_termination()
