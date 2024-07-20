package service

import (
	"errors"
	"math/rand"
)

func RandomElement() string {
	var elements = []string{"element1", "element2", "element3, element4"}
	randomElement := elements[rand.Intn(len(elements)-1)]
	return randomElement
}

func GetElementById(id int) (string, error) {
	var elements = []string{"element1", "element2", "element3", "element4"}
	if id < 0 || id >= len(elements) {
		return "", errors.New("element not found")
	}
	return elements[id], nil
}
