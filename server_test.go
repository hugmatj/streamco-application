package main

import(
  //"github.com/go-martini/martini"
  "os"
  "bufio"
  "testing"
  "reflect"
  "net/http"
  "net/http/httptest"
  "encoding/json"
)

func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func TestDrmHandler(t * testing.T) {
  m := NewMartiniServer()

  // open input file
  file, err := os.Open("spec/request.json")
  if err != nil { panic(err) }

  // close file on exit and check for its returned error
  defer func() {
      if err := file.Close(); err != nil {
          panic(err)
      }
  }()

  // make a read buffer
  reader := bufio.NewReader(file)

  res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", reader)

  m.ServeHTTP(res, req)

  expect(t, res.Code, 200)
  expect(t, res.Header().Get("Content-Type"), "application/json; charset=UTF-8")

  result := make(map[string] []DrmResponse)
  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil { panic(err) }

  expect(t, len(result["response"]), 7)
}

func TestInvalidJsonParams(t * testing.T) {
  m := NewMartiniServer()

  // open input file
  file, err := os.Open("spec/invalid.json")
  if err != nil { panic(err) }

  // close file on exit and check for its returned error
  defer func() {
      if err := file.Close(); err != nil {
          panic(err)
      }
  }()

  // make a read buffer
  reader := bufio.NewReader(file)

  res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", reader)
  req.Header.Add("Content-Type", "application/json")

  m.ServeHTTP(res, req)

  expect(t, res.Code, 400)
  expect(t, res.Header().Get("Content-Type"), "application/json; charset=UTF-8")

  result := make(map[string] string)
  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil { panic(err) }

  expect(t, result["error"], "Could not decode request")
}
