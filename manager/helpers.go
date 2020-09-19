package manager

import (
	"github.com/v8platform/errors"
	"github.com/v8platform/runner"
	v8 "github.com/v8platform/v8"
)

func Run(where runner.Infobase, what runner.Command, opts ...interface{}) error {

	err := v8.Run(where, what, opts...,
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

func syncInfobase(connString, user, password string) v8.Infobase {

	if len(connString) == 0 {
		return v8.NewTempIB()
	}
	// TODO Сделать получение базы для выполнения синхронизации
	return v8.NewTempIB()

}
