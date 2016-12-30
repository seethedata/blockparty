//blockparty.go handles the webserver functions for the blockparty demo
package main

import (
	"encoding/json"
	"fmt"
	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/satori/go.uuid"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	pool    *redis.Pool
	mainURL string       = "https://blockparty.local.pcfdev.io"
	store   *sessions.CookieStore = sessions.NewCookieStore([]byte("BlockParty"))
)

type cfServices struct {
	Services []cfService `json:"services"`
}

type cfService struct {
	Platform string `json:"platform"`
	Name     string `json:"service"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

//House that will be listed
type House struct {
	Name        string  `json:"name" redis:"name"`
	Address     string  `json:"address" redis:"address"`
	Price       float64 `json:"price" redis:"price"`
	Image       string  `json:"image" redis:"image"`
	Contract    string  `json:"contract" redis:"contract"`
	Description string  `json:"description" redis:"description"`
	Bedrooms    float64 `json:"bedrooms" redis:"bedrooms"`
	Bathrooms   float64 `json:"bathrooms" redis:"bathrooms"`
	Status      string  `json:"status" redis:"status"`
	Quality     int     `json:"quality" redis:"quality"`
}

type Bid struct {
	User     string  `json:"user" redis:"user"`
	Amount   float64 `json:"amount" redis:"amount"`
	Contract string  `json:"contract" redis:"contract"`
	Status   string  `json:"status" redis:"status"`
}

type bidPayload struct {
	Data []Bid  `json:"data" redis:"data"`
	User string `json:"user" redis:"user"`
	Url  string `json:"url" redis:"url"`
}

//JSONPayload is a generic Container to hold JSON
type JSONPayload struct {
	Data []House `json:"data" redis:"data"`
	User string  `json:"user" redis:"user"`
	Url  string  `json:"url" redis:"url"`
	Message  string  `json:"message" redis:"message"`
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

func newPayload() *JSONPayload {
	return &JSONPayload{Url: mainURL}
}

func newBidPayload() *bidPayload {
	return &bidPayload{Url: mainURL}
}

func getHouses() []House {
	var h House
	var houses []House = make([]House, 0)

	c := pool.Get()
	defer c.Close()

	n, err := redis.Strings(c.Do("SMEMBERS", "houses"))
	check("SMEMBERS", err)

	for _, v := range n {
		r, err := redis.Values(c.Do("HGETALL", v))
		err = redis.ScanStruct(r, &h)
		check("ScanStruct", err)
		houses = append(houses, House{Name: h.Name, Address: h.Address, Price: h.Price, Image: h.Image, Contract: h.Contract,
			Description: h.Description, Bedrooms: h.Bedrooms, Bathrooms: h.Bathrooms, Status: h.Status, Quality: h.Quality})
	}
	return houses
}

func setDefaultHouses() {
	c := pool.Get()
	defer c.Close()

	n, err := redis.Int(c.Do("EXISTS", "houses"))
	check("EXISTS", err)
	if n == 0 {
		fmt.Println("Creating default houses.")
		file, err := ioutil.ReadFile("./houses.json")
		check("Read JSON", err)

		var listings JSONPayload
		err = json.Unmarshal(file, &listings)
		check("Unmarshal", err)

		for _, v := range listings.Data {
			contract := getContractId()
			_, err := c.Do("SADD", "houses", "house:"+contract)
			check("LPUSH", err)
			_, err = c.Do("HMSET", "house:"+contract, "name", v.Name, "address", v.Address, "price",
				v.Price, "image", v.Image, "contract", contract, "description", v.Description, "bedrooms", v.Bedrooms, "bathrooms", v.Bathrooms, "status", v.Status)
			check("HMSET", err)
		}
	} else {
		fmt.Println("Default houses already exist. Skipping house creation.")
	}
}

func initialize() {
	var cfServices cfServices
	fmt.Println("Starting")
	file, err := ioutil.ReadFile("./services.json")
	check("Read JSON", err)

	err = json.Unmarshal(file, &cfServices)
	check("Unmarshal", err)

	env, _ := cfenv.Current()
	services := env.Services

	var credentials map[string]interface{}
	var host string
	var password string
	var port string

	for _, service := range cfServices.Services {
		if _, ok := services[service.Name]; ok {
			credentials = services[service.Name][0].Credentials
			if _, ok := credentials[service.Host]; ok {
				host = credentials[service.Host].(string)
			} else {
				log.Fatal("Unable to identify Redis host from config. Platform attempted:" + service.Platform)
			}
			if _, ok := credentials[service.Password]; ok {
				password = credentials[service.Password].(string)
			} else {
				log.Fatal("Unable to identify Redis password from config. Platform attempted:" + service.Platform)
			}
			if _, ok := credentials[service.Port]; ok {
				switch credentials[service.Port].(type) {
				case string:
					port = credentials[service.Port].(string)
				case float64:
					port = strconv.FormatFloat(credentials[service.Port].(float64), 'f', -1, 64)
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

func getContractId() string {
	return uuid.NewV4().String()
}

func getUserId() string {
	return uuid.NewV4().String()
}

func getHouse(id string) House {
	var h House
	c := pool.Get()
	defer c.Close()

	n, err := redis.Values(c.Do("HGETALL", "house:"+id))
	check("HGETALL", err)
	err = redis.ScanStruct(n, &h)
	check("ScanStruct", err)
	return h
}

func getBids(i string) []Bid {
	var bids []Bid = make([]Bid, 0)
	c := pool.Get()
	defer c.Close()
	n, err := redis.Values(c.Do("ZREVRANGEBYSCORE", "bids:"+i, "+inf", "-inf"))
	check("ZREVRANGEBYSCORE", err)

	for _, v := range n {
		var bid Bid
		err = json.Unmarshal(v.([]byte), &bid)
		check("Unmarshal", err)
		bids = append(bids, bid)
	}
	return bids
}

func getMyBids(u string) []Bid {
	var bids []Bid = make([]Bid, 0)
	var myBids []Bid = make([]Bid, 0)
	var houses []House
	houses = getHouses()

	for _, h := range houses {
		bids = getBids(h.Contract)
		for _, b := range bids {
			if b.User == u {
				myBids = append(myBids, b)
			}
		}
	}
	return myBids

}

func listingsHandler(w http.ResponseWriter, r *http.Request) {
	var u string
	session, err := store.Get(r, "BlockPartySession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, ok := session.Values["user"]; ok {
		u = session.Values["user"].(string)
	} else {
		u = getUserId()
		session.Values["user"] = u
	}

	session.Save(r, w)
	t, err := template.ParseFiles("templates/listings.tmpl", "templates/navbar.tmpl")
	check("Parse template", err)
	listings := newPayload()
	log.Print(listings.Url)
	listings.Data = getHouses()
	listings.User = u
	t.Execute(w, listings)
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	var u string
	session, err := store.Get(r, "BlockPartySession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u = session.Values["user"].(string)

	vars := mux.Vars(r)
	i := vars["contract-id"]

	c := pool.Get()
	defer c.Close()

	var h House
	h = getHouse(i)

	var payload = newPayload()
	payload.Data = append(payload.Data, h)
	payload.User = u
	t, err := template.ParseFiles("templates/details.tmpl", "templates/navbar.tmpl")
	check("Parse template", err)
	t.Execute(w, payload)
}

func bidsHandler(w http.ResponseWriter, r *http.Request) {
	var u string
	vars := mux.Vars(r)
	i := vars["contract-id"]

	session, err := store.Get(r, "BlockPartySession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u = session.Values["user"].(string)

	p := newBidPayload()
	p.Data = getBids(i)
	p.User=u

	t, err := template.ParseFiles("templates/bids.tmpl", "templates/navbar.tmpl")
	check("Parse template", err)
	t.Execute(w, p)

}

func bidHandler(w http.ResponseWriter, r *http.Request) {
	var u string
	bidExists:=false
	message:=""
	session, err := store.Get(r, "BlockPartySession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u = session.Values["user"].(string)

	vars := mux.Vars(r)
	i := vars["contract-id"]
	a := r.PostFormValue("bidAmount")
	a = strings.Replace(a,",","",-1)
	a = strings.Replace(a,".","",-1)


	c := pool.Get()
	defer c.Close()

	n:=getBids(i)

	for _,v:=range n {
		if v.User == u {
			bidExists=true
			break
		}
	}
	if bidExists {
		message="User " + u + " has already bid on house " + i + "."
	} else {
		_, err = c.Do("ZADD", "bids:" + i, a, "{\"user\":\"" + u + "\", \"amount\":" + a + ", \"contract\":\"" + i + "\", \"status\":\"submitted\"}")
		check("ZADD", err)
		message="Bid on " + i + " submitted."
	}
	var h House
	h = getHouse(i)
	var payload = newPayload()
	payload.User = u
	payload.Message=message
	payload.Data = append(payload.Data, h)

	t, err := template.ParseFiles("templates/bid.tmpl", "templates/navbar.tmpl")
	check("Parse template", err)
	t.Execute(w, payload)
}


func applyHandler(w http.ResponseWriter, r *http.Request) {
	var u string
	session, err := store.Get(r, "BlockPartySession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u = session.Values["user"].(string)

	vars := mux.Vars(r)
	i := vars["contract-id"]

	var h House
	h = getHouse(i)
	t, err := template.ParseFiles("templates/apply.tmpl", "templates/navbar.tmpl")
	check("Parse template", err)
	var payload = newPayload()
	payload.Data = append(payload.Data, h)
	payload.User = u
	t.Execute(w, payload)
}

func mortgageHandler(w http.ResponseWriter, r *http.Request) {
	var u string
	session, err := store.Get(r, "BlockPartySession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u = session.Values["user"].(string)

	vars := mux.Vars(r)
	i := vars["contract-id"]

	var h House
	h = getHouse(i)
	t, err := template.ParseFiles("templates/mortgage.tmpl", "templates/navbar.tmpl")
	check("Parse template", err)
	var payload = newPayload()
	payload.Data = append(payload.Data, h)
	payload.User = u
	t.Execute(w, payload)
}

func realtorHandler(w http.ResponseWriter, r *http.Request) {
	var u string
	session, err := store.Get(r, "BlockPartySession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u = session.Values["user"].(string)

	t, err := template.ParseFiles("templates/realtor.tmpl", "templates/navbar.tmpl")
	check("Parse template", err)
	var payload = newPayload()
	payload.Data = getHouses()
	payload.User = u
	t.Execute(w, payload)
}

func myBidsHandler(w http.ResponseWriter, r *http.Request) {
	var u string
	session, err := store.Get(r, "BlockPartySession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u = session.Values["user"].(string)

	log.Print("User from form is : " + u)
	myBids := getMyBids(u)
	var payload = newBidPayload()
	payload.Data = myBids
	payload.User = u

	t, err := template.ParseFiles("templates/mybids.tmpl", "templates/navbar.tmpl")
	check("Parse template", err)
	t.Execute(w, payload)
}

func main() {
	initialize()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", listingsHandler)
	router.HandleFunc("/house/{contract-id}", detailsHandler)
	router.HandleFunc("/house/{contract-id}/bid", bidHandler)
	router.HandleFunc("/house/{contract-id}/bids", bidsHandler)
	router.HandleFunc("/house/{contract-id}/applyForMortgage", applyHandler)
	router.HandleFunc("/house/{contract-id}/mortgage", mortgageHandler)
	router.HandleFunc("/realtor", realtorHandler)
	router.HandleFunc("/mybids", myBidsHandler)

	http.Handle("/images/", http.FileServer(http.Dir("/app")))
	http.Handle("/css/", http.FileServer(http.Dir("/app")))
	http.Handle("/fonts/", http.FileServer(http.Dir("/app")))
	http.Handle("/js/", http.FileServer(http.Dir("/app")))
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
