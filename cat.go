package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	for _, v := range os.Args[1:] {
		r, error := os.Open(v)

		check(error)

		if _, err := io.Copy(os.Stdout, r); err != nil {
			log.Fatal(err)
		}
	}

}
