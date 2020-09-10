package flow

import "github.com/khorevaa/r2gitsync/plugin/subscription"

var _ Flow = (*subscribeTasker)(nil)

type subscribeTasker struct {
	tasker
	pm *subscription.SubscribeManager
}

func (t subscribeTasker) UpdateCfg(v8end V8Endpoint, workDir string, number int64) (err error) {

	UpdateCfg := t.pm.UpdateCfg

	err = UpdateCfg.Before(v8end, workDir, number)

	if err != nil {
		return
	}

	standartHandler := true

	err = UpdateCfg.On(v8end, workDir, number, &standartHandler)

	if err != nil {
		return
	}

	if standartHandler {
		err = t.tasker.UpdateCfg(v8end, workDir, number)
	}

	err = UpdateCfg.After(v8end, workDir, number)

	return nil
}
