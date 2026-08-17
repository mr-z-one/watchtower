package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"watchtower/Server/Model/program"
)

func ReadeConfig(name string) []program.Wt_Program {

	r := regexp.MustCompile(`[^.]+\.json$`)

	if !r.MatchString(name) {
		fmt.Println("[-]", "this not a json file !!")
		return nil
	}
	config, err := os.Open(name)

	if err != nil {
		fmt.Println("[-]", err)
		return nil
	}
	defer config.Close()
	data, err := io.ReadAll(config)
	if err != nil {
		fmt.Println("[-]", err)
		return nil
	}

	var userConfig []program.Wt_Program

	err = json.Unmarshal(data, &userConfig)

	if err != nil {
		fmt.Println("[-]", err)
		return nil
	}

	return userConfig
}
