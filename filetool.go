package main

import (
	"flag"
	"fmt"
	"github.com/zhangyuchen0411/goutil/printer"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

var (
	depth   int
	sizeStr string
	size    int64
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.IntVar(&depth, "d", 1, "tree depth")
	flag.StringVar(&sizeStr, "size", "0K", "min size of printed tree node")
}

func main() {
	flag.Parse()

	var err error
	size, err = parseSize(sizeStr)
	if err != nil {
		panic(err)
	}

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
