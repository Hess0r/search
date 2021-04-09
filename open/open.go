package open

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func Open(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	default:
		fmt.Println("unsupported platform", runtime.GOOS)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
