package main

import "go-urlsaver/internal/app"

const configPath = "configs/main"

func main() {
	app.Run(configPath)
}
