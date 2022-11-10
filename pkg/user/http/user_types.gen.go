// Package http provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package http

// NewUser defines model for NewUser.
type NewUser struct {
	// Name of the user
	Name string `json:"name"`

	// Type of the user
	Tag *string `json:"tag,omitempty"`
}

// User defines model for User.
type User struct {
	// Unique id of the user
	Id int64 `json:"id"`

	// Name of the user
	Name string `json:"name"`

	// Type of the user
	Tag *string `json:"tag,omitempty"`
}

// Error defines model for error.
type Error = interface{}

// FindUsersParams defines parameters for FindUsers.
type FindUsersParams struct {
	// maximum number of results to return
	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`

	// tags to filter by
	Tags *[]string `form:"tags,omitempty" json:"tags,omitempty"`
}

// CreateUserJSONBody defines parameters for CreateUser.
type CreateUserJSONBody = NewUser

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody = CreateUserJSONBody