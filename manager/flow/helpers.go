package flow

import (
	"fmt"
	"golang.org/x/text/encoding"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type repositoryVersion struct {
	version string
	author  string
	date    time.Time
	comment string
	number  int64
}

func (r repositoryVersion) Version() string {
	return r.version
}

func (r repositoryVersion) Author() string {
	return r.author
}

func (r repositoryVersion) Date() time.Time {
	return r.date
}

func (r repositoryVersion) Comment() string {
	return r.comment
}

func (r repositoryVersion) Number() int64 {
	return r.number
}

func parseRepositoryReport(file string) (versions []RepositoryVersion, err error) {

	err, bytes := readFile(file, nil)
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
					versionInfo.number = int64(ver)
				}
			case "Версия конфигурации:":
				versionInfo.version = array[id+1]
			case "Пользователь:":
				versionInfo.author = array[id+1]
			case "Комментарий:":
				// Комментария может не быть, по этому вот такой костыльчик
				if array[id+1] != "Изменены:" {
					versionInfo.comment = strings.Replace(array[id+1], "\n", " ", -1)
					versionInfo.comment = strings.Replace(array[id+1], "\r", "", -1)
				}
			case "Дата создания:":
				if t, err := time.Parse("02.01.2006", array[id+1]); err == nil {
					versionInfo.date = t
				}
			case "Время создания:":
				if !versionInfo.date.IsZero() {
					str := versionInfo.date.Format("02.01.2006") + " " + array[id+1]
					if t, err := time.Parse("02.01.2006 15:04:05", str); err == nil {
						versionInfo.date = t
					}
				}
			}
		}
		versions = append(versions, versionInfo)
	}

	return []RepositoryVersion{}, nil
}

func readFile(filePath string, Decoder *encoding.Decoder) (error, *[]byte) {
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
