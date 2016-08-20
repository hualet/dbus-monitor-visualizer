package main

import (
	"dbus/org/freedesktop/dbus"
)

type processID uint32

var manager *dbus.DBusDaemon

func init() {
	var err error

	manager, err = dbus.NewDBusDaemon("org.freedesktop.DBus", "/")
	if err != nil {
		panic(err)
	}
}

func processFromBusAddress(addr string) processID {
	id, err := manager.GetConnectionUnixProcessID(addr)
	if err != nil {
		return 0
	}

	return processID(id)
}
