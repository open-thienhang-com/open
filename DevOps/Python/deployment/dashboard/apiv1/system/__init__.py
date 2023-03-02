from flask_restx import Namespace, Resource
from python.modules.checker import get_submodules, add_submodule
from python.modules.adapter.http import response_with_json
from flask import request
ns = Namespace("repos", description="Get system information")


@ns.route("")
@ns.doc(
    responses={
        200: "Success",
        400: "Bad request",
    },
)
class NamespaceSystem(Resource):
    def get(self):
        """get submodules infomation"""
        try:
            return response_with_json(200, get_sub_modules(), "")
        except Exception as err:
            return response_with_json(400, err, "")
    def put(self):
        """reconfigure"""
        try:
            return response_with_json(200, get_system_info(), "")
        except Exception as err:
            return 400
    @ns.doc(
        responses={
            200: "",
            400: "",
        },
        params={
            "repo_path": "https://gitlab.data.Openbook.dev/data-science/ai_mark",
            "module_name": "ai_mark"
        }
    )
    def post(self):
        """add new submodule"""
        try:
            repo_path = request.args.get('repo_path')
            module_name = request.args.get('module_name')
            return response_with_json(200, add_submodule(gitlab_url=repo_path, module_name = module_name), "")
        except Exception as err:
            print("*****")
            print(err)
            return response_with_json(400, str(err), "")
    @ns.doc(
        responses={
            200: "",
            400: "",
        },
        params={
            "repo_path": "https://thienhang.dev/"
        }
    )
    def delete(self):
        """remove submodule"""
        try:
            repo_path = request.args.get('repo_path')
            return response_with_json(200, add_submodule(gitlab_url=repo_path), "")
        except Exception as err:
            return response_with_json(400, str(err), "")
    
