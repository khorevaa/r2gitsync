package manager

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/khorevaa/r2gitsync/manager/types"
	"path/filepath"
	"time"
)

const (
	VERSION_FILE = "VERSION"
	AUTHORS_FILE = "AUTHORS"
)

func CommitVersionFile(workdir string) error {

	r, err := git.PlainOpen(workdir)

	if err != nil {
		return err
	}

	filename := filepath.Join(workdir, VERSION_FILE)

	w, err := r.Worktree()

	if err != nil {
		return err
	}

	w.Add(filename)
	c, err := w.Commit("Изменена версия в файле VERSION", &git.CommitOptions{})

	if err != nil {
		return err
	}

	_, err = r.CommitObject(c)

	return err
}

func commitFile(filename string, msg string, opts *git.CommitOptions) error {

	dir := filepath.Dir(filename)

	r, err := git.PlainOpen(dir)

	if err != nil {
		return err
	}

	//filename := filepath.Join(workdir, VERSION_FILE)

	w, err := r.Worktree()

	if err != nil {
		return err
	}

	if opts == nil {
		opts = &git.CommitOptions{}
	}

	w.Add(filename)
	c, err := w.Commit(msg, opts)

	if err != nil {
		return err
	}

	_, err = r.CommitObject(c)

	return err
}

func commitFiles(dir string, author types.RepositoryAuthor, when time.Time, comment string) error {

	g, err := git.PlainOpen(dir)

	if err != nil {
		return err
	}

	w, err := g.Worktree()

	if err != nil {
		return err
	}

	err = w.RemoveGlob(".")

	if err != nil {
		return err
	}

	//pattern := "**/!(*.git)*.*"
	err = w.AddGlob(".")

	if err != nil {
		return err
	}

	c, err := w.Commit(comment, &git.CommitOptions{
		All: true,
		Author: &object.Signature{
			Name:  author.Name(),
			Email: author.Email(),
			When:  when,
		},
	})

	if err != nil {
		return err
	}

	_, err = g.CommitObject(c)

	return err

}
