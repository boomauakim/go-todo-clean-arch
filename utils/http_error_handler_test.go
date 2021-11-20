package utils_test

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/boomauakim/go-todo-clean-arch/domain"
	"github.com/boomauakim/go-todo-clean-arch/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestErrBadRequest(t *testing.T) {
	app := fiber.New()
	app.Post("/todos", func(c *fiber.Ctx) error {
		return utils.ErrorHandler(c, fiber.ErrBadRequest)
	})

	res, err := app.Test(httptest.NewRequest(fiber.MethodPost, "/todos", nil))

	body, _ := ioutil.ReadAll(res.Body)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
	assert.Equal(t, `{"error":{"message":"The request was unacceptable, often due to missing a required parameter."}}`, string(body))
}

func TestErrUnprocessableEntity(t *testing.T) {
	app := fiber.New()
	app.Post("/todos", func(c *fiber.Ctx) error {
		return utils.ErrorHandler(c, fiber.ErrUnprocessableEntity)
	})

	res, err := app.Test(httptest.NewRequest(fiber.MethodPost, "/todos", nil))

	body, _ := ioutil.ReadAll(res.Body)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
	assert.Equal(t, `{"error":{"message":"The request was unacceptable, often due to missing a required parameter."}}`, string(body))
}

func TestErrNotFound(t *testing.T) {
	app := fiber.New()
	app.Get("/todos", func(c *fiber.Ctx) error {
		return utils.ErrorHandler(c, domain.ErrNotFound)
	})

	res, err := app.Test(httptest.NewRequest(fiber.MethodGet, "/todos", nil))

	body, _ := ioutil.ReadAll(res.Body)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusNotFound, res.StatusCode)
	assert.Equal(t, `{"error":{"message":"The requested resource doesn't exist."}}`, string(body))
}

func TestErrInternalServerError(t *testing.T) {
	app := fiber.New()
	app.Get("/todos", func(c *fiber.Ctx) error {
		return utils.ErrorHandler(c, fiber.ErrInternalServerError)
	})

	res, err := app.Test(httptest.NewRequest(fiber.MethodGet, "/todos", nil))

	body, _ := ioutil.ReadAll(res.Body)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, res.StatusCode)
	assert.Equal(t, `{"error":{"message":"Something went wrong."}}`, string(body))
}
