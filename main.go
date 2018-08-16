package main

import (
	"github.com/o1egl/govatar"
	"fmt"
	"log"
)

func main(){
	var gender = govatar.FEMALE
	var username = "enrobbins0523"
	var message = createAvatar(gender, username)
	fmt.Println(message)
}

func createAvatar(gender govatar.Gender, username string) string{
	var err error
	var message string
	var path = fmt.Sprintf("C:/Users/All Users/%s.jpg", username)
	err = govatar.GenerateFileFromUsername(gender, username, path)
	if err == nil{
		message = fmt.Sprintf("%s's govatar has been created! Check C:/Users/All Users/%s.jpg!", username, username)
	}
	if err != nil {
		log.Fatal(err)
		message = err.Error()
	}
	return message
}