from modules.adapter.flask import OpenNamespace, OpenResource, OpenHttpRespone

ns_model = OpenNamespace(
    name="model2",
    description="Test V2"
)


@ns_model.route("")
@ns_model.doc(
    responses={
        200: "Success",
        400: "Bad request",
    },
)
class NamespaceTestV2(OpenResource):
    def post(self):
        """Open namespace"""
        try:
            return 200
        except Exception as err:
            return 400
