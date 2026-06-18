package util

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
)

func ClickToExit() {
	fmt.Println("Press any key for exit")
	keyboard.Open()
	defer keyboard.Close()

	keyboard.GetKey()
	os.Exit(1)
}
