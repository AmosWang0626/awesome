package exec_utils

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func ExecCmd(path string, cmd string) (result string) {
	if len(path) > 0 {
		cmd = fmt.Sprintf("cd %s | %s", path, cmd)
		fmt.Printf("RUN CMD [%s]\n", cmd)
	}

	var buf []byte
	if runtime.GOOS == "windows" {
		buf = ExecWindows(cmd)
	} else {
		buf = ExecLinux(cmd)
	}
	result = strings.TrimSpace(string(buf))
	fmt.Println("", result)

	return
}

func ExecLinux(cmd string) (result []byte) {
	result, err := exec.Command("/bin/bash", "-c", cmd).Output()
	if err != nil {
		println("ERROR ::", err.Error())
	}
	return
}

func ExecWindows(cmd string) (result []byte) {
	result, err := exec.Command("cmd", "/c", cmd).Output()
	if err != nil {
		println("ERROR ::", err.Error())
	}
	return
}
