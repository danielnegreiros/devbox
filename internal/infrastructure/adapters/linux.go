package adapters

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"path"

	"github.com/danielnegreiros/go-proxmox-cli/internal/infrastructure/ports"
	"golang.org/x/crypto/ssh"
)

type LinuxAdapter struct {
	host string
	port int
	user string
	// password string
	conn *ssh.Client
}

func NewLinuxAdapter(host string, port int, user string, password string) *LinuxAdapter {

	// Create SSH configuration with the given credentials
	sshConfig := createSSHConfig(user, password)

	// Connect to the remote host with the given host and port
	conn, err := connectToHost(host, port, sshConfig)
	if err != nil {
		log.Fatalf("unable to connect to remote host: %v", err)
	}

	return &LinuxAdapter{
		host: host,
		port: port,
		user: user,
		// password: password,
		conn: conn,
	}
}

func (a *LinuxAdapter) Close() {
	a.conn.Close()
}

func (a *LinuxAdapter) ExecuteCommand(command string) ports.CommandOutputDTO {
	log.Printf("Starting command: %s", command)
	// Run the command on the remote host
	// Create a new session on the connection
	session, err := createSession(a.conn)
	if err != nil {
		log.Fatalf("unable to create session: %v", err)
	}
	defer session.Close()

	stdout, stderr, err := runCommand(session, command)
	log.Println("Done")
	return ports.NewCommandOutputDTO(stdout, stderr, err == nil, err)

}

// Creates an SSH configuration with the given user and password
func createSSHConfig(user string, password string) *ssh.ClientConfig {
	key, err := parsePrivateKey()
	if err != nil {
		fmt.Printf("Failed to load private key: %v\n", err)
		os.Exit(1)
	}

	return &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

// Connects to the remote host with the given host, port, and SSH configuration
func connectToHost(host string, port int, sshConfig *ssh.ClientConfig) (*ssh.Client, error) {
	return ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), sshConfig)
}

// Creates a new session on the given SSH connection
func createSession(conn *ssh.Client) (*ssh.Session, error) {
	return conn.NewSession()
}

// Runs the given command on the given session and returns the stdout, stderr, and error
func runCommand(session *ssh.Session, command string) (string, string, error) {
	stdout, err := session.StdoutPipe()
	if err != nil {
		return "", "", fmt.Errorf("unable to setup stdout for session: %v", err)
	}

	stderr, err := session.StderrPipe()
	if err != nil {
		return "", "", fmt.Errorf("unable to setup stderr for session: %v", err)
	}

	err = session.Start(command)
	if err != nil {
		return "", "", fmt.Errorf("unable to run command on remote host: %v", err)
	}

	stdoutBytes, err := io.ReadAll(stdout)
	if err != nil {
		return "", "", fmt.Errorf("unable to read stdout from remote host: %v", err)
	}

	stderrBytes, err := io.ReadAll(stderr)
	if err != nil {
		return "", "", fmt.Errorf("unable to read stderr from remote host: %v", err)
	}

	err = session.Wait()
	if err != nil {
		return "", "", fmt.Errorf("error: %v", string(stderrBytes))
	}

	return string(stdoutBytes), string(stderrBytes), nil
}

// Parse a private key from a file
func parsePrivateKey() (ssh.Signer, error) {
	user, err := user.Current()
	if err != nil {
		return nil, err
	}

	path := path.Join(user.HomeDir, ".ssh/id_rsa")

	keyBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	key, err := ssh.ParsePrivateKey(keyBytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}
