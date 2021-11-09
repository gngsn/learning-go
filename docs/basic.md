## 1. Variable & Constant

<br />

### 변수
``` go
var [변수명] [타입(생략가능/go가 타입추론)]
// Short Assignment Statement ( := ) 
[변수명] := [초기값]
```

- 재할당 가능.

- 변수 선언 후 사용하지 않으면 ERROR.

- 초기값을 지정하지 않으면, Go는 Zero Value를 할당. 

    - 숫자형: 0
    - bool: false 
    - string: "" (빈문자열)

복수 개의 동일한 타입의 변수를 한 번에 선언 가능

``` go
var i, j, k int

var i, j, k int = 1, 2, 3
```

<br/><br/>

### 상수

``` go
const [변수명] [타입(생략가능/go가 타입추론)] = 초기값
// Short Assignment Statement ( := ) 
[변수명] := [초기값]
```

여러 상수 한 번에 선언

``` go
const (
    Visa = "Visa"
    Master = "MasterCard"
    Amex = "American Express"
)
```

`iota` identifier: 상수값을 0부터 순차적으로 부여.

``` go
// iota가 지정된 Apple에는 0이 할당되고, 나머지 상수들을 순서대로 1씩 증가된 값을 부여
const (
    Apple = iota // 0
    Grape        // 1
    Orange       // 2
)
```



## 2. Type

<br/>

### 불린 타입 `bool`

<br/>

### 문자열 타입 `string`

- string은 한번 생성되면 수정될 수 없는 Immutable 타입

- 문자열 리터럴은 Back Quote(` `) 혹은 이중인용부호(" ")를 사용하여 표현할 수 있다.

- Back Quote ` `` `: Raw String Literal
 : 이 안에 있는 문자열은 별도로 해석되지 않고 Raw String 그대로의 값을 갖는다. 예를 들어, 문자열 안에 `\n` 이 있을 경우 이는 NewLine으로 해석되지 않는다. 또한, Back Quote은 복수 라인의 문자열을 표현할 때 자주 사용된다.

- 이중인용부호 `" "`: Interpreted String Literal
 : 복수 라인에 걸쳐 쓸 수 없으며, 인용부호 안의 Escape 문자열들은 특별한 의미로 해석된다. 예를 들어, 문자열 안에 `\n` 이 있을 경우 이는 NewLine으로 해석된다. 이중인용부호를 이용해 문자열을 여러 라인에 걸쳐 쓰기 위해서는 + 연산자를 이용해 결합하여 사용한다.


<br/>

### 정수형 타입
`int` `int8` `int16` `int32` `int64`
`uint` `uint8` `uint16` `uint32` `uint64` `uintptr`

<br/>

### Float 및 복소수 타입
`float32` `float64` `complex64` `complex128`

### 기타
- `byte`: uint8과 동일하며 바이트 코드에 사용

- `rune`: int32과 동일하며 유니코드 코드포인트에 사용한다

<br/>

### Type Conversion

타입(`T`)로의 변환을 `T(v)` 와 같이 표현

<aside>
Example 

정수 100을 float로 변경 -> float32(100)

문자열을 바이트배열로 변경 -> []byte("ABC")
</aside>

<br />

#### ⚠️ Go에서 타입간 변환은 명시적으로 지정해 주어야 함

<aside>
  Example
정수형 int에서 uint로 변환할 때, 암묵적(implicit) 변환이 일어나지 않으므로 uint(i) 처럼 반드시 변환을 지정해 주어야 한다. 
🚫 명시적 지정이 없이 변환이 일어나면 <b>런타임 에러</b> 🚫 

</aside>





``` go
func main() {
    var i int = 100
    var u uint = uint(i)
    var f float32 = float32(i)  
    println(f, u)
 
    str := "ABC"
    bytes := []byte(str)
    str2 := string(bytes)
    println(bytes, str2)
}
```

<br /><br />

## 3. Conditional Statement



### if statement

✔️ if 조건문은 아래 예제에서 보듯이 조건식을 괄호( )로 둘러 싸지 않아도 됨

