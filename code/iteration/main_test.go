package iteration

import "testing"

func TestRepeat(t *testing.T) {
    repeated := Repeat("a")
    expected := "aaaaa"
    if repeated != expected {
        t.Errorf("expected %q but got %q", expected, repeated)
    }
}

/* 코딩을 하다보면 내가 짠 함수의 성능이 좋은지 안좋은지 궁금할때가 많다. 
Golang에서는 testing 라이브러리를 기본적으로 제공하면서 유닛테스트를 지원하면서도 성능을 측정할 수 있는 벤치마크 테스트도 제공한다.

벤치마크를 수행하는 함수는 몇몇 규칙을 꼭 지켜야한다.

Test함수는 Benchmark로 시작하는 이름을 가진다.
Benchmark뒤에는 대문자가 옵니다. ex) BenchmarkSum
*testing.B 타입의 매개 변수를 받는다.

-benchmem 옵션을 사용하면 작업당 메모리 사용량과 작업당 할당량을 확인할 수 있다
*/

func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}

/* Result Check!

goos: darwin
goarch: amd64
pkg: github.com/gngsn/learning-go/iteration
cpu: Intel(R) Core(TM) i5-8279U CPU @ 2.40GHz
BenchmarkRepeat-8   	 7754646	       154.3 ns/op
PASS

*/