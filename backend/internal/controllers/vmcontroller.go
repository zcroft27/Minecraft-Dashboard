package controllers

import (
	"mcdashboard/internal/service/ssh"

	"github.com/gofiber/fiber/v2"
)

type VMController struct {
	SSHClient *ssh.SSHClient
}

func NewVMController() *VMController {
	return &VMController{
		ssh.NewSSHClient(),
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
