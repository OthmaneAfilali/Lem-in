package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

var startRoom, endRoom *string
var antsNumber string
var visitedRooms []string
var roomsAndTunnels = make(map[string][]string)

func main() {
	readInput()
	fmt.Println("Ants number: ", antsNumber)
	fmt.Println("Start room: ", *startRoom)
	fmt.Println("End room: ", *endRoom)
	for room, tunnelTo := range roomsAndTunnels {
		fmt.Printf("room: %s, connected rooms: %s\n", room, tunnelTo)
	}
	var currentRoom = *startRoom
	var possiblePath []string
	foundPaths := searchForPath(currentRoom, visitedRooms, possiblePath)
	fmt.Println("Found path: ", foundPaths)
	moveAnts(foundPaths)
}

func readInput() {
	input, err := os.ReadFile("input.txt")
	errorCheck("Error reading the input file", err)
	data := strings.Split(string(input), "\n")
	antsNumber = data[0]
	for i := 0; i < len(data); i++ {
		splittedData := strings.Split(data[i], " ")
		if len(splittedData) == 3 {
			roomsAndTunnels[splittedData[0]] = []string{}
		}
		if data[i] == "##start" {
			splittedData = strings.Split(data[i+1], " ")
			startRoom = &splittedData[0]
		}
		if data[i] == "##end" && startRoom != nil {
			splittedData = strings.Split(data[i+1], " ")
			endRoom = &splittedData[0]
		}
		if endRoom != nil {
			splittedData = strings.Split(data[i], "-")
			if len(splittedData) == 2 {
				for i := range roomsAndTunnels {
					if i == splittedData[0] {
						roomsAndTunnels[i] = append(roomsAndTunnels[i], splittedData[1])
					}
					if i == splittedData[1] {
						roomsAndTunnels[i] = append(roomsAndTunnels[i], splittedData[0])
					}
				}
			}
		}
	}
	checkInput()
}

func checkInput() {
	if startRoom == nil || endRoom == nil || len(roomsAndTunnels) == 0 {
		fmt.Println("Input data is corrupted")
		os.Exit(1)
	}
	for rooms := range roomsAndTunnels {
		if rooms[0] == '#' || rooms[0] == 'L' {
			fmt.Println("Wrong naming of the rooms")
			os.Exit(1)
		}
	}
}

func searchForPath(currentRoom string, visitedRooms, possiblePath []string) [][]string {
	var succesfulPaths [][]string
	for i := 0; i < len(roomsAndTunnels[currentRoom]); i++ {
		if currentRoom == *endRoom {
			possiblePath = append(possiblePath, currentRoom)
			succesfulPaths = append(succesfulPaths, possiblePath)
		}
		nextRoom := roomsAndTunnels[currentRoom][i]
		if !isVisited(nextRoom, visitedRooms) {
			possiblePath = append(possiblePath, currentRoom)
			visitedRooms = append(visitedRooms, nextRoom)
			results := searchForPath(nextRoom, visitedRooms, possiblePath)
			succesfulPaths = append(succesfulPaths, results...)
			if len(visitedRooms) > 0 && len(possiblePath) > 0 {
				visitedRooms = visitedRooms[:len(visitedRooms)-1]
				possiblePath = possiblePath[:len(possiblePath)-1]
			}
		}
	}
	return isDublicate(succesfulPaths)
}

func isVisited(currentRoom string, visitedRooms []string) bool {
	for _, room := range visitedRooms {
		if room == currentRoom {
			return true
		}
	}
	return false
}

func isDublicate(succesfulPaths [][]string) [][]string {
	var uniquePaths [][]string
	for _, path := range succesfulPaths {
		keepTrack := make(map[string]bool)
		isDublicate := false
		for _, room := range path {
			if keepTrack[room] {
				isDublicate = true
				break
			}
			keepTrack[room] = true
		}
		if !isDublicate {
            uniquePaths = append(uniquePaths, path)
        }
	}
	return uniquePaths
}

func moveAnts(foundPaths [][]string) {
	ants, err := strconv.Atoi(antsNumber)
	errorCheck("Error converting ants number", err)
	for _, path := range foundPaths {
		for i := 0; i < ants; i++ {
			for rooms := range path {
				fmt.Printf("L%v-%v ", i+1, rooms)
			}
		}
	}
}

/*func isFound(possiblePath []string) bool {
	for _, room := range possiblePath {
		if room == *endRoom {
			return true
		}
	}
	return false
}*/

func errorCheck(msg string, err error) {
	if err != nil {
		log.Println(msg, err)
		os.Exit(1)
	}
}
