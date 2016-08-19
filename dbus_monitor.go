package main

import (
	"strings"
	"fmt"
	"strconv"
)

type MessageType string

const ( 
    TypeMethodCall MessageType = "mc"
    TypeMethodReturn  MessageType = "mr"
    TypeSignal MessageType = "sig"
)

type Message struct {
    Type MessageType
    Timestamp float64
    Serial int
    Sender string
    Destination string
    Path string 
    Interface string 
    Member string
}

func parseLine(line string) *Message {
    fields := strings.Fields(line)
    if (len(fields) != 8) {
        fmt.Println("parse line error")
        return nil
    }

    t, err := strconv.ParseFloat(fields[1], 32)
    if err != nil {
        fmt.Println("parse timestamp error")
    }

    s, err := strconv.Atoi(fields[2])
    if err != nil {
        fmt.Println("parse serial error")
    }

    return &Message {
        Type: MessageType(fields[0]),
        Timestamp: t,
        Serial: s,
        Sender: fields[3],
        Destination: fields[4],
        Path: fields[5],
        Interface: fields[6],
        Member: fields[7],
    }
}