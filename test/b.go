package main

import "strings"

func main() {
	path := strings.Split("/test/a/adsdds","")
	lastKey := path[len(path)-1]
	var backupmap = make(map[string]interface{})
	deepestMap := deepSearch(backupmap, path[0:len(path)-1])

	deepestMap[lastKey] = string("a")
}


func deepSearch(m map[string]interface{}, path []string) map[string]interface{} {
	for _, k := range path {
		m2, ok := m[k]
		if !ok {
			// intermediate key does not exist
			// => create it and continue from there
			m3 := make(map[string]interface{})
			m[k] = m3
			m = m3
			continue
		}
		m3, ok := m2.(map[string]interface{})
		if !ok {
			// intermediate key is a value
			// => replace with a new map
			m3 = make(map[string]interface{})
			m[k] = m3
		}
		// continue search from here
		m = m3
	}
	return m
}

