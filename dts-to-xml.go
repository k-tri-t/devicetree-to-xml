package main

import (
	"fmt"
	"os"
	"bufio"
	// "regexp"
	"strings"
	// "unicode"
)

var state int = 0
var depth int = 0

type Node struct {
	Name string
	NumChilds int
	Parent *Node
	Child *Node
//	Childs []Node
}

func main() {
	var top Node
	var p *Node
	p = &top
	state := ""

	top.Name = "top"
	fmt.Println("init: ", top)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		trimedInput := strings.TrimSpace(input.Text())
		if(trimedInput == "") {
			// blank line
			state = "-"
		} else if(strings.HasSuffix(trimedInput, "{")) {
			// node start
			p.NumChilds++
			p.Child = new(Node)
			p.Child.Name = "child"
			p.Child.NumChilds = 0
			fmt.Println(p)
			p.Child.Parent = p
			p = p.Child
			state = "S"
			depth++
		} else if(strings.HasSuffix(trimedInput, "};")) {
			// node end
			state = "E"
			depth--
		} else {
			// node body
			state = "B"
		}
		fmt.Printf("%s [%d] %s\n", state, depth, input.Text())
	}
}
