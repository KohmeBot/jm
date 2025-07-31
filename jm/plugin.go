package jm

import (
	"fmt"
	"github.com/kohmebot/pkg/command"
	"github.com/kohmebot/pkg/version"
	"github.com/kohmebot/plugin"
	zero "github.com/wdvxdr1123/ZeroBot"
	"time"
)

type PluginJM struct {
	env  plugin.Env
	conf Config
	t    *TaskDuration
	svr  *Service
}

func NewPlugin() plugin.Plugin {
	return new(PluginJM)
}

func (p *PluginJM) Init(engine *zero.Engine, env plugin.Env) error {
	p.env = env

	err := p.env.GetConf(&p.conf)
	if err != nil {
		return err
	}

	p.svr = NewService(p.conf.Address)
	err = p.svr.Ping()
	if err != nil {
		return err
	}

	p.t = NewTaskDuration(time.Second * time.Duration(p.conf.CD))

	p.SetOnJM(engine)

	return nil
}

func (p *PluginJM) Name() string {
	return "jm"
}

func (p *PluginJM) Description() string {
	return "你的副机长"
}

func (p *PluginJM) Commands() fmt.Stringer {
	return command.NewCommands(
		command.NewCommand("下载jm pdf", "jm"),
	)
}

func (p *PluginJM) Version() uint64 {
	return uint64(version.NewVersion(0, 0, 12))
}

func (p *PluginJM) OnBoot() {

}
