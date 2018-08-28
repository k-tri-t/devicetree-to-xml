package main

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
)

const (
	NodeStart = 0 << iota
	NodeBody
	NodeEnd
)

var state int = 0
var depth int = 0

type Node struct {
	Name string
	NumChilds int
	Child *Node
//	Childs []Node
}

func main() {
	regNodeStart := regexp.MustCompile(`\{`)
	regNodeEnd := regexp.MustCompile(`\};`)
	var top Node
	var p *Node
	p = &top

	top.Name = "top"
	fmt.Println("init: ", top)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if(regNodeStart.MatchString(input.Text())) {
			p.NumChilds++
			fmt.Println(p.NumChilds)
			p.Child = new(Node)
			p = p.Child
			p.Name = "child"
			p.NumChilds = 0
			state = 0
			depth++
		} else if(regNodeEnd.MatchString(input.Text())) {
			state = 1
			depth--
		} else {
			state = 2
		}
		fmt.Printf("%d [%d] %s\n", state, depth, input.Text())
	}
}
