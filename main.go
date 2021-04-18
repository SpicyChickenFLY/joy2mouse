package main

import "github.com/SpicyChickenFLY/xinput2mouse/manager"

func main() {
	manager := manager.NewManager()
	err := manager.Loop()
	panic(err)
}
