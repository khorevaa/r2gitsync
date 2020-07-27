package main

import (
	"encoding/xml"
	"fmt"
	cli "github.com/jawher/mow.cli"
	"github.com/khorevaa/r2gitsync/internal/args"
	"github.com/khorevaa/r2gitsync/internal/opts"
	"io/ioutil"
	"os"
	"path"
)

// Sample use: vault list OR vault config list
func cmdSetVersion(cmd *cli.Cmd) {

	var (
		doCommit   = opts.BoolOpt("c commit", false, "закоммитить изменения в git").Opt(cmd)
		setVersion = args.StringArg("VERSION", "", "Номер версии для записи в файл.").Arg(cmd)
		workdir    = WorkdirArg(cmd)
	)

	cmd.Spec = "[OPTIONS] VERSION [WORKDIR]"

	cmd.Action = func() {

		err := writeVersionFile(*workdir, *setVersion)

		failOnErr(err)

		if *doCommit {
			err = commitVersionFile(*workdir)
			failOnErr(err)
		}

	}
}

type versionReader struct {
	currentVersion int64 `xml:"VERSION"`
}

func readVersion(workdir string) {

	fileVesrion := path.Join(workdir, VERSION_FILE)

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

	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, r)

	fmt.Println(r.currentVersion)

}
