// serializing_bson project main.go
package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	type Person struct{
		name 	string
		age		int
	}
	
	bob := Person{
		name: 	"Bob",
		age:	23,	
	}
	data,err := bson.Marshal(bob)
	if err != nil{
		fmt.Printf("Errors occured -- %v\n",err)
	}
	fmt.Printf("%v\n",bob)
	fmt.Printf("%v\n",data)
	
	var p Person
	err = bson.Unmarshal(data,&p)
	if err != nil{
		fmt.Printf("Errors occured -- %v", err)
	}
	fmt.Printf("Unserialized...<%v>\n",p)
}
