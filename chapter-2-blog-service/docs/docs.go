// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://github.com/ChrisLeejing/go-programming-tour-book-exercise",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/articles": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Article 文章"
                ],
                "summary": "新增文章",
                "parameters": [
                    {
                        "type": "string",
                        "description": "JWT Token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "CreateArticleRequest",
                        "name": "CreateArticleRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/validate.CreateArticleRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.TagSwagger"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/tags": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tag 标签"
                ],
                "summary": "获取多个标签",
                "parameters": [
                    {
                        "type": "string",
                        "description": "JWT Token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "标签名称",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "type": "integer",
                        "default": 1,
                        "description": "状态",
                        "name": "state",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.TagSwagger"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tag 标签"
                ],
                "summary": "新增标签",
                "parameters": [
                    {
                        "type": "string",
                        "description": "JWT Token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "CreateTagRequest",
                        "name": "CreateTagRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/validate.CreateTagRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.TagSwagger"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/tags/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tag 标签"
                ],
                "summary": "获取标签",
                "parameters": [
                    {
                        "type": "string",
                        "description": "JWT Token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "标签 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tag 标签"
                ],
                "summary": "更新标签",
                "parameters": [
                    {
                        "type": "string",
                        "description": "JWT Token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "UpdateTagRequest",
                        "name": "UpdateTagRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/validate.UpdateTagRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.TagSwagger"
                            }
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tag 标签"
                ],
                "summary": "删除标签",
                "parameters": [
                    {
                        "type": "string",
                        "description": "JWT Token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "标签 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            }
        },
        "/auth": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token 鉴权"
                ],
                "summary": "获取认证token",
                "parameters": [
                    {
                        "description": "GetAuthRequest",
                        "name": "GetAuthRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/validate.GetAuthRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            }
        },
        "/upload/file": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "File 文件"
                ],
                "summary": "文件上传",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "文件类型",
                        "name": "type",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errcode.Error": {
            "type": "object"
        },
        "model.TagSwagger": {
            "type": "object"
        },
        "validate.CreateArticleRequest": {
            "type": "object",
            "required": [
                "created_by",
                "title"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "cover_img_url": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "desc": {
                    "type": "string"
                },
                "state": {
                    "type": "integer"
                },
                "tag_ids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "validate.CreateTagRequest": {
            "type": "object",
            "required": [
                "created_by"
            ],
            "properties": {
                "created_by": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "state": {
                    "type": "integer"
                }
            }
        },
        "validate.GetAuthRequest": {
            "type": "object",
            "required": [
                "app_key",
                "app_secret"
            ],
            "properties": {
                "app_key": {
                    "description": "https://github.com/gin-gonic/gin#model-binding-and-validation",
                    "type": "string"
                },
                "app_secret": {
                    "type": "string"
                }
            }
        },
        "validate.UpdateTagRequest": {
            "type": "object",
            "required": [
                "id",
                "modified_by"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "modified_by": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "state": {
                    "type": "integer"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "v1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "博客系统",
	Description: "Go 语言编程之旅: 一起用 Go 做项目",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
