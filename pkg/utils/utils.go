package utils

import "sync"

func CheckNameIsUnique(name string, clientsMap *sync.Map) bool {
	_, ok := clientsMap.Load(name)
	return !ok
}
