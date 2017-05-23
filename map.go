package main

import "fmt"

func main() {
	leagueTitles := make(map[string]int)
	leagueTitles["Sunderlang"] = 6
	leagueTitles["Newcastle"] = 4

	recentHead2Head := map[string]int{
		"Sunderland": 5,
		"Newcaastle": 0,
	}

	fmt.Printf("\nLeague titles: %v\nRecent head to heads: %v\n",
		leagueTitles, recentHead2Head)
}
