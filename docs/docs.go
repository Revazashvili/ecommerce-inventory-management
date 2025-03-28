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
        "/api/product": {
            "get": {
                "description": "Get product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Get product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/database.Product"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/product/count": {
            "get": {
                "description": "Get product count",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Get product count",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/stock": {
            "get": {
                "description": "Get Stock",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stocks"
                ],
                "summary": "Get Stock",
                "parameters": [
                    {
                        "description": "getStockRequest",
                        "name": "getStockRequest",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/handlers.GetStocksRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/database.Stock"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/stock/add": {
            "post": {
                "description": "Add Stock",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stocks"
                ],
                "summary": "Add Stock",
                "parameters": [
                    {
                        "description": "addStockRequest",
                        "name": "addStockRequest",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/handlers.AddStockRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/stock/reserve": {
            "post": {
                "description": "Reserve Stock",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stocks"
                ],
                "summary": "Reserve Stock",
                "parameters": [
                    {
                        "description": "reserveRequest",
                        "name": "reserveRequest",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/handlers.ReserveRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/stock/unreserve": {
            "post": {
                "description": "Unreserve Stock",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stocks"
                ],
                "summary": "ReserUnreserveve Stock",
                "parameters": [
                    {
                        "description": "unreserveRequest",
                        "name": "unreserveRequest",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/handlers.UnreserveRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "database.Product": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "database.Stock": {
            "type": "object",
            "properties": {
                "createDate": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lastUpdateDate": {
                    "type": "string"
                },
                "productID": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                },
                "reservedQuantity": {
                    "type": "integer"
                },
                "version": {
                    "type": "integer"
                }
            }
        },
        "handlers.AddStockRequest": {
            "type": "object",
            "properties": {
                "productID": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "handlers.GetStocksRequest": {
            "type": "object",
            "properties": {
                "from": {
                    "type": "string"
                },
                "productID": {
                    "type": "string"
                },
                "to": {
                    "type": "string"
                }
            }
        },
        "handlers.ReserveRequest": {
            "type": "object",
            "properties": {
                "orderNumber": {
                    "type": "string"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handlers.ReserveRequestProduct"
                    }
                }
            }
        },
        "handlers.ReserveRequestProduct": {
            "type": "object",
            "properties": {
                "productId": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "handlers.UnreserveRequest": {
            "type": "object",
            "properties": {
                "orderNumber": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3456",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "ECommerce Inventory Management System",
	Description:      "ECommerce Inventory Management System used to add, reserve, unreserve and get stocks in your system",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
