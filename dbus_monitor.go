package main

import (
	"log"
	"strconv"
	"strings"
)

type MessageType string
type Address string

const (
	TypeMethodCall   MessageType = "mc"
	TypeMethodReturn MessageType = "mr"
	TypeSignal       MessageType = "sig"
)

type Message struct {
	Type        MessageType
	Timestamp   float64
	Serial      int
	Sender      Address
	Destination Address
	Path        string
	Interface   string
	Member      string
}

func parseLine(line string) *Message {
	fields := strings.Fields(line)
	if len(fields) != 8 {
		// TODO(hualet): or maybe it's type is method return
		log.Println("parse line error: ", line)
		return nil
	}

	t, err := strconv.ParseFloat(fields[1], 32)
	if err != nil {
		log.Println("parse timestamp error")
		return nil
	}

	s, err := strconv.Atoi(fields[2])
	if err != nil {
		log.Println("parse serial error")
		return nil
	}

	return &Message{
		Type:        MessageType(fields[0]),
		Timestamp:   t,
		Serial:      s,
		Sender:      Address(fields[3]),
		Destination: Address(fields[4]),
		Path:        fields[5],
		Interface:   fields[6],
		Member:      fields[7],
	}
}
