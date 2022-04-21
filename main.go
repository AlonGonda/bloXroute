package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/xuri/excelize"
)

func main() {
	source := "S"
	destination := "T"
	f, err := excelize.OpenFile("bloXgraph.xlsx")
	if err != nil { // Error if the file wasn't opened correctly
		log.Fatal(err)
	}
	graph := newGraph() // Building the Graph
	rows, err := f.GetRows("bloXroute") // Parse the graph file (xlsx) as matrix. Prints error if exists
	if err != nil {
		log.Fatal(err)
	}
	prices := make(map[string]int) // HashMap to get the price of each edge by O(1)
	for i := 2; i < len(rows); i++ { // Runs all the rows and puts the data into the graph
		intVar, err := strconv.Atoi(rows[i][2]) // Convert string to int, due to Graph settings
		if err != nil {
			log.Fatal(err)
		}
		graph.addEdge(rows[i][0], rows[i][1], intVar)
		prices[rows[i][0] + " " + rows[i][1]] = intVar // Add each edge to the HashMap, later we will use it to get the price
	}
	_, route := graph.getPath(source, destination) // Get the route after Dijkstra
	value := prices[route[0] + " " + route[1]]
	concatenated := fmt.Sprintf("%s -- (%d) --> %s", route[0], value, route[1]) // concatenate the needs string for printing
	fmt.Print(concatenated) // Print Concatenated
	for i := 1; i < len(route) - 1; i++ {
		value := prices[route[i] + " " + route[i+1]]
		concatenated := fmt.Sprintf(" -- (%d) --> %s", value, route[i + 1]) // concatenate the needs string for printing
		fmt.Print(concatenated) // Print Concatenated
	}
}
