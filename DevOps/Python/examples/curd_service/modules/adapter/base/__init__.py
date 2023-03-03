import json


class Config:
    _host: str = "0.0.0.0"
    _port: int = 80

    def __init__(self, host: str = "0.0.0.0", port: int = 80):
        self._host = host
        self._port = port
        pass

    def getConfig(self) -> json:
        return {
            "host": self._host,
            "port": self._port
        }


class Service:

    def __init__(self, config: Config = any):
        pass

    # def __init__(self, config: Config = Config("0.0.0.0", 80)):
    #     pass

    # def __init__(self, host: str = "0.0.0.0", port: int = 80):
    #     pass

    def run():
        pass
