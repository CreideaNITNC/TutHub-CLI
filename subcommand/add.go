package subcommand

import (
	"flag"
	"fmt"
	"os"
	"tut/repository"
)

const MaxSourceFileByteSize = 100_000 /* 100[KB] */

func Add(args []string) {
	flags := flag.NewFlagSet("add", flag.ExitOnError)
	flags.Parse(args)
	if flags.NArg() == 0 {
		panic(fmt.Errorf("引数は最低でも１つ必要です。"))
	}
	var stage []repository.SourceCode
	readJson(".tut/stage.json", &stage)
outer:
	for _, arg := range flags.Args() {
		sourceFile, err := os.Open(arg)
		if err != nil {
			sourceFile.Close()
			panic(err)
		}
		content := make([]byte, MaxSourceFileByteSize)
		var newSourceCode repository.SourceCode
		if count, err := sourceFile.Read(content); err != nil {
			panic(err)
		} else {
			newSourceCode = repository.SourceCode{Name: arg, Content: string(content[:count])}
		}
		for i, _ := range stage {
			if arg == stage[i].Name {
				stage[i] = newSourceCode
				err := sourceFile.Close()
				if err != nil {
					panic(err)
				}
				continue outer
			}
		}
		stage = append(stage, newSourceCode)
		err = sourceFile.Close()
		if err != nil {
			panic(err)
		}
	}
	writeJson(".tut/stage.json", stage)
}
