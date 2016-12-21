//blockparty.go handles the webserver functions for the blockparty demo
package main

import (
	"encoding/json"
	"fmt"
	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	_ "os"
	"strconv"
	"time"
)

var (
	pool *redis.Pool
)

func newPool(addr string, port string, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr+":"+port)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
}

func check(function string, e error) {
	if e != nil {
		log.Fatal(function, e)
	}
}

//House that will be listed
type House struct {
	Name     string `json:"name" redis:"name"`
	Address  string `json:"address" redis:"address"`
	Price    string `json:"price" redis:"price"`
	Image    string `json:"image" redis:"image"`
	Contract string `json:"contract" redis:"contract"`
	Status   string `json:"status" redis:"status"`
}

//JSONPayload is a generic Container to hold JSON
type JSONPayload struct {
	Data []House
}

func getHouses() []byte {
	var payload JSONPayload
	var listings []House
	var h House
	listings = make([]House, 0)

	c := pool.Get()
	defer c.Close()

	n, err := redis.Strings(c.Do("SMEMBERS", "houses"))
	check("SMEMBERS", err)

	for _, v := range n {
		r, err := redis.Values(c.Do("HGETALL", v))
		err = redis.ScanStruct(r, &h)
		check("ScanStruct", err)
		listings = append(listings,
			House{Name: h.Name, Address: h.Address, Price: h.Price, Image: h.Image, Contract: h.Contract, Status: h.Status})
	}

	payload.Data = listings
	houses, err := json.Marshal(payload)
	check("getHouses()", err)
	return houses
}

func setDefaultHouses() {
	var listings []House
	listings = make([]House, 0)

	listings = append(listings,
		House{Name: "House1", Address: "123 Main Street", Price: "$300,000", Image: "house1.jpg", Contract: "abc-123", Status: "sold"},
		House{Name: "House2", Address: "66 Pine Street", Price: "$150,000", Image: "house2.jpg", Contract: "abc-234", Status: "sold"},
		House{Name: "House3", Address: "8500 Rue Avenue", Price: "$900,000", Image: "house3.jpg", Contract: "abc-345", Status: "sold"},
		House{Name: "House4", Address: "1250 Maple Road", Price: "$450,000", Image: "house4.jpg", Contract: "abc-456", Status: "sold"},
		House{Name: "House5", Address: "34A Bridge Street", Price: "$90,000", Image: "house5.jpg", Contract: "abc-567", Status: "sold"})

	c := pool.Get()
	defer c.Close()

	for _, v := range listings {
		_, err := c.Do("SADD", "houses", "house:"+v.Contract)
		check("LPUSH", err)
		_, err = c.Do("HMSET", "house:"+v.Contract, "name", v.Name, "address", v.Address, "price", v.Price, "image", v.Image, "contract", v.Contract, "status", v.Status)
		check("HMSET", err)
	}
}

func initialize() {
	fmt.Println("Starting")
	env, _ := cfenv.Current()
	services := env.Services
	credentials := services["p-redis"][0].Credentials
	host := credentials["host"].(string)
	password := credentials["password"].(string)
	port := strconv.FormatFloat(credentials["port"].(float64), 'f', -1, 64)

	pool = newPool(host, port, password)
	setDefaultHouses()
}

func listingHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/listings.tmpl")
	check("Parse template", err)
	var listings JSONPayload
	err = json.Unmarshal(getHouses(), &listings)
	check("Unmarshal", err)
	t.Execute(w, listings)
}

func main() {
	initialize()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", listingHandler)
	http.Handle("/images/", http.FileServer(http.Dir("")))
	http.Handle("/css/", http.FileServer(http.Dir("")))
	http.Handle("/fonts/", http.FileServer(http.Dir("")))
	http.Handle("/js/", http.FileServer(http.Dir("")))
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
