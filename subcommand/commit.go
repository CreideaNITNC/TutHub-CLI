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
	imageNames := imageNames{}
	flags.Var(&imageNames, "p", "解説に添付する画像")
	flags.Parse(args)
	if flags.NArg() > 0 {
		panic(fmt.Errorf("引数が多すぎます。"))
	}
	var data repository.Data
	readJson(".tut/data.json", &data)
	if len(data.Tags) == 0 {
		panic(fmt.Errorf("セクションが１つも定義されていません。tut start-sectionコマンドで最初のセクションを作成してください。"))
	}
	var sourceCodes []repository.SourceCode
	readJson(".tut/stage.json", &sourceCodes)
	var pictures []repository.Picture
	for _, imageName := range imageNames {
		content := make([]byte, 1_000_000_000)
		file, err := os.Open(imageName)
		//defer file.Close()
		if err != nil {
			file.Close()
			panic(err)
		}
		if count, error := file.Read(content); error != nil {
			file.Close()
			panic(error)
		} else {
			pictures = append(pictures, repository.Picture{Name: imageName, Binary: base64.StdEncoding.EncodeToString(content[:count])})
		}
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}
	if id, error := uuid.NewRandom(); error != nil {
		panic(error)
	} else {
		data.Tags[len(data.Tags)-1].Commits = append(data.Tags[len(data.Tags)-1].Commits, repository.Commit{Id: id.String(), Message: *message, Codes: sourceCodes, Pictures: pictures})
	}
	writeJson(".tut/data.json", data)
	writeJson(".tut/stage.json", []repository.SourceCode{})
}
