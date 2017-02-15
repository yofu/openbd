package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/yofu/openbd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: book {isbn...}")
		os.Exit(0)
	}
	isbn := strings.Join(os.Args[1:], ",")
	obd, err := openbd.GetOpenBD(isbn, 10*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	for i, o := range obd {
		fmt.Println("Book", i+1)
		fmt.Println("    Title:", o.Summary.Title)
		fmt.Println("    Author:", o.Summary.Author)
	}
}
