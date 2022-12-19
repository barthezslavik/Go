package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/moby/moby/client"
	"github.com/moby/moby/pkg/term"
)

func main() {
	// Create a Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println("Error creating Docker client:", err)
		return
	}

	// Pull an image from a registry
	out, err := cli.ImagePull(context.Background(), "alpine", client.WithRegistryAuth(""))
	if err != nil {
		fmt.Println("Error pulling image:", err)
		return
	}
	defer out.Close()
	io.Copy(os.Stdout, out)

	// Run a container from the image
	resp, err := cli.ContainerCreate(context.Background(), &client.Config{
		Image: "alpine",
		Cmd:   []string{"echo", "Hello, World!"},
	}, nil, nil, "")
	if err != nil {
		fmt.Println("Error creating container:", err)
		return
	}

	// Start the container
	if err := cli.ContainerStart(context.Background(), resp.ID, client.WithStart); err != nil {
		fmt.Println("Error starting container:", err)
		return
	}

	// Attach to the container and print the output
	out, err = cli.ContainerAttach(context.Background(), resp.ID, client.ContainerAttachOptions{
		Stream: true,
		Stdout: true,
		Stderr: true,
	})
	if err != nil {
		fmt.Println("Error attaching to container:", err)
		return
	}
	defer out.Close()
	if err := term.TTY{}.SetRawTerminal(os.Stdin.Fd()); err != nil {
		fmt.Println("Error setting terminal mode:", err)
		return
	}
	io.Copy(os.Stdout, out.Reader)
}