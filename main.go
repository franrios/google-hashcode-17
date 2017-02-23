package main

func main() {
	readFile("kittens.in")

}

// Heuristic parameter for a first viable solution
func ponderate() {
	Ponderations = make([]int, V)
	VideoReproductions = make([]int, V)
	for i := 0; i < V; i++ {
		Ponderations[i] = VideoSizes[i] * VideoReproductions[i]
	}
	// Sort
}

//start := time.Now()
//elapsed := time.Since(start)
/*for x := 0; x < c; x++ {
	for y := 0; y < r; y++ {

	}
}*/
