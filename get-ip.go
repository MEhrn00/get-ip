package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "log"
    "encoding/json"
)

type data struct {
    IP string `json:"ip"`
}

func main() {

    url := "https://api.ipify.org?format=json"
    var response data

    resp, err := http.Get(url)

    if err != nil {
        log.Fatalln(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }

    json.Unmarshal([]byte(body), &response)
    fmt.Println(response.IP)

}
