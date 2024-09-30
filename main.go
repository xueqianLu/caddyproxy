package main

import "caddyproxy/command/root"

func main() {
	root.NewRootCommand().Execute()
}
