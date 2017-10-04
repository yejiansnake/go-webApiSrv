package utility

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Process struct {
}

var CurrentProcessInstance *Process = new(Process)

func (ptr *Process) GetCurrentFileName() (string, error) {
	fullPath, err := ptr.GetCurrentFullPath()

	if err != nil {
		return "", err
	}

	index := strings.LastIndex(fullPath, "/")

	return string(fullPath[index+1:]), nil
}

func (ptr *Process) GetCurrentDir() (string, error) {
	fullPath, err := ptr.GetCurrentFullPath()

	if err != nil {
		return "", err
	}

	index := strings.LastIndex(fullPath, "/")

	return string(fullPath[:index]), nil
}

func (ptr *Process) GetCurrentFullPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])

	if err != nil {
		return "", err
	}

	path, err := filepath.Abs(file)

	if err != nil {
		return "", err
	}

	i := strings.LastIndex(path, "/")

	if i < 0 {
		i = strings.LastIndex(path, "\\")
		if i >= 0 {
			path = strings.Replace(path, "\\", "/", -1)
		}
	}

	if i < 0 {
		return "", errors.New("get path failed")
	}
	return string(path[0:]), nil
}
