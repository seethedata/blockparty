//blockparty.go handles the webserver functions for the blockparty demo
package main

import (
	"fmt"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"github.com/cloudfoundry-community/go-cfenv"
	"html/template"
	"log"
	"net/http"
	"strconv"
	_"os"
)

func check(function string, e error) {
	if e != nil {
		log.Fatal(function, e)
	}
}

//House that will be listed
type House struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Price    string `json:"price"`
	Image    string `json:"image"`
	Contract string `json:"contract"`
}

//JSONPayload is a generic Container to hold JSON
type JSONPayload struct {
	Data []House
}

type serviceList struct {
	service  map[string]interface{}
}

type serviceInfo struct {
	Credentials []string
	Label string
	Name string
	Plan string
	Provider string
	Drain string
	Tags []string
	VolumeMounts []string
}

func getHouses() []byte {
	var payload JSONPayload
	var listings []House
	listings = make([]House, 0)

	listings = append(listings,
		House{Name: "House1", Address: "123 Main Street", Price: "$300,000", Image: "house1.jpg", Contract: "abc-123"},
		House{Name: "House2", Address: "66 Pine Street", Price: "$150,000", Image: "house2.jpg", Contract: "abc-234"},
		House{Name: "House3", Address: "8500 Rue Avenue", Price: "$900,000", Image: "house3.jpg", Contract: "abc-345"},
		House{Name: "House4", Address: "1250 Maple Road", Price: "$450,000", Image: "house4.jpg", Contract: "abc-456"},
		House{Name: "House5", Address: "34A Bridge Street", Price: "$90,000", Image: "house5.jpg", Contract: "abc-567"})

	payload.Data = listings
	houses, err := json.Marshal(payload)
	check("getHouses()", err)
	return houses
}

func listingHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/listings.tmpl")
	check("Parse template", err)
	var listings JSONPayload
	err = json.Unmarshal(getHouses(), &listings)
	check("Unmarshal", err)
	t.Execute(w, listings)
}

func initialize() {
	fmt.Println("Starting")
	connect2Redis()
}

func connect2Redis() {
	env,_:=cfenv.Current()
	services:=env.Services
	credentials:=services["p-redis"][0].Credentials
	host:=credentials["host"]
	password:=credentials["password"]
	port:=strconv.FormatFloat(credentials["port"].(float64),'f',-1,64)
	c, err := redis.DialURL(fmt.Sprintf("redis://%s:%s",host,port))
	check("Connect", err)
	n,err:=c.Do("AUTH",password)
	check("AUTH", err)
	defer c.Close()



}
func main() {
	initialize()
	http.HandleFunc("/", listingHandler)
	http.Handle("/images/", http.FileServer(http.Dir("")))
	http.Handle("/css/", http.FileServer(http.Dir("")))
	http.Handle("/fonts/", http.FileServer(http.Dir("")))
	http.Handle("/js/", http.FileServer(http.Dir("")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
