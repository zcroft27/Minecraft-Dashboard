package services

import (
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
)

type SSHClient struct{}

func NewSSHClient() *SSHClient {
	return &SSHClient{}
}

func (s *SSHClient) ConnectAndExecute(cmd string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return "", err
	}

	hostname := os.Getenv("vmIP")
	port := "22"
	username := "zcroft27"
	privateKeyPath := "./mckey.pem"

	privateKeyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatalf("Failed to read private key: %v", err)
		return "", err
	}

	signer, err := ssh.ParsePrivateKey(privateKeyBytes)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
		return "", err
	}

	clientConfig := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //TODO: Switch to FixedHostKey(VMPublicKey)
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", hostname, port), clientConfig)
	if err != nil {
		log.Fatalf("Failed to connect to remote VM: %v", err)
		return "", err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create SSH session: %v", err)
		return "", err
	}
	defer session.Close()

	// Run the command and capture the output
	output, err := session.Output(cmd)
	if err != nil {
		log.Fatalf("Failed to execute command: %v", err)
		return "", err
	}

	return string(output), nil
}
