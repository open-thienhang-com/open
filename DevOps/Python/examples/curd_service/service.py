

from modules.adapter.flask import OpenHttpService
from apiv1 import api_v1
# from apiv2 import api_v2


if __name__ == "__main__":
    # import pkg_resources
    # from subprocess import call

    # packages = [dist.project_name for dist in pkg_resources.working_set]
    # call("pip install --upgrade " + ' '.join(packages), shell=True)
    # import flask_restx
    # from importlib import metadata
    # metadata.version("pip")
    # print(metadata.version("flask_restx"))
    # pip_metadata = metadata.metadata("pip")
    # print(list(pip_metadata))
    # import importlib
    # spam_spec = importlib.util.find_spec("importlib_resources")
    # found = spam_spec is not None
    # print(found)

    # ***********************
    service = OpenHttpService()
    service.add_api(api_v1.get_api())
    # service.add_api(api_v2.get_api())
    service.run()
