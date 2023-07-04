package pasawutil

import (
	"errors"
	"fmt"

	grp "github.com/elvisgastelum/pasawutil/src/generate-random-password"
	"github.com/go-cmd/cmd"
	"github.com/urfave/cli/v2"
)

const (
	passwordLength = 20
)

func GeneratePassword(args cli.Args) {
	entry := args.Get(0)
	possibleUser := args.Get(1)

	if entry == "" {
		fmt.Print("No query provided")
		return
	}

	passwordGenerated, err := createEntry(entry, possibleUser)

	if err != nil {
		fmt.Print(err.Error())
		return
	}

	fmt.Print(passwordGenerated)
}

func createEntry(entry string, user string) (string, error) {
	if entry == "" {
		return "", errors.New("no entry provided")
	}

	randomPassStr, err := grp.GenerateRandomPassword(passwordLength)

	if err != nil {
		return "", err
	}

	var echoValue string
	if user == "" {
		echoValue = fmt.Sprintf("echo \"%s\"", randomPassStr)
	} else {
		echoValue = fmt.Sprintf("echo \"%s\nlogin: %s\"", randomPassStr, user)
	}

	passValue := fmt.Sprintf("pass insert -m %s", entry)

	c := cmd.NewCmd("bash", "-c", fmt.Sprintf("%s | %s", echoValue, passValue))

	<-c.Start()

	return entry, nil
}
