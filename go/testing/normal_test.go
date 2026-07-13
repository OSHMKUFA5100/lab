package testing

import (
	"sort"
	"testing"
	"time"
)

func TestNormal(t *testing.T) {
	numbers := map[int][]int{
		1: {8, 2, 4, 4},
		2: {0, 5, 3},
		3: {1},
	}
	for _, numbers := range numbers {
		middle := GetMiddNumber(numbers)
		t.Log(middle)
	}

}

func GetMiddNumber(numbers []int) float64 {
	sort.Ints(numbers)
	length := len(numbers)
	isOddNUmber := length%2 == 1
	if !isOddNUmber {
		return float64(numbers[length/2-1]+numbers[length/2]) / 2.0
	}
	return float64(numbers[length/2])
}

func TestChannel(t *testing.T) {
	var ball = make(chan string)
	kickBall := func(playerName string) {
		for {
			t.Log(<-ball, "传球", "\n")
			time.Sleep(time.Second)
			ball <- playerName
		}
	}
	go kickBall("张三")
	go kickBall("李四")
	go kickBall("王二麻子")
	go kickBall("刘大")
	ball <- "裁判"    // 开球
	var c chan bool // 一个零值nil通道
	<-c             // 永久阻塞在此
}
