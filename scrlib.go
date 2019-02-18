package scrlib

import (
	"os"
	"os/exec"
	"runtime"
)

// Clear コンソールをクリアする
func Clear() error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
