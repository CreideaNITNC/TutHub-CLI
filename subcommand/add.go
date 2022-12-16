package subcommand

import (
	"flag"
	"fmt"
	"os"
	"tut/repository"
)

func Add(args []string) {
	flags := flag.NewFlagSet("add", flag.ExitOnError)
	flags.Parse(args)
	if flags.NArg() == 0 {
		panic(fmt.Errorf("引数は最低でも１つ必要です。"))
	}
	var stage []repository.File
	readJson(".tut/stage.json", &stage)
outer:
	for _, arg := range flags.Args() {
		file, error := os.Open(arg)
		defer file.Close()
		if error != nil {
			panic(error)
		}
		content := make([]byte, 1024)
		var newFile repository.File
		if count, error := file.Read(content); error != nil {
			panic(error)
		} else {
			newFile = repository.File{Name: arg, Type: "text", Content: string(content[:count])}
		}
		for i, _ := range stage {
			if arg == stage[i].Name {
				stage[i] = newFile
				continue outer
			}
		}
		stage = append(stage, newFile)
	}
	writeJson(".tut/stage.json", stage)
}
