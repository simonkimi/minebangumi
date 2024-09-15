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
        "/api/v1/config/init_user": {
            "post": {
                "description": "Initialize the first user with a username and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Initialize the first user",
                "parameters": [
                    {
                        "description": "User initialization form",
                        "name": "InitUserForm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.InitUserForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.TokenResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid form data",
                        "schema": {
                            "$ref": "#/definitions/errno.ApiError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/errno.ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/errno.ApiError"
                        }
                    }
                }
            }
        },
        "/api/v1/config/system": {
            "get": {
                "description": "Get the current system version, first run status, and login status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "system"
                ],
                "summary": "Get system information",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.SystemInfo"
                        }
                    }
                }
            }
        },
        "/api/v1/proxy/poster": {
            "get": {
                "description": "Retrieve a poster image based on the target type and target",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "image/jpeg",
                    " image/png",
                    " image/gif",
                    " image/webp",
                    " image/bmp",
                    " image/svg+xml",
                    " image/jp2"
                ],
                "tags": [
                    "proxy"
                ],
                "summary": "Get poster image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Target type",
                        "name": "target_type",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Target",
                        "name": "target",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Poster image",
                        "schema": {
                            "type": "file"
                        }
                    }
                }
            }
        },
        "/api/v1/source/parse": {
            "post": {
                "description": "Parse the source using the specified parser",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "parser"
                ],
                "summary": "Parse source",
                "parameters": [
                    {
                        "description": "Parse Source Form",
                        "name": "form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.ParseSourceForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ParseSourceResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/source/scrape": {
            "post": {
                "description": "Scrape the source using the specified scraper",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "scraper"
                ],
                "summary": "Scrape source",
                "parameters": [
                    {
                        "description": "Scrape Form",
                        "name": "form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.ScrapeForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ScrapeResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/user/login": {
            "post": {
                "description": "Authenticate user and return JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "Login Form",
                        "name": "form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.LoginForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JWT Token",
                        "schema": {
                            "$ref": "#/definitions/api.TokenResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.InitUserForm": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "api.LoginForm": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "api.ParseSourceForm": {
            "type": "object",
            "required": [
                "parser",
                "source"
            ],
            "properties": {
                "parser": {
                    "type": "string"
                },
                "source": {
                    "type": "string"
                }
            }
        },
        "api.ParseSourceResponse": {
            "type": "object",
            "properties": {
                "files": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "season": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "api.ScrapeForm": {
            "type": "object",
            "required": [
                "scraper",
                "title"
            ],
            "properties": {
                "language": {
                    "type": "string"
                },
                "scraper": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "api.ScrapeResponse": {
            "type": "object",
            "properties": {
                "background": {
                    "type": "string"
                },
                "first_air_year": {
                    "type": "string"
                },
                "original_title": {
                    "type": "string"
                },
                "overview": {
                    "type": "string"
                },
                "poster_path": {
                    "type": "string"
                },
                "scraper": {
                    "type": "string"
                },
                "seasons": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.ScrapeSeasonResponse"
                    }
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "api.ScrapeSeasonResponse": {
            "type": "object",
            "properties": {
                "overview": {
                    "type": "string"
                },
                "poster": {
                    "type": "string"
                },
                "season_id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "api.SystemInfo": {
            "type": "object",
            "properties": {
                "is_first_run": {
                    "type": "boolean"
                },
                "is_login": {
                    "type": "boolean"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "api.TokenResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "errno.ApiError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
