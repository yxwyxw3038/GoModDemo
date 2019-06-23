// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-06-23 12:32:15.4877175 +0800 CST m=+0.043850601

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "目前仅仅是一个demo",
        "title": "Rest API",
        "contact": {},
        "license": {},
        "version": "1.0"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/Login": {
            "post": {
                "description": "验证用户名密码有效性 accountName=aaa\u0026passWord=base64(bbb) 成功输出Token",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "登录验证服务",
                "parameters": [
                    {
                        "type": "string",
                        "default": "yxw",
                        "description": "用户名",
                        "name": "accountName",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "123",
                        "description": "密码",
                        "name": "passWord",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"Code\":1,\"Data\":{Token},\"Message\":\"\"} or {\"Code\":-1,\"Data\":{},\"Message\":\"错误提示\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/helper": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "State",
                        "name": "state",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"Code\":,\"Data\":{},\"Message\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
