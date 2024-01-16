package dictionary

import "errors"


type Dictionary map[string]string

var ErrNotFound = errors.New(" word not found")

func (d Dictionary)Search(word string) (string, error){
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary)Add(k ,v string){
	d[k] = v
}