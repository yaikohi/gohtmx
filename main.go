package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Post struct {
	Id       string
	Username string
	Content  string
}

func main() {

	fmt.Println("Hello, world!")

	firstHandler := func(w http.ResponseWriter, r *http.Request) {
		posts := map[string][]Post{
			"Posts": {
				{
					Id:       "8eeeffde-9c6a-4b6d-9ad2-7274b6b70c0a",
					Username: "user1",
					Content:  "This is the first tweet in the array.",
				},
				{
					Id:       "126d3678-7362-4a65-b716-92c94df34a12",
					Username: "user2",
					Content:  "The second tweet is even more exciting!",
				},
				{
					Id:       "b2d7a6f4-62d2-4869-9787-1e2f352ad141",
					Username: "user3",
					Content:  "Tweet number three, sharing some interesting news.",
				},
				{
					Id:       "4f35e100-6d2e-48e5-b500-25f73e0e066f",
					Username: "user4",
					Content:  "Fourth tweet here, getting the hang of it!",
				},
				{
					Id:       "e92af500-89cc-4772-a58e-336f8e200f3f",
					Username: "user5",
					Content:  "Fifth tweet, reaching out to the Twitterverse!",
				},
				{
					Id:       "d81c9b0d-15ae-4ab1-a501-9af0bda9b09e",
					Username: "user6",
					Content:  "This is tweet number six, halfway there!",
				},
				{
					Id:       "301f8e89-7f0c-49dd-80b7-0d82a6becc94",
					Username: "user7",
					Content:  "Seventh tweet, sharing an interesting article.",
				},
				{
					Id:       "3fc83155-9ff4-456d-9c2b-b37a8e457a69",
					Username: "user8",
					Content:  "Eighth tweet, just expressing some thoughts.",
				},
				{
					Id:       "5d3e94b7-3525-4a3d-9d23-c04856719934",
					Username: "user9",
					Content:  "Ninth tweet, connecting with the Twitter community!",
				},
				{
					Id:       "f1fb3074-4c97-4c0f-90d5-0e968eef1268",
					Username: "user10",
					Content:  "Tenth tweet, double digits already!",
				},
				{
					Id:       "1dd26a52-8cb0-4156-9262-83ae8a4e78fb",
					Username: "user11",
					Content:  "Eleventh tweet, just sharing a quick update.",
				},
				{
					Id:       "73d92ed6-80db-4f73-96d0-7cdd6876ed92",
					Username: "user12",
					Content:  "Twelfth tweet, what an amazing day!",
				},
				{
					Id:       "a9ffed06-7647-4591-8b3f-0c3a0c9ae5c7",
					Username: "user13",
					Content:  "Thirteenth tweet, lucky number!",
				},
			},
		}

		template := template.Must(template.ParseFiles("index.html"))
		template.Execute(w, posts)
		// io.WriteString(w, r.Method)
	}

	http.HandleFunc("/", firstHandler)

	http.ListenAndServe(":8080", nil)
}
