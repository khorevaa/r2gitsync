package flow

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
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

type repositoryAuthor struct {
	name  string
	email string
}

func (a repositoryAuthor) Name() string {
	return a.name
}

func (a repositoryAuthor) Email() string {
	return a.email
}

func (a repositoryAuthor) Desc() string {

	return fmt.Sprintf("%s <%s> ", a.name, a.email)
}

func NewAuthor(name, email string) RepositoryAuthor {

	return repositoryAuthor{
		name:  name,
		email: email,
	}

}

func parseRepositoryReport(file string) (versions []RepositoryVersion, err error) {

	err, bytes := readFile(file, nil)
	if err != nil {
		return
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

	return
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

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}
func IsNoExist(name string) (bool, error) {

	ok, err := Exists(name)
	return !ok, err
}

func clearDir(dir string, skipFiles ...string) error {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, f := range files {

		_, file := filepath.Split(f.Name())

		if isSkipFile(file, skipFiles) {
			continue
		}

		os.Remove(f.Name())
		//fmt.Println(f.Name())
	}

	return nil
}

func isSkipFile(file string, skipFiles []string) bool {

	for _, skipFile := range skipFiles {

		if strings.Contains(file, skipFile) {
			return true
		}

	}

	return false
}

func CopyFile(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()

	_, err = io.Copy(out, in)
	if err != nil {
		return
	}

	err = out.Sync()
	if err != nil {
		return
	}

	si, err := os.Stat(src)
	if err != nil {
		return
	}
	err = os.Chmod(dst, si.Mode())
	if err != nil {
		return
	}

	return
}

// CopyDir recursively copies a directory tree, attempting to preserve permissions.
// Source directory must exist,
// Symlinks are ignored and skipped.
func CopyDir(src string, dst string) (err error) {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	si, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !si.IsDir() {
		return fmt.Errorf("source is not a directory")
	}

	//_, err = os.Stat(dst)
	//if err != nil && !os.IsNotExist(err) {
	//	return
	//}

	if ok, _ := Exists(dst); !ok {

		err = os.MkdirAll(dst, si.Mode())
		if err != nil {
			return
		}
		//return fmt.Errorf("destination already exists")
	}

	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = CopyDir(srcPath, dstPath)
			if err != nil {
				return
			}
		} else {
			// Skip symlinks.
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}

			err = CopyFile(srcPath, dstPath)
			if err != nil {
				return
			}
		}
	}

	return
}

// Decode decodes a byte slice into a signature
func decodeAuthor(b []byte) (string, string) {
	open := bytes.LastIndexByte(b, '<')
	closeSym := bytes.LastIndexByte(b, '>')
	if open == -1 || closeSym == -1 {
		return "", ""
	}

	if closeSym < open {
		return "", ""
	}

	Name := string(bytes.Trim(b[:open], " "))
	Email := string(b[open+1 : closeSym])

	return Name, Email

}
