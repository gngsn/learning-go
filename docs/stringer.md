### Stringer
https://pkg.go.dev/fmt#Stringer

Stringer는 String Method를 구현한 어떤 값에서도 구현할 수 있음.

String Method는 문자열 값으로 받아들여지는 데이터나, 포맷되지 않은 형식의 데이터 등의 전달받은 값을 출력하는데 사용되곤 한다. 

> Stringer is implemented by any value that has a String method, which defines the “native” format for that value. The String method is used to print values passed as an operand to any format that accepts a string or to an unformatted printer such as Print.

``` go
package main

import "fmt"

type Bitcoin int

func (b Bitcoin) String() string {
    return fmt.Sprintf("%d BTC", b)
}

func main() {
    b := Bitcoin(10)
    fmt.Printf("%s", b)
}
```