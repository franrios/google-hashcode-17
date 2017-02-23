package main

/*
VIDEOS			V
ENDPOINTS		E
Cache servers	C
--
V video sizes

--

LD				Latency of serving a video request from the data center to this endpoint
K				number of cache servers that an endpoint is connected to

c				cahce server ID
Lc				latency to this endpoint from this cache server

Rv				ID of the requested video
Re				ID of the endpoint from which the requests are coming from
Rn				the number of requests
*/

/*
	1. V E R C X
	2. V1 V2...			// V Sizes of videos
	--------------
	LD  K				// HEADER: E ENDPOINTS				/E times\
	c	Lc				// K lines							\E times/
	--------------
	Rv Re Rn					// R Request descriptions


*/

// ---- INPUT DATA FROM FILE ----

// Counts:
var V int // Count of videos
var E int // Count of endpoints
var R int // Number of request descriptions
var C int // Count of cache servers
var X int // Capacity of each cache server in megabytes

var VideoSizes []int
var Endpoints []Endpoint
var Requests []Request

type Endpoint struct {
	LD           int           // Latency of serving a video request from the data center to this endpoint
	K            int           // number of cache servers that an endpoint is connected to
	CacheServers []CacheServer // K servers
}

type CacheServer struct {
	C  int // ID
	Lc int //latency to this endpoint from this cache server
}

type Request struct {
	Rv int // ID of the requested video
	Re int // ID of the endpoint from which the requests are coming from
	Rn int // the number of requests
}

type Possibility struct {
	Rv     int // ID of the requested video
	Re     int // ID of the endpoint from which the requests are coming from
	Rn     int // the number of requests
	Score  int //= Rn * VideoSizes[thisone]  == (the number of requests * size)
	Server *CacheServer
}

type EndpointSlice []Endpoint

var freeSpace []int

/*
func (a EndpointSlice)seek(i int){
	length := len(a)
	for index := 0; index < length; index++ {
		if i==a[i]. {

		}
	}
	return -1
}
*/
// MORE DATA
var Possibilities PossibilitySlice
var ChosenPossibilities PossibilitySlice

type PossibilitySlice []Possibility

var VideoReproductions []int
