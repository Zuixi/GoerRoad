package basic

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"log"
)

func CreateFile(filepath string) bool{

	if filepath == "" {
		fmt.Println("filepath is empty")
		return false
	}

	source, err := os.Create(filepath)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	defer source.Close()

	// Write something in the file
	source.WriteString("Hello, use go to operate file")
	return true

}

func OpenFile(filepath string) bool {

	if filepath == "" {
		log.Println("filepath is empty")
		return false
	}

	source, err := os.OpenFile(filepath, os.O_APPEND | os.O_CREATE | os.O_RDWR, 0644)

	if err != nil {
		log.Println(err.Error())
		return false
	}
	defer source.Close()

	source.WriteString("\nwrite something in new line")
	return true

}

func ReadStreamFromFile(filepath string) bool {
	if filepath == "" {
		log.Println("filepath is empty")
		return false
	}

	source, err := os.OpenFile(filepath, os.O_CREATE | os.O_RDWR, 0644)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	defer source.Close()

	// get stream from file with bufio
	reader := bufio.NewReader(source)

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		fmt.Println(line)
	}
	return true

}