package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Hess0r/search/open"
)

func main() {
	link := flag.Bool("l", false, "Flag to open the link directly")
	flag.Parse()
	args := flag.Args()
	var url string
	if !*link {
		keyword := strings.Join(args, "+")
		prefix := "https://www.google.com/search?q="
		url = fmt.Sprintf("%s%s", prefix, keyword)
	} else {
		if len(args) > 1 {
			fmt.Println("Too many arguments")
			os.Exit(1)
		}
		url = args[0]
	}
	open.Open(url)
}
