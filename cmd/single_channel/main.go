package main

import (
	"fmt"
	"github.com/GoosvandenBekerom/go-sync/util"
	"os"
	"time"
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
	start := time.Now()
	total, err := util.GetFileSize(path)
	check(err)

	n := total / load
	i := int64(0)
	c := make(chan []byte)
	for i <= n {
		go worker(int(i), i*load, load, c)
		i++
	}

	var r int64
	for r < total {
		b, open := <-c
		if !open {
			return
		}
		r += int64(len(b))
		println(fmt.Sprintf("bytes received from channel: %s, total bytes read: %d", string(b), r))
	}
	close(c)
	t := time.Now()
	println(fmt.Sprintf("processed file %s in %d", path, t.Sub(start)))
}

func worker(id int, start, size int64, c chan<- []byte) {
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
		println(fmt.Sprintf("worker %d total %d", id, r))
		_, err = f.Seek(start+r, 0)
		check(err)
	}
	err = f.Close()
	check(err)
}
