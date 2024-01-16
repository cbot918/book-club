package main

import "fmt"

func main(){
	
	data := map[string]string{}

	editedMap := editMap(data, "a","b")

	fmt.Println(data)
	fmt.Println(editedMap)
}

func editMap(m map[string]string, k,v string) map[string]string{

	nMap := map[string]string{}

	for k,v := range m {
		nMap[k] = v
	}

	nMap[k] = v

	return nMap
}