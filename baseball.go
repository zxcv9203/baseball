package main

import (
	"fmt"
	"math/rand"
	"time"
)

func initBaseBall(cnt int) (ball []int) {
	var exist [10]bool
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < cnt; {
		tmp := rand.Intn(10)
		if exist[tmp] || (i == 0 && cnt > 1 && tmp == 0) {
			continue
		}
		ball = append(ball, tmp)
		exist[tmp] = true
		i++
	}
	return
}

func contain(num int, baseBall []int) bool {
	for _, value := range baseBall {
		if num == value {
			return true
		}
	}
	return false
}

func main() {
	n := 0

	fmt.Print("플레이할 숫자를 선택해주세요 : ")
	fmt.Scan(&n)

	if n <= 0 || n > 9 {
		println("0 ~ 9의 숫자만 입력가능합니다.")
		return
	}
	baseBall := initBaseBall(n)
	for {
		ball := 0
		strike := 0
		str := ""
		fmt.Scan(&str)
		for i, value := range str {
			value -= '0'
			if value == rune(baseBall[i]) {
				strike++
			} else if contain(int(value), baseBall) {
				ball++
			}
		}
		if strike == n {
			fmt.Println("축하합니다! 정답을 맞추셨습니다.", str)
			break
		} else {
			fmt.Printf("%dB%dS\n", ball, strike)
		}
	}
}
