// frequentlyUsedFunctions project frequentlyUsedFunctions.go
package frequentlyUsedFunctions
import(
	"net/http"
	"time"
	"html"
	"strings"
)
//Später könnten hier eventuell noch mehr sachen abgefangen werden
/*
	Eig sollte diese Funktion einen Fehler ausgeben und
	das Registrieren blockieren wenn die Vorgaben nicht 
	gegeben sind
*/
func SanitizeInput(input string) string{
	output := ""
	//Leerzeichen rausnehmen
	output = html.EscapeString(input)
	output = strings.Replace(output," ","",-1)
	
	return html.EscapeString(output)
}
func MakeCookie(w http.ResponseWriter, cname string, val string){
	c := http.Cookie{
		Name: cname,
		Value: val,
		Expires: time.Now().AddDate(0,0,1),
	}
	http.SetCookie(w,&c)
	
}
func DeleteCookie(w http.ResponseWriter, cookie *http.Cookie){
	c := http.Cookie{
		Name: cookie.Name,
		Value: "",
		MaxAge: -1,
		Expires: time.Unix(1,0),
	}
	
	http.SetCookie(w,&c)
}