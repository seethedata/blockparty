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
	"strconv"
	"time"
)

var (
	pool *redis.Pool
)
var	redis_keys = `{ "services" : [
			{
                                "platform": "pws",
                                "service": "rediscloud",
                                "host": "hostname",
                                "port": "port",
                                "password": "password"
                        },
			{
                                "platform": "pcfdev",
                                "service": "p-redis",
                                "host": "host",
                                "port": "port",
                                "password": "password"
                         },
			 {
                                "platform": "bluemix6",
                                "service": "redis-2.6",
                                "host": "hostname",
                                "port": "port",
                                "password": "password"
                         }]
            }`
type cfServices struct {
	Services	[]cfService `json:"services"`
}

type cfService struct {
	Platform string `json:"platform"`
	Name	string `json:"service"`
	Host	string	`json:"host"`
	Port	string	`json:"port"`
	Password	string `json:"password"`
}

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
	Description string `json:"description" redis:"description"`
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
		House{Name: h.Name, Address: h.Address, Price: h.Price, Image: h.Image, Contract: h.Contract, Description: h.Description, Status: h.Status})
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
	House{Name: "House1", Address: "123 Main Street", Price: "$300,000", Image: "house1.jpg", Contract: "abc-123", Description: "A good house with 3 bd/2br", Status: "sold"},
	House{Name: "House2", Address: "66 Pine Street", Price: "$150,000", Image: "house2.jpg", Contract: "abc-234", Description: "Great starter house! 3bd/1br", Status: "listed"},
	House{Name: "House3", Address: "8500 Rue Avenue", Price: "$900,000", Image: "house3.jpg", Contract: "abc-345", Description: "Gorgeous home. 4bd/3br", Status: "sold"},
	House{Name: "House4", Address: "1250 Maple Road", Price: "$450,000", Image: "house4.jpg", Contract: "abc-456", Description:"Beautiful home 4bd/3br and a pool", Status: "sold"},
	House{Name: "House5", Address: "34A Bridge Street", Price: "$90,000", Image: "house5.jpg", Contract: "abc-567", Description: "Great location! No HOA fees!", Status: "sold"})

	c := pool.Get()
	defer c.Close()

	for _, v := range listings {
		_, err := c.Do("SADD", "houses", "house:"+v.Contract)
		check("LPUSH", err)
		_, err = c.Do("HMSET", "house:"+v.Contract, "name", v.Name, "address", v.Address, "price", 
				v.Price, "image", v.Image, "contract", v.Contract, "description", v.Description, "status", v.Status)
		check("HMSET", err)
	}
}

func initialize() {
	fmt.Println("Starting")
	env, _ := cfenv.Current()
	services := env.Services
	var cfServices cfServices
	var credentials map[string] interface{}
	var host string
	var password string
	var port string
	err:=json.Unmarshal([]byte(redis_keys),&cfServices)
	check("Unmarshal",err)

	for _,service:= range cfServices.Services  {
		if _,ok:=services[service.Name]; ok {
			credentials= services[service.Name][0].Credentials
			if _,ok:= credentials[service.Host] ; ok {
				host= credentials[service.Host].(string)
			} else {
				log.Fatal("Unable to identify Redis host from config. Platform attempted:" + service.Platform)
			}
			if _,ok:= credentials[service.Password]; ok {
				password= credentials[service.Password].(string)
			} else {
				log.Fatal("Unable to identify Redis password from config. Platform attempted:" + service.Platform)
			}
			if _,ok:= credentials[service.Port]; ok {
				switch credentials[service.Port].(type) {
					case string:
						port= credentials[service.Port].(string)
					case float64:
						port= strconv.FormatFloat(credentials[service.Port].(float64), 'f', -1, 64)
					default:
						log.Fatal("Redis port value is of unexpected type.")
				}
			} else {
				log.Fatal("Unable to identify Redis port from config. Platform attempted:" + service.Platform)
			}
			break
		}
	}

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

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	vars:=mux.Vars(r)
	id:=vars["contract-id"]

	c := pool.Get()
	defer c.Close()

	var h House
	n, err := redis.Values(c.Do("HGETALL", "house:" + id))
	err = redis.ScanStruct(n, &h)
	check("ScanStruct", err)
	t, err := template.ParseFiles("templates/details.tmpl")
	check("Parse template", err)
	t.Execute(w, h)
}

func approveHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/approve.tmpl")
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
	router.HandleFunc("/details/{contract-id}",detailsHandler)
	router.HandleFunc("/approve",approveHandler)
	http.Handle("/images/", http.FileServer(http.Dir("/app")))
	http.Handle("/css/", http.FileServer(http.Dir("/app")))
	http.Handle("/fonts/", http.FileServer(http.Dir("/app")))
	http.Handle("/js/", http.FileServer(http.Dir("/app")))
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
