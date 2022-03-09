package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)
	fmt.Println(states)
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"
	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)

	delete(states, "OR")
	fmt.Println(states)

	states["NY"] = "New York"
	// unordered collection, so can't count on order
	fmt.Println(states)
	
	// to iterate over a map
	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	// can extract slice
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)

	// to get the output of map to be alphabetical
	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}