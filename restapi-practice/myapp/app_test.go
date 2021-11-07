package myapp

import "testing"

func TestIndex(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler()) // http mock server
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal("Hello World", string(data))
}