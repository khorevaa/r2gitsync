package main

import (
	"fmt"
	cli "github.com/jawher/mow.cli"
	"github.com/khorevaa/go-v8platform/errors"
	"github.com/khorevaa/r2gitsync/internal/args"
	"github.com/khorevaa/r2gitsync/internal/env"
	"github.com/v8platform/runner"
	v8 "github.com/v8platform/v8"
	"golang.org/x/text/encoding"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func WorkdirArg(cmd *cli.Cmd) *string {
	return args.StringArg(cmd, "WORKDIR", pwd, "Каталог исходников внутри локальной копии git-репозитория.").
		HideValue(true).
		Env(env.WorkDir).
		Arg()
}

func WorkdirArgPtr(into *string, cmd *cli.Cmd) {
	args.StringArg(cmd, "WORKDIR", pwd, "Каталог исходников внутри локальной копии git-репозитория.").
		HideValue(true).
		Env(env.WorkDir).
		Ptr(into)
}

func failOnErr(err error) {
	if err != nil {
		fmt.Printf("Ошибка выполненния программы: %v \n", err.Error())
		cli.Exit(1)
	}
}

func Run(where runner.Infobase, what runner.Command, opts *SyncOptions) error {

	err := v8.Run(where, what, v8.WithPath(opts.v8Path),
		v8.WithVersion(opts.v8version),
		//	v8.WithTempDir(opts.tempDir), // TODO Сделать для запуска временный катиалог
	)

	errorContext := errors.GetErrorContext(err)

	out, ok := errorContext["message"]
	if ok {
		err = errors.Internal.Wrap(err, out)
	}

	//TODO Сделать несколько попыток при отсутсвии лицензиии

	return err

}

func parseRepositoryReport(file string) (versions []repositoryVersion, err error) {

	err, bytes := ReadFile(file, nil)
	if err == nil {
		return nil, err
	}

	report := string(*bytes)
	// Двойные кавычки в комментарии мешают, по этому мы заменяем из на одинарные
	report = strings.Replace(report, "\"\"", "'", -1)

	tmpArray := [][]string{}
	reg := regexp.MustCompile(`[{]"#","([^"]+)["][\}]`)
	matches := reg.FindAllStringSubmatch(report, -1)
	for _, s := range matches {
		if s[1] == "Версия:" {
			tmpArray = append(tmpArray, []string{})
		}

		if len(tmpArray) > 0 {
			tmpArray[len(tmpArray)-1] = append(tmpArray[len(tmpArray)-1], s[1])
		}
	}

	for _, array := range tmpArray {
		versionInfo := repositoryVersion{}
		for id, s := range array {
			switch s {
			case "Версия:":
				if ver, err := strconv.Atoi(array[id+1]); err == nil {
					versionInfo.Number = int64(ver)
				}
			case "Версия конфигурации:":
				versionInfo.Version = array[id+1]
			case "Пользователь:":
				versionInfo.Author = array[id+1]
			case "Комментарий:":
				// Комментария может не быть, по этому вот такой костыльчик
				if array[id+1] != "Изменены:" {
					versionInfo.Comment = strings.Replace(array[id+1], "\n", " ", -1)
					versionInfo.Comment = strings.Replace(array[id+1], "\r", "", -1)
				}
			case "Дата создания:":
				if t, err := time.Parse("02.01.2006", array[id+1]); err == nil {
					versionInfo.Date = t
				}
			case "Время создания:":
				if !versionInfo.Date.IsZero() {
					str := versionInfo.Date.Format("02.01.2006") + " " + array[id+1]
					if t, err := time.Parse("02.01.2006 15:04:05", str); err == nil {
						versionInfo.Date = t
					}
				}
			}
		}
		versions = append(versions, versionInfo)
	}

	return []repositoryVersion{}, nil
}

func ReadFile(filePath string, Decoder *encoding.Decoder) (error, *[]byte) {
	//dec := charmap.Windows1251.NewDecoder()

	if fileB, err := ioutil.ReadFile(filePath); err == nil {
		// Разные кодировки = разные длины символов.
		if Decoder != nil {
			newBuf := make([]byte, len(fileB)*2)
			Decoder.Transform(newBuf, fileB, false)

			return nil, &newBuf
		} else {
			return nil, &fileB
		}
	} else {
		return fmt.Errorf("Ошибка открытия файла %q:\n %v", filePath, err), nil
	}
}
