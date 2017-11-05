// An echo web application.
// https://ianmcloughlin.github.io :: 2017-10-05

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func count(w http.ResponseWriter, r *http.Request) {

	count := 0

	var cookie, err = r.Cookie("count")
	if err == nil {
		count, _ = strconv.Atoi(cookie.Value)

		//	fmt.Fprintf(w, "Cookie is: ", cookievalue)
	}
	count++

	countstr := strconv.Itoa(count)

	cookie := &http.Cookie{
		Name:  "count",
		Value: countstr,
	}

	http.SetCookie(w, cookie)
	fmt.Fprintf(w, "You have visited %d times,", count)

}

func main() {
	http.HandleFunc("/", count)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
