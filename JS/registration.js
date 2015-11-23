/**
 * Hier werden alle Aufrufe getätigt  definition der jeweiligen Funktion ist weiter unten zu finden..
 */
window.onload = function(){
    registerHandler();
}

function registerHandler() {
    var regForm = document.getElementById("regForm");

    regForm.onsubmit = function(event){
        event.preventDefault();
        startRegistration();
    }

    //________________________

    var logForm = $("logForm");

    logForm.addEventListener("submit", function(event){
        event.preventDefault();
        find();
    });

    //________________________

    var getAllUserBtn = $("getAllUser");

    getAllUserBtn.addEventListener("click", function(){
        requestAllUser();
    });

    //________________________

    var deleteUserBtn = $("deleteUserBtn");
    deleteUserBtn.addEventListener("click",function(){
        deleteUser();
    });
}
function find(){
    var logUsername = $("log_username");

    var xhr = new XMLHttpRequest();
    var params = "username=" + logUsername.value;

    xhr.open("POST", "http://localhost:4242/getUser", true);
    xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhr.send(params);

    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var response = xhr.responseText;
            postResponse("System: " +response,"notification", "green");
        }
    }
}

function requestAllUser(){
    var xhr = new XMLHttpRequest();

    xhr.open("POST", "http://localhost:4242/getAllUser",true);
    xhr.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    xhr.send(null);

    xhr.onreadystatechange = function(){
        if(xhr.readyState === 4 && xhr.status === 200){
            var response = xhr.responseText;
            postResponse(response, "notification","green");
        }

    }
}
function deleteUser(){
    var xhr = new XMLHttpRequest();

    xhr.open( "POST","http://localhost:4242/deleteUser",true );
    xhr.withCredentials = true;
    xhr.send();

    xhr.onreadystatechange = function(){
        if(xhr.readyState === 4 && xhr.status === 200){
            var response = xhr.responseText;
            postResponse(response, "notification","green");
        }
    }
}
function makeCookie(){
    var xhr = XMLHttpRequest();

    xhr.withCredentials = true;
    xhr.open( "POST", "http://localhost:4242/makeCookie",true);
    xhr.send();
}
function startRegistration() {
    var regUsername = document.getElementById("username");
    var regPasswordOne = document.getElementById("passwordOne");
    var regPasswordTwo = document.getElementById("passwordTwo");


    //  Ist Passwort_1 und Passwort_2 gleich ? check

        //  Ja  --> test ob Eingaben Vorgaben genügen

            //  Ja  --> Ajax-Request Usermanagment check

            //  Nein    --> Fehlermeldung ausgeben

        //  Nein    --> Fehlermeldung ausgeben

    if(regPasswordOne.value === regPasswordTwo.value){
        if(validateInputs()) {
            var xhr = new XMLHttpRequest();
            var params = "username=" + regUsername.value + "&passwordOne=" + regPasswordOne.value;

            xhr.open("POST", "http://localhost:4242/addUser", true);
            xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
            xhr.send(params);

            xhr.onreadystatechange = function () {
                if (xhr.readyState === 4 && xhr.status === 200) {
                    var response = xhr.responseText;
                    postResponse("Herzlich willkommen an board " +response,"notification", "green");
                }
            }
        }
    }
    else {
        postResponse("Die eingegebenen Passwörter stimmen nicht überein","notification", "#f00");
    }
}
function validateInputs(){
    var unMin = 4;
    var unMax = 20;
    var pwMin = 6;
    var pwMax = 35;
    var name = $("username").value;
    var pw = $("passwordOne").value;
    if( name.length >= unMin && name.length <= unMax )
    {
        if( pw.length >= pwMin && pw.length <= pwMax ){
            return true;
        }
        else{
            postResponse("Ein Passwort sollte zwischen " + pwMin + " und " + pwMax + " Ziffern beinhalten", "notification","#f00");
        }
    }
    else
    {
        postResponse("Ein Benutzername sollte zwischen" + unMin + " und " + unMax + " Ziffern beinhalten", "notification","#f00");
    }
    return false;
}


    //Helper
function $(id)
{
    return document.getElementById(id);
}

function postResponse(msg, container, color){
    $(container).innerHTML = msg;
    $(container).style.color = color;
}