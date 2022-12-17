package subcommand

import (
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"os"
	"strings"
	"tut/repository"
)

type imageNames []string

func (this *imageNames) String() string {
	return fmt.Sprintf("[%s]", strings.Join(*this, ","))
}
func (this *imageNames) Set(arg string) error {
	*this = append(*this, arg)
	return nil
}

func Commit(args []string) {
	flags := flag.NewFlagSet("commit", flag.ExitOnError)
	message := flags.String("m", "", "コミットに対する解説")
	pictures := imageNames{}
	flags.Var(&pictures, "p", "解説に添付する画像")
	flags.Parse(args)
	if flags.NArg() > 0 {
		panic(fmt.Errorf("引数が多すぎます。"))
	}
	var data repository.Data
	readJson(".tut/data.json", &data)
	if len(data.Tags) == 0 {
		panic(fmt.Errorf("セクションが１つも定義されていません。tut start-sectionコマンドで最初のセクションを作成してください。"))
	}
	var files []repository.File
	readJson(".tut/stage.json", &files)
	for _, picture := range pictures {
		content := make([]byte, 1024)
		file, error := os.Open(picture)
		defer file.Close()
		if error != nil {
			panic(error)
		}
		if count, error := file.Read(content); error != nil {
			panic(error)
		} else {
			files = append(files, repository.File{Name: picture, Type: "image", Content: base64.StdEncoding.EncodeToString(content[:count])})
		}
	}
	if id, error := uuid.NewRandom(); error != nil {
		panic(error)
	} else {
		data.Tags[len(data.Tags)-1].Commits = append(data.Tags[len(data.Tags)-1].Commits, repository.Commit{Id: id.String(), Message: *message, Files: files})
	}
	writeJson(".tut/data.json", data)
	writeJson(".tut/stage.json", []repository.File{})
}
