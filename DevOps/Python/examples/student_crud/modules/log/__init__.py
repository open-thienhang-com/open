import logging
import datetime
import sys


class OpenLog(logging.Formatter):
    grey = '\x1b[38;21m'
    blue = '\x1b[38;5;39m'
    yellow = '\x1b[38;5;226m'
    red = '\x1b[38;5;196m'
    bold_red = '\x1b[36m'
    reset = '\x1b[0m'

    def __init__(self, fmt):
        super().__init__()
        self.fmt = fmt
        self.FORMATS = {
            logging.DEBUG: self.grey + self.fmt + self.reset,
            logging.INFO: self.blue + self.fmt + self.reset,
            logging.WARNING: self.yellow + self.fmt + self.reset,
            logging.ERROR: self.red + self.fmt + self.reset,
            logging.CRITICAL: self.bold_red + self.fmt + self.reset
        }

    def format(self, record):
        log_fmt = self.FORMATS.get(record.levelno)
        formatter = logging.Formatter(log_fmt)
        return formatter.format(record)


# Create custom logger logging all five levels
logger = logging.getLogger(__name__)
logger.setLevel(logging.DEBUG)

# Define format for logs
fmt = '%(levelname)8s: %(asctime)s |  %(message)s'

# Create stdout handler for logging to the console (logs all five levels)
stdout_handler = logging.StreamHandler()
stdout_handler.setLevel(logging.DEBUG)
stdout_handler.setFormatter(OpenLog(fmt))

# Create file handler for logging to a file (logs all five levels)
today = datetime.date.today()
# file_handler = logging.FileHandler(
#     'my_app_{}.log'.format(today.strftime('%Y_%m_%d')))
# file_handler.setLevel(logging.DEBUG)
# file_handler.setFormatter(logging.Formatter(fmt))

# Add both handlers to the logger
logger.addHandler(stdout_handler)
# logger.addHandler(file_handler)


print(f"{sys.version_info.major}.{sys.version_info.minor}")
logger.debug('This is a debug-level message')
logger.info('This is an info-level message')
logger.warning('This is a warning-level message')
logger.error('This is an error-level message')
logger.critical('This is a critical-level message')
