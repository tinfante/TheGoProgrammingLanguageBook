package main

import (
    "fmt"
    "os"
    "net/http"
    "net/url"
    "strings"
    "io/ioutil"
    "io"
    "encoding/json"
    "github.com/joho/godotenv"
)

func check(err error) {
    if err != nil {
        panic(err)
    }
}

type Movie struct {
    Title string
    Year string
    Poster string
}

func main() {
    godotenv.Load()
    key := os.Getenv("api_key")
    search := os.Args[1:]
    url := "http://www.omdbapi.com/?apikey=" + key + "&t=" + url.QueryEscape(strings.Join(search, " "))
    resp, err := http.Get(url)
    check(err)
    body, err := ioutil.ReadAll(resp.Body)
    check(err)
    defer resp.Body.Close()
    var movie Movie
    json.Unmarshal(body, &movie)
    if movie == (Movie{}) {
        fmt.Println("Couldn't find movie.")
        os.Exit(0)
    }
    if movie.Poster == "N/A" {
        fmt.Printf("%s (%s)\n%s\n", movie.Title, movie.Year, movie.Poster)
        os.Exit(0)
    }
    splitPoster := strings.Split(movie.Poster, ".")
    ftype := splitPoster[len(splitPoster)-1]
    fname := "posters/" + strings.ReplaceAll(movie.Title, " ", "_") + "-" +
        movie.Year + "." + ftype
    f, err := os.Create(fname)
    check(err)
    defer f.Close()
    resp, err = http.Get(movie.Poster)
    check(err)
    defer resp.Body.Close()
    _, err2 := io.Copy(f, resp.Body)
    check(err2)
    fmt.Printf("%s (%s)\nOK => %s\n", movie.Title, movie.Year, fname)
}
