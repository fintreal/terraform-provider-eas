// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eas

import (
    "fmt"
    "os/exec"
    "bytes"
    "os"
	"math/rand"
	"time"
)

func RunCommand(dir string, command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)

	var out bytes.Buffer
	cmd.Dir = dir
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf(out.String())
	}
	return out.String(), nil
}

func createDir(name string) {
    os.MkdirAll(name, os.ModePerm)
}

func createFile(dir string, name string, content string) {
    file, _:= os.Create(dir + "/" + name)
    file.WriteString(content)
    file.Close()
}

const letters = "abcdefghijklmnopqrstuvwxyz0123456789"

func RandomString() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
