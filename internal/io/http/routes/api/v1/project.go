package v1

import (
	"github.com/gofiber/fiber/v2"
)

func (ctr *Controller) GetProjects(c *fiber.Ctx) error {

	projects, err := ctr.bl.ProjectsLogic.GetProjects(c.UserContext())
	if err != nil {
		return err
	}

	return c.JSON(projects)
}
