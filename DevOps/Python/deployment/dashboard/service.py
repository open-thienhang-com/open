from Python.modules.adapter.flask import OpenHttpService
from apiv1 import api_v1
from views import ns_index


if __name__ == "__main__":
    service = OpenHttpService()
    service.add_api(api_v1.get_api())
    service.add_ui(ns_index),
    service.run()
