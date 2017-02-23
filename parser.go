package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) {

	// Open file and such
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln("Couldn't open file")
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	bytes, _, err := reader.ReadLine()
	if err != nil {
		log.Fatalln("Couldn't read")
	}

	// 1. Parse first line
	line := string(bytes)
	numbers := strings.Split(line, " ")
	// Fill variables
	V, _ = strconv.Atoi(numbers[0])
	E, _ = strconv.Atoi(numbers[1])
	R, _ = strconv.Atoi(numbers[2])
	C, _ = strconv.Atoi(numbers[3])
	X, _ = strconv.Atoi(numbers[4])

	// - Allocate memory
	VideoSizes = make([]int, V)
	Endpoints = make([]Endpoint, E)
	Requests = make([]Request, R)
	freeSpace = make([]int, C)

	// 2. Parse second line
	line, _ = reader.ReadString('\n')
	line = line[:len(line)-1]
	numbers = strings.Split(line, " ")

	// Go readLine gets ioly upto 1053 fields!!!
	// Fill variables
	for i := 0; i < V; i++ {
		VideoSizes[i], _ = strconv.Atoi(numbers[i])
	}

	// 3. Read Endpoints
	for i := 0; i < E; i++ {
		bytes, _, _ = reader.ReadLine()
		line = string(bytes)
		numbers = strings.Split(line, " ")
		Endpoints[i].LD, _ = strconv.Atoi(numbers[0])
		Endpoints[i].K, _ = strconv.Atoi(numbers[1])
		// Allocate memory for cache Servers in this Endpoint
		Endpoints[i].CacheServers = make([]CacheServer, Endpoints[i].K)
		for n := 0; n < Endpoints[i].K; n++ {
			bytes, _, _ = reader.ReadLine()
			line = string(bytes)
			numbers = strings.Split(line, " ")
			Endpoints[i].CacheServers[n].C, _ = strconv.Atoi(numbers[0])
			Endpoints[i].CacheServers[n].Lc, _ = strconv.Atoi(numbers[1])
		}
	}
	// 4. Read Requests
	for i := 0; i < R; i++ {
		bytes, _, _ = reader.ReadLine()
		line = string(bytes)
		numbers = strings.Split(line, " ")
		Requests[i].Rv, _ = strconv.Atoi(numbers[0])
		Requests[i].Re, _ = strconv.Atoi(numbers[1])
		Requests[i].Rn, _ = strconv.Atoi(numbers[2])
	}
}
