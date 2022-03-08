package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func init() {
	fmt.Println(os.Args[0])
	execDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	var splitStr = "/"
	if runtime.GOOS == "windows" {
		splitStr = `\`
	}
	pathArr := strings.Split(execDir, splitStr)
	BifrostDirName := pathArr[len(pathArr)-1]
	switch BifrostDirName {
	case "bin", "sbin":
		BifrostDir = filepath.Dir(execDir) + splitStr
		break
	default:
		BifrostDir = execDir + splitStr
		break
	}
}
