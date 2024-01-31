package main

import (
	"fmt"
	"os"
)

func getEnvString(key, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}

func main() {
	githubActor := getEnvString("GITHUB_ACTOR", "santiagozuluaga")

	fmt.Printf("Hello %s!", githubActor)
}
