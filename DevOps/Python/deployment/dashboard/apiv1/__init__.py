from python.modules.adapter.http import OpenAPI
from .system import ns as ns_system


api_v1 = OpenAPI(
    api_name="api",
    pre_path="/api/v1",
    doc_path="/ui",
    doc_title="Open",
    doc_version="1.0",
    doc_description="No thing here",
    license_url="thienhang"
)

api_v1.add_custom_namespace(ns_system)
