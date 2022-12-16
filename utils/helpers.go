package utils

import (
	"math"
	"math/rand"
	"time"
)

func RandomElements(len int) []int {
	var nums []int

	for i := 0; i < len; i++ {
		nums = append(nums, RandomIntBetween(0, 9))
	}

	return nums
}

func RandomIntBetween(first, second int) int {
	rand.Seed(time.Now().UnixNano())

	return int(math.Floor(rand.Float64()*float64(second-first+1)) + float64(first))
}
