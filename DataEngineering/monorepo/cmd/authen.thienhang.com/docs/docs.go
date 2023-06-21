// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://thienhang.com",
        "contact": {
            "name": "API Support",
            "url": "http://thienhang.com",
            "email": "me@thienhang.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/notify/email": {
            "post": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Lấy thông tin của user, nếu không có thì đồng bộ từ firebase",
                "parameters": [
                    {
                        "type": "string",
                        "description": "With the bearer started",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/residential": {
            "get": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Residential"
                ],
                "parameters": [
                    {
                        "maxLength": 30,
                        "minLength": 0,
                        "type": "string",
                        "description": "string valid",
                        "name": "province",
                        "in": "query"
                    },
                    {
                        "maxLength": 30,
                        "minLength": 0,
                        "type": "string",
                        "description": "string valid",
                        "name": "district",
                        "in": "query"
                    },
                    {
                        "maxLength": 30,
                        "minLength": 0,
                        "type": "string",
                        "description": "string valid",
                        "name": "ward",
                        "in": "query"
                    },
                    {
                        "maxLength": 30,
                        "minLength": 0,
                        "type": "string",
                        "description": "string valid",
                        "name": "building",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Province"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "put": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Residential"
                ],
                "summary": "Cập nhật thông tin cho người dùng",
                "parameters": [
                    {
                        "description": "Add InputCreateListener",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Address"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Address"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "post": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Residential"
                ],
                "summary": "Lấy thông tin của user, nếu không có thì đồng bộ từ firebase",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Address"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/residential/address": {
            "get": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Residential"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Address"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/user": {
            "put": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Cập nhật thông tin cho người dùng",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Điền token firebase",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Add InputCreateListener",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "post": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Lấy thông tin của user, nếu không có thì đồng bộ từ firebase",
                "parameters": [
                    {
                        "type": "string",
                        "description": "With the bearer started",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Address": {
            "type": "object",
            "required": [
                "building",
                "id"
            ],
            "properties": {
                "block": {
                    "type": "string"
                },
                "building": {
                    "type": "string"
                },
                "floor": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "website": {
                    "type": "string"
                }
            }
        },
        "entity.Award": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entity.Building": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "address": {
                    "description": "becase can be \"present\"",
                    "type": "string"
                },
                "addresses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Address"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "desctiption": {
                    "description": "Ward        primitive.ObjectID ` + "`" + `bson:\"ward\" json:\"ward\"` + "`" + `\nProvince    primitive.ObjectID ` + "`" + `bson:\"province\" json:\"province\"` + "`" + `\nDistrict    primitive.ObjectID ` + "`" + `bson:\"district\" json:\"district\"` + "`" + `",
                    "type": "string"
                },
                "icon": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "size": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "website": {
                    "type": "string"
                }
            }
        },
        "entity.District": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "pre": {
                    "type": "string"
                },
                "ward": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Ward"
                    }
                }
            }
        },
        "entity.Education": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "from": {
                    "type": "string"
                },
                "grade": {
                    "type": "string",
                    "example": "0"
                },
                "id": {
                    "type": "string"
                },
                "major": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "to": {
                    "description": "becase can be \"present\"",
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entity.Experience": {
            "type": "object",
            "properties": {
                "company": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "from": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "responsibility": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "to": {
                    "description": "becase can be \"present\"",
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entity.Province": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "code": {
                    "type": "string"
                },
                "district": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.District"
                    }
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entity.Qualification": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "expireAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "issueDate": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entity.Reference": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entity.Skill": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "level": {
                    "description": "from 1 to 5",
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entity.User": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "about": {
                    "type": "string"
                },
                "address": {
                    "type": "string"
                },
                "awards": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Award"
                    }
                },
                "courses": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "dob": {
                    "type": "string"
                },
                "educations": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Education"
                    }
                },
                "email": {
                    "type": "string"
                },
                "experiences": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Experience"
                    }
                },
                "firstname": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "lives_in": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "occupation": {
                    "type": "string"
                },
                "pages": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "penalty": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "phone": {
                    "type": "string"
                },
                "photo_url": {
                    "type": "string"
                },
                "qualifications": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Qualification"
                    }
                },
                "references": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Reference"
                    }
                },
                "score": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "skills": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Skill"
                    }
                },
                "status": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "usernames": {
                    "type": "string"
                },
                "website": {
                    "type": "string"
                }
            }
        },
        "entity.Ward": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "building": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Building"
                    }
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "pre": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "OpenKey",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "AUTHENTICATION OPEN API - thienhang.com",
	Description:      "This is a sample server Petstore server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
