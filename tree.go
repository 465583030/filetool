package main

import (
	"os"
	"path/filepath"
	"fmt"
)

type treeNode struct {
	fullPath string
	info     os.FileInfo
	size int64
	parent   *treeNode
	children []*treeNode
}

func NewTreeNode(fullPath string) *treeNode {
	info, err := os.Stat(fullPath)
	if err != nil {
		panic(err)
	}
	return &treeNode{
		fullPath: fullPath,
		info:     info,
	}
}

func (n *treeNode) AppendChild(c *treeNode) {
	if c != nil {
		c.parent = n
		n.children = append(n.children, c)
	}
}

func (n *treeNode) String() string {
	p := n.fullPath
	if n.parent != nil {
		p = filepath.Base(n.fullPath)
	}

	typ := "F"
	if n.info.IsDir() {
		typ = "D"
	}

	return fmt.Sprintf("[%s] %5s %s %s", typ, readableBytes(n.size), n.info.Mode().String(), p)
}

func (n *treeNode) Children() []interface{} {
	kids := make([]interface{}, len(n.children))
	for i, c := range n.children {
		kids[i] = c
	}
	return kids
}
