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
        "/balance": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "balance"
                ],
                "summary": "Show balance",
                "responses": {
                    "200": {
                        "description": "OK"
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
        "/debt": {
            "get": {
                "description": "Lists all debts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "debts"
                ],
                "summary": "Show debts",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Debt id",
                        "name": "overdue",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
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
                "description": "Updates an existing debt",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "debts"
                ],
                "summary": "Update a debt",
                "parameters": [
                    {
                        "description": "Request body to update debt",
                        "name": "debt",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/debt.BodyUpdateDebt"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
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
                "description": "Creates a new debt",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "debts"
                ],
                "summary": "Create a debt",
                "parameters": [
                    {
                        "description": "Request body to add a debt",
                        "name": "debt",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/debt.BodyNewDebt"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
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
        "/debt/payment/": {
            "post": {
                "description": "Sets an existing debt paid/unpaid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "debts"
                ],
                "summary": "Set a debt paid/unpaid",
                "parameters": [
                    {
                        "description": "Request body to update a debt payment status",
                        "name": "debt",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/debt.BodyUpdateDebtPaymentStatus"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
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
        "/debt/{id}": {
            "delete": {
                "description": "Deletes an existing debt",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "debts"
                ],
                "summary": "Delete a debt",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Debt id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
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
        "/income": {
            "get": {
                "description": "Lists all incomes",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "incomes"
                ],
                "summary": "Show incomes",
                "responses": {
                    "200": {
                        "description": "OK"
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
                "description": "Updates an existing income",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "incomes"
                ],
                "summary": "Update an income",
                "parameters": [
                    {
                        "description": "Request body to update an income",
                        "name": "income",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/income.BodyUpdateIncome"
                        }
                    }
                ],
                "responses": {
                    "20": {
                        "description": ""
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
                "description": "Creates a new income",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "incomes"
                ],
                "summary": "Create an income",
                "parameters": [
                    {
                        "description": "Request body to add an income",
                        "name": "income",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/income.BodyNewIncome"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
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
        "/income/{id}": {
            "delete": {
                "description": "Deletes an existing income",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "incomes"
                ],
                "summary": "Delete an income",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Income id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
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
        "debt.BodyNewDebt": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "dueDate": {
                    "type": "string",
                    "example": "2024-12-25"
                },
                "name": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "debt.BodyUpdateDebt": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "dueDate": {
                    "type": "string",
                    "example": "2024-12-25"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "paymentDate": {
                    "type": "string",
                    "example": "2024-12-25"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "debt.BodyUpdateDebtPaymentStatus": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "paid": {
                    "type": "boolean"
                }
            }
        },
        "income.BodyNewIncome": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "receivedDate": {
                    "type": "string",
                    "example": "2024-12-25"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "income.BodyUpdateIncome": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "receivedDate": {
                    "type": "string",
                    "example": "2024-12-25"
                },
                "value": {
                    "type": "number"
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
