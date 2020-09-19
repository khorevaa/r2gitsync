package cmd

import (
	"encoding/xml"
	"fmt"
	cli "github.com/jawher/mow.cli"
	"github.com/khorevaa/r2gitsync/cmd/flags"
	"github.com/khorevaa/r2gitsync/manager"
	"io/ioutil"
	"os"
	"path"
)

// Sample use: vault list OR vault config list
func (app *Application) cmdSetVersion(cmd *cli.Cmd) {

	var (
		doCommit            bool
		setVersion, workdir string
	)

	flags.BoolOpt("c commit", false, "закоммитить изменения в git").
		Ptr(&doCommit).
		Apply(cmd, app.ctx)
	flags.StringArg("VERSION", "", "Номер версии для записи в файл.").
		Ptr(&setVersion).
		Apply(cmd, app.ctx)
	WorkdirArg.Ptr(&workdir).Apply(cmd, app.ctx)

	cmd.Spec = "[OPTIONS] VERSION [WORKDIR]"

	cmd.Action = func() {

		err := manager.WriteVersionFile(workdir, setVersion)

		failOnErr(err)

		if doCommit {
			err = manager.CommitVersionFile(workdir)
			failOnErr(err)
		}

		readVersion(workdir)

	}
}

type versionReader struct {
	CurrentVersion int64 `xml:"VERSION"`
}

func readVersion(workdir string) {

	fileVesrion := path.Join(workdir, manager.VERSION_FILE)

	// Open our xmlFile
	xmlFile, err := os.Open(fileVesrion)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	var r = versionReader{}

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)
	fmt.Println(string(byteValue))

	// xmlFiles content into 'users' which we defined above
	err = xml.Unmarshal(byteValue, &r.CurrentVersion)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fmt.Sprintf("Write version: <%d>", r.CurrentVersion))

}
