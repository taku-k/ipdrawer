package serverpb

const (
	swagger = `{
  "swagger": "2.0",
  "info": {
    "title": "server/serverpb/server.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/api/v0/ip/{ip}/activate": {
      "post": {
        "operationId": "ActivateIP",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/serverpbActivateIPResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "ip",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "IPServiceV0"
        ]
      }
    },
    "/api/v0/ip/{ip}/network": {
      "get": {
        "operationId": "GetNetworkIncludingIP",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/serverpbGetNetworkIncludingIPResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "ip",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "IPServiceV0"
        ]
      }
    },
    "/api/v0/network/{ip}/{mask}": {
      "get": {
        "operationId": "GetNetwork",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/serverpbGetNetworkResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "ip",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "mask",
            "in": "path",
            "required": true,
            "type": "integer",
            "format": "int32"
          }
        ],
        "tags": [
          "NetworkServiceV0"
        ]
      }
    },
    "/api/v0/network/{ip}/{mask}/create": {
      "post": {
        "operationId": "CreateNetwork",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/serverpbCreateNetworkResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "ip",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "mask",
            "in": "path",
            "required": true,
            "type": "integer",
            "format": "int32"
          }
        ],
        "tags": [
          "NetworkServiceV0"
        ]
      }
    },
    "/api/v0/network/{ip}/{mask}/drawip": {
      "get": {
        "operationId": "DrawIP",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/serverpbDrawIPResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "ip",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "mask",
            "in": "path",
            "required": true,
            "type": "integer",
            "format": "int32"
          }
        ],
        "tags": [
          "NetworkServiceV0"
        ]
      }
    },
    "/api/v0/network/{ip}/{mask}/pool/create": {
      "post": {
        "operationId": "CreatePool",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/serverpbCreatePoolResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "ip",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "mask",
            "in": "path",
            "required": true,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/serverpbCreatePoolRequest"
            }
          }
        ],
        "tags": [
          "NetworkServiceV0"
        ]
      }
    }
  },
  "definitions": {
    "serverpbActivateIPResponse": {
      "type": "object"
    },
    "serverpbCreateNetworkResponse": {
      "type": "object"
    },
    "serverpbCreatePoolRequest": {
      "type": "object",
      "properties": {
        "ip": {
          "type": "string"
        },
        "mask": {
          "type": "integer",
          "format": "int32"
        },
        "pool": {
          "$ref": "#/definitions/serverpbPool"
        }
      }
    },
    "serverpbCreatePoolResponse": {
      "type": "object"
    },
    "serverpbDrawIPResponse": {
      "type": "object",
      "properties": {
        "ip": {
          "type": "string"
        }
      }
    },
    "serverpbGetNetworkIncludingIPResponse": {
      "type": "object"
    },
    "serverpbGetNetworkResponse": {
      "type": "object",
      "properties": {
        "ipnet": {
          "type": "string"
        },
        "default_gateways": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "broadcast": {
          "type": "string"
        },
        "netmask": {
          "type": "string"
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/serverpbTag"
          }
        }
      }
    },
    "serverpbPool": {
      "type": "object",
      "properties": {
        "start": {
          "type": "string"
        },
        "end": {
          "type": "string"
        },
        "status": {
          "type": "integer",
          "format": "int32"
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/serverpbTag"
          }
        }
      }
    },
    "serverpbTag": {
      "type": "object",
      "properties": {
        "key": {
          "type": "string"
        },
        "value": {
          "type": "string"
        }
      }
    }
  }
}
`
)