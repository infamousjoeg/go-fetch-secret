package main

import (
	"fmt"
	"github.com/cyberark/conjur-api-go/conjurapi"
	"github.com/cyberark/conjur-api-go/conjurapi/authn"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		panic("Missing required secret variable as an argument. e.g. fetchsecret db/password")
	} else if len(os.Args) > 2 {
		panic("Too many arguments detected.  Please only provide one secret variable at a time.")
	}
	variableIdentifier := os.Args[1]

	config := conjurapi.LoadConfig()

	conjur, err := conjurapi.NewClientFromKey(config,
		authn.LoginPair{
			Login:  os.Getenv("CONJUR_AUTHN_LOGIN"),
			APIKey: os.Getenv("CONJUR_AUTHN_API_KEY"),
		},
	)
	if err != nil {
		panic(err)
	}

	// Retrieve a secret into []byte
	secretValue, err := conjur.RetrieveSecret(variableIdentifier)
	if err != nil {
		panic(err)
	}
	fmt.Println("The secret value is: ", string(secretValue))
}
