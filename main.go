package main

import (
	"fmt"
	"os"
	"strings"
)

var lemIn = &lem //input file info

var lem struct {
	NrOfAnts        int
	StartPointName  string
	EndPointName    string
	Coordinates     map[string][]int
	OriginalTunnels []string
	Tunnels         map[string][]string
	RoutesByLength  map[int][][]string
	LengthByLength  []int
}

func main() {
	//startTime := time.Now()//start stopwatch
	args := os.Args[1:]
	if len(args) == 1 {
		data, err := os.ReadFile(args[0])
		check(err)

		getInfoFromFile(strings.Split(string(data), "\n"))
		findAllPath()
		printInput()
		sendAntsOnTheirWay()

		//fmt.Println("\nThis program took", time.Since(startTime), "to run.")//print elapsed time
	}
}

// Function to deal with errors
func check(e error) {
	if e != nil {
		returnError("anything")
	}
}

func returnError(a string) {
	fmt.Println("ERROR: Unable to get " + a + " from input file! Check if it exists.")
	os.Exit(0)
}
