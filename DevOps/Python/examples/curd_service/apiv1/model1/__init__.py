from flask_restx import Namespace, Resource

ns_model = Namespace(
    "model1",
    description="Test V1")


@ns_model.route("")
@ns_model.doc(
    responses={
        200: "Success",
        400: "Bad request",
    },
)
class NamespaceTestV1(Resource):
    def post(self):
        """Flask namespace"""
        try:
            return 200
        except Exception as err:
            return 400
