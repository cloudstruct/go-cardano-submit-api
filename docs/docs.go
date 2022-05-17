// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "CloudStruct",
            "url": "https://cloudstruct.net",
            "email": "support@cloudstruct.net"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/submit/tx": {
            "post": {
                "description": "Submit an already serialized transaction to the network.",
                "produces": [
                    "application/json"
                ],
                "summary": "Submit Tx",
                "parameters": [
                    {
                        "enum": [
                            "application/cbor"
                        ],
                        "type": "string",
                        "description": "Content type",
                        "name": "Content-Type",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "3.1.0",
	Host:             "localhost:8090",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "go-cardano-submit-api",
	Description:      "Cardano Submit API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
