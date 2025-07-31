package jm

import (
	"fmt"
	"github.com/kohmebot/pkg/chain"
	"github.com/kohmebot/pkg/gopool"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/extension"
	"github.com/wdvxdr1123/ZeroBot/message"
	"strconv"
	"strings"
)

func (p *PluginJM) SetOnJM(engine *zero.Engine) {
	engine.OnCommand("jm", p.env.Groups().Rule()).SetBlock(true).Handle(func(ctx *zero.Ctx) {
		var cmd extension.CommandModel
		err := ctx.Parse(&cmd)
		if err != nil {
			p.env.Error(ctx, err)
			return
		}
		uid := ctx.Event.UserID

		ok, mid := p.t.AddTask(uid)
		if !ok {
			var msg chain.MessageChain
			if mid.ID() == 0 {
				msg.SplitEmpty(
					message.At(uid),
					message.Text("你还有任务在下载中哦..."),
				)
				ctx.Send(msg)
			} else {
				msg.Join(message.Reply(mid))
				msg.SplitEmpty(
					message.At(uid),
					message.Text("你在短时间内已经请求过一次了，待会再来吧..."),
				)
				ctx.Send(msg)
			}
			return
		}

		args := strings.TrimSpace(cmd.Args)
		aid, err := strconv.ParseInt(args, 10, 64)
		if err != nil {

			p.env.Error(ctx, fmt.Errorf(`无法解析"%s": %w`, args, err))
			p.t.Fail(uid)
			return
		}
		var msg chain.MessageChain
		msg.SplitEmpty(
			message.At(uid),
			message.Text(fmt.Sprintf("正在下载%d...", aid)),
		)
		ctx.Send(msg)

		gopool.Go(func() {
			var err error
			var mid message.ID
			defer func() {
				if err != nil {
					p.env.Error(ctx, err)
					p.t.Fail(uid)
					return
				}
				p.t.Done(uid, mid)
			}()

			mid = ctx.Send(message.File(p.svr.DownloadUrl(aid), fmt.Sprintf("%d.pdf", aid)))
			if mid.ID() == 0 {
				err = fmt.Errorf("文件发送失败")
				return
			}

			var msg chain.MessageChain
			msg.SplitEmpty(
				message.Reply(mid),
				message.At(uid),
				message.Text(fmt.Sprintf("%d 下载完成，航班启航🛫", aid)),
			)
			ctx.Send(msg)
		})

	})
}
