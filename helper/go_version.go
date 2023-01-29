package helper

import (
	"os/exec"
	"runtime"
	"strings"
)

func GoVersion() string {

	if runtime.GOOS == "windows" {
		output, err := exec.Command("cmd", "/C", "go version").Output()
		Check(err)
		return splitGoVersion(string(output))
	} else {
		output, err := exec.Command("bash", "-c", "go version").Output()
		Check(err)
		return splitGoVersion(string(output))
	}
}
func splitGoVersion(version string) string {
	split := strings.Split(version, " ")
	removeGo := strings.Trim(split[2], "go")
	return removeGo[:len(removeGo)-2]

}
