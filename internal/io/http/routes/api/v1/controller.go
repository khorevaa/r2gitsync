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
		r.Post("projects", ctr.PostProjects)
		r.Post("projects/:uuid", ctr.PostProject)

		// r.Post("projects/:uuid/hook", ctr.PostProjectWebhook)

		// Получает мастер репозитория для проекта
		r.Get("projects/:uuid/master", ctr.GEtProjectMaster)
		// Создает мастер репозитория для проекта
		r.Post("projects/:uuid/master", ctr.PostProjectMaster)
		// Изменяет мастер репозитория для проекта
		r.Put("projects/:uuid/master", ctr.PutProjectMaster)
		// Отключает мастер репозитория для проекта
		r.Delete("projects/:uuid/master", ctr.DeleteProjectMaster)

	}
}
