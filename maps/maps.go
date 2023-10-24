package main

import "fmt"

func main() {
	items := getItems()
	fmt.Printf("The duplicate items are %v", findDuplicates(items))
}

func findDuplicates(items []string) string {
	duplicateItems := make(map[string]bool)
	for i := 0; i < len(items); i++ {
		if !duplicateItems[items[i]] {
			duplicateItems[items[i]] = true
		} else {
			return items[i]
		}
	}
	return ("")
}

func getItems() []string {
	items := []string{
		"banana",
		"apple",
		"grape",
		"helicopter",
		"plant",
		"plant",
		"avocado",
	}
	return items
}
