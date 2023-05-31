package pasawutil

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func GetLoginField(entry string) {
	if entry == "" {
		fmt.Print("No query provided")
		return
	}

	out, err := exec.Command("pass", "show", entry).Output()

	if err != nil {
		fmt.Print(err.Error())
		return
	}

	login := strings.Split(string(out), "\n")

	loginField, err := findLoginField(login, []string{
		"login:",
		"email:",
		"username:",
		"user:",
	})

	if err != nil {
		fmt.Print(err.Error())
		return
	}

	fmt.Print(loginField)
}

func findLoginField(lines []string, posibleFieldNames []string) (string, error) {
	for _, posibleFieldName := range posibleFieldNames {
		for _, line := range lines {
			if strings.HasPrefix(line, posibleFieldName) {
				return strings.Trim(strings.TrimPrefix(line, posibleFieldName), " "), nil
			}
		}
	}
	return "", errors.New("login field not found")
}
