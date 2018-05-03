package main 

import (
	"flag"
	"fmt"
	"os"
	"log"
)

func main() {
	var filename string
	var filecontext string

	textPtr := flag.String("controller", "", "Text to parse.")
	flag.Parse()
		
	fmt.Printf("textPtr: %s", *textPtr)
	filecontext = "<?php\nnamespace " + *textPtr + "\\Controller;\n use Application\\Mvc\\Controller;\nclass " +*textPtr + "Controller extends Controller\n{\n//methods goes here...\n}\n"
	filename = *textPtr + "Controller.php"
	file, err := os.Create(filename)
    if err != nil {
        log.Fatal("Cannot create file", err)
    }
    defer file.Close()

	fmt.Fprintf(file, filecontext)
}