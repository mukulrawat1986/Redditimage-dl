package main

import (
    "log"
    "io/ioutil"
    "net/http"
)

const url = "https://i.imgur.com/EdcinRv.jpg"

func main() {
    // Just a simple GET request to the image URL
    // We get back a *Response, and an error
    res, err := http.Get(url)

    if err != nil {
        log.Fatalf("http.Get -> %v", err)
    }

    // We read all the bytes of the image
    // Types: data []byte
    data, err := ioutil.ReadAll(res.Body)

    if err != nil {
        log.Fatalf("ioutil.ReadAll -> %v", err)
    }
    
    // Close the response body
    defer res.Body.Close()

    // You can now save it to disk or whatever...
    log.Println("Saving Image.........")
    ioutil.WriteFile("Dog.jpg", data, 0666)

    log.Println("Image Saved")
}
