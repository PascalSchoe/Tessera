// usermanagment project main.go
package main

import (
	"fmt"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	dbm "Databasemanager"
	"time"
)
var(
	ActiveUser User
)

type User struct{
	Id bson.ObjectId `bson:"_id,omitempty"`
	Name string
	Password string
}
type Image struct{
	Name string
	Pixels int
	Photografer string
}

func addUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")

	U := User{}
	
	dbm.GetRecordByName(&U,r.FormValue("username"))
	// Es gibt keinen User mit diesem Namen bisher, Registrierung kann weitergehen
	if U == (User{}){
		/*U = User{
			bson.NewObjectId(),
			r.FormValue("username"),
			r.FormValue("passwordOne"),
		}
		*/
		U.Id = bson.NewObjectId()
		U.Name = r.FormValue("username")
		U.Password = r.FormValue("passwordOne")
		
		dbm.InsertRecord(&U)
		// Hier  könnte noch ein redirect auf kanal.html erfolgen
		// Außerdem könnte hier schon der Cookie für die Id gesetzt werden 
	}else{
		fmt.Fprintf(w,"Leider gibt es schon einen Nutzer mit dem Namen: %s, wähle bitte einen anderen Namen...", r.FormValue("username"))
	}
}
func getAllUser(w http.ResponseWriter, r *http.Request){
	results := make([]User,0)	
	w.Header().Set("Access-Control-Allow-Origin","*")
	dbm.GetAllRecords(&results)
	for _, user := range results{
		fmt.Fprintf(w,"%v\n%v\n%v \n\n",user.Id,user.Name,user.Password)
	}	
}

func getUser(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Access-Control-Allow-Origin","http://localhost")
	w.Header().Set("Access-Control-Allow-Methods","GET,OPTIONS")
	w.Header().Set("Access-Control-Allow-Credentials","true")
	
	ActiveUser := User{}
	dbm.GetRecordByName(&ActiveUser,r.FormValue("username"))
	if ActiveUser != (User{}){
		cookie := http.Cookie{Name:"userId",Value: ActiveUser.Id.Hex(),Expires: time.Now().AddDate(0,0,1)}
		http.SetCookie(w,&cookie)
	}else{
		fmt.Fprint(w,"No such User found")
	}
}
func setcookie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin","http://localhost")
	w.Header().Set("Access-Control-Allow-Methods","GET,OPTIONS")
	w.Header().Set("Access-Control-Allow-Credentials","true")
	
	cookie := http.Cookie{Name: "test2", Value: "nech", Expires: time.Now().AddDate(0,0,1)}
	http.SetCookie(w, &cookie)
}
func readcookie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin","http://localhost")
	w.Header().Set("Access-Control-Allow-Credentials","true")
	
	c, err := req.Cookie("ersterCookie")
	fmt.Fprintf(w, "Keks: %v %v\n", c, err)
	fmt.Fprintf(w, "all:\n")
	for _, c := range req.Cookies() {
		fmt.Fprintf(w, "\t%s\n", c)
	}
}
func editUser(w http.ResponseWriter, r *http.Request){
	/*w.Header().Set("Access-Control-Allow-Origin","http://localhost")
	w.Header().Set("Access-Control-Allow-Credentials","true")
	User := User{}
	c,err := r.Cookie("userId")
	cval := c.Value
	uId := bson.ObjectIdHex(cval)
	if err !=nil{
		panic(err)
	}
	*/
}
func deleteUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin","http://localhost")
	w.Header().Set("Access-Control-Allow-Credentials","true")
	User := User{}
	c,err := r.Cookie("userId")
	cval := c.Value
	uId := bson.ObjectIdHex(cval)
	if err !=nil{
		panic(err)
	}
	
	dbm.DeleteRecord(&User, uId)
	
	// löschen des Cookie!
	DeleteCookie(w,c)
	fmt.Fprintf(w,"User wurde gelöscht!")
}

func DeleteCookie(w http.ResponseWriter, cookie *http.Cookie){
	c := &http.Cookie{
		Name: cookie.Name,
		Value: "",
		MaxAge: -1,
		Expires: time.Unix(1,0),
	}
	
	http.SetCookie(w,c)
}

func main() {
	http.HandleFunc("/makeCookie", setcookie)
	http.HandleFunc("/giveCookie", readcookie)
	http.HandleFunc("/addUser",addUser)
	http.HandleFunc("/getUser",getUser)
	http.HandleFunc("/getAllUser", getAllUser)
	http.HandleFunc("/deleteUser",deleteUser)
	http.HandleFunc("/edit",editUser)
	
	http.ListenAndServe(":4242",nil)
}
