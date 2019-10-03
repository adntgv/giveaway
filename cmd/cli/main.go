package main

import (
	"fmt"
	"log"
	"os"

	ga "github.com/adntgv/giveaway"
)

func main() {
	login := os.Getenv("INSTAGRAM_USERNAME")
	password := os.Getenv("INSTAGRAM_PASSWORD")
	shortcode := "B3Ew-WRiZsj"
	app, err := ga.DefaultApp(login, password, shortcode)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(app.GetCommenters())

	return
}
