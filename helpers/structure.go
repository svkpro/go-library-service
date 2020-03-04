package helpers

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

func FromJson(jsonFile string, s interface{}) {
	file, err := filepath.Abs(jsonFile)
	die(err)

	jsonContent, err := os.Open(file)
	die(err)
	defer func() {
		if err := jsonContent.Close(); err != nil {
			panic(err)
		}
	}()

	structData, err := ioutil.ReadAll(jsonContent)
	die(err)

	err = json.Unmarshal(structData, &s)
}

func die(err error) {
	if err != nil {
		panic(err)
	}
}
