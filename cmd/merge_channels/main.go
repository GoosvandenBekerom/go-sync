package main

import (
	"fmt"
	"github.com/GoosvandenBekerom/go-sync/util"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	path   = "resources/large_file.txt"
	buffer = 32
	load   = buffer * 30
)

func main() {
	util.ExecuteMeasured(example, fmt.Sprintf("processed file %s", path))
}

func example() {
	total, err := util.GetFileSize(path)
	check(err)

	n := total / load
	i := int64(0)
	var cs []<-chan []byte
	for i <= n {
		c := make(chan []byte)
		cs = append(cs, c)
		go worker(int(i), i*load, load, c)
		i++
	}

	for b := range util.Merge(cs...) {
		println(fmt.Sprintf("bytes received: %d", len(b)))
	}
}

func worker(id int, start, size int64, c chan<- []byte) {
	defer close(c)
	f, err := os.Open(path)
	check(err)
	_, err = f.Seek(start, 0)
	check(err)

	var r int64
	for r < size {
		b := make([]byte, buffer)
		n, err := f.Read(b)
		check(err)
		c <- b
		println(fmt.Sprintf("worker %d read %d bytes", id, n))
		if n < buffer {
			break
		}
		r += int64(n)
		_, err = f.Seek(start+r, 0)
		check(err)
	}
	err = f.Close()
	check(err)
}
