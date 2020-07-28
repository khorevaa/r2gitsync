package main

import cli "github.com/jawher/mow.cli"

// Sample use: vault creds reddit.com
func (app *Application) CmdInit2(cmd *cli.Cmd) {

	cmd.LongDesc = `Данный режим работает по HTTP (REST API) с базой данных.
		Возможности:
		* самостоятельно получает список информационных баз к обновления;
		* поддержание нескольких потоков обновления
		* переодический/разовый опрос необходимости обновления
		* отправка журнала обновления на url.`

	cmd.Action = func() {
		//fmt.Printf("display account info for %s\n", *account)
	}
}
