package main

import (
	"bufio"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("dbus-monitor", "--profile")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}

	go func() {
		bufReader := bufio.NewReader(stdout)
		graphObj := newGraph()
		for i := 0; i < 20; i++ {
			lineStr, _, err := bufReader.ReadLine()
			if err == nil {
				msg := parseLine(string(lineStr))
				if msg != nil {
					fromNode := node(serviceNameFromBusAddress(msg.Sender))
					if !isValidServiceName(string(fromNode)) {
						fromNode = node(processNameFromBusAddress(msg.Sender))
					}
					toNode := node(serviceNameFromBusAddress(msg.Destination))
					if !isValidServiceName(string(toNode)) {
						toNode = node(processNameFromBusAddress(msg.Destination))
					}

					if fromNode == "" || toNode == "" {
						continue
					}

					label := msg.Member
					graphObj.addLine(line{fromNode, toNode, label})
				}
			}
		}
		graphObj.generateDotFile("/tmp/test.dot")
	}()

	if err = cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
