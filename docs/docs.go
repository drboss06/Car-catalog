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
        "/cars": {
            "get": {
                "description": "Получение списка автомобилей с возможностью фильтрации по номеру и марке.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cars"
                ],
                "summary": "Получение списка автомобилей",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Номер автомобиля для фильтрации",
                        "name": "regNum",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Марка автомобиля для фильтрации",
                        "name": "mark",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Номер страницы",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Размер страницы",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешный ответ",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Car_model"
                            }
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера"
                    }
                }
            },
            "post": {
                "description": "Добавляет новый автомобиль.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cars"
                ],
                "summary": "Добавление нового автомобиля",
                "parameters": [
                    {
                        "description": "Данные нового автомобиля",
                        "name": "model.CarAdd",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CarAdd"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Car added successfully"
                    },
                    "400": {
                        "description": "Ошибка запроса"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/cars/{id}": {
            "put": {
                "description": "Изменяет информацию об автомобиле по его идентификатору.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cars"
                ],
                "summary": "Изменение информации об автомобиле",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор автомобиля",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Данные для обновления автомобиля",
                        "name": "model.CarUpdate",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CarUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Car updated successfully"
                    },
                    "400": {
                        "description": "Ошибка запроса"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            },
            "delete": {
                "description": "Удаляет автомобиль по его идентификатору.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cars"
                ],
                "summary": "Удаление автомобиля",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор автомобиля",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Car deleted successfully"
                    },
                    "400": {
                        "description": "Ошибка запроса"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.CarAdd": {
            "type": "object",
            "properties": {
                "regNum": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.CarUpdate": {
            "type": "object",
            "properties": {
                "mark": {
                    "type": "string"
                },
                "model": {
                    "type": "string"
                },
                "regNum": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "model.Car_model": {
            "type": "object",
            "properties": {
                "mark": {
                    "type": "string"
                },
                "model": {
                    "type": "string"
                },
                "regNum": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
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
