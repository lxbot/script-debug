package main

import (
	"github.com/lxbot/lxlib/v2"
	"github.com/lxbot/lxlib/v2/common"
)

type M = map[string]interface{}

func main() {
	script, messageCh := lxlib.NewScript()

	for {
		message := <-*messageCh

		common.TraceLog("(script)", "lxlib.listen()", "event received", "message:", message)

		response, err := message.Copy()
		if err != nil {
			common.ErrorLog("message copy error:", err)
			continue
		}

		key := "last_msg_" + message.User.ID

		lastMessage := script.GetStorage(key)
		if lastMessage == nil {
			lastMessage = ""
		}
		script.SetStorage(key, message.Contents[0].Text)

		script.SendMessage(response.Reply().ResetContents().AddContent("last message: '" + lastMessage.(string) + "'"))
	}
}
