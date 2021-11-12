package myapp

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	t *testing.T -> Test Code Convention
*/
func TestIndexPathHandler(t *testing.T) {
	assert := assert.New(t)
	
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	/*
		if res.Code != http.StatusOK {
			t.Fatal("Fail : ", res.Code)
		}
		
		위의 코드를 아래로 대체
		assert.Equal(expected, actual) : res.Code가 StatusOK 값과 다르면 assert
	*/
	assert.Equal(http.StatusOK, res.Code)
}


func TestIndexPathHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/query", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)
	
	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello world", string(data))
}

func TestIndexPathHandler_WithName(t *testing.T) {
	assert := assert.New(t)
	name := "gngsn"

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/query?name="+name, nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)
	
	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello " + name, string(data))
}

func TestJsonPathHandler_WithJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/json", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)
	
	assert.Equal(http.StatusBadRequest, res.Code)
	// data, _ := ioutil.ReadAll(res.Body)
	// assert.Equal("EOF", string(data))
	// assert.Equal(http.StatusBadRequest, string(data))
}