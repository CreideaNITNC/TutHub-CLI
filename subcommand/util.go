package subcommand

import (
	"encoding/json"
	"os"
)

func readJson(fileName string, ret any) {
	file, error := os.Open(fileName)
	defer file.Close()
	if error != nil {
		panic(error)
	}
	if error := json.NewDecoder(file).Decode(ret); error != nil {
		panic(error)
	}
}
func writeJson(fileName string, content any) {
	file, error := os.Create(fileName)
	defer file.Close()
	if error != nil {
		panic(error)
	}
	if error := json.NewEncoder(file).Encode(content); error != nil {
		panic(error)
	}
}
