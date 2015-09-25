package main

import (
    "errors"
    "encoding/json"
	"fmt"
	_"io"
	"io/ioutil"
	"log"
	"net/http"
	_"os"
    "time"
)

// const url = "https://i.imgur.com/EdcinRv.jpg"

// Create our own type here to store the JSON response
// and to store our data
type Item struct {
	Title string
	URL   string
}

type Response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

// Add a string method to our Item type so that it implements the Stringer interface
// and makes it easier for us to print
func (i Item)String() string{
    return fmt.Sprintf("%s\n%s", i.Title, i.URL)
}


func downloadImage(url string, title string) {
	// Just a simple GET request to the image URL
	// We get back a *Response, and an error
	res, err := http.Get(url)

	if err != nil {
		log.Fatalf("http.Get -> %v", err)
	}

//	fmt.Printf("%v", res.Body)

	// We read all the bytes of the image
	// Types: data []byte
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("ioutil.ReadAll -> %v", err)
	}

	// Close the response body
	defer res.Body.Close()

	// You can now save it to disk or whatever...
	//log.Println("Saving Image.........")
    filename := fmt.Sprintf("%s.jpg", title)
    if title == "bunny"{
        filename = fmt.Sprintf("%s.png", title)
    }
    log.Println("Saving image ", filename)
	ioutil.WriteFile(filename, data, 0666)
	log.Println("Image Saved")
}

// Function to get link and title from a subreddit
func Get(subreddit string) ([]Item, error){

    url := fmt.Sprintf("http://www.reddit.com/r/%s.json", subreddit)
    r, err := http.Get(url)
    if err != nil{
        return nil, err
    }

    defer r.Body.Close()
    if r.StatusCode != http.StatusOK{
        return nil, errors.New(r.Status)
    }

    resp := new(Response)
    err = json.NewDecoder(r.Body).Decode(resp)
    if err != nil{
        return nil, err
    }

    items := make([]Item, len(resp.Data.Children))
    for i, child := range resp.Data.Children{
        items[i] = child.Data
    }
    return items, nil
}

func main() {
    items, err := Get("aww")
    if err != nil{
        log.Fatal(err)
    }

    for _, item := range items{
        fmt.Println(item)
    }

    // Download first 5 images
    count := 0
    for _, item := range items{
        if item.URL[len(item.URL)-3:] != "jpg"{
            continue
        }else{
            downloadImage(item.URL, item.Title[0:5])
            count += 1
        }
        if count == 5 {
            break
        }
    }
    time.Sleep(3)

    downloadImage("http://i.imgur.com/GQcvWqx.jpg", "bunny")
}
