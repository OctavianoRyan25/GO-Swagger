// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/car/{id}": {
            "get": {
                "description": "get detail of car corresponding to the input ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "Get Car for a given ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the car",
                        "name": "Id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Car"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update detail of car corresponding to the input ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "Update Car for a given ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the car",
                        "name": "Id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Car"
                        }
                    }
                }
            }
        },
        "/cars": {
            "get": {
                "description": "Get Detail All object",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "Get All Object",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Car"
                        }
                    }
                }
            },
            "post": {
                "description": "Post detail of car corresponding to the input ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "Post Details for a given id",
                "parameters": [
                    {
                        "description": "create car",
                        "name": "model.Car",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Car"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Car"
                        }
                    }
                }
            }
        },
        "/cars/{id}": {
            "delete": {
                "description": "Delete detail of car corresponding to the input ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "Delete Car for a given ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the car",
                        "name": "Id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Car"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Car": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "merk": {
                    "type": "string"
                },
                "pemilik": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}