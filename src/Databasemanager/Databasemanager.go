// Databasemanager project Databasemanager.go
package Databasemanager

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"fmt"
	"strings"
)

var (
	session *mgo.Session
	server	string 	//Später mit configfile?
	dbname  string	//Später mit configfile?
	
	//Fehler die auftreten können
)

// wird vor dem jeweils aufrufendem Main ausgeführt
func init(){
	server 	= "localhost"
	dbname 	= "test"
	
	sess,err := mgo.Dial(server)
	//ersetzen durch fuf.CheckErr oder so...
	if err != nil{
		panic(err)
	}
	
	session = sess
	
}

//wird nur paket-intern genutzt daher nicht global
func getCollection(sess *mgo.Session, coll string) *mgo.Collection{
	return sess.DB(dbname).C(coll)	
}

func InsertRecord(record interface{}){
	sess := session.Clone()
	defer sess.Close()
	
	coll := getCollection(sess, GetType(record))
	err := coll.Insert(record)
	if err != nil{
		panic(err)
	}
}
func UpdateRecord(i interface{}, id bson.ObjectId){
	// hier brauche ich eine Id um den jeweiligen Datensatz anzusprechen
		// habe ich den datensatz muss ich eine bson map übergeben 
		// die aus keys des zu ändernden keys und values des neuen values bestehen
	//__________________Ich weiß nicht ob ich überhaupt ein Update benötigen werde_____________________
	sess := session.Clone()
	defer sess.Close()
	
	//coll := getCollection(sess, GetType(i))
	
	/* MÖGLICHKEITEN
	variante 1)err := coll.UpdateId(id, bson.M{"$set": bson.M{"name": "updated name"}})
	variante 2)change := bson.M{"$set": bson.M{"phone": "+86 99 8888 7777", "timestamp": time.Now()}}
				err = coll.Update(colQuerier, change)
	*/
}
func DeleteRecord(i interface{}, id bson.ObjectId){
	//fmt.Printf("User ist: %v\nid ist: %v\nder typ ist: %v",i,id,t)
	
	sess := session.Clone()
	defer sess.Close()
	
	coll := getCollection(sess, GetType(i))
	
	coll.RemoveId(id)
}

func GetRecordById(i interface{},id string){
	GetRecord(i,bson.M{"_id": bson.ObjectIdHex(id)},true)
}

func GetRecordByName(i interface{},name string){
	GetRecord(i,bson.M{"name": name},true)
}

func GetRecord(i interface{},q bson.M,once bool){
	//fmt.Printf("%v %v type: %s\n", i, q, GetType(i))
	sess := session.Clone()
	defer sess.Close()
	
	coll := getCollection(sess, GetType(i))
	
	query := coll.Find(q)
	
	if once{
		query.One(i)
	}else{
		query.All(i)
	}
}
func GetAllRecords(i interface{}){
	GetRecord(i,nil,false)
}
//	Gibt den Typ des jeweiligen Inputs zurück
//	Wird zb. verwendet um automatisch die Collection 
//	für user zu finden wenn InsertRecord ein User übergeben wird 
func GetType(i interface{})string{
	t := reflect.TypeOf(i)
	
	//wenn i ein Pointer ist
	if t.Kind() == reflect.Ptr{
		t = t.Elem()
		
		// Tritt ein wenn Pointer auf ein Slice verweist
		if t.Kind() == reflect.Slice{
			t = t.Elem()
		}
	}
	return strings.ToLower(t.Name())
}
func GetStructField(i interface{},fieldname string) (interface{}){
	val := reflect.ValueOf(i)
	if val.Kind() == reflect.Ptr{
		val = val.Elem()
	}
	fmt.Printf("eingabe: %v\nsuchfeld: %v\nval: %v\n",i,fieldname,val);
	f := val.FieldByName(fieldname)
	fmt.Println(f)
	return f
}