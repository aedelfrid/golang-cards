package main

type playset struct {
	name string
	handSize int
	cardValues []int
}

func gameImports(game string) playset{
	// in the future import games from json file?

	if game == "Texas Hold'em" {
		return playset{
			"Texas Hold'em",
			5,
			[]int{13, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		}
	}

	return playset{}
}