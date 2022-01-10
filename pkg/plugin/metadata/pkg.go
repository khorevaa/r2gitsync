package metadata

//
// type PkgMetadata struct {
// 	name    string
// 	version string
// 	plugins []PluginSymbol
// 	file    string
// 	hash    crypto.Hash
//
// 	check func() error
// }
//
// type PkgSymbol interface {
// 	Name() string
// 	Version() string
// 	Plugins() []PluginSymbol
// }
//
// func (m *PkgMetadata) Checker(fn func() error) {
//
// 	m.check = fn
//
// }
//
// func (m PkgMetadata) Check() error {
//
// 	if m.check != nil {
// 		return m.check()
// 	}
// 	return nil
// }
//
// func (m PkgMetadata) File() string {
//
// 	return m.file
// }
//
// func (m PkgMetadata) LoadFile() (pm.Symbol, error) {
//
// 	var pSymbol pm.Symbol
//
// 	pluginFile, err := pm.Open(m.file)
// 	if err != nil {
// 		return pSymbol, err
// 	}
//
// 	pSymbol, err = pluginFile.Lookup(PkgSymbolName)
//
// 	if err == nil {
// 		return pSymbol, nil
// 	}
//
// 	pSymbol, err = pluginFile.Lookup(PluginSymbolName)
//
// 	return pSymbol, err
// }
//
// func NewPkgMetadata(file string, sym PkgSymbol) PkgMetadata {
//
// 	return PkgMetadata{
// 		name:    sym.Name(),
// 		version: sym.Version(),
// 		plugins: sym.Plugins(),
// 		file:    file,
// 		hash:    0,
// 		check:   nil,
// 	}
//
// }
