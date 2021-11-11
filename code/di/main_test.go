package di

import (
	"bytes"
	"testing"
)

/*
    출력이 어떻게 되는지 그 동작을 알 필요는 없으니, interface를 구현하고
	fmt.Printf는 내부적으로 os.Stdout값을 넣은 Fprintf를 호출함
*/
func TestGreet(t *testing.T) {
    buffer := bytes.Buffer{}
    Greet(&buffer, "Chris")

    got := buffer.String()
    want := "Hello, Chris"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}

func TestMain(t *testing.T) {
    main()
}