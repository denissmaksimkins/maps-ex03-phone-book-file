package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func createPhoneBook(names []string, numbers []string) map[string]string {
	m := make(map[string]string)
	for i := range names {
		m[names[i]] = numbers[i]
	}
	return m
}

func main() {
	fmt.Println("This program creates a phone boook by reading a txt file")

	f, err := os.Open("phone_numbers.txt")
	if err != nil {
		log.Fatalf("open: %v", err)
	}
	defer f.Close()

	var names, numbers []string
	s := bufio.NewScanner(f)
	isname := true
	for s.Scan() {
		if isname {
			names = append(names, s.Text())
			isname = false
		} else {
			numbers = append(numbers, s.Text())
			isname = true
		}
	}
	if err := s.Err(); err != nil {
		log.Fatalf("scanning: %v", err)
	}
	fmt.Println(createPhoneBook(names, numbers))
}
