package clientRequest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func CliAndReq() {
	// only single Client to deal with various concurrent reqs from go routines is needed
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Generating request
	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet, "https://jsonplaceholder.typicode.com/todos/1",
		nil) //io.Reader (Body)
	if err != nil {
		panic(err)
	}
	// Setting header
	req.Header.Add("X-My-Client", "Let's go")

	// Requesting and processing response
	res, err := client.Do(req)
	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("unexpected status: got %v", res.Status))
	}
	defer res.Body.Close()

	fmt.Println(res.Header.Get("Content-Type"))
	var data struct {
		UserID    int    `json:"userId"`
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", data)
}
