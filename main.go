package main

import (
	"log"
	"sort"
	"time"
)

func main() {
	readFile("kittens.in")
	ponderate()
	assign()

}

/*
func fillVideoReproductions(){
VideoReproductions = make([]int, V)
for i := 0; i < V; i++ {
	for n := 0; i < R; i++ {
		Requests[n].Re*Requests[n].Rv
	}
		VideoReproductions[i]=count
	}
}*/

// Heuristic parameter for a first viable solution
func ponderate() {
	Possibilities = make([]Possibility, R)
	for i := 0; i < R; i++ {
		Possibilities[i].Re = Requests[i].Re
		Possibilities[i].Rv = Requests[i].Rv
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
		buffer := Possibility{Possibilities[i].Re, Possibilities[i].Rv, Possibilities[i].Score, nil}
		// Seek the perfect option among possible CacheServers for this Endpoint. nil means "Data Center""
		// Iterate in Endpoints[Possibilities[i].Re].CacheServers
		for n := 0; n < len(Endpoints[Possibilities[i].Re].CacheServers); n++ {
			if true {
				buffer.Server = &Endpoints[Possibilities[i].Re].CacheServers[n]
			}
			ChosenPossibilities = append(ChosenPossibilities, buffer)
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

//start := time.Now()
//elapsed := time.Since(start)
/*for x := 0; x < c; x++ {
	for y := 0; y < r; y++ {

	}
}*/
