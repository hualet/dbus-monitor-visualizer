package main

import (
	"testing"
)

func TestParseLine(t *testing.T) {
	m1 := parseLine("sig	1471584508.056224	2	org.freedesktop.DBus	:1.177	/org/freedesktop/DBus	org.freedesktop.DBus	NameAcquired")
	m2 := Message{
		Type:        TypeSignal,
		Timestamp:   1471584508.056224,
		Serial:      2,
		Sender:      "org.freedesktop.DBus",
		Destination: ":1.177",
		Path:        "/org/freedesktop/DBus",
		Interface:   "org.freedesktop.DBus",
		Member:      "NameAcquired",
	}

	if m1.Type != m2.Type {
		t.Error("Type not correct")
	}
}
