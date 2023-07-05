package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Post struct {
	Id       string
	Username string
	Content  string
	// How to use dates in golang
	// source: https://www.digitalocean.com/community/tutorials/how-to-use-dates-and-times-in-go
	Timestamp string
}

func getCurrentTime() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05-0700")
}
func main() {

	fmt.Println("Hello, world!")

	firstHandler := func(w http.ResponseWriter, r *http.Request) {
		posts := map[string][]Post{
			"Posts": {
				{
					Id:        "8eeeffde-9c6a-4b6d-9ad2-7274b6b70c0a",
					Username:  "user1",
					Content:   "This is the first post in the array.",
					Timestamp: "2023-07-05T10:30+0000",
				},
				{
					Id:        "126d3678-7362-4a65-b716-92c94df34a12",
					Username:  "user2",
					Content:   "The second post is even more exciting!",
					Timestamp: "2023-07-05T11:15+0000",
				},
				{
					Id:        "b2d7a6f4-62d2-4869-9787-1e2f352ad141",
					Username:  "user3",
					Content:   "post number three, sharing some interesting news.",
					Timestamp: "2023-07-05T12:00+0000",
				},
				{
					Id:        "4f35e100-6d2e-48e5-b500-25f73e0e066f",
					Username:  "user4",
					Content:   "Fourth post here, getting the hang of it!",
					Timestamp: "2023-07-05T13:30+0000",
				},
				{
					Id:        "e92af500-89cc-4772-a58e-336f8e200f3f",
					Username:  "user5",
					Content:   "Fifth post, reaching out to the Twitterverse!",
					Timestamp: "2023-07-05T14:45+0000",
				},
				{
					Id:        "d81c9b0d-15ae-4ab1-a501-9af0bda9b09e",
					Username:  "user6",
					Content:   "This is post number six, halfway there!",
					Timestamp: "2023-07-05T15:20+0000",
				},
				{
					Id:        "301f8e89-7f0c-49dd-80b7-0d82a6becc94",
					Username:  "user7",
					Content:   "Seventh post, sharing an interesting article.",
					Timestamp: "2023-07-05T16:00+0000",
				},
				{
					Id:        "3fc83155-9ff4-456d-9c2b-b37a8e457a69",
					Username:  "user8",
					Content:   "Eighth post, just expressing some thoughts.",
					Timestamp: "2023-07-05T17:10+0000",
				},
				{
					Id:        "5d3e94b7-3525-4a3d-9d23-c04856719934",
					Username:  "user9",
					Content:   "Ninth post, connecting with the Twitter community!",
					Timestamp: "2023-07-05T18:30+0000",
				},
				{
					Id:        "f1fb3074-4c97-4c0f-90d5-0e968eef1268",
					Username:  "user10",
					Content:   "Tenth post, double digits already!",
					Timestamp: "2023-07-05T19:05+0000",
				},
				{
					Id:        "1dd26a52-8cb0-4156-9262-83ae8a4e78fb",
					Username:  "user11",
					Content:   "Eleventh post, just sharing a quick update.",
					Timestamp: "2023-07-05T20:20+0000",
				},
				{
					Id:        "73d92ed6-80db-4f73-96d0-7cdd6876ed92",
					Username:  "user12",
					Content:   "Twelfth post, what an amazing day!",
					Timestamp: "2023-07-05T21:45+0000",
				},
				{
					Id:        "a9ffed06-7647-4591-8b3f-0c3a0c9ae5c7",
					Username:  "user13",
					Content:   "Thirteenth post, lucky number!",
					Timestamp: "2023-07-05T22:10+0000",
				},
				{
					Id:        "c0829ce9-4b35-415a-bc61-4e0b0f826b6e",
					Username:  "user14",
					Content:   "Fourteenth post, sharing a funny meme!",
					Timestamp: "2023-07-05T23:30+0000",
				},
				{
					Id:        "9e4b7e22-7be4-4b3a-9345-49ee0707af36",
					Username:  "user15",
					Content:   "Fifteenth post, time flies on Twitter!",
					Timestamp: "2023-07-06T00:15+0000",
				},
				{
					Id:        "4af3cda0-6f0a-4af4-af4e-1c0af4a6af4a",
					Username:  "user16",
					Content:   "Sixteenth post, engaging with followers.",
					Timestamp: "2023-07-06T01:00+0000",
				},
				{
					Id:        "b049ebdb-152f-4e6f-b872-0133f9b6525a",
					Username:  "user17",
					Content:   "Seventeenth post, sharing a cool photo!",
					Timestamp: "2023-07-06T02:20+0000",
				},
				{
					Id:        "f4f785ed-344f-42e7-bb96-9728da40a17b",
					Username:  "user18",
					Content:   "Eighteenth post, let's start a conversation!",
					Timestamp: "2023-07-06T03:05+0000",
				},
				{
					Id:        "2f8676e0-9202-46b3-b7f0-ff1da535eb6b",
					Username:  "user19",
					Content:   "Nineteenth post, almost there!",
					Timestamp: "2023-07-06T04:40+0000",
				},
				{
					Id:        "a67b9c1a-8bea-4ea7-9d22-c9f10b2f351f",
					Username:  "user20",
					Content:   "Twentieth post, what an incredible journey!",
					Timestamp: "2023-07-06T05:15+0000",
				},
			},
		}

		template := template.Must(template.ParseFiles("index.html"))
		template.Execute(w, posts)
	}

	createPostHandler := func(w http.ResponseWriter, r *http.Request) {
		log.Print("Post received!")
		content := r.PostFormValue("content")
		username := r.PostFormValue("username")
		log.Print("\n")

		htmlString := fmt.Sprintf(`
		<div class="bg-blue-50 rounded-xl px-4 py-2">
		<div class="flex flex-row place-items-center justify-between">
			<p class="text-sm font-light">@%s</p>
			<p class="text-xs">%s</p>
		</div>
		<div class="my-2 rounded-md">
			<p class="text-base"">%s</p>
		</div>
	</div>`, username, getCurrentTime(), content)

		postTemplate, err := template.New("postTemplate").Parse(htmlString)
		if err != nil {
			log.Fatal(err)
		}

		postTemplate.Execute(w, nil)
	}

	http.HandleFunc("/", firstHandler)
	http.HandleFunc("/create-post/", createPostHandler)
	// Enables CSS styling
	// source: https://stackoverflow.com/a/43601392
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", nil)
}
