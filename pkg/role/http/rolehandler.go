// Package http implement the oapi-codegen ServerInterface
package http

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

// RoleHandler handle users related apis.
type RoleHandler struct{}

// FindRoles - Returns all roles
// (GET /users)
func (h *RoleHandler) FindRoles(c *fiber.Ctx, params FindRolesParams) error {
	var limit int32
	var tags []string
	if params.Limit != nil {
		limit = *params.Limit
	}
	if params.Tags != nil {
		tags = *params.Tags
	}
	log.Printf("FindRoles with limit = %d and tags = %v\n", limit, tags)
	return c.SendStatus(http.StatusOK)
}

// CreateRole - Creates a new role
// (POST /users)
func (h *RoleHandler) CreateRole(c *fiber.Ctx) error {
	requestBody := new(NewRole)
	log.Println("CreateRole")

	err := c.BodyParser(requestBody)
	if err != nil {
		return errors.Wrap(err, "failed to bind the request body")
	}
	log.Printf("  request body +v = %+v", requestBody)

	return c.SendStatus(http.StatusOK)
}

// FindRoleByID - Returns a role by ID
// (GET /users/{id})
func (h *RoleHandler) FindRoleByID(c *fiber.Ctx, id int64) error {
	log.Printf("FindRoleByID id = %d\n", id)
	return c.SendStatus(http.StatusOK)
}
