package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"watchtower/Server/Model/program"
	"watchtower/utils"
)

func ReadeConfig(name string) []program.Wt_Program {

	r := regexp.MustCompile(`[^.]+\.json$`)

	if !r.MatchString(name) {
		fmt.Println("[-]", "this not a json file !!")
		return nil
	}
	config, err := os.Open(name)

	e := utils.FailOnError(err, "", nil)
	if e {

		return nil
	}

	defer config.Close()
	data, err := io.ReadAll(config)

	e = utils.FailOnError(err, "", nil)
	if e {

		return nil
	}

	var userConfig []program.Wt_Program

	err = json.Unmarshal(data, &userConfig)

	e = utils.FailOnError(err, "", nil)
	if e {

		return nil
	}

	return userConfig
}
