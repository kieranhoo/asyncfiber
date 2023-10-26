package controller

import (
	"qrcheckin/internal/handler/schema"
	"qrcheckin/internal/handler/service"
	"qrcheckin/internal/validator"

	"github.com/gofiber/fiber/v2"
)

// SignUp.
// @Description Register account for admin.
// @Tags auth
// @Accept json
// @Produce json
// @Param sign_up body schema.SignInRequest true "Sign Up"
// @Success 200 {object} schema.SignUpResponse
// @Router /api/v1/auth/signup [POST]
func SignUp(c *fiber.Ctx) error {
	var req schema.SignUpRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     err.Error(),
		})
	}
	if _validate := validator.Validator(req); _validate != "" {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     _validate,
		})
	}
	if err := service.SignUp(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     err.Error(),
		})
	}
	return c.JSON(schema.SignUpResponse{
		Success: true,
		Msg:     "your account was created successfully",
	})
}

// SignIn.
// @Description Sign in account for admin.
// @Tags auth
// @Accept json
// @Produce json
// @Param sign_in body schema.SignInRequest true "Sign In"
// @Success 200 {object} schema.SignInResponse
// @Router /api/v1/auth/signin [POST]
func SignIn(c *fiber.Ctx) error {
	var req schema.SignInRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     err.Error(),
		})
	}
	if _validate := validator.Validator(req); _validate != "" {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     _validate,
		})
	}
	token, err := service.SignIn(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     err.Error(),
		})
	}
	return c.JSON(schema.SignInResponse{
		Success: true,
		Token:   token,
		Email:   req.Email,
	})
}
