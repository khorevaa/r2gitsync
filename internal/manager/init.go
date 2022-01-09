package manager

import _ "embed"

//go:embed tempExtension.cfe
var tempExtension []byte

func (r *SyncRepository) initWorkdir(opts *Options) (err error) {

	// if !opts.ForceInit {
	//	// TODO Проверить папку и ругнуться если там что-то есть
	//	// TODO При принудительной инициализации - очистить целевой каталог от лишего
	// }
	//
	// err = r.init(opts)
	// if err != nil {
	//	return err
	// }
	//
	// r.log.Infow("Start init repository data",
	//	zap.String("name", r.Name),
	//	zap.String("path", r.Repository.Path),
	// )
	//
	// r.log.Infow("Using infobase for init repository data",
	//	zap.String("path", opts.infobase.ConnectionString()))
	//
	// err = r.GetRepositoryAuthors()
	//
	// if err != nil {
	//	return err
	// }
	//
	// err = r.GetRepositoryHistory()
	// if err != nil {
	//	return err
	// }
	//
	// // TODO Сделать обход версии получение всех пользователей хранилища
	// // TODO Сделать запись файла AUTHORS
	// // TODO В плагины добавить поддержку записи файла AUTHORS
	//
	// err = r.WriteVersionFile(0)
	//
	// if err != nil {
	//	return err
	// }

	return

}
