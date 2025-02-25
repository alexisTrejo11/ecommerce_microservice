package routes

import (
	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/input/http/v1/handlers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(r fiber.Router, authHandler *handlers.AuthHandler) {
	r.Get("/home", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"Welcome": "Welcome to User Service",
		})
	})

	authPath := r.Group("v1/api")
	authPath.Post("/register", authHandler.Register)
	authPath.Post("/login", authHandler.Login)
	authPath.Post("/activate-account/:token", authHandler.ActivateAccount)
	authPath.Post("/logout/:refresh_token", authHandler.Logout)
	authPath.Post("/logout-all", authHandler.LogoutAll)
	authPath.Get("/refresh_acces_token/:refresh_token", authHandler.RefreshAccesToken)
}

func UserAddressRoutes(r fiber.Router, addresHandler *handlers.UserAddressHandler) {
	addressPath := r.Group("v1/api/users/address")
	addressPath.Get("", addresHandler.MyAddresses)
	addressPath.Post("", addresHandler.AddAddress)
	addressPath.Put("/:id", addresHandler.UpdateMyAddress)
	addressPath.Delete("/:id", addresHandler.DeleteAddress)
}

func UserMFARoutes(r fiber.Router, mfaHandler *handlers.UserMfaHandler) {
	addressPath := r.Group("v1/api/users/mfa")
	addressPath.Post("", mfaHandler.EnableMfa)
	addressPath.Delete("", mfaHandler.DisableMfa)
	addressPath.Get("", mfaHandler.GetMfa)
	addressPath.Put("", mfaHandler.VerifyMfa)
}

func SessionRoutes(r fiber.Router, sessionHandler *handlers.SessionHandler) {
	path := r.Group("v1/api/sessions")
	path.Get("/:id", sessionHandler.GetSessionByUserId)
	path.Delete("/:id/user/:user_id", sessionHandler.DeleteSessionById)
}
