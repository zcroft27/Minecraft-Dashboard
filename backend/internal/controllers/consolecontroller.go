package controllers

import (
	"mcdashboard/internal/service/ssh"

	"github.com/gofiber/fiber/v2"
)

type ConsoleController struct {
	SSHClient *ssh.SSHClient
}

func NewConsoleController() *ConsoleController {
	return &ConsoleController{
		ssh.NewSSHClient(),
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
