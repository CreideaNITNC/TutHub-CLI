package subcommand

import (
	"flag"
	"fmt"
	"tut/repository"
)

func RemoteAdd(args []string) {
	flags := flag.NewFlagSet("remote add", flag.ExitOnError)
	flags.Parse(args)
	if len(flags.Args()) != 2 {
		panic(fmt.Errorf("引数は１つである必要があります。"))
	}
	var config repository.Config
	readJson(".tut/config.json", &config)
	config.Remote = append(config.Remote, repository.RemoteRepository{Name: flags.Arg(0), Uri: flags.Arg(1)})
	writeJson(".tut/config.json", config)
}
