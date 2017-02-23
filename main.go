package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	readFile("kittens.in")
	ponderate()
	assign()
	calculateFinalScore()
}

// Heuristic parameter for a first viable solution
func ponderate() {
	Possibilities = make([]Possibility, R)
	for i := 0; i < R; i++ {
		Possibilities[i].Re = Requests[i].Re
		Possibilities[i].Rv = Requests[i].Rv
		Possibilities[i].Rn = Requests[i].Rn
		Possibilities[i].Score = VideoSizes[Possibilities[i].Rv] * Requests[i].Rn
	}
	// Sort
	start := time.Now()
	sort.Sort(Possibilities)
	elapsed := time.Since(start)
	log.Println("Sorted in: ", elapsed)
}

func assign() {
	for i := 0; i < R; i++ {
		buffer := Possibility{Possibilities[i].Rv, Possibilities[i].Re, Possibilities[i].Rn, Possibilities[i].Score, nil}
		// Seek the perfect option among possible CacheServers for this Endpoint. nil means "Data Center""
		// Iterate in Endpoints[Possibilities[i].Re].CacheServers
		chosenCacheServer := 0
		for n := 0; n < len(Endpoints[Possibilities[i].Re].CacheServers); n++ {
			if (Endpoints[Possibilities[i].Re].CacheServers[n].Lc < Endpoints[Possibilities[i].Re].CacheServers[chosenCacheServer].Lc) && (freeSpace[Endpoints[Possibilities[i].Re].CacheServers[chosenCacheServer].C] > 0) {
				buffer.Server = &Endpoints[Possibilities[i].Re].CacheServers[n]

			}
			freeSpace[Endpoints[Possibilities[i].Re].CacheServers[chosenCacheServer].C] = freeSpace[Endpoints[Possibilities[i].Re].CacheServers[chosenCacheServer].C] + VideoSizes[Possibilities[i].Rv]
		}
		ChosenPossibilities = append(ChosenPossibilities, buffer)
	}
}

func calculateFinalScore() {
	savings := 0
	nChosen := len(ChosenPossibilities)
	for i := 0; i < nChosen; i++ {
		if ChosenPossibilities[i].Server != nil {
			savings = savings + (Endpoints[ChosenPossibilities[i].Re].LD-ChosenPossibilities[i].Server.Lc)*ChosenPossibilities[i].Rn
		}
	}
	log.Println("SCORE: ", savings)
}

func outPut() {

	usedServers := C
	var ChosenServers []int
	f, _ := os.Create("output")
	defer f.Close()
	w := bufio.NewWriter(f)
	w.WriteString(strconv.Itoa(usedServers) + "\n")

	for index := 0; index < len(ChosenPossibilities); index++ {
		ChosenServers = append(ChosenServers, ChosenPossibilities[index].Rv)
	}
	for index := 0; index < len(ChosenServers); index++ {
		for i := 0; i < len(ChosenServers); i++ {
			w.WriteString(strconv.Itoa(ChosenServers[i]) + " ")
		}

	}
}

// Sorting
func (slice PossibilitySlice) Len() int {
	return len(slice)
}

func (slice PossibilitySlice) Less(i, j int) bool {
	return (slice[i].Score) > (slice[j].Score)
}

func (slice PossibilitySlice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
