package main

import (
	"github.com/kohmebot/jm/jm"
	"github.com/kohmebot/plugin"
)

func NewPlugin() plugin.Plugin {
	return jm.NewPluginJM()
}
