package main

import (
	"fgc/cmd"
)

const version = "0.0.0"

func main() {
	//cfg := config.New("./config/sdk.yaml")
	//_ = cfg
	c := cmd.New()
	c.Version(version)
	c.Execute()
}