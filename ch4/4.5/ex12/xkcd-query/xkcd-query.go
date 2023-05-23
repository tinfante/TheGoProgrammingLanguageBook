package main

import (
    "fmt"
    "os"
    "strings"
    "encoding/json"
    "io/ioutil"
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
    Day string
}


func check(err error) {
    if err != nil {
        panic(err)
    }
}


func main() {
    terms := os.Args[1:]
    dname := "../xkcd-db/data"
    d, err := os.Open(dname)
    check(err)
    defer d.Close()
    files, err := d.Readdir(0)
    check(err)
    for _, v := range files {
        fname := v.Name()
        if strings.HasSuffix(fname, ".json") {
            f, err := os.Open(dname + "/" + fname)
            check(err)
            defer f.Close()
            fbytes, err := ioutil.ReadAll(f)
            check(err)
            var jdata xkcdJson
            json.Unmarshal(fbytes, &jdata)
            for _, t := range  terms {
                if strings.Contains(
                        strings.ToLower(jdata.Transcript),
                        strings.ToLower(t)) {
                    fmt.Printf("%s\n%s\n\n", jdata.Img, jdata.Transcript)
                    break
                }
            }
        }
    }
}
