package subcommand

import (
	"flag"
	"fmt"
	"github.com/google/uuid"
	"tut/repository"
)

func StartSection(args []string) {
	flags := flag.NewFlagSet("start-section", flag.ExitOnError)
	flags.Parse(args)
	if flags.NArg() != 1 {
		panic(fmt.Errorf("引数の数が合いません。"))
	}
	var data repository.Data
	readJson(".tut/data.json", &data)
	if id, error := uuid.NewRandom(); error != nil {
		panic(error)
	} else {
		data.Sections = append(data.Sections, repository.Section{Id: id.String(), Name: flags.Arg(0), Commits: []repository.Commit{}})
	}
	writeJson(".tut/data.json", data)
}
