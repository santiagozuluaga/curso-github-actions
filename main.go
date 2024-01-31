package main

import (
	"fmt"
	"os"
)

func main() {
	githubActor := os.Getenv("GITHUB_ACTOR")

	if githubActor == "" {
		githubActor = "unknown"
	}

	fmt.Printf("Hello %s!", githubActor)
}
