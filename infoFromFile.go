package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Function to get info from input file
func getInfoFromFile(lines []string) {

	start := false // Start coordinates
	end := false   // End coordinates

	coordinates := make(map[string][]int)
	tunnels := make(map[string][]string)

	for i, line := range lines {
		if len(line) > 0 {
			firstLetter := string(line[0])
			xy := strings.Split(line, " ")     // All coordinates are separated with " "
			tunnel := strings.Split(line, "-") // All tunnels are separated with "-"
			if i == 0 {
				lemIn.NrOfAnts = toNbr(line) // Get number of ants from the first line
			} else if firstLetter != "L" && firstLetter != "#" { // A room will never start with the letter L or with # and must have no spaces
				if len(xy) == 3 { // Room name, x-coordinate, y-coordinate
					name := xy[0]
					x := toNbr(xy[1])
					y := toNbr(xy[2])
					ans := []int{x, y}

					if start {
						lemIn.StartPointName = name
					} else if end {
						lemIn.EndPointName = name
					}
					coordinates[name] = ans
					start = false
					end = false
				} else if len(tunnel) == 2 {
					lemIn.OriginalTunnels = append(lemIn.OriginalTunnels, line)
					tunnels[tunnel[0]] = appendIfNotContains(tunnels[tunnel[0]], tunnel[1])
					sort.Strings(tunnel)
					tunnels[tunnel[0]] = appendIfNotContains(tunnels[tunnel[0]], tunnel[1])
				}
			} else if line == "##start" { // If this, get the next line as StartPointName coordinates
				start = true
			} else if line == "##end" { // If this, get the next line as EndPointName coordinates
				end = true
			}
		}
	}
	lemIn.Coordinates = coordinates
	lemIn.Tunnels = tunnels
	checkInput()
}

// Check if input file has everything
func checkInput() {
	if lemIn.NrOfAnts > 0 {
		if len(lemIn.StartPointName) != 0 {
			if len(lemIn.EndPointName) != 0 {
				if len(lemIn.Coordinates) != 0 {
					if len(lemIn.Tunnels) == 0 {
						returnError("tunnels")
					}
				} else {
					returnError("rooms")
				}
			} else {
				returnError("start")
			}
		} else {
			returnError("start")
		}
	} else {
		returnError("NrOfAnts")
	}
}

// String to Int
func toNbr(a string) int {
	b, err := strconv.Atoi(a)
	check(err)
	return b
}

// If array doesn't already contain this str, then add it
func appendIfNotContains(s []string, str string) []string {
	if ifContains(s, str) {
		return s
	} else {
		return append(s, str)
	}
}

// If array already contains this string
func ifContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func printInput() {
	fmt.Println(lemIn.NrOfAnts)
	fmt.Println("##start")
	fmt.Println(lemIn.StartPointName, lemIn.Coordinates[lemIn.StartPointName][0], lemIn.Coordinates[lemIn.StartPointName][1])
	for name, coordinate := range lemIn.Coordinates {
		if name != lemIn.StartPointName && name != lemIn.EndPointName {
			fmt.Println(name, coordinate[0], coordinate[1])
		}
	}
	fmt.Println("##end")
	fmt.Println(lemIn.EndPointName, lemIn.Coordinates[lemIn.EndPointName][0], lemIn.Coordinates[lemIn.EndPointName][1])
	for i := 0; i < len(lemIn.OriginalTunnels); i++ {
		fmt.Println(lemIn.OriginalTunnels[i])
	}
	fmt.Println()
}
