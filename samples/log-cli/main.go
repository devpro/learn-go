package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	path := flag.String("path", "myapp.log", "The path to the log that should be analyzed")
	level := flag.String("level", "ERROR", "Log level to search for. Options are DEBUG, INFO, ERROR, and CRITICAL")

	flag.Parse()

	file, error := os.Open(*path)
	if error != nil {
		log.Fatal(error)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, error := reader.ReadString('\n')
		if error != nil {
			break
		}
		if strings.Contains(line, *level) {
			fmt.Println(line)
		}
	}
}
