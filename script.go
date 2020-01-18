package main

import (
	"fmt"
	"plugin"
	"strings"
)

type M = map[string]interface{}

var store *plugin.Plugin

func Boot(s *plugin.Plugin, c *chan M) {
	store = s
}

func OnMessage() []func(M) M {
	return []func(M) M{
		func(msg M) M {
			text := msg["message"].(M)["text"].(string)
			last := get(msg)
			set(msg, text)
			if last != "" {
				msg["message"].(M)["text"] = "last: "+last+"\n"+text
			}
			if strings.HasPrefix(text, "reply ") {
				msg["mode"] = "reply"
			} else {
				msg["mode"] = "send"
			}
			return msg
		},
	}
}

func get(m M) string {
	fn, err := store.Lookup("Get")
	if err != nil {
		_ = fmt.Errorf("%v", err)
		return ""
	}
	result := fn.(func(string) interface{})(key(m))
	if result == nil {
		return ""
	}
	return result.(string)
}

func set(m M, value string) {
	fn, err := store.Lookup("Set")
	if err != nil {
		_ = fmt.Errorf("%v", err)
		return
	}
	fn.(func(string, interface{}))(key(m), value)
}

func key(m M) string {
	roomID := m["room"].(M)["id"].(string)
	userID := m["user"].(M)["id"].(string)
	return "last_msg_" + roomID + "_" + userID
}