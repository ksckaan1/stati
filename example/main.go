package main

import (
	"log"
	"time"

	"github.com/ksckaan1/stati"
)

func main() {
	s := stati.New().
		WithAddr(":3000").
		WithChartBuffer(100).
		WithInterval(time.Second).
		WithTitle("Stati")

	log.Println("test server started on port :3000")
	err := s.Start()
	if err != nil {
		log.Fatalln(err)
	}
}
