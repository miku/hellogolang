package main

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	client, err := client.NewEnvClient()
	if err != nil {
		log.Fatal(err)
	}

	images, err := client.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, image := range images {
		fmt.Println(image.ID)
	}
}
