// Package http implement the oapi-codegen ServerInterface
package http

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

// UserHandler handle users related apis.
type UserHandler struct{}

// FindUsers - Returns all users
// (GET /users)
func (h *UserHandler) FindUsers(c *fiber.Ctx, params FindUsersParams) error {
	var limit int32
	var tags []string
	if params.Limit != nil {
		limit = *params.Limit
	}
	if params.Tags != nil {
		tags = *params.Tags
	}
	log.Printf("FindUsers with limit = %d and tags = %v\n", limit, tags)
	return c.SendStatus(http.StatusOK)
}

// CreateUser - Creates a new user
// (POST /users)
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	requestBody := new(NewUser)
	log.Println("CreateUser")

	err := c.BodyParser(requestBody)
	if err != nil {
		return errors.Wrap(err, "failed to bind the request body")
	}
	log.Printf("  request body +v = %+v", requestBody)

	return c.SendStatus(http.StatusOK)
}

// FindUserByID - Returns a user by ID
// (GET /users/{id})
func (h *UserHandler) FindUserByID(c *fiber.Ctx, id int64) error {
	log.Printf("FindUserByID id = %d\n", id)
	return c.SendStatus(http.StatusOK)
}
