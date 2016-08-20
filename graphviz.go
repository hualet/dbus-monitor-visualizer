package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

type node string

type line struct {
	from  node
	to    node
	label string
}

type graph struct {
	nodes []node
	lines []line
}

const (
	header string = "digraph G {\n"
	foot   string = "}"
)

func newGraph() *graph {
	return &graph{}
}

func (g *graph) addNode(nd node) {
	for _, n := range g.nodes {
		if n == nd {
			return
		}
	}

	g.nodes = append(g.nodes, nd)
}

func (g *graph) addLine(ln line) {
	g.addNode(ln.from)
	g.addNode(ln.to)
	g.lines = append(g.lines, ln)
}

func generateNodeString(n node) string {
	return fmt.Sprintf("\"%v\";\n", n)
}

func generateLineString(l line) string {
	return fmt.Sprintf("\"%v\" -> \"%v\" [label =\"%v\"];\n", l.from, l.to, l.label)
}

func (g *graph) generateDotFile(path string) error {
	buf := bytes.NewBufferString(header)

	for _, n := range g.nodes {
		_, err := buf.WriteString(generateNodeString(n))
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	buf.WriteString("\n")

	for _, l := range g.lines {
		_, err := buf.WriteString(generateLineString(l))
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	_, err := buf.WriteString(foot)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return ioutil.WriteFile(path, buf.Bytes(), 0664)
}

func (g *graph) render() {

}
