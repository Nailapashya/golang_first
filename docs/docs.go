// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "lifelinejar@mail.com"
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
        "/index": {
            "get": {
                "description": "Returns a user's information in JSON format",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get user information",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/site": {
            "get": {
                "description": "get site page.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Site"
                ],
                "summary": "site page",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/utils.RequestError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/utils.RequestError"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "$ref": "#/definitions/utils.RequestError"
                        }
                    },
                    "422": {
                        "description": "Data validation failed",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/utils.DataValidationError"
                            }
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "$ref": "#/definitions/utils.RequestError"
                        }
                    }
                }
            }
        },
        "/strict/urusan": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Menampilkan list master data urusan dengan id daerah yang sama dengan user yang mengaksesnya.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Urusan"
                ],
                "summary": "Menampilkan List Master Data Urusan.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Tahun yang ditampilkan",
                        "name": "tahun",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Filter kode urusan (match)",
                        "name": "kode_urusan",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter nama urusan (like)",
                        "name": "nama_urusan",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Halaman yang ditampilkan",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Jumlah data per halaman, maksimal 50 data",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.UrusanModel"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/utils.RequestError"
                        }
                    },
                    "404": {
                        "description": "Data not found",
                        "schema": {
                            "$ref": "#/definitions/utils.RequestError"
                        }
                    },
                    "422": {
                        "description": "Data validation failed",
                        "schema": {
                            "$ref": "#/definitions/utils.RequestError"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "$ref": "#/definitions/utils.RequestError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.UrusanModel": {
            "type": "object",
            "properties": {
                "id_daerah": {
                    "type": "integer",
                    "example": 371
                },
                "id_unik": {
                    "description": "ID unik urusan",
                    "type": "string",
                    "example": "0076b9dc-02eb-4d05-990a-44b16a9ff76a"
                },
                "id_urusan": {
                    "type": "integer",
                    "example": 12
                },
                "is_locked": {
                    "description": "Urusan ini dikunci atau tidak",
                    "type": "integer",
                    "example": 0
                },
                "kode_urusan": {
                    "type": "string",
                    "example": "2"
                },
                "nama_urusan": {
                    "type": "string",
                    "example": "URUSAN PEMERINTAHAN WAJIB YANG TIDAK BERKAITAN DENGAN PELAYANAN DASAR"
                },
                "tahun": {
                    "type": "integer",
                    "example": 2022
                }
            }
        },
        "utils.DataValidationError": {
            "type": "object",
            "properties": {
                "field": {
                    "type": "string",
                    "example": "email"
                },
                "message": {
                    "type": "string",
                    "example": "Invalid email address"
                }
            }
        },
        "utils.RequestError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 422
                },
                "fields": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/utils.DataValidationError"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "Invalid email address"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "SIPD Service Boilerpate",
	Description:      "SIPD Service Boilerpate Rest API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
