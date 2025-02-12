// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/account/sessions/login": {
            "post": {
                "description": "Authenticate user credentials and start a user session.",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Login user for a project",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "X-Kaho-Project",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User Email",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User Password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login success response",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "X-Kaho-Project is required",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Invalid credentials",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/account/sessions/register": {
            "post": {
                "description": "Authenticate user credentials and start a user session.",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Register user for a project",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "X-Kaho-Project",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User Email",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User Password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User Name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login success response",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "201": {
                        "description": "Login success response",
                        "schema": {
                            "$ref": "#/definitions/models.Session"
                        }
                    },
                    "400": {
                        "description": "X-Kaho-Project is required",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Invalid credentials",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Session": {
            "type": "object",
            "properties": {
                "$createdAt": {
                    "type": "string"
                },
                "$id": {
                    "type": "string"
                },
                "$updatedAt": {
                    "type": "string"
                },
                "clientCode": {
                    "type": "string"
                },
                "clientEngine": {
                    "type": "string"
                },
                "clientEngineVersion": {
                    "type": "string"
                },
                "clientName": {
                    "type": "string"
                },
                "clientType": {
                    "type": "string"
                },
                "clientVersion": {
                    "type": "string"
                },
                "countryCode": {
                    "type": "string"
                },
                "countryName": {
                    "type": "string"
                },
                "current": {
                    "type": "boolean"
                },
                "deviceBrand": {
                    "type": "string"
                },
                "deviceModel": {
                    "type": "string"
                },
                "deviceName": {
                    "type": "string"
                },
                "expire": {
                    "type": "string"
                },
                "factors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "ip": {
                    "type": "string"
                },
                "mfaUpdatedAt": {
                    "type": "string"
                },
                "osCode": {
                    "type": "string"
                },
                "osName": {
                    "type": "string"
                },
                "osVersion": {
                    "type": "string"
                },
                "provider": {
                    "type": "string"
                },
                "providerAccessToken": {
                    "type": "string"
                },
                "providerAccessTokenExpiry": {
                    "type": "string"
                },
                "providerRefreshToken": {
                    "type": "string"
                },
                "providerUid": {
                    "type": "string"
                },
                "secret": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Kaho BaaS API Documentation",
	Description:      "API documentation for Kaho BaaS",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
