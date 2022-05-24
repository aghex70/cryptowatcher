import logging

logger = logging.getLogger("fetcher.main")


if __name__ == "__main__":
    logger.info("Running fetcher main.")
    from serve import run_server

    run_server()
