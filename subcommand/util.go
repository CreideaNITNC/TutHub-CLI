package subcommand

import (
	"encoding/json"
	"os"
)

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
