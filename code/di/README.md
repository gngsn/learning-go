### Di

<br/>

<small>ย ๐ย <a href="https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection" alt="Dependency Injection">Dependency Injection</a></small>



<br/>์ค์ํ์์๋ ์ด๋ค ๋ด์ฉ์ ์ถ๋ ฅํ  ๋ `stdout`์ ์ฌ์ฉํจ.

์ค์  `fmt` package๋ฅผ ํ์ธํด๋ณด์ ~

<br/>

`fmt` package์ `Printf`๋ ๋ด๋ถ์ ์ผ๋ก `fmt` package์ `Fprintf`์ `os.Stdout`๋ฅผ ๋ด์ ํธ์ถํจ

``` go
// It returns the number of bytes written and any write error encountered.
func Printf(format string, a ...interface{}) (n int, err error) {
    return Fprintf(os.Stdout, format, a...)
}
```

<br/>

์ค,, ๊ทธ๋ผ `os.Stdout`์ ์ด๋ค ์ญํ ๋ก ์ฌ์ฉ๋๊ธธ๋ ์ฒซ ๋ฒ์งธ์ธ์๋ก ๋ฐ๋ ๊ฒ์ผ๊น,,

`Fprintf`์ ๊ตฌ์กฐ๋ ํ์ธํด๋ณด์. 


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

`io.Writer`์ ์ธ์๋ก ๋ฐ์ ๊ฒ์ ํ์ธํ  ์ ์๋ค. `io.Writer` ์ ์ ์๋ ๋ณด์ !

``` go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

๊ทธ๋ ๋ค๋ฉด `Write(p []byte) (n int, err error)` Method๋ฅผ ๊ตฌํํ ๊ฐ์ฒด๋ฅผ ๋ฐ์ ์ ์๋ ๊ฑด๋ฐ,

์ค์ ๋ก ๊ทธ๋ฐ์ง ํ์ธํด๋ณด์ ~ (์ด๊ฑด ๊ทธ๋ฅ ํผ์๋ง์ ๊ถ๊ธ์ฆ... ๋ฐ๋ก ๋ด๋ ค๊ฐ๋ ์๊ด ใดใด)

<br/>

``` go
// NewFile์ *File๋ฅผ ๋ฐํ
Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
```

`Stdout`์ `NewFile`๋ก ์์ฑ๋๋๋ฐ, ๋ฐํ๋๋ `*File` ํ์์.

<br/>

`go/src/os/types.go`๋ฅผ ๋ณด๋ฉด ์๋์ ๊ฐ์ด ์ ์๋์ด ์์.

``` go
// os/types.go
type File struct {
	*file // os specific
}
```

<br/>

๋ง์ง๋ง!  `go/src/os/file.go`๋ฅผ ๋ณด๋ฉด ์๋์ ๊ฐ์ด `Write` Method๊ฐ ์ ์๋์ด ์์.

``` go
func (f *File) Write(b []byte) (n int, err error) {
  ...
}
```

