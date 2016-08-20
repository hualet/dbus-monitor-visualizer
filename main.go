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
		for {
			line, _, err := bufReader.ReadLine()
			if err == nil {
				msg := parseLine(string(line))
				if msg != nil {
					log.Println("%#v", processFromBusAddress(msg.Sender))
				}
			}
		}
	}()

	if err = cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
