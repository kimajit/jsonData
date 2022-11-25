package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

// type posts struct {
// 	Id    int    `json:"id"`
// 	Title string `json:"title"`
// }

// type comments struct {
// 	Id     int    `json:"id"`
// 	Body   string `json:"body"`
// 	PostId int    `json:"postid"`
// }

// type profile struct {
// 	Name string `json:"name"`
// }

func main() {
	resp, err := http.Get("https://my-json-server.typicode.com/typicode/demo/db")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Print(sb)

	_ = ioutil.WriteFile("file.json", body, 0644)
	// var result profile
	// if err := json.Unmarshal(body, &result); err != nil {
	// 	fmt.Println("cannot unmarshal")
	// }

	// for _, rec := range result.name {
	// 	fmt.Println(rec.name)
}
