package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func SumAtoB(a, b int) {
	sum := 0
  	for i := a; i <= b; i++ {
  		 sum += i
  	}
  	fmt.Printf("%d부터 %d까지 합계는 %d입니다\n", a, b, sum)
  	wg.Done()			// 10번의 Done을 실행하면 Wait()가 실행이 됨
}

func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go SumAtoB(1, 1000000000)
	}

	wg.Wait()
	fmt.Println("모든 계산이 완료되었습니다.")
}