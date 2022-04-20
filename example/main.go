package main

import (
	"fmt"
	"os"

	blizzapi "github.com/jgrancell/go-battlenet"
)

func main() {
	client := blizzapi.Client{
		Config: blizzapi.Config{
			ClientID:     os.Getenv("BNET_CLIENT_ID"),
			ClientSecret: os.Getenv("BNET_CLIENT_SECRET"),
			Region:       "us",
		},
	}

	// To pull a single character
	_, err := client.Wow().Character("sargeras", "foobarbaz").Equipment()
	if err != nil {
		fmt.Println("error fetching character profile:", err.Error())
	}
}
