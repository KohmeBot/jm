package jm

import (
	"github.com/kohmebot/pkg/chain"
	"github.com/kohmebot/plugin/v2"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
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

func (p *PluginJM) ConfigModel() any {
	return new(Config)
}

func (p *PluginJM) OnInit(engine plugin.Engine, env plugin.Env) error {
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

func (p *PluginJM) OnHelp(ctx *zero.Ctx) {
	var msg chain.MessageChain

	msg.Split(
		message.Text("jm 插件所有命令"),
		message.Text("jm <jm code>：下载jm pdf"),
	)

	ctx.Send(msg)
}

func (p *PluginJM) Version() string {
	return "v0.1.1"

}

func (p *PluginJM) OnBoot() {

}
