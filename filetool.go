package main

import (
	"fmt"
	"github.com/zhangyuchen0411/goutil/printer"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"flag"
)

var (
	depth int
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.IntVar(&depth, "d", 1, "tree depth")
}

func main() {
	flag.Parse()

	d, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	root := NewTreeNode(d)

	walk(root)

	s := printer.TreeStringDepth(root, "\n", depth)
	fmt.Println(s)
}

func walk(n *treeNode) {
	files, err := ioutil.ReadDir(n.fullPath)
	if err != nil {
		panic(err)
	}
	childSize := int64(0)
	for _, f := range files {
		c := NewTreeNode(filepath.Join(n.fullPath, f.Name()))
		if f.IsDir() {
			walk(c)
		} else {
			c.size = f.Size()
		}
		childSize += c.size
		n.AppendChild(c)
	}
	n.size = childSize
}
