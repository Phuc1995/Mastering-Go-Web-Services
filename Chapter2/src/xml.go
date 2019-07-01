package main

import (
	"encoding/xml"
	"fmt"
)

type User2 struct {
	Name string `xml:"name"`
	Email string `xml:"email"`
	ID int `xml:"id"`
}

func main()  {
	ourUser := User{}
	ourUser.Name = "Bill Smith"
	ourUser.Email = "bill.smith@example.com"
	ourUser.ID = 100
	output,_ := xml.Marshal(&ourUser)
	fmt.Println(string(output))
}
