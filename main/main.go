package main

import (
	"log"

	"github.com/isaacd9/multi"
)

type printer int

func (p printer) inc(n int) {
	for i := 0; i < n; i++ {
		log.Printf("%v", i)
	}
}

func (p printer) Run() error {
	p.inc(10)
	return nil
}

func main() {
	p := multi.New(printer(0))
	err := p.Start()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("proc: %v", p)

	state, err := p.Join()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("proc: %v", state)
}
