// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Ngô Hoài Phong",
            "url": "http://www.swagger.io/support",
            "email": "hoaiphong01012002@gmai.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/category/create": {
            "post": {
                "description": "create category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "create category",
                "responses": {}
            }
        },
        "/admin/category/update": {
            "put": {
                "description": "create category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "update category",
                "responses": {}
            }
        },
        "/admin/product/create": {
            "post": {
                "description": "create product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "create product",
                "responses": {}
            }
        },
        "/admin/product/enable": {
            "patch": {
                "description": "update enable product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "update enable product",
                "responses": {}
            }
        },
        "/admin/product/get-update-date": {
            "get": {
                "description": "get product by updates date",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "get product by updates date",
                "responses": {}
            }
        },
        "/admin/product/get-who-update": {
            "get": {
                "description": "get product by who updates",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "get product by who updates",
                "responses": {}
            }
        },
        "/admin/product/remove-first": {
            "patch": {
                "description": "update product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "update product",
                "responses": {}
            }
        },
        "/admin/product/update": {
            "post": {
                "description": "update product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "update product",
                "responses": {}
            }
        },
        "/admin/product/update-enable": {
            "patch": {
                "description": "update product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "update-enable product",
                "responses": {}
            }
        },
        "/client/category/delete": {
            "patch": {
                "description": "delete category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "delete category",
                "responses": {}
            }
        },
        "/client/category/delete-second": {
            "delete": {
                "description": "delete category with admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "delete category with admin",
                "responses": {}
            }
        },
        "/client/category/get-all": {
            "get": {
                "description": "get all categories",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "get all categories",
                "responses": {}
            }
        },
        "/client/category/get-enable/:enable": {
            "get": {
                "description": "get category by update date",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "get category by update date",
                "responses": {}
            }
        },
        "/client/category/get-update-date/:update_date": {
            "get": {
                "description": "get category by update date",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "get category by update date",
                "responses": {}
            }
        },
        "/client/category/get/:category_id": {
            "get": {
                "description": "get category by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "get category by id",
                "responses": {}
            }
        },
        "/client/category/get2/:category_id": {
            "get": {
                "description": "get category by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "get category by id",
                "responses": {}
            }
        },
        "/client/category/resolve": {
            "patch": {
                "description": "resolve category with admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "resolve category with admin",
                "responses": {}
            }
        },
        "/client/get-all-user": {
            "get": {
                "description": "get all user item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "get all user",
                "responses": {}
            }
        },
        "/client/get-user": {
            "get": {
                "description": "get a new user item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "get user",
                "responses": {}
            }
        },
        "/client/get-user-role": {
            "get": {
                "description": "get user by role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "get user role",
                "responses": {}
            }
        },
        "/client/get-user/email": {
            "get": {
                "description": "get user by email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "get user email",
                "responses": {}
            }
        },
        "/client/get-user/role": {
            "get": {
                "description": "get user by role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "get user role",
                "responses": {}
            }
        },
        "/client/get-user/username": {
            "get": {
                "description": "get a new user item by username",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "get user",
                "responses": {}
            }
        },
        "/client/login/email": {
            "post": {
                "description": "login user item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "login user",
                "parameters": [
                    {
                        "description": "login user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SignInInput"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/client/login/username": {
            "post": {
                "description": "login user with username",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "login user",
                "parameters": [
                    {
                        "description": "login user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/client/logout": {
            "get": {
                "description": "logout item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "logout user",
                "responses": {}
            }
        },
        "/client/product/get-all": {
            "get": {
                "description": "get all products",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "get all products",
                "responses": {}
            }
        },
        "/client/product/get-id": {
            "get": {
                "description": "get product by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "get product by id",
                "responses": {}
            }
        },
        "/client/product/get-name": {
            "get": {
                "description": "get product by name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "get product by name",
                "responses": {}
            }
        },
        "/client/product/get-price": {
            "get": {
                "description": "get product by price",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "get product by price",
                "responses": {}
            }
        },
        "/client/refresh": {
            "get": {
                "description": "refresh token item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "refresh token user",
                "responses": {}
            }
        },
        "/client/register": {
            "post": {
                "description": "Create a new user item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "register user",
                "parameters": [
                    {
                        "description": "register user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/client/update": {
            "put": {
                "description": "update a new user item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "login user",
                "parameters": [
                    {
                        "description": "login user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "model.SignInInput": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "avatarUser": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "enable": {
                    "type": "integer"
                },
                "fullName": {
                    "type": "string"
                },
                "id": {
                    "description": "Tên thuộc tính được đặt trong golang phải là ID nếu kiểu dữ liệu là uuid.UUID",
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "photo": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "provider": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                },
                "verified": {
                    "type": "boolean"
                },
                "who_updates": {
                    "type": "string"
                }
            }
        },
        "model.UserRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000/",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Cỏ Studio API",
	Description:      "Đây là API của Cỏ Studio,\nViệc sử dụng API này phải có sự đồng ý của Mr. Phong",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
