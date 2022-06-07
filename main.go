package main

// define packages
import (
	"fmt"
	"net/http"
)

// define variables
var (
	admin_u    = "admin"
	admin_p    = "password123!"
	port       = ":9000"
	handlepath = "/cisco"
)

// ^ base64 value: YWRtaW46cGFzc3dvcmQxMjMh

// define main:

func main() {

	//Define api handler
	apihandler := http.HandlerFunc(requestHandler)

	//Register api handler in the DefaultServeMux.
	http.Handle(handlepath, apihandler)

	//Start HTTP server with a given address and handler
	http.ListenAndServe(port, nil)
}

//define request handler
// ResponseWriter interface is used by an HTTP handler to construct an HTTP response.
// Request represents an HTTP request received by a server or to be sent by a client.
func requestHandler(w http.ResponseWriter, r *http.Request) {

	//BasicAuth returns the username and password provided in the request's Authorization header
	u, p, ok := r.BasicAuth()
	// if unable to parse auth header, return 401
	if !ok {
		fmt.Println("Login Failed; Error parsing auth header")
		w.WriteHeader(401)
		return
	}
	// if unable to validate username, return 401
	if u != admin_u {
		fmt.Println("Login Failed; cannot find username:", u)
		w.WriteHeader(401)
		return
	}
	// if unable to validate password, return 401
	if p != admin_p {

		fmt.Println("Login Failed; wrong password, user:", u)
		w.WriteHeader(401)
		return
	}
	// able to parse auth header and validate username, password and return 200
	fmt.Println("Login Successful; username", u)
	w.WriteHeader(200)
	return
}
