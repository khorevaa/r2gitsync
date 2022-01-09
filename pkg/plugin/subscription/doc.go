package subscription

//
//type Subscriber struct {
//
//	UpdateCfg  UpdateCfgSubscriber
//	Handlers   []interface{}
//
//}
//
//func WithUpdateCfg(u UpdateCfgSubscriber) func(s *Subscriber) {
//
//	return func(s *Subscriber) {
//		s.UpdateCfg = u
//	}
//}
//
//func NSubscriber(handlers... interface{}) Subscriber {
//
//	s := Subscriber{
//		Handlers: handlers,
//	}
//	//s.options(opts...)
//
//	return s
//
//}
//
//
//func NewSubscriber(opts... func(s *Subscriber)) Subscriber {
//
//	s := Subscriber{}
//	s.options(opts...)
//
//	return s
//
//}
//
//func (s *Subscriber) options(opts... func(s *Subscriber))  {
//
//	for _, opt := range opts {
//		opt(s)
//	}
//
//}
//
//func (s Subscriber) Empty() bool  {
//
//	return s.UpdateCfg.Empty()
//
//}
//
//type UpdateCfgSubscriber struct {
//	On OnUpdateCfgFn
//	Start BeforeUpdateCfgFn
//	Finish AfterUpdateCfgFn
//}
//
//func (s UpdateCfgSubscriber) Empty() bool {
//	return s.On == nil &&
//		s.Finish == nil &&
//		s.Start == nil
//}
//
//func (s UpdateCfgSubscriber) OnFn(fn OnUpdateCfgFn) UpdateCfgSubscriber {
//	s.On = fn
//	return s
//}
//
//func (s UpdateCfgSubscriber) BeforeFn(fn BeforeUpdateCfgFn) UpdateCfgSubscriber {
//	s.Start = fn
//	return s
//}
//
//func (s UpdateCfgSubscriber) AfterFn(fn AfterUpdateCfgFn) UpdateCfgSubscriber {
//	s.Finish = fn
//	return s
//}
