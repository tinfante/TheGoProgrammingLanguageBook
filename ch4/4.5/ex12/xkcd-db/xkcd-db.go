package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "os"
    "strconv"
)

type xkcdJson struct {
    Month string
    Num int
    Link string
    Year string
    News string
    SafeTitle string `json:"safe_title"`
    Transcript string
    Alt string
    Img string
    Title string
    day string
}


func makeUrl(id int) string {
    return fmt.Sprintf("https://xkcd.com/%d/info.0.json", id)
}


func getUrl(url string) (*xkcdJson, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getUrl failed: %s", resp.Status)
    }
    var result xkcdJson
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    for id := 1; id <= 100; id++ {
        xj, err := getUrl(makeUrl(id))
        check(err)
        data, err := json.MarshalIndent(xj, "",  "    ")
        check(err)
        fmt.Printf("%s\n", data)
        f, err := os.Create("data/"+strconv.Itoa(id)+".json")
        check(err)
        defer f.Close()
        _, err = f.WriteString(string(data))
        check(err)
    }
}
