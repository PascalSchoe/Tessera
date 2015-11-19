// usermanagment project main.go
package main

import (
	"fmt"
	"net/http"
)

func addUser(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	
	fmt.Println(r.Form)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, r.FormValue("username"))
	
}

func main() {
	
	http.HandleFunc("/addUser",addUser)
	http.ListenAndServe(":4242",nil)
}
