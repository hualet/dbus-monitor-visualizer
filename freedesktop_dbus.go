package main

import (
	"bytes"
	"dbus/org/freedesktop/dbus"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type processID uint32

var manager *dbus.DBusDaemon
var addressProcMap map[Address]processID
var procServiceMap map[processID]string

func init() {
	var err error

	manager, err = dbus.NewDBusDaemon("org.freedesktop.DBus", "/")
	if err != nil {
		panic(err)
	}

	addressProcMap = make(map[Address]processID)
	procServiceMap = make(map[processID]string)

	allNames, err := manager.ListNames()
	if err != nil {
		log.Print("ListNames failed :", err)
	}
	for _, name := range allNames {
		id := processIDFromBusAddress(Address(name))
		addressProcMap[Address(name)] = processID(id)

		tryUpdateProcServiceMap(processID(id), name)
	}
}

func tryUpdateProcServiceMap(pid processID, name string) {
	// we just want service names, not dbus addresses.
	if isValidServiceName(name) {
		procServiceMap[pid] = name
	}
}

func serviceNameFromProcessID(pid processID) string {
	v, ok := procServiceMap[pid]
	if ok {
		return v
	}

	return ""
}

func serviceNameFromBusAddress(addr Address) string {
	pid := processIDFromBusAddress(addr)
	name := serviceNameFromProcessID(pid)
	if name != "" {
		return name
	}

	return string(addr)
}

func processIDFromBusAddress(addr Address) processID {
	v, ok := addressProcMap[addr]
	if ok {
		return v
	}

	id, err := manager.GetConnectionUnixProcessID(string(addr))
	if err != nil {
		log.Print(err)
		return 0
	}

	result := processID(id)
	addressProcMap[addr] = result
	tryUpdateProcServiceMap(result, string(addr))

	return result
}

func processNameFromProcessID(pid processID) string {
	name, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", pid))
	if err != nil {
		log.Print("get process name error: ", err)
	}

	result := bytes.Trim(name, "\x00")
	return string(result)
}

func processNameFromBusAddress(addr Address) string {
	pid := processIDFromBusAddress(addr)
	return processNameFromProcessID(pid)
}

func isValidServiceName(name string) bool {
	return name != "" && !strings.HasPrefix(name, ":")
}
