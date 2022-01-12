package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khorevaa/r2gitsync/internal/bl"
	v1 "github.com/khorevaa/r2gitsync/internal/io/http/routes/api/v1"
)

func SetUpRoutes(app *fiber.App, bl *bl.BL) {

	apiV1 := v1.New(bl)
	apiV1.InitRouter(app)

}
