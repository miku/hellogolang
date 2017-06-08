package main

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// StatusMessage from Image Pull.
type StatusMessage struct {
	Status string `json:"status"`
	ID     string `json:"id"`
}

func main() {
	ctx := context.Background()
	client, err := client.NewEnvClient()
	if err != nil {
		log.Fatal(err)
	}

	rc, err := client.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	if err != nil {
		log.Fatal(err)
	}

	br := bufio.NewReader(rc)
	for {
		bb, err := br.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		var msg StatusMessage
		if err := json.Unmarshal(bb, &msg); err != nil {
			log.Fatal(err)
		}
		log.Println(msg.Status)
	}

	resp, err := client.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"echo", "hello world"},
	}, nil, nil, "")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("created: %s", resp.ID)

	if err := client.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		log.Fatal(err)
	}
	if _, err = client.ContainerWait(ctx, resp.ID); err != nil {
		log.Fatal(err)
	}
	out, err := client.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(os.Stdout, out); err != nil {
		log.Fatal(err)
	}

	log.Printf("removing: %s", resp.ID)

	if err := client.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{Force: true}); err != nil {
		log.Fatal(err)
	}
}
