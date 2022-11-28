package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type comments struct {
	Comments []comment `json:"comments"`
}

type comment struct {
	ID     int    `json:"id"`
	Body   string `json:"body"`
	PostId int    `json:"postid"`
}

func main() {
	resp, err := http.Get("https://my-json-server.typicode.com/typicode/demo/db")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Print(sb)

	_ = ioutil.WriteFile("file.json", body, 0644)

	// Reading only comments
	file, _ := ioutil.ReadFile("file.json")
	data := comments{}
	_ = json.Unmarshal(file, &data)
	for i := 0; i < len(data.Comments); i++ {
		fmt.Println("comments ID: ", data.Comments[i].ID)
		fmt.Println("Body: ", data.Comments[i].Body)
		fmt.Println("PostId: ", data.Comments[i].PostId)
	}

}
