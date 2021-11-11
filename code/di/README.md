### Di

<br/>

<small> 🔗 <a href="https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection" alt="Dependency Injection">Dependency Injection</a></small>



<br/>실생활에서는 어떤 내용을 출력할 때 `stdout`을 사용함.

실제 `fmt` package를 확인해보자 ~

<br/>

`fmt` package의 `Printf`는 내부적으로 `fmt` package의 `Fprintf`에 `os.Stdout`를 담아 호출함

``` go
// It returns the number of bytes written and any write error encountered.
func Printf(format string, a ...interface{}) (n int, err error) {
    return Fprintf(os.Stdout, format, a...)
}
```

<br/>

오,, 그럼 `os.Stdout`은 어떤 역할로 사용되길래 첫 번째인자로 받는 것일까,,

`Fprintf`의 구조도 확인해보자. 


``` go
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
    p := newPrinter()
    p.doPrintf(format, a)
    n, err = w.Write(p.buf)
    p.free()
    return
}
```

<br/>

`io.Writer`의 인자로 받은 것을 확인할 수 있다. `io.Writer` 의 정의도 보자 !

``` go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

그렇다면 `Write(p []byte) (n int, err error)` Method를 구현한 개체를 받을 수 있는 건데,

실제로 그런지 확인해보자 ~ (이건 그냥 혼자만의 궁금증... 바로 내려가도 상관 ㄴㄴ)

<br/>

``` go
// NewFile은 *File를 반환
Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
```

`Stdout`은 `NewFile`로 생성되는데, 반환되는 `*File` 타입임.

<br/>

`go/src/os/types.go`를 보면 아래와 같이 정의되어 있음.

``` go
// os/types.go
type File struct {
	*file // os specific
}
```

<br/>

마지막!  `go/src/os/file.go`를 보면 아래와 같이 `Write` Method가 정의되어 있음.

``` go
func (f *File) Write(b []byte) (n int, err error) {
  ...
}
```

