package v1

import (
	"github.com/gofiber/fiber/v2"
	uuid2 "github.com/google/uuid"
	"github.com/khorevaa/r2gitsync/internal/dto"
)

func (ctr *Controller) GetProjects(c *fiber.Ctx) error {

	projects, err := ctr.bl.ProjectsLogic.GetProjects(c.UserContext())
	if err != nil {
		return err
	}

	return c.JSON(projects)
}

func (ctr *Controller) PostProjects(c *fiber.Ctx) error {

	var project dto.Project

	if err := c.BodyParser(&project); err != nil {
		return err
	}

	projects, err := ctr.bl.ProjectsLogic.CreateProject(c.UserContext(), &project)
	if err != nil {
		return err
	}

	return c.JSON(projects)
}

func (ctr *Controller) PostProject(c *fiber.Ctx) error {

	var project dto.Project

	if err := c.BodyParser(&project); err != nil {
		return err
	}

	uuid, err := uuid2.Parse(c.Params("uuid"))
	if err != nil {
		return err
	}

	projects, err := ctr.bl.ProjectsLogic.UpdateProject(c.UserContext(), uuid, &project)
	if err != nil {
		return err
	}

	return c.JSON(projects)
}

func (ctr *Controller) GEtProjectMaster(ctx *fiber.Ctx) error {
	return nil
}

func (ctr *Controller) PostProjectMaster(ctx *fiber.Ctx) error {
	return nil
}

func (ctr *Controller) PutProjectMaster(ctx *fiber.Ctx) error {
	return nil
}

func (ctr *Controller) DeleteProjectMaster(ctx *fiber.Ctx) error {
	return nil
}
