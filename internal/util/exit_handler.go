package util

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func ClickToExit() {
	fmt.Println("Press any key for exit")
	keyboard.Open()
	defer keyboard.Close()

	keyboard.GetKey()
}
