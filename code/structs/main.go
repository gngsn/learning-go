package structs

import "math"

/*
func Perimeter(width float64, height float64) float64 {
    return 2 * (width + height)
}

func Area(width float64, height float64) float64 {
    return width * height
}

-> 모두 도형과 관련된 함수
Shape - Rectangle, Circle, Triangle 구조체를 만들어서 Method로 구현

Wait, What..?
흔히 Interface이라고 생각하는 구조 -> Go에서는 Interface를 내재적으로 가지고 있음.
컴파일을 할 때 '전달되는 개체'와 '요구되는 interface'의 유형이 일치하면 성공적으로 실행.
*/

type Shape interface {
    Area() float64
}

type Rectangle struct {
    Width float64
    Height float64
}

type Circle struct {
    Radius float64
}

type Triangle struct {
    Base   float64
    Height float64
}

/*
	다른 언어는 아래와 같이 사용할 수 있지만 Go는 'Area redeclared in this block' Error를 띄움
	func Area(circle Circle) float64 { ... }
	func Area(rectangle Rectangle) float64 { ... }

	위와 같이 구현하고 싶으면 아래의 2가지 방법이 있음
	1. 다른 Package에 정의
	2. methods를 정의
	   -> func (receiverName ReceiverType) MethodName(args)

	receiverName은 Go의 Convention에 따라 ReceiverType의 첫 번째 문자의 소문자를 사용
*/

func (r Rectangle) Area() float64  {
    return r.Width * r.Height
}

func (c Circle) Area() float64  {
    return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
    return (t.Base * t.Height) * 0.5
}