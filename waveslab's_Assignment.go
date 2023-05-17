package main

import (
	"fmt"
	"math"
)

func TimeTaken(t [][]int, n, k int) int { //The function returns an integer,
	//which is the minimum time it takes for all the nodes to receive the signal.

	d := make([]int, n+1) // This ensures that the source node has the shortest distance to itself.
	for i := range d {
		d[i] = math.MaxInt32
	}
	d[k] = 0

	for i := 0; i < n; i++ {
		updated := false
		for _, time := range t {
			u, v, w := time[0], time[1], time[2]
			if d[u]+w < d[v] {
				d[v] = d[u] + w
				updated = true //This step is responsible for finding the shortest paths in the graph.
			}
		}
		if !updated {
			break // If no distances were updated in the current iteration
			//It means that no further improvements can be made,and the loop breaks early to save unnecessary iterations.
		}
	}

	mD := 0
	for i := 1; i <= n; i++ {
		if d[i] == math.MaxInt32 {
			return -1 //If any node's distance is still infinity it means that the signal couldn't reach that node so we return -1
		}
		mD = max(mD, d[i])
	}

	return mD //we return the maximum distance as the minimum time it takes for all nodes to receive the signal.
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	t := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	n := 4
	k := 3
	result := TimeTaken(t, n, k)
	fmt.Println(result)
}
