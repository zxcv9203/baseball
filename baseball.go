package main

import (
	"fmt"
	"math/rand"
	"time"
)

//숫자야구의 답이 들어가는 타입
type Baseball []int

// 숫자야구에 초기값을 설정하는 함수
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

// 현재 예측한 수의 자리수의 수와 일치하는 숫자야구의 수가 있는지 확인하는 함수
func (b *Baseball) contain(num int) bool {
	for _, value := range *b {
		if num == value {
			return true
		}
	}
	return false
}

// 숫자야구에서 현재 자리수가 스트라이크인지 볼인지 아웃인지 확인하는 함수
func (b Baseball) callJudge(strike, ball chan int, num int, idx int) {
	if b[idx] == num {
		strike <- 1
		ball <- 0
		return
	} else if b.contain(num) {
		ball <- 1
		strike <- 0
		return
	}
	strike <- 0
	ball <- 0
}

func main() {
	var baseball Baseball
	n := 0
	ball := make(chan int, 1)
	strike := make(chan int, 1)
	fmt.Print("플레이할 숫자를 선택해주세요 : ")
	fmt.Scan(&n)

	if n <= 0 || n > 9 {
		println("1 ~ 9의 숫자만 입력가능합니다.")
		return
	}
	baseball.init(n)
	for {
		s := 0
		b := 0
		str := ""
		fmt.Print("답을 맞춰주세요 : ")
		fmt.Scan(&str)
		startTime := time.Now()
		if len(str) != n {
			fmt.Println("잘못된 값을 입력하셨습니다.")
			continue
		}
		for i, value := range str {
			go baseball.callJudge(strike, ball, int(value-'0'), i)
			s += <-strike
			b += <-ball
		}
		fmt.Println(time.Since(startTime))
		if s == n {
			fmt.Println("축하합니다! 정답을 맞추셨습니다.", str)
			break
		} else {
			fmt.Printf("%dB%dS\n", b, s)
		}
	}
}

/* 왜 루틴으로 나눴는데 루틴으로 안나눈게 훨씬 빠르지?
테스트가 잘못된건가?
아니면 구현 자체가 이상해서 문제가 생긴건가?
잘모르겠다.. 일단 루틴은 그대로 두고
웹서버 구동 함수 나누기로 넘어가자
*/
