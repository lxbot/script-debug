package main

import (
	"github.com/lxbot/lxlib/v2"
	"github.com/lxbot/lxlib/v2/common"
)

func main() {
	script, messageCh := lxlib.NewScript()

	for {
		message := <-*messageCh
		response, err := message.Copy()
		if err != nil {
			common.ErrorLog("message copy error:", err)
			continue
		}
		script.SendMessage(response.Reply().AddContent("script-debug"))
	}
}
