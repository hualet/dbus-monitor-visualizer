package main

import (
	"dbus/org/freedesktop/dbus"
	"fmt"
	"io/ioutil"
	"log"
)

type processID uint32

var manager *dbus.DBusDaemon
var knownAddress map[Address]processID

func init() {
	var err error

	manager, err = dbus.NewDBusDaemon("org.freedesktop.DBus", "/")
	if err != nil {
		panic(err)
	}

	knownAddress = make(map[Address]processID)
	allNames, err := manager.ListNames()
	if err != nil {
		log.Print("ListNames failed :", err)
	}
	for _, name := range allNames {
		id := processIDFromBusAddress(Address(name))
		knownAddress[Address(name)] = processID(id)
	}
}

func processIDFromBusAddress(addr Address) processID {
	v, ok := knownAddress[addr]
	if ok {
		return v
	}

	id, err := manager.GetConnectionUnixProcessID(string(addr))
	if err != nil {
		return 0
	}

	result := processID(id)
	knownAddress[addr] = result

	return result
}

func processNameFromBusAddress(addr Address) string {
	pid := processIDFromBusAddress(addr)
	name, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", pid))
	if err != nil {
		log.Print("get process name error: ", err)
	}

	return string(name)
}
