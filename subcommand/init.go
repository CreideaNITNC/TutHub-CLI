package subcommand

import (
	"flag"
	"fmt"
	"os"
	"tut/repository"
)

func Init(args []string) {
	var flags = flag.NewFlagSet("init", flag.ExitOnError)
	flags.Parse(args)
	if len(flags.Args()) > 0 {
		panic(fmt.Errorf("引数が多すぎます。"))
	}
	if repository, error := os.Stat(".tut"); error == nil && repository.IsDir() {
		panic(fmt.Errorf("リポジトリが既に存在しています。"))
	}
	if error := os.Mkdir(".tut", os.ModePerm); error != nil {
		panic(error)
	}
	writeJson(".tut/config.json", repository.Config{Remote: []repository.RemoteRepository{}})
	writeJson(".tut/data.json", repository.Data{Tags: []repository.Tag{}})
	writeJson(".tut/stage.json", []repository.File{})
}
