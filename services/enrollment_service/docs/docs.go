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
        "/enrollments/me": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieves all enrollments of the authenticated user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Enrollments"
                ],
                "summary": "Get My Enrollments",
                "responses": {
                    "200": {
                        "description": "Enrollments retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    }
                }
            }
        },
        "/progress/lesson/{lesson_id}/complete": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Marks a specific lesson as complete for the authenticated user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Progress"
                ],
                "summary": "Mark Lesson as Complete",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Lesson ID",
                        "name": "lesson_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Lesson marked as complete successfully",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid lesson ID",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    }
                }
            }
        },
        "/progress/lesson/{lesson_id}/incomplete": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Marks a specific lesson as incomplete for the authenticated user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Progress"
                ],
                "summary": "Mark Lesson as Incomplete",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Lesson ID",
                        "name": "lesson_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Lesson marked as incomplete successfully",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid lesson ID",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    }
                }
            }
        },
        "/progress/me": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieves the progress of all lessons for the authenticated user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Progress"
                ],
                "summary": "Get My Course Progress",
                "responses": {
                    "200": {
                        "description": "Course progress retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    }
                }
            }
        },
        "/v1/api/certificates/my": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve all certificates associated with the authenticated user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Certificates"
                ],
                "summary": "Get user's certificates",
                "responses": {
                    "200": {
                        "description": "Certificates successfully retrieved",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dtos.CertificateDTO"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "404": {
                        "description": "No certificates found",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    }
                }
            }
        },
        "/v1/api/certificates/{enrollment_id}": {
            "get": {
                "description": "Retrieve a certificate associated with a specific enrollment ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Certificates"
                ],
                "summary": "Get certificate by enrollment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Enrollment ID",
                        "name": "enrollment_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Certificate successfully retrieved",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dtos.CertificateDTO"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid enrollment ID",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "404": {
                        "description": "Certificate not found",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    }
                }
            }
        },
        "/v1/api/courses/{course_id}/enrollments": {
            "get": {
                "description": "Retrieves all enrollments for a specific course.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Enrollments"
                ],
                "summary": "Get Enrollments for a Course",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Course ID",
                        "name": "course_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Enrollments retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid course ID",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "404": {
                        "description": "No enrollments found for this course",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    }
                }
            }
        },
        "/v1/api/enrollments/user/{user_id}/course/{course_id}": {
            "get": {
                "description": "Retrieves a specific enrollment using the user ID and course ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Enrollments"
                ],
                "summary": "Get Enrollment by User and Course",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Course ID",
                        "name": "course_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Enrollment retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid user or course ID",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "404": {
                        "description": "Enrollment not found",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    }
                }
            }
        },
        "/v1/api/enrollments/{enrollent_id}/complete": {
            "post": {
                "description": "Complete a course and generate certificate for the enrollment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Enrollment Management"
                ],
                "summary": "Mark course as completed",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Enrollment ID",
                        "name": "enrollent_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Course completed successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dtos.CertificateDTO"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid enrollment ID",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "404": {
                        "description": "Enrollment not found",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    }
                }
            }
        },
        "/v1/api/enrollments/{enrollment_id}": {
            "get": {
                "description": "Retrieves a specific enrollment using its ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Enrollments"
                ],
                "summary": "Get Enrollment by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Enrollment ID",
                        "name": "enrollment_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Enrollment retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid enrollment ID",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "404": {
                        "description": "Enrollment not found",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    }
                }
            }
        },
        "/v1/api/subscriptions": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Creates a new subscription for a user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subscriptions"
                ],
                "summary": "Create Subscription",
                "parameters": [
                    {
                        "description": "Subscription details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.SubscriptionInsertDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Subscription created successfully",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request data",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    }
                }
            }
        },
        "/v1/api/subscriptions/cancel/{lesson_id}": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Cancels the subscription for the authenticated user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subscriptions"
                ],
                "summary": "Cancel My Subscription",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Lesson ID",
                        "name": "lesson_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Subscription cancelled successfully",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid parameter",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    }
                }
            }
        },
        "/v1/api/subscriptions/me": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieves the subscription details for the authenticated user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subscriptions"
                ],
                "summary": "Get My Subscription",
                "responses": {
                    "200": {
                        "description": "Subscription retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    }
                }
            }
        },
        "/v1/api/subscriptions/{subscription_id}": {
            "delete": {
                "description": "Deletes a subscription by its ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subscriptions"
                ],
                "summary": "Delete Subscription",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Subscription ID",
                        "name": "subscription_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Subscription deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid subscription ID",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    }
                }
            }
        },
        "/v1/api/subscriptions/{user_id}/change/{sub_type}": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Changes the subscription type for the authenticated user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subscriptions"
                ],
                "summary": "Change My Subscription Type",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Subscription Type",
                        "name": "sub_type",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Subscription type updated successfully",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid subscription type or user ID",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    }
                }
            }
        },
        "/v1/api/users/{user_id}/courses/{course_id}/enroll": {
            "post": {
                "description": "Create a new enrollment and track record for a user-course combination",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Enrollment Management"
                ],
                "summary": "Enroll user in course",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Course ID",
                        "name": "course_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Enrollment created successfully",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid user/course ID",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "409": {
                        "description": "User already enrolled",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    }
                }
            }
        },
        "/v1/api/users/{user_id}/enrollments/{enrollent_id}": {
            "delete": {
                "description": "Cancel an existing enrollment for a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Enrollment Management"
                ],
                "summary": "Cancel enrollment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Enrollment ID",
                        "name": "enrollent_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Enrollment cancelled successfully",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid user/enrollment ID",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "404": {
                        "description": "Enrollment not found",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.ApiResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.CertificateDTO": {
            "type": "object",
            "properties": {
                "certificate_url": {
                    "description": "The URL where the certificate can be accessed.\nExample: https://example.com/certificates/12345",
                    "type": "string"
                },
                "enrollment_id": {
                    "description": "The ID of the enrollment associated with the certificate.\nExample: 550e8400-e29b-41d4-a716-446655440001",
                    "type": "string"
                },
                "expires_at": {
                    "description": "The date and time when the certificate expires (optional).\nFormat: date-time\nExample: 2024-10-01T12:34:56",
                    "type": "string"
                },
                "id": {
                    "description": "The unique identifier of the certificate.\nExample: 550e8400-e29b-41d4-a716-446655440000",
                    "type": "string"
                },
                "issued_at": {
                    "description": "The date and time when the certificate was issued.\nFormat: date-time\nExample: 2023-10-01T12:34:56",
                    "type": "string"
                }
            }
        },
        "dtos.SubscriptionInsertDTO": {
            "type": "object",
            "required": [
                "payment_id",
                "plan_name",
                "status",
                "type",
                "user_id"
            ],
            "properties": {
                "payment_id": {
                    "description": "The ID of the payment associated with the subscription.\nRequired: true\nExample: 550e8400-e29b-41d4-a716-446655440002",
                    "type": "string"
                },
                "plan_name": {
                    "description": "The name of the subscription plan.\nRequired: true\nMinimum length: 3\nMaximum length: 50\nExample: Premium Monthly",
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 3
                },
                "status": {
                    "description": "The status of the subscription.\nRequired: true\nEnum: active, inactive, cancelled\nExample: active",
                    "type": "string"
                },
                "type": {
                    "description": "The type of the subscription.\nRequired: true\nEnum: monthly, yearly\nExample: monthly",
                    "type": "string"
                },
                "user_id": {
                    "description": "The ID of the user associated with the subscription.\nRequired: true\nExample: 550e8400-e29b-41d4-a716-446655440001",
                    "type": "string"
                }
            }
        },
        "response.ApiResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "correlationId": {
                    "type": "string"
                },
                "data": {},
                "error_code": {},
                "message": {},
                "page": {
                    "type": "integer"
                },
                "perPage": {
                    "type": "integer"
                },
                "success": {
                    "type": "boolean"
                },
                "timestamp": {
                    "type": "string"
                },
                "totalCount": {
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
