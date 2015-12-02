package main

import "io/ioutil"

// Error checking
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Get if 2 slices of people's name match, order doesn't matter
// Not by me
func matchingPersons(persons1 []string, persons2 []string) bool {
	diffStr := []string{}
	m := map[string]int{}

	for _, s1Val := range persons1 {
		m[s1Val] = 1
	}
	for _, s2Val := range persons2 {
		m[s2Val] = m[s2Val] + 1
	}

	for mKey, mVal := range m {
		if mVal == 1 {
			diffStr = append(diffStr, mKey)
		}
	}

	return len(diffStr) == 0
}

// Check if an stack contains a needle
func contains(stack []string, needle string) bool {
	for _, element := range stack {
		if element == needle {
			return true
		}
	}

	return false
}

// Get the content of a file
func GetFileContent(file string) string {
	reader, err := ioutil.ReadFile(file)
	check(err)

	return string(reader)
}

func DeleteElementFromSlice(elements []string, delete string) []string {
	newElements := make([]string, 0)
	for _, element := range elements {
		if element != delete {
			newElements = append(newElements, element)
		}
	}
	return newElements
}
