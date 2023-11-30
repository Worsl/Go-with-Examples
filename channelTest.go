package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "context"
    "time"
)

func checkError(err error){
    if err != nil{
        log.Fatal(err)
        os.Exit(1)
    }
}

func main()  {

    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    // Create an HTTP request with the context.
    req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/pokedex/kanto/", nil)
    checkError(err)
    req = req.WithContext(ctx)

    // Send the request and get the response.
    client := &http.Client{}
    response, err := client.Do(req)
    checkError(err)

    defer response.Body.Close()
    // Read and print the response data.
    responseData, err := io.ReadAll(response.Body)
    checkError(err)

    fmt.Println(string(responseData))
    
}