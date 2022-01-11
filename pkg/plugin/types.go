package plugin

type subscriber struct {
	handlers []interface{}
}

func (s subscriber) Handlers() []interface{} {
	return s.handlers
}
