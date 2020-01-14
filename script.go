package main

import (
	"strings"
)

type M = map[string]interface{}

func OnMessage() []func(M) M {
	return []func(M) M{
		func(msg M) M {
			if strings.HasPrefix(msg["message"].(M)["text"].(string), "reply ") {
				msg["mode"] = "reply"
			} else {
				msg["mode"] = "send"
			}
			return msg
		},
	}
}
