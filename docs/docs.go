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
            "name": "Turing Zhu",
            "email": "qishiwenjun@163.com"
        },
        "license": {
            "name": "MIT License",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/common/currencys": {
            "get": {
                "description": "This endpoint returns all Huobi's supported trading currencies.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reference Data"
                ],
                "summary": "Get all Supported Currencies",
                "operationId": "SupportCurrenciesV1",
                "responses": {
                    "200": {
                        "description": "supported currency",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "type": "string"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/common/symbols": {
            "get": {
                "description": "This endpoint returns all Huobi's supported trading symbol.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reference Data"
                ],
                "summary": "Get all Supported Trading Symbol",
                "operationId": "SupportedTradingSymbolV1",
                "responses": {
                    "200": {
                        "description": "Supported Trading Symbol List",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.ResponseV1"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.Symbol"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/common/timestamp": {
            "get": {
                "description": "This endpoint returns the current timestamp, i.e. the number of milliseconds that have elapsed since 00:00:00 UTC on 1 January 1970.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reference Data"
                ],
                "summary": "Get Current Timestamp",
                "operationId": "TimestampV1",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/v2/reference/currencies": {
            "get": {
                "description": "API user could query static reference information for each currency, as well as its corresponding chain(s). (Public Endpoint)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reference Data"
                ],
                "summary": "Currency \u0026 Chains",
                "operationId": "CurrencyAndChainsV2",
                "parameters": [
                    {
                        "type": "string",
                        "description": "btc, ltc, etc ...(available currencies in Huobi Global)",
                        "name": "currency",
                        "in": "path"
                    },
                    {
                        "type": "boolean",
                        "description": "true or false (if not filled, default value is true)",
                        "name": "authorizedUser",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Currency \u0026 Chains list",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/models.Currency"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "chains": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/models.Chain"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Chain": {
            "type": "object",
            "properties": {
                "baseChain": {
                    "type": "string"
                },
                "baseChainProtocol": {
                    "type": "string"
                },
                "chain": {
                    "type": "string"
                },
                "deposit_status": {
                    "type": "string"
                },
                "displayName": {
                    "type": "string"
                },
                "isDynamic": {
                    "type": "boolean"
                },
                "maxTransactFeeWithdraw": {
                    "type": "number"
                },
                "maxWithdrawAmt": {
                    "type": "number"
                },
                "minDepositAmt": {
                    "type": "number"
                },
                "minTransactFeeWithdraw": {
                    "type": "number"
                },
                "minWithdrawAmt": {
                    "type": "number"
                },
                "numOfConfirmations": {
                    "type": "integer"
                },
                "numOfFastConfirmations": {
                    "type": "integer"
                },
                "transactFeeRateWithdraw": {
                    "type": "number"
                },
                "transactFeeWithdraw": {
                    "type": "number"
                },
                "withdrawFeeType": {
                    "type": "string"
                },
                "withdrawPrecision": {
                    "type": "integer"
                },
                "withdrawQuotaPerDay": {
                    "type": "number"
                },
                "withdrawQuotaPerYear": {
                    "type": "number"
                },
                "withdrawQuotaTotal": {
                    "type": "number"
                },
                "withdrawStatus": {
                    "type": "string"
                }
            }
        },
        "models.Currency": {
            "type": "object",
            "properties": {
                "chains": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Chain"
                    }
                },
                "currency": {
                    "type": "string"
                },
                "instStatus": {
                    "type": "string"
                }
            }
        },
        "models.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.ResponseV1": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "models.Symbol": {
            "type": "object",
            "properties": {
                "amount-precision": {
                    "type": "integer"
                },
                "base-currency": {
                    "type": "string"
                },
                "buy-market-max-order-value": {
                    "type": "number"
                },
                "charge-time": {
                    "type": "string"
                },
                "init-nav": {
                    "type": "number"
                },
                "leverage_ratio": {
                    "type": "number"
                },
                "limit-order-max-order-amt": {
                    "type": "number"
                },
                "limit-order-min-order-amt": {
                    "type": "number"
                },
                "max-order-amt": {
                    "type": "number"
                },
                "max-order-value": {
                    "type": "number"
                },
                "mgmt-fee-rate": {
                    "type": "number"
                },
                "min-order-amt": {
                    "type": "number"
                },
                "min-order-value": {
                    "type": "number"
                },
                "price-precision": {
                    "type": "integer"
                },
                "quote-currency": {
                    "type": "string"
                },
                "rebal-threshold": {
                    "type": "number"
                },
                "rebal-time": {
                    "type": "string"
                },
                "sell-market-min-order-amt": {
                    "type": "number"
                },
                "sell_market_max_order_amt": {
                    "type": "number"
                },
                "state": {
                    "type": "string"
                },
                "symbol": {
                    "type": "string"
                },
                "symbol-partition": {
                    "type": "string"
                },
                "underlying": {
                    "type": "string"
                },
                "value-precision": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "AccessKeyId",
            "in": "path"
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
	Version:     "0.1",
	Host:        "api.testnet.huobi.pro",
	BasePath:    "/",
	Schemes:     []string{"https"},
	Title:       "Huobi API",
	Description: "Huobi API",
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