⭐️ 반드시 조건 블럭 시작 브레이스({)를 if문과 같은 라인에 적어야 함. 다음 라인에 두게 되면 에러를 발생 🚫

✔️ if 문의 조건식은 반드시 <b>Boolean 식</b>으로 표현 (C/C++ 같은 다른 언어들이 조건식에 1, 0 과 같은 숫자를 쓸 수 있는 것과 대조적)




``` go
if k == 1 {  //같은 라인, 괄호 생략, 반드시 bool 조건
    println("One")
}
```



✔️ if 문에서 조건식을 사용하기 이전에 간단한 문장(<b>Optional Statement</b>)을 함께 실행할 수 있음 

<small>즉, 아래 예제처럼 val := i*2 라는 문장을 조건식 이전에 실행할 수 있는데, 여기서 주의할 점은 이때 정의된 변수 val는 if문 블럭 (혹은 if-else 블럭 scope) 안에서만 사용할 수 있다는 것이다. 이러한 Optional Statement 표현은 아래의 switch문, for문 등 Go의 여러 문법에서 사용할 수 있다.</small>

``` go
if val := i * 2; val < max {
    println(val)
}

// ERROR : Scope 벗어남
val++ 
```



### Switch statement

여러 값을 비교해야 하는 경우 혹은 다수의 조건식을 체크해야 하는 경우 switch 문을 사용한다. 다른 언어들과 비슷하게 switch 문 뒤에 하나의 변수(혹은 Expression)를 지정하고, case 문에 해당 변수가 가질 수 있는 값들을 지정하여, 각 경우에 다른 문장 블럭들을 실행할 수 있다. 복수개의 case 값들이 있을 경우는 아래 예제에서 보듯이 <b>case 3,4 처럼 콤마를 써서 나열</b>할 수 있다.



``` go
switch [expression] {
    case [case1]:
        name = "Paper Book"
    case [case2]:
        name = "eBook"
    case [case3], [case4]:
        name = "Blog"
    default:
        name = "Other"
}
```





#### Go의 Switch문 특이점

✔️ **expression이 없을 수 있음** 👉🏻  if...else if..문 단순화에 용이

다른 언어는 switch 키워드 뒤에 변수나 expression 반드시 두지만, Go는 이를 쓰지 않아도 된다. 이 경우 Go는 switch expression을 true로 생각하고 첫번째 case문으로 이동하여 검사한다

``` go
switch {
  ...
}
```



✔️  **case문의 복잡한 expression**

다른 언어의 case문은 일반적으로 리터럴 값만을 갖지만, Go는 case문에 복잡한 expression을 가질 수 있다



✔️  **No default fall through**

다른 언어의 case문은 break를 쓰지 않는 한 다음 case로 이동하지만, Go는 다음 case로 가지 않는다



✔️ **Type switch**

다른 언어의 switch는 일반적으로 변수의 값을 기준으로 case로 분기하지만, Go는 그 변수의 Type에 따라 case로 분기할 수 있다. Go의 또 다른 용법은 switch 변수의 타입을 검사하는 타입 switch가 있다. 

``` go
// 변수 v의 타입 검사
switch v.(type) {
case int:
    println("int")
case bool:
    println("bool")
case string:
    println("string")
default:
    println("unknown")
}
```



#### fallthrough 

Go의 Switch는 break를 걸지 않아도 해당 case만 실행함.

<small>(👉🏻 Go Compiler가 자동으로 break 문을 각 case문 블럭 마지막에 추가함)</small>



C or C# 처럼 case 사이를 연속해서 실행하고 싶다면 (break를 쓰지 않는 C/C#처럼) **fallthrough** 사용

``` go
package main

import "fmt"

func main() {
    check(2)  // "2 이하/3 이하/default 도달"을 모두 출력
}
 
func check(val int) {
    switch val {
    case 1:
        fmt.Println("1 이하")
        fallthrough
    case 2:
        fmt.Println("2 이하")
        fallthrough
    case 3:
        fmt.Println("3 이하")
        fallthrough
    default:
        fmt.Println("default 도달")
    }
}
```

<br /><br />

## 4. Iteration Statement

### For statement

✔️ Go는 반복문에 for 하나 밖에 없음. 👉🏻 while 없음

✔️ 다른 언어와 비슷하게 `for 초기값; 조건식; 증감 { ... }` 의 형식

⭐️ 다만, "초기값; 조건식; 증감"을 둘러싸는 괄호 ( )를 생략하는데, **괄호를 쓰면 에러**



✔️ **조건식만 쓰는 for 루프 (초기값, 증감 생략)**

``` go
for n < 100 {
	n *= 2
}
```



✔️ **무한 루프 (초기값, 조건식, 증감 생략)**

``` go
for {
	println("Infinite loop")
}
```



✔️ **for - range 문**

for range 문은 컬렉션으로 부터 한 요소(element)씩 가져와 차례로 for 블럭의 문장들을 실행. 다른 언어의 foreach와 비슷.



`for 인덱스,요소값 := range 컬렉션`

👉🏻 range 키워드 다음의 컬렉션으로부터 하나씩 요소를 리턴해서 그 요소의 **위치인덱스**와 **값**을 for 키워드 다음의 2개의 변수에 각각 할당.

``` go
names := []string{"홍길동", "이순신", "강감찬"}
 
for index, name := range names {
    println(index, name)
}
```



✔️ **break, continue, goto 문**



``` go
package main
func main() {
    var a = 1
    for a < 15 {
        if a == 5 {
            a += a
            continue // for루프 시작으로
        }
        a++
        if a > 10 {
            break  //루프 빠져나옴
        }
    }
    if a == 11 {
        goto END //goto 사용예
    }
    println(a)
 
END:
    println("End")
}
```

break문은 **break *label***과 같이 사용하여 지정된 레이블로 이동할 수도 있음. 



<br /><br />

## 5. Function



### Function

**Multiple Return Parameter**

Go는 여러 값을 리턴할 수 있음. 👉🏻 리턴 타입들을 괄호 ( ) 안에 적어 줌

예를 들어, 처음 리턴값이 int이고 두번째 리턴값이 string 인 경우 (int, string) 과 같이 적어 줌

``` go
func sum(nums ...int) (int, int) {
    s := 0      // 합계
    count := 0  // 요소 갯수
    for _, n := range nums {
        s += n
        count++
    }
    return count, s
}
```



**Named Return Parameter**

Go에서 Named Return Parameter들에 리턴값들을 할당하여 리턴할 수 있는데, 코드 가독성을 높이는 장점이 있음. 



- 함수 내에서 count, total에 결과값을 직접 할당
- 리턴되는 값이 있을 경우에는 빈 return 문을 반드시 써 주어야 함 (이를 생략하면 에러 발생).

``` go
package main
 
func main() {
    count, total := sum(1, 7, 3, 5, 9)
    println(count, total)   
}

func sum(nums ...int) (count int, total int) {
    for _, n := range nums {
        total += n
    }
    count = len(nums)
    return					// 안쓰면 ERROR!
}
```





**Variadic Function (가변인자함수)**

함수에 고정된 수의 파라미터 수가 아닌 **개수가 정해지지 않은 파라미터를 전달하고자 할 때** 가변 파라미터를 나타내는 `...` 을 사용

``` go
func say(msg ...string) {
    for _, s := range msg {
        println(s)
    }
}
```



### Anonymous Function

``` go
package main
 
func main() {
    sum := func(n ...int) int { //익명함수 정의
        s := 0
        for _, i := range n {
            s += i
        }
        return s
    }
 
    result := sum(1, 2, 3, 4, 5) //익명함수 호출
    println(result)
}
```



#### 일급함수

Go 프로그래밍 언어에서 함수는 일급함수로서 Go의 기본 타입과 동일하게 취급되며, 따라서 다른 함수의 파라미터로 전달하거나 다른 함수의 리턴값으로도 사용될 수 있다. 즉, 함수의 입력 파라미터나 리턴 파라미터로서 함수 자체가 사용될 수 있다. 함수를 다른 함수의 파라미터로 전달하기 위해서는 익명함수를 변수에 할당한 후 이 변수를 전달하는 방법과 직접 다른 함수 호출 파라미터에 함수를 적는 방법이 있다.

``` go
package main
 
func main() {
    //변수 add 에 익명함수 할당
    add := func(i int, j int) int {
        return i + j
    }
 
    // add 함수 전달
    r1 := calc(add, 10, 20)
    println(r1)
 
    // 직접 첫번째 파라미터에 익명함수를 정의함
    r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
    println(r2)
 
}
 
func calc(f func(int, int) int, a int, b int) int {
    result := f(a, b)
    return result
}
```





#### type문을 사용한 함수 원형 정의

type 문은 구조체(struct), 인터페이스 등 Custom Type(혹은 User Defined Type)을 정의하기 위해 사용된다. type 문은 또한 함수 원형을 정의하는데 사용.

``` go
// 원형 정의
type calculator func(int, int) int
 
// calculator 원형 사용
func calc(f calculator, a int, b int) int {
    result := f(a, b)
    return result
}
```

이렇게 함수의 원형을 정의하고 함수를 타 메서드에 전달하고 리턴받는 기능을 타 언어에서 흔히 **델리게이트(Delegate)**라 부른다. Go는 이러한 Delegate 기능을 제공하고 있다.



### 클로저 (Closure)

Closure : 함수 바깥에 있는 변수를 참조하는 함수값(function value)를 일컫는데, 이때의 함수는 바깥의 변수를 마치 함수 안으로 끌어들인 듯이 그 변수를 읽거나 쓸 수 있게 된다.

아래 예제에서 nextValue() 함수는 int를 리턴하는 익명함수(func() int)를 리턴하는 함수이다. Go 언어에서 함수는 일급함수로서 다른 함수로부터 리턴되는 리턴값으로 사용될 수 있다. 그런데 여기서 이 익명함수가 그 함수 바깥에 있는 변수 i 를 참조하고 있다. 익명함수 자체가 로컬 변수로 i 를 갖는 것이 아니기 때문에 (만약 그렇게 되면 함수 호출시 i는 항상 0으로 설정된다) 외부 변수 i 가 상태를 계속 유지하는 즉 값을 계속 하나씩 증가시키는 기능을 하게 된다.
예제에서 next := nextValue() 에서 Closure 함수를 next라는 변수에 할당한 후에, 계속 next()를 3번 호출하는데 이때마다 Clouse 함수내의 변수 i는 계속 증가된 값을 가지고 있게 된다. 이것은 마치 next 라는 함수값이 변수 i 를 내부에 유지하고 있는 모양새이다. 그러나 만약 anotherNext := nextValue()와 같이 새로운 Closure 함수값을 생성한다면, 변수 i는 초기 0을 갖게 되므로 다시 1부터 카운팅을 하게 된다.

``` go
package main
 
func nextValue() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
 
func main() {
    next := nextValue()
 
    println(next())  // 1
    println(next())  // 2
    println(next())  // 3
 
    anotherNext := nextValue()
    println(anotherNext()) // 1 다시 시작
    println(anotherNext()) // 2
}
```


<br /><br />

## 6. Arrays 

미리 말하자면 Go에서는 Arrays를 잘 사용하지 않음.

ㅎㅎㅎ위의 말이 의아하겠지만 **Arrays** 말고 **Slices** 사용을 Go에서도 권장함!

배열은 지정된 형식의 메모리를 만들 때 유용하며 할당을 피하는 데 도움이 될 수 있지만, 주로 Slices(아래 8번째 섹션)을 구성하는 Block으로 사용 됨. 



### Go Arrays VS C Array s

✔️ Go에서 배열은 **값** 👉🏻 하나의 Arrays를 다른 Arrays에 할당하면 **모든 요소가 복사**된다.

✔️ 특히, 배열을 함수에 전달하면 해당 함수에 대한 **포인터가 아니라 배열의 복사본**을 받게 되니 주의

✔️배열의 크기는 배열 타입의 일부 👉🏻 `[10]int`와 `[20]int` 유형은 서로 다름.



**Value**로 전달되는 건 유용할 수도 있지만 비용이 많이 들 수도 있음. 

만약 효율성을 높이기 위해 C의 배열처럼 사용하고 싶다면 포인터를 전달할 수 있음.  👉🏻 근데 비추. Slices 참고

```
func Sum(a *[3]float64) (sum float64) {
    for _, v := range *a {
        sum += v
    }
    return
}

array := [...]float64{7.0, 8.5, 9.1}
x := Sum(&array)  // Note the explicit address-of operator
```



## 7. Slices

슬라이스는 배열을 래핑하여 Data sequences에 보편적이고 강력하며 편리한 인터페이스를 제공. 

Go에서는 Array를 특정 경우(변환 행렬과 같이 규격이 필요한 경우...)를 제외하고 대부분 Slices를 사용.



슬라이스는 기본 배열에 대한 참조값을 보관.

한 슬라이스를 다른 슬라이스에 할당하는 경우 둘 다 동일한 배열을 나타낸다.



#### make로 Slices 만들기

``` go
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:] // len(b)=4, cap(b)=4
```



#### func append(s []T, vs ...T) []T

``` go
s = append(s, 3, 12)
s1 = append(s1, s2...)
```



#### len(slices) & cap(slices)

`len(s)` : slices 의 길이

`cap(s)` : slices 의 용량

용량은 아래에서 자세히 확인!



#### Slices의 내부 구조

배열의 첫 번째 요소를 가리키는 포인터 + slice의 길이 + slice의 용량

<img src="https://media.vlpt.us/post-images/kimmachinegun/57629370-cbb4-11e8-a4ed-41eb39296329/%EC%8A%AC%EB%9D%BC%EC%9D%B4%EC%8B%B11.png" alt="img" style="zoom:48%;" />



``` go 
s := make([]int, 0, 3)
for i := 0; i < 5; i++ {
    s = append(s, i)
    fmt.Printf("cap %v, len %v, %p\n", cap(s), len(s), s)
}

/* result
cap 3, len 1, 0x1040e130
cap 3, len 2, 0x1040e130
cap 3, len 3, 0x1040e130
cap 6, len 4, 0x10432220
cap 6, len 5, 0x10432220
*/
```



#### Two-dimensional slices





<br /><br />

## 8. Package
