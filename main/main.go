package main

import (
	"log"

	"github.com/isaacd9/multi"
)

type printer struct {
	in multi.Pipe
}

func (p printer) inc(n int) {
	for i := 0; i < n; i++ {
		log.Printf("%v", i)
	}
}

func (p printer) Run() error {
	p.inc(100)
	return nil
}

func main() {
	pipe, err := multi.NewPipe()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	p := multi.NewProcess(printer{in: pipe})
	if err := p.Start(); err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("proc: %v", p)

	state, err := p.Join()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("proc: %v", state)
}
