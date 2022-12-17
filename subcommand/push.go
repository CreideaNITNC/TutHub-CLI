package subcommand

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"tut/repository"
)

func Push(args []string) {
	flags := flag.NewFlagSet("push", flag.ExitOnError)
	flags.Parse(args)
	if flags.NArg() != 1 {
		panic(fmt.Errorf("引数の数が合いません。"))
	}
	var config repository.Config
	readJson(".tut/config.json", &config)
	uri := ""
	for _, remoteRepository := range config.Remote {
		if remoteRepository.Name == args[0] {
			uri = remoteRepository.Uri
		}
	}
	if uri == "" {
		panic(fmt.Errorf("指定された識別名のリモートリポジトリは登録されていません。"))
	}
	file, error := os.Open(".tut/data.json")
	if error != nil {
		panic(error)
	}
	request, _ := http.NewRequest("POST", uri, file)
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	if _, error := client.Do(request); error != nil {
		panic(error)
	}
}
