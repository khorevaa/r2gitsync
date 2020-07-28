package main

import (
	"errors"
	"github.com/v8platform/designer/repository"
	v8 "github.com/v8platform/v8"
	"sync"
)

type SubscribeManager struct {
	mu            sync.Mutex
	Subscriptions map[string][]*Subscription
	ssid          int
}

type Subscription struct {
	sid   int
	topic string
	fn    interface{}
}

func (sm *SubscribeManager) Subscribe(topic string, fn interface{}) error {

	err := checkSubscription(topic, fn)
	if err != nil {
		return err
	}

	sub := &Subscription{
		topic: topic,
		fn:    fn,
	}

	sm.mu.Lock()
	sm.ssid++
	sub.sid = sm.ssid
	sm.Subscriptions[topic] = append(sm.Subscriptions[topic], sub)
	sm.mu.Unlock()

	return nil

}

var ErrorFuncType = errors.New("error subscribe func type")

type BeforeUpdateCfgFunc func(workdir string, infobase v8.Infobase, repository repository.Repository,
	version int64, extention string) error

func checkSubscription(topic string, fn interface{}) error {

	switch {
	case topic == "BeforeUpdateCfg":

		if _, ok := fn.(BeforeUpdateCfgFunc); !ok {
			return ErrorFuncType
		}

	default:
		return errors.New("unknown topic to subscribe")
	}

	return nil
}

func (m *SubscribeManager) BeforeUpdateCfgHandler(workdir string, infobase v8.Infobase, repository repository.Repository,
	version int64, extention string) error {

	subs := m.Subscriptions["BeforeUpdateCfg"]

	for _, sub := range subs {

		fn := sub.fn.(BeforeUpdateCfgFunc)
		err := fn(workdir, infobase, repository, version, extention)

		if err != nil {
			return err
		}
	}

	return nil

}
