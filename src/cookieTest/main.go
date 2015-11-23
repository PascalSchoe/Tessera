// cookieTest project main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin","http://localhost")
	c, err := r.Cookie("timestamp")
	cookie := &http.Cookie{
		Name:  "timestamp",
		Value: strconv.FormatInt(time.Now().Unix(), 10),
	}
	w.Header().Add("Set-Cookie", cookie.String())
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Hello, you're here for the first time")
	} else {
		timeint, _ := strconv.ParseInt(c.Value, 10, 0)
		fmt.Fprintf(w, "Hi, your last visit was at "+time.Unix(timeint, 0).Format("15:04:05"))
	}
}

func main() {
	http.HandleFunc("/deleteUser", rootHandler)
	log.Fatal(http.ListenAndServe(":4242", nil))
}