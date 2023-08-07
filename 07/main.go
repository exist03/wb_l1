package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type storage struct {
	mu    sync.Mutex
	myMap map[int]int
}

func newStorage() *storage {
	return &storage{myMap: make(map[int]int)}
}

func (s *storage) create() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.myMap[rand.Int()] = rand.Int()
}

func main() {
	s := newStorage()
	s.create()
	fmt.Println(s.myMap)
}
