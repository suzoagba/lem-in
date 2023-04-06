package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

var routes [][]string

// Find all possible path from start point
func findAllPath() {
	startPoint := lemIn.Tunnels[lemIn.StartPointName]
	for i := 0; i < len(startPoint); i++ {
		firstVisited := []string{lemIn.StartPointName}
		firstRoute := []string{lemIn.StartPointName}
		possibleRoutes(startPoint[i], firstVisited, firstRoute)
	}
	if len(routes) == 0 {
		returnError("any possible path")
	}
	removeCrossroads()
	routesByLength()
}

// Find all possible routes
func possibleRoutes(point string, visited []string, route []string) {
	if !ifContains(visited, point) {
		if point == lemIn.EndPointName {
			blackMagic(append(append(route, point), point))
		} else {
			visited = append(visited, point)
			for a := 0; a < len(lemIn.Tunnels[point]); a++ {
				possibleRoutes(lemIn.Tunnels[point][a], visited, append(route, point))
			}
		}
	}
}

// To get rid of the dead end errors
func blackMagic(r []string) {
	routes = append(routes, r[:len(r)-1])
}

// Find the best possible combination and remove any crossing paths
func removeCrossroads() {
	var stepCount []float64
	var routeCollections [][][]string
	for firstRouteNum, firstRoute := range routes {
		var routeCollection [][]string
		var visited []string
		roomCount := len(firstRoute)
		routeCount := 1
		routeCollection = append(routeCollection, firstRoute)
		for _, room := range firstRoute {
			if room != lemIn.StartPointName && room != lemIn.EndPointName {
				visited = append(visited, room)
			}
		}
		for routeNum, route := range routes {
			if routeNum != firstRouteNum {
				possible := true
				for _, room := range route {
					if possible {
						if room != lemIn.StartPointName && room != lemIn.EndPointName {
							for _, visitedRoom := range visited {
								if room == visitedRoom {
									possible = false
								}
							}
						}
					}
				}
				if possible {
					routeCount++
					routeCollection = append(routeCollection, route)
					for _, room := range route {
						if room != lemIn.StartPointName && room != lemIn.EndPointName {
							visited = append(visited, room)
							roomCount++
						}
					}
				}
			}
		}
		stepCount = append(stepCount, math.Round(float64((lemIn.NrOfAnts+roomCount)/routeCount-1)))
		routeCollections = append(routeCollections, routeCollection)
	}
	min := stepCount[0]
	minNum := 0
	for num, v := range stepCount {
		if v < min {
			min = v
			minNum = num
		}
	}
	routes = routeCollections[minNum]
}

// Group routes by length
func routesByLength() {
	routesMap := make(map[int][][]string)
	var length []int
	for a := 0; a < len(routes); a++ {
		routesMap[len(routes[a])] = append(routesMap[len(routes[a])], routes[a])
	}
	lemIn.RoutesByLength = routesMap
	for i := range lemIn.RoutesByLength {
		length = append(length, i)
	}
	sort.Ints(length)
	lemIn.LengthByLength = length
}

func sendAntsOnTheirWay() {
	var availableRouteLength []int
	var availableRoutes [][]string
	pathForAnts := make(map[int][]int)
	for _, a := range lemIn.RoutesByLength {
		for _, b := range a {
			availableRouteLength = append(availableRouteLength, len(b))
		}
	}
	sort.Ints(availableRouteLength)
	// Put length of routes to an array
	for a := 0; a < len(availableRouteLength); a++ {
		pathForAnts[a] = []int{}
		availableRoutes = append(availableRoutes, lemIn.RoutesByLength[availableRouteLength[a]]...)
	}

	// Find which way to send the ants according to path length and waiting line length
	var result [][]string
	for a := 1; a <= lemIn.NrOfAnts; a++ {
		smallestStack := smallestLine(availableRouteLength, pathForAnts, a)
		pathForAnts[smallestStack] = append(pathForAnts[smallestStack], a)
	}

	for n1, i := range pathForAnts {
		for n2, ant := range i {
			for n3, room := range availableRoutes[n1] {
				if n3 != 0 {
					position := n2 + n3 - 1
					newString := "L" + strconv.Itoa(ant) + "-" + room
					if len(result) > position {
						result[position] = append(result[position], newString)
					} else {
						result = append(result, []string{newString})
					}
				}
			}
		}
	}
	for _, a := range result {
		fmt.Println(strings.Join(a, " "))
	}
}

func smallestLine(availableRouteLength []int, pathForAnts map[int][]int, ant int) int {
	smallestNr := len(pathForAnts[0]) + availableRouteLength[0]
	smallest := 0
	for b := 0; b < len(availableRouteLength); b++ {
		if len(pathForAnts[b])+availableRouteLength[b] < smallestNr {
			smallestNr = len(pathForAnts[smallest])
			smallest = b
		}
	}
	return smallest
}
