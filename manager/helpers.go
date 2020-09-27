package manager

import (
	v8 "github.com/v8platform/v8"
)

type syncInfobase struct {
	connString string
}

func (ib syncInfobase) Path() string {

	return ""
}

func (ib syncInfobase) ConnectionString() string {

	return ib.connString
}

func (ib syncInfobase) Values() (v []string) {
	return
}

func getSyncInfobase(connString string) v8.Infobase {

	if len(connString) == 0 {
		return v8.NewTempIB()
	}
	return syncInfobase{
		connString: connString,
	}

}
