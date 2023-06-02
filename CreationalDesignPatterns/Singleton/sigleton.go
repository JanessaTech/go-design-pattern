package singleton

import (
	"fmt"
	"sync"
)

type target struct{}

var instance target
var once sync.Once

func getIntstance() *target {
	once.Do(func() {
		instance = target{}
		fmt.Println("instance is created")
	})
	return &instance
}

// this pattern is used when we want to create one global instance only
func Main() {
	inst := getIntstance()
	inst = getIntstance()
	_ = inst
}
