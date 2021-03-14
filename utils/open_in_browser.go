package utils

import (
	"fmt"
	"os/exec"
	"runtime"
)

func OpenUrlInBrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url, "-a", "Google Chrome.app").Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	return err
}
