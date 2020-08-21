package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type data struct {
	IP string `json:"ip"`
}

func main() {

	var url string
	protocol := flag.Int("protocol", 4, "Format to return the IP address either 6 for IPv6 or 4 for IPv4")
	flag.Parse()

	switch *protocol {
	case 4:
		url = "https://api.ipify.org?format=json"
	case 6:
		url = "https://api6.ipify.org?format=json"
	default:
		fmt.Println("Invalid protocol. Must be either 6 or 4.")
		os.Exit(1)
	}
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
