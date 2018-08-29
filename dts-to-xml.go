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
	// NumChilds int
	// Parent *Node
	Children []*Node
//	Childs []Node
}

var root *Node

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

func showNode(node *Node, prefix string) {
    if prefix == "" {
        fmt.Printf("%v\n\n", node.name)
    } else {
        fmt.Printf("%v %v\n\n", prefix, node.name)
    }
    for _, n := range node.children {
        showNode(n, prefix+"--")
    }
}

func show() {
    if root == nil {
        fmt.Printf("show: root node not found\n")
        return
    }
    fmt.Printf("RESULT:\n")
    showNode(root, "")
}
