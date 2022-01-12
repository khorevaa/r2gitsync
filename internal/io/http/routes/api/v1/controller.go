package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khorevaa/r2gitsync/internal/bl"
)

type Controller struct {
	bl *bl.BL
}

func New(bl *bl.BL) *Controller {
	return &Controller{bl: bl}
}

func (ctr *Controller) InitRouter(router fiber.Router) {
	r := router.Group("api/v1")

	{
		r.Get("projects", ctr.GetProjects)
	}
}
