package controllers

import (
	"github.com/gofiber/fiber/v2"
	"mcdashboard/internal/services"
)

type ConsoleController struct {
	SSHClient *services.SSHClient
}

func NewConsoleController() *ConsoleController {
	return &ConsoleController{
		services.NewSSHClient(),
	}
}

func (cc *ConsoleController) GetPlayerList(c *fiber.Ctx) error {
	cmd := "cd minecraft-server ; ./player-list.sh"

	output, err := cc.SSHClient.ConnectAndExecute(cmd)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": string(output),
	})
}
