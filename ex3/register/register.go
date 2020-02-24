package register

import (
	"fmt"
	"net/http"
	"os/user"
	"strconv"
	"time"
)

// init
//
// Register the handlers when loading the module
func init() {

	http.HandleFunc("/", greeter)
	http.HandleFunc("/timestamp", timer)
	http.HandleFunc("/user", userName)
}

// greeter handler
func greeter(w http.ResponseWriter, r *http.Request) {

	// greet the user!
	fmt.Fprintln(w, "Welcome to our tiny experimental server!")
}

// show the current date and time
func timer(w http.ResponseWriter, r *http.Request) {

	// show the current date and time
	now := time.Now()
	fmt.Fprintln(w, now.String())
}

// show the current user name
func userName(w http.ResponseWriter, r *http.Request) {

	// show the current user name
	if username, err := user.Current(); err != nil {
		fmt.Fprint(w, "It was not possible to retrieve the username")
	} else {
		fmt.Fprint(w, username.Username)
	}
}

// serve requests once all handlers have been registered
func Serve(port int) {

	httpAddress := "localhost:" + strconv.Itoa(port)
	http.ListenAndServe(httpAddress, nil)
}
