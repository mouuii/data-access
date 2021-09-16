package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

// catch the first err
func test() {
	var eg errgroup.Group
	for i := 0; i < 100; i++ {
		i := i
		eg.Go(func() error {
			time.Sleep(2 * time.Second)
			if i > 90 {
				fmt.Println("Error:", i)
				return fmt.Errorf("Error occurred: %d", i)
			}
			fmt.Println("End:", i)
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		fmt.Println("xxxxxxxxxxxxx")
		log.Fatal(err)
	}
}
