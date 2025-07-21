package main

import (
	"fmt"
	"strings"
	"encoding/csv"
	"os"
	"log"
)
var emails []string = []string{
	"user1@domain1.com",
	"user2@domain1.com",
	"user3@domain1.com",
	"user4@domain1.com"}

func createAlias(email string) string{
	username := strings.SplitAfter(email,"@")
//	newAlias := username[0] + "domain2.com"
	newAlias := fmt.Sprintf("%sdomain2.com", username[0])
//	fmt.Println(newEmail)
	return newAlias
}


func main() {
	//I want to keep a local record of all the written emails, currently I am going
	//to use a csv file, as a database.
	file, err := os.OpenFile("emails.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	w := csv.NewWriter(file)

	for _,  email := range emails {

		alias := createAlias(email)
		if err := w.Write([]string{email,alias}); err != nil {
			log.Fatalln("error writing record to csv: ", err)
		}
//		fmt.Printf("\nOriginal Email: %v, Alias: %s", v, alias)
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
