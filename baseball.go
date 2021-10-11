package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Baseball []int

var baseball Baseball

func (b *Baseball) init(cnt int) {
	var exist [10]bool
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < cnt; {
		tmp := rand.Intn(10)
		if exist[tmp] || (i == 0 && cnt > 1 && tmp == 0) {
			continue
		}
		*b = append(*b, tmp)
		exist[tmp] = true
		i++
	}
}

func contain(num int, baseBall []int) bool {
	for _, value := range baseBall {
		if num == value {
			return true
		}
	}
	return false
}

func call_judge(strike, ball chan int, num int) {

	for _, value := range baseball {
		if value == num {
			strike <- 1
		} else if contain(num, baseball) {
			ball <- 1
		}
	}
	strike <- 0
	ball <- 0
}

func main() {
	n := 0
	ball := make(chan int)
	strike := make(chan int)
	fmt.Print("플레이할 숫자를 선택해주세요 : ")
	fmt.Scan(&n)

	if n <= 0 || n > 9 {
		println("0 ~ 9의 숫자만 입력가능합니다.")
		return
	}
	baseball.init(n)
	for {
		s := 0
		b := 0
		str := ""
		fmt.Print("답을 맞춰주세요 : ")
		fmt.Scan(&str)
		if len(str) != n {
			fmt.Println("잘못된 값을 입력하셨습니다.")
			continue
		}
		fmt.Println(baseball)
		for _, value := range str {
			go call_judge(strike, ball, int(value-'0'))
			s += <-strike
			b += <-ball
		}
		if s == n {
			fmt.Println("축하합니다! 정답을 맞추셨습니다.", str)
			break
		} else {
			fmt.Printf("%dB%dS\n", b, s)
		}
	}
}
