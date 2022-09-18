package main

import (
	"fmt"

	"github.com/pathak107/coderahi-learn/pkg/services"
)

func main() {
	fmt.Print("asdasd")
	db, err := services.NewDatabaseService()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db)
}
