package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type loginSchema struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Return `true`, nil if given user and password exists in database.
func loginUser(username string, password string) (bool, error) {
	return true, nil
}

// Use as a wrapper around the handler functions.
type rootHandler func(http.ResponseWriter, *http.Request) error

func loginHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return NewHTTPError(nil, 405, "Method not allowed.")
	}

	body, err := ioutil.ReadAll(r.Body) // Read request body.
	if err != nil {
		return fmt.Errorf("Request body read error : %v", err)
	}

	// Parse body as json.
	var schema loginSchema
	if err = json.Unmarshal(body, &schema); err != nil {
		return NewHTTPError(err, 400, "Bad request : invalid JSON.")
	}

	ok, err := loginUser(schema.Username, schema.Password)
	if err != nil {
		return fmt.Errorf("loginUser DB error : %v", err)
	}

	if !ok { // Authentication failed.
		return NewHTTPError(nil, 401, "Wrong password or username.")
	}
	w.WriteHeader(200) // Successfully authenticated. Return access token?
	return nil
}

// rootHandler implements http.Handler interface.
func (fn rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r) // Call handler function
	if err == nil {
		return
	}
	// This is where our error handling logic starts.
	log.Printf("An error accured: %v", err) // Log the error.

	clientError, ok := err.(ClientError) // Check if it is a ClientError.
	if !ok {
		// If the error is not ClientError, assume that it is ServerError.
		w.WriteHeader(500) // return 500 Internal Server Error.
		return
	}

	body, err := clientError.ResponseBody() // Try to get response body of ClientError.
	if err != nil {
		log.Printf("An error accured: %v", err)
		w.WriteHeader(500)
		return
	}
	status, headers := clientError.ResponseHeaders() // Get http status code and headers.
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(status)
	w.Write(body)
}

type ClientError interface {
	Error() string
	// ResponseBody returns response body.
	ResponseBody() ([]byte, error)
	// ResponseHeaders returns http status code and headers.
	ResponseHeaders() (int, map[string]string)
}

func NewHTTPError(err error, status int, detail string) error {
	return &HTTPError{
		Cause:  err,
		Detail: detail,
		Status: status,
	}
}

type HTTPError struct {
	Cause  error  `json:"-"`
	Detail string `json:"detail"`
	Status int    `json:"-"`
}

func (e *HTTPError) Error() string {
	if e.Cause == nil {
		return e.Detail
	}
	return e.Detail + " : " + e.Cause.Error()
}

// ResponseBody returns JSON response body.
func (e *HTTPError) ResponseBody() ([]byte, error) {
	body, err := json.Marshal(e)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing response body: %v", err)
	}
	return body, nil
}

// ResponseHeaders returns http status code and headers.
func (e *HTTPError) ResponseHeaders() (int, map[string]string) {
	return e.Status, map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
}

func main() {
	// Notice rootHandler.
	http.Handle("/login/", rootHandler(loginHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
