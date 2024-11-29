package controllers

import (
	"github.com/gofiber/fiber/v2"
	"mcdashboard/internal/services"
)

type VMController struct {
	SSHClient *services.SSHClient
}

func NewVMController() *VMController {
	return &VMController{
		services.NewSSHClient(),
	}
}

func (vmc *VMController) ToggleServer(c *fiber.Ctx, startServer bool) error {
	var cmd string

	if startServer {
		cmd = "cd minecraft-server ; ./start-server.sh"
	} else {
		cmd = "cd minecraft-server ; ./stop-server.sh"
	}

	output, err := vmc.SSHClient.ConnectAndExecute(cmd)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": string(output),
	})
}
