package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type user struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}


var person = []user{
	{ID: "1", Name: "Rakshya", Email: "rakshyak@gmail.com", Password: "123456"},
	{ID: "2", Name: "Rakshya", Email: "rakshyak@gmail.com", Password: "123456"},
	{ID: "3", Name: "Rakshya", Email: "rakshyak@gmail.com", Password: "123456"},
}


func main(){
	request := gin.Default()
	request.GET("/users", getUsers)
	request.Run("localhost:8000")
}

func getUsers(context *gin.Context){
	context.IndentedJSON(http.StatusOK, person)
}