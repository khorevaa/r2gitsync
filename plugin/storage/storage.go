package storage

import (
	"crypto"
	"errors"
	"github.com/khorevaa/r2gitsync/log"
	"github.com/khorevaa/r2gitsync/plugin"
	"github.com/khorevaa/r2gitsync/plugin/metadata"
	"github.com/recoilme/pudge"
	"io/ioutil"
	"os"
	"path/filepath"
	pm "plugin"
	"regexp"
	"sync"
)

const databaseFilename = "cache"
const PkgSymbolName = "PluginPackage"

type PluginStorage interface {

	Load() error // Загружает плагины в внутренний индекс
	Reload() error // Перечитывае плагины в внутренний индекс

	Store(filename string) error
	Delete(name string) error
	List() ([]metadata.PluginMetadata, error)

	Clear() error
	Registry(m plugin.Manager) error

}

type FilePluginStorage struct {
	mu  sync.Mutex
	Dir string
	PkgIdx map[string]metadata.PkgMetadata
	PluginsIdx map[string][]metadata.PluginMetadata

	loaded bool
	database  string
}

func NewFilePluginStorage(dir string) *FilePluginStorage {

	return &FilePluginStorage{
		Dir: dir,
		database: filepath.Join(dir, databaseFilename),
		PkgIdx: make(map[string]metadata.PkgMetadata),
		PluginsIdx: make(map[string][]metadata.PluginMetadata),
	}
}

func (s *FilePluginStorage) Store(filename string) error {


	// TODO Копирование файла
	// TODO Проверка файла на плагины и зашрузка их метаданных

	return nil

}

func (s *FilePluginStorage) Delete(name string) error {

	return nil

}

func (s *FilePluginStorage) Registry(mng plugin.Manager) error {

	err := s.Load()
	if err != nil {
		return err
	}
	var pl []plugin.Symbol

	for _, pkgMetadata := range s.PkgIdx {

		err := pkgMetadata.Check()

		if err != nil {
			continue
		}

		pkgPlugins, _ := loadPkgPlugins(pkgMetadata.)

		pl = append(pl, pkgPlugins...)

	}

	for _, storagePlugin := pl {
		_ = mng.Register(storagePlugin)
	}

	return nil
}

func (s *FilePluginStorage) openDb() (*pudge.Db, error) {

	cfg := &pudge.Config{
		SyncInterval: 30,
	} // every second fsync
	db, err := pudge.Open(s.database, cfg)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (s *FilePluginStorage) List() (ls []pluginMetadata, err error) {

	err = s.Load()
	if err != nil {
		return ls, err
	}

	for _, metadata := range s.PkgIdx {

		ls = append(ls, metadata.plugins...)

	}

	return
}

func (s *FilePluginStorage) Reload() error {

	err := s.updatePluginsFromDir()

	return err
}


func (s *FilePluginStorage) Load() error {

	if s.loaded {
		return nil
	}
	
	err := s.loadData()
	return err
}

func (s *FilePluginStorage) Clear() error {

	return clearDir(s.Dir)
}

func (s *FilePluginStorage) loadData() error {

	db, err := s.openDb()
	if err != nil {
		return err
	}

	defer db.Close()

	keys, err := db.Keys(nil, 0, 0, true)
	if err == nil {
		for _, key := range keys {

			var metadata pkgMetadata
			err = db.Get(key, &metadata)
			if err != nil {
				return err
			}

			s.PkgIdx[string(key)] = metadata
		}
	}

	s.loaded = true

	err = s.checkPkgIdx()
	if err != nil {
		return err
	}

	return nil

}

func (s *FilePluginStorage) checkPkgIdx() error {

	err := s.Load()
	if err != nil {
		return err
	}

	for _, metadata := range s.PkgIdx {
		// TODO Провека наличия и хеша файла пакета
		println(metadata)
	}

	return nil
}

func (s *FilePluginStorage) savePkgIdx() error {

	db, err := s.openDb()
	if err != nil {
		return err
	}

	defer db.Close()

	for n, k := range s.PkgIdx {

		err = db.Set(n, k)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *FilePluginStorage) updatePluginsFromDir() error {

	if len(s.Dir) == 0 {
		return nil
	}
	if _, err := os.Stat(s.Dir); err != nil {
		return err
	}

	files, err := listFiles(s.Dir, `*.so`)
	if err != nil {
		return err
	}

	s.PkgIdx = make(map[string]pkgMetadata)

	for _, file := range files {
		filename := filepath.Join(s.Dir, file.Name())

		pkg, err := loadPkgMetadata(filename)
		if err != nil {
			log.Error(err)
			continue
		}

		s.PkgIdx[pkg.name] = pkg

	}

	err = s.savePkgIdx()

	return err
}

func loadPluginFile(filename string) (pm.Symbol, error) {

	var pSymbol pm.Symbol

	pluginFile, err := pm.Open(filename)
	if err != nil {
		return pSymbol, err
	}

	pSymbol, err = pluginFile.Lookup(PkgSymbolName)

	if err == nil {
		return pSymbol, nil
	}

	pSymbol, err = pluginFile.Lookup(plugin.SymbolName)

	return pSymbol, err

}

func loadPkgPlugins(filename string) (pl []plugin.Symbol, err error) {

	pSymbol, err := loadPluginFile(filename)

	if err != nil {
		return pl, err
	}

	switch sym := pSymbol.(type) {

	case PkgSymbol:

		return sym.Plugins(), nil

	case plugin.Symbol:

		pl = append(pl, sym)

	default:
		return nil, errors.New("plugin is not implement <plugin.Symbol>")
	}

	return
}

func loadPkgMetadata(filename string) (pkg pkgMetadata, err error) {

	pSymbol, err := loadPluginFile(filename)

	if err != nil {
		return pkg, err
	}

	switch sym := pSymbol.(type) {

	case PkgSymbol:

		pkg = pkgMetadata{
			name: sym.Name(),
			file: filename,
			hash: crypto.MD5, // TODO Расчет хеша файла
		}

		pkg.plugins = pkg.getPluginMetadata(sym.Plugins()...)

	case plugin.Symbol:

		pkg = pkgMetadata{
			name: sym.Name(),
			file: filename,
			hash: crypto.MD5, // TODO Расчет хеша файла
		}
		pkg.plugins = pkg.getPluginMetadata(sym)

	default:
		return pkgMetadata{}, errors.New("plugin is not impliment <plugin.Symbol>")
	}

	return
}

func (pkg pkgMetadata) getPluginMetadata(p... plugin.Symbol) (pm []pluginMetadata) {

	for _, symbol := range p {
		pm = append(pm, pluginMetadata{
			name: symbol.Name(),
			file: pkg.file,
			desc: symbol.Desc(),
			modules: symbol.Modules(),
			ver:  symbol.ShortVersion(),
			pkg:  pkg.name,
			hash: pkg.hash,
		})
	}
	return
}

func clearDir(dir string) error {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, f := range files {

		err = os.Remove(f.Name())
		if err != nil {
			return err
		}
		//fmt.Println(f.Name())
	}

	return nil
}


func listFiles(dir, pattern string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var filteredFiles []os.FileInfo
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		matched, err := regexp.MatchString(pattern, file.Name())
		if err != nil {
			return nil, err
		}
		if matched {
			filteredFiles = append(filteredFiles, file)
		}
	}
	return filteredFiles, nil
}