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
        "contact": {
            "name": "Trading AI",
            "url": "https://trading-ai.id",
            "email": "info@trading-ai.id"
        },
        "license": {
            "name": "Trading AI License"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/trading-ai/v1/earning/qrstring": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pos Service"
                ],
                "summary": "Request QR String from Hellobil",
                "operationId": "ReqQrString",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UUID Partner (example = 'PSM0002')",
                        "name": "Institutionid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Hardcoded",
                        "name": "Deviceid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": " ",
                        "name": "Geolocation",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Hardcoded (example = 'H2H')",
                        "name": "Channelid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Generated system by partner, all request POST validate with signature (example = 'a0acd4cc5bb2')",
                        "name": "Signature",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "(Example = 'pos-customer')",
                        "name": "Appsid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": " ",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "(Example = '1579666534')",
                        "name": "Timestamp",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Request Body",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ReqQrString"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": " ",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.BaseResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.ResQrString"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/trading-ai/v1/earning/request": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pos Service"
                ],
                "summary": "hit from apps mobile / partner user apps",
                "operationId": "ReqEarning",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UUID Partner (example = 'PSM0002')",
                        "name": "Institutionid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Hardcoded",
                        "name": "Deviceid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": " ",
                        "name": "Geolocation",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Hardcoded (example = 'H2H')",
                        "name": "Channelid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Generated system by partner, all request POST validate with signature (example = 'a0acd4cc5bb2')",
                        "name": "Signature",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "(Example = 'pos-customer')",
                        "name": "Appsid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": " ",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "(Example = '1579666534')",
                        "name": "Timestamp",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Request Body",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ReqEarningHandler"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": " ",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.BaseResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.ResQrString"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/trading-ai/v1/earning/status": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pos Service"
                ],
                "summary": "Check Status Hit by ottopoint-stamp",
                "operationId": "CheckStatus",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UUID Partner (example = 'PSM0002')",
                        "name": "Institutionid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Hardcoded",
                        "name": "Deviceid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": " ",
                        "name": "Geolocation",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Hardcoded (example = 'H2H')",
                        "name": "Channelid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Generated system by partner, all request POST validate with signature (example = 'a0acd4cc5bb2')",
                        "name": "Signature",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "(Example = 'pos-customer')",
                        "name": "Appsid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": " ",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "(Example = '1579666534')",
                        "name": "Timestamp",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": " ",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ReqCheckStatus"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": " ",
                        "schema": {
                            "$ref": "#/definitions/model.BaseResp"
                        }
                    }
                }
            }
        },
        "/trading-ai/v1/product/list": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pos Service"
                ],
                "summary": "Product List Hit by admin portal",
                "operationId": "ProductList",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UUID Partner (example = 'PSM0002')",
                        "name": "Institutionid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Hardcoded",
                        "name": "Deviceid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": " ",
                        "name": "Geolocation",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Hardcoded (example = 'H2H')",
                        "name": "Channelid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Generated system by partner, all request POST validate with signature (example = 'a0acd4cc5bb2')",
                        "name": "Signature",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "(Example = 'pos-customer')",
                        "name": "Appsid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": " ",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "(Example = '1579666534')",
                        "name": "Timestamp",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": " ",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ReqProductList"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": " ",
                        "schema": {
                            "$ref": "#/definitions/model.BaseResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.BaseResp": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "http status at header response api-gateway",
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "responseCode": {
                    "type": "string"
                },
                "responseDesc": {
                    "type": "string"
                }
            }
        },
        "model.PosDetailProduct": {
            "type": "object",
            "required": [
                "productId",
                "productName",
                "quantity",
                "unitPrice"
            ],
            "properties": {
                "productId": {
                    "type": "string"
                },
                "productName": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                },
                "unitPrice": {
                    "type": "integer"
                }
            }
        },
        "model.ReqCheckStatus": {
            "type": "object",
            "required": [
                "earningId"
            ],
            "properties": {
                "earningId": {
                    "type": "string"
                }
            }
        },
        "model.ReqEarningHandler": {
            "type": "object",
            "required": [
                "accountNo",
                "earningId",
                "requestAt"
            ],
            "properties": {
                "accountNo": {
                    "type": "string"
                },
                "earningId": {
                    "type": "string"
                },
                "requestAt": {
                    "type": "string"
                }
            }
        },
        "model.ReqProductList": {
            "type": "object",
            "required": [
                "invoiceNo"
            ],
            "properties": {
                "invoiceNo": {
                    "type": "string"
                }
            }
        },
        "model.ReqQrString": {
            "type": "object",
            "required": [
                "detailProduct",
                "invoiceNo",
                "merchantId",
                "staffId",
                "staffName",
                "storeId",
                "subTotalAmount",
                "totalAmount",
                "transactionTime"
            ],
            "properties": {
                "detailProduct": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.PosDetailProduct"
                    }
                },
                "earningCode": {
                    "type": "string"
                },
                "invoiceNo": {
                    "type": "string"
                },
                "merchantId": {
                    "type": "string"
                },
                "serviceCharge": {
                    "type": "integer"
                },
                "staffId": {
                    "type": "string"
                },
                "staffName": {
                    "type": "string"
                },
                "storeId": {
                    "type": "string"
                },
                "subTotalAmount": {
                    "type": "integer"
                },
                "tax": {
                    "type": "integer"
                },
                "totalAmount": {
                    "type": "integer"
                },
                "transactionTime": {
                    "type": "string"
                }
            }
        },
        "model.ResQrString": {
            "type": "object",
            "properties": {
                "qrString": {
                    "type": "string"
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
