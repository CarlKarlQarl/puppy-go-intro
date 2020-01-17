package main

import "github.com/gin-gonic/gin"

// structs are objects/classes for Go
// H below is making a struct

func main() {
	// capital D means it's exported (public method)
	// := is the go funky syntx for creating a new variable
	// but only inside a function
	// go likes single letter variables
	router := gin.Default()
	router.GET("/puppy", handlePuppy)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
