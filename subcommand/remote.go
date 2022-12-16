package subcommand

import (
	"flag"
	"fmt"
)

func Remote(args []string) {
	var flags = flag.NewFlagSet("remote", flag.ExitOnError)
	flags.Parse(args)
	if flags.Args()[0] == "add" {
		RemoteAdd(args[1:])
	} else {
		panic(fmt.Errorf("未知のサブコマンド「%s」", flags.Args()[0]))
	}
}
