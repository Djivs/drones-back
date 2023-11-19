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
        "/book": {
            "put": {
                "description": "Creates a new flight and adds current region in it",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "general"
                ],
                "summary": "Book region",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/flight": {
            "get": {
                "description": "Returns flight with given parameters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "flights"
                ],
                "summary": "Get flight",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/flight/delete/:flight_id": {
            "put": {
                "description": "Changes flight status to \"Удалён\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "flights"
                ],
                "summary": "Deletes flight",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/flight/edit": {
            "put": {
                "description": "Finds flight and updates it fields",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "flights"
                ],
                "summary": "Edits flight",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/flight/status_change/moderator": {
            "put": {
                "description": "Changes flight status to any available status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "flights"
                ],
                "summary": "Changes flight status as moderator",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/flight/status_change/user": {
            "put": {
                "description": "Changes flight status as allowed to user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "flights"
                ],
                "summary": "Changes flights status as user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/flight_to_region/delete": {
            "put": {
                "description": "Deletes region from flight",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "flights"
                ],
                "summary": "Deletes flight_to_region connection",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/flights": {
            "get": {
                "description": "Returns list of all available flights",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "flights"
                ],
                "summary": "Get flights",
                "responses": {
                    "302": {
                        "description": "Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/region/:region": {
            "get": {
                "description": "Returns region with given name",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "regions"
                ],
                "summary": "Get region",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/region/add": {
            "put": {
                "description": "Creates a new reigon with parameters, specified in json",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "regions"
                ],
                "summary": "Adds region to database",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/region/delete/:region_name": {
            "put": {
                "description": "Finds region by name and changes its status to \"Недоступен\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "regions"
                ],
                "summary": "Deletes region",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/region/delete_restore/:region_name": {
            "get": {
                "description": "Switches region status from \"Действует\" to \"Недоступен\" and back",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "regions"
                ],
                "summary": "Deletes or restores region",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/region/edit": {
            "put": {
                "description": "Finds region by name and updates its fields",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "regions"
                ],
                "summary": "Edits region",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/regions": {
            "get": {
                "description": "Returns all existing regions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "regions"
                ],
                "summary": "Get all existing regions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": ""
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0-0",
	Host:             "127.0.0.1:8000",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "drones",
	Description:      "UAV route control applications",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
