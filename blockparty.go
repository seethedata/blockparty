//blockparty.go handles the webserver functions for the blockparty demo
package main

import (
	"encoding/json"
	"fmt"
	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/satori/go.uuid"
	"html/template"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	pool    *redis.Pool
	mainURL string
	store   = sessions.NewCookieStore([]byte("BlockParty"))
	maxGas  = big.NewInt(4700000)
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
	ID             string  `json:"id" redis:"id"`
	Name           string  `json:"name" redis:"name"`
	Owner          string  `json:"owner" redis:"owner"`
	Address        string  `json:"address" redis:"address"`
	Price          float64 `json:"price" redis:"price"`
	Image          string  `json:"image" redis:"image"`
	Contract       string  `json:"contract" redis:"contract"`
	Description    string  `json:"description" redis:"description"`
	Bedrooms       float64 `json:"bedrooms" redis:"bedrooms"`
	Bathrooms      float64 `json:"bathrooms" redis:"bathrooms"`
	Status         string  `json:"status" redis:"status"`
	Quality        int     `json:"quality" redis:"quality"`
	InspectionDate string  `json:"inspectionDate" redis:"inspectionDate"`
}

func newHouse() House {
	var h House
	h.ID = getNewHouseID()
	return h
}

// A Bid on a house
type Bid struct {
	UserID       string  `json:"userID" redis:"userID"`
	UserAddress  string  `json:"userAddress" redis:"userAddress"`
	User         User    `json:"user" redis:"user"`
	Amount       float64 `json:"amount" redis:"amount"`
	HouseID      string  `json:"houseID" redis:"houseID"`
	HouseAddress string  `json:"houseAddress" redis:"houseAddress"`
	Status       string  `json:"status" redis:"status"`
}

// A Mortgage on a house
type Mortgage struct {
	UserID        string  `json:"user" redis:"user"`
	Amount        float64 `json:"amount" redis:"amount"`
	HouseID       string  `json:"houseID" redis:"houseID"`
	Lender        string  `json:"lender" redis:"lender"`
	Appraisal     float64 `json:"appraisal" redis:"appraisal"`
	AppraisalDate string  `json:"appraisalDate" redis:"appraisalDate"`
	Status        string  `json:"status" redis:"status"`
	Override      string  `json:"override" redis:"override"`
}

func (m Mortgage) getKey() string {
	return "mortgage:" + m.HouseID + ":" + m.UserID
}

func (b Bid) getKey() string {
	return "bid:" + b.HouseID + ":" + b.UserID
}

//Payload is a generic Container to hold data for templates
type Payload struct {
	Houses     []House    `json:"data" redis:"data"`
	Bids       []Bid      `json:"bids" redis:"bids"`
	Mortgages  []Mortgage `json:"mortgages" redis:"mortgages"`
	Users      []User     `json:"users" redis:"users"`
	URL        string     `json:"URL" redis:"URL"`
	Parameters []byte     `json:"parameters" redis:"parameters"`
}

// User is a user from ethereum
type User struct {
	ID      string `json:"id" redis:"id"`
	Name    string `json:"name" redis:"name"`
	Address string `json:"address" redis:"address"`
	Type    string `json:"type" redis:"type"`
}

// Address is a contract address in ethereum
type Address struct {
	ID     string `json:"id" redis:"id"`
	UserID string `json:"userID" redis:"userID"`
	Status string `json:"status" redis:"status"`
}

// AddressList is a list of ethereum users
type AddressList struct {
	Data []Address `json:"data" redis:"data:`
}

// UserList is a list of ethereum users
type UserList struct {
	Data []User `json:"data" redis:"data"`
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

func newPayload() *Payload {
	return &Payload{URL: mainURL}
}

func getHouses() []House {
	var houses []House

	c := pool.Get()
	defer c.Close()

	n, err := redis.Strings(c.Do("KEYS", "house:*"))
	check("Keys", err)

	for _, v := range n {
		var h House
		r, err := redis.Values(c.Do("HGETALL", v))
		err = redis.ScanStruct(r, &h)
		check("ScanStruct", err)
		houses = append(houses, h)
	}
	return houses
}

func setDefaultHouses() {
	c := pool.Get()
	defer c.Close()

	n, err := redis.Strings(c.Do("KEYS", "house:*"))
	check("KEYS", err)
	if len(n) == 0 {
		fmt.Println("Creating default houses.")
		file, err := ioutil.ReadFile("./houses.json")
		check("Read houses JSON", err)

		var listings Payload
		err = json.Unmarshal(file, &listings)
		check("Unmarshal", err)

		for _, v := range listings.Houses {
			h := newHouse()
			_, err = c.Do("HMSET", "house:"+h.ID, "id", h.ID, "name", v.Name, "address", v.Address, "price",
				v.Price, "image", v.Image, "contract", v.Contract, "description", v.Description, "bedrooms", v.Bedrooms, "bathrooms", v.Bathrooms, "status", v.Status, "quality", v.Quality, "inspectionDate", "")
			check("HMSET", err)
		}
	} else {
		fmt.Println("Default houses already exist. Skipping house creation.")
	}
}

func setDefaultUsers() {
	c := pool.Get()
	defer c.Close()

	var users UserList
	n, err := redis.Strings(c.Do("KEYS", "users"))
	check("KEYS", err)
	if len(n) == 0 {
		fmt.Println("Creating default users.")
		file, err := ioutil.ReadFile("./users.json")
		check("Read houses JSON", err)

		err = json.Unmarshal(file, &users)
		check("Unmarshal", err)

		for _, v := range users.Data {
			key := "user:" + v.Name
			_, err = c.Do("HMSET", key, "address", v.Address, "type", "System")
			check("HSET", err)
		}
	} else {
		fmt.Println("Default system users already exist. Skipping system user creation.")
	}

}

func setDefaultAddresses() {
	c := pool.Get()
	defer c.Close()

	var addresses AddressList
	n, err := redis.Strings(c.Do("KEYS", "address:*"))
	check("KEYS", err)
	if len(n) == 0 {
		fmt.Println("Creating contracts addresses.")
		file, err := ioutil.ReadFile("./addresses.json")
		check("Read addresses JSON", err)

		err = json.Unmarshal(file, &addresses)
		check("Unmarshal", err)

		for _, v := range addresses.Data {
			key := "address:" + v.ID
			_, err = c.Do("HMSET", key, "id", v.ID, "status", "Unassigned")
			check("HSET", err)
		}
	} else {
		fmt.Println("Default addresses already exist. Skipping addresses creation.")
	}
}
func changeHouseStatus(i string, status string) error {
	c := pool.Get()
	defer c.Close()

	if status == "sold" {
		err := changeHousePrice(i, 0)
		check("ChangeHousePrice", err)
	}

	key := "house:" + i
	_, err := c.Do("HSET", key, "status", status)
	return err
}

func changeHousePrice(i string, price float64) error {
	c := pool.Get()
	defer c.Close()

	key := "house:" + i
	_, err := c.Do("HSET", key, "price", price)
	return err
}
func changeHouseQuality(i string, q int) error {
	c := pool.Get()
	defer c.Close()

	key := "house:" + i
	_, err := c.Do("HSET", key, "quality", q)
	return err
}
func changeBidStatus(i string, u string, status string) error {
	c := pool.Get()
	defer c.Close()

	key := "bid:" + i + ":" + u
	_, err := c.Do("HSET", key, "status", status)
	return err
}

func rejectOtherBids(i string, u string) error {
	var err error
	c := pool.Get()
	defer c.Close()

	bids := getBids(i + ":*")
	for _, v := range bids {
		if v.UserID != u {
			key := "bid:" + i + ":" + v.UserID
			_, err := c.Do("HSET", key, "status", "Rejected")
			check("HSET", err)
		}
	}
	return err
}

func changeMortgageStatus(i string, u string, status string) error {
	c := pool.Get()
	defer c.Close()

	key := "mortgage:" + i + ":" + u
	_, err := c.Do("HSET", key, "status", status)
	return err
}

func rejectOtherMortgages(i string, u string) error {
	var err error
	c := pool.Get()
	defer c.Close()

	mortgages := getMortgages(i + ":*")
	for _, v := range mortgages {
		if v.UserID != u {
			key := "mortgage:" + i + ":" + v.UserID
			_, err := c.Do("HSET", key, "status", "Rejected")
			check("HSET", err)
		}
	}
	return err
}

func setInspectionDate(i string, d string) error {
	c := pool.Get()
	defer c.Close()

	key := "house:" + i
	_, err := c.Do("HSET", key, "inspectionDate", d)
	return err
}

func setAppraisalDate(i string, u string, d string) error {
	c := pool.Get()
	defer c.Close()
	key := "mortgage:" + i + ":" + u
	_, err := c.Do("HSET", key, "appraisalDate", d)
	check("HSET", err)
	return err
}

func setAppraisal(i string, u string, a string) error {
	c := pool.Get()
	defer c.Close()
	key := "mortgage:" + i + ":" + u
	_, err := c.Do("HSET", key, "appraisal", a)
	check("HSET", err)
	return err
}
func initialize() {
	var cfServices cfServices
	fmt.Println("Starting")
	file, err := ioutil.ReadFile("./services.json")
	check("Read services JSON", err)

	err = json.Unmarshal(file, &cfServices)
	check("Unmarshal", err)

	env, _ := cfenv.Current()
	mainURL = "http://" + env.ApplicationURIs[0]
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

	store.Options = &sessions.Options{MaxAge: 0}
	pool = newPool(host, port, password)
	setDefaultHouses()
	setDefaultUsers()
	setDefaultAddresses()
}

func getNewHouseID() string {
	return uuid.NewV4().String()
}

func getUser(i string) User {
	var user User
	c := pool.Get()
	defer c.Close()

	key := "user:" + i
	n, err := redis.Values(c.Do("HGETALL", key))
	check("HGETALL", err)
	err = redis.ScanStruct(n, &user)
	check("ScanStruct", err)
	return user
}

func getNewUserID() string {
	return uuid.NewV4().String()
}

func createUser(u string) error {
	c := pool.Get()
	defer c.Close()

	addresses, err := redis.Strings(c.Do("KEYS", "address:*"))
	check("Strings", err)

	var a Address
	for _, v := range addresses {
		n, err := redis.Values(c.Do("HGETALL", v))
		check("HGETALL", err)
		err = redis.ScanStruct(n, &a)
		if a.Status == "Unassigned" {
			_, err := c.Do("HMSET", v, "userID", u, "status", "Assigned")
			check("HMSET", err)
			key := "user:" + u
			_, err = c.Do("HMSET", key, "id", u, "address", a.ID, "type", "User")
			check("HMSET", err)
			break
		}
	}

	if getUser(u).Address == "" {
		return fmt.Errorf(" - No addresses available for new user")
	}
	return err

}

func getHouse(i string) House {
	var h House
	c := pool.Get()
	defer c.Close()

	n, err := redis.Values(c.Do("HGETALL", "house:"+i))
	check("HGETALL", err)
	err = redis.ScanStruct(n, &h)
	check("ScanStruct", err)
	return h
}

func getBid(ID string) Bid {
	var bid Bid
	c := pool.Get()
	defer c.Close()

	n, err := redis.Values(c.Do("HGETALL", ID))
	check("HGETALL", err)
	err = redis.ScanStruct(n, &bid)
	check("ScanStruct", err)
	bid.User = getUser(bid.UserID)
	return bid
}

func getBids(filter string) []Bid {
	var bids []Bid

	c := pool.Get()
	defer c.Close()

	bidKeys, err := redis.Strings(c.Do("KEYS", "bid:"+filter))
	check("Strings", err)

	for _, v := range bidKeys {
		bids = append(bids, getBid(v))
	}

	return bids
}

func getHouseBids(i string) []Bid {
	return getBids(i + ":*")
}

func getMyBids(u string) []Bid {
	return getBids("*:" + u)
}

func getMyBid(i string, u string) Bid {
	return getBid("bid:" + i + ":" + u)

}

func deleteAllHouseBids(i string) error {
	var err error
	c := pool.Get()
	defer c.Close()

	bids := getHouseBids(i)

	for _, b := range bids {
		_, err = c.Do("DEL", b.getKey())
		check("DEL", err)
	}
	return err
}

func getMortgage(ID string) Mortgage {
	var mortgage Mortgage
	c := pool.Get()
	defer c.Close()

	n, err := redis.Values(c.Do("HGETALL", ID))
	check("HGETALL", err)
	err = redis.ScanStruct(n, &mortgage)
	check("ScanStruct", err)
	return mortgage
}

func getMortgages(filter string) []Mortgage {
	var mortgages []Mortgage

	c := pool.Get()
	defer c.Close()

	mortgageKeys, err := redis.Strings(c.Do("KEYS", "mortgage:"+filter))
	check("Strings", err)

	for _, v := range mortgageKeys {
		mortgages = append(mortgages, getMortgage(v))
	}

	return mortgages
}

func getMyMortgages(u string) []Mortgage {
	return getMortgages("*:" + u)
}

func getMyMortgage(i string, u string) Mortgage {
	return getMortgage("mortgage:" + i + ":" + u)
}

func getHouseMortgages(i string) []Mortgage {
	return getMortgages(i + ":*")
}

func deleteAllHouseMortgages(i string) error {
	var err error
	c := pool.Get()
	defer c.Close()

	mortgages := getHouseMortgages(i)

	for _, m := range mortgages {
		_, err = c.Do("DEL", m.getKey())
		check("DEL", err)
	}
	return err
}

func addressListHandler(w http.ResponseWriter, r *http.Request) {
	c := pool.Get()
	defer c.Close()

	keys, err := redis.Strings(c.Do("KEYS", "address:*"))
	check("Strings", err)

	var a Address
	var al AddressList
	for _, k := range keys {
		r, err := c.Do("HGETALL", k)
		err = redis.ScanStruct(r.([]interface{}), &a)
		check("ScanStruct", err)
		if a.Status == "Assigned" {
			al.Data = append(al.Data, a)
		}
	}
	t, err := template.ParseFiles("templates/addressList.tmpl")
	check("Parse template", err)
	t.Execute(w, al)
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
		user := getUser(u)
		if user.ID == "" {
			session.Options = &sessions.Options{MaxAge: 0}
			err = createUser(u)
			check("createUser", err)
		}
	} else {
		u = getNewUserID()
		session.Values["user"] = u
		err = createUser(u)
		check("createUser", err)
	}

	session.Save(r, w)
	t, err := template.ParseFiles("templates/listings.tmpl", "templates/housePanel.tmpl", "templates/head.tmpl", "templates/navbar.tmpl")
	check("Parse template", err)
	listings := newPayload()
	listings.Houses = getHouses()
	listings.Users = append(listings.Users, getUser(u))
	check("Marshal", err)
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
	i := vars["houseID"]

	c := pool.Get()
	defer c.Close()

	var h House
	h = getHouse(i)

	var payload = newPayload()
	payload.Houses = append(payload.Houses, h)
	payload.Users = append(payload.Users, getUser(u))
	t, err := template.ParseFiles("templates/details.tmpl", "templates/head.tmpl", "templates/navbar.tmpl")
	check("Parse template", err)
	t.Execute(w, payload)
}

func listHouseHandler(w http.ResponseWriter, r *http.Request) {
	var u string
	session, err := store.Get(r, "BlockPartySession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u = session.Values["user"].(string)

	vars := mux.Vars(r)
	i := vars["houseID"]

	c := pool.Get()
	defer c.Close()

	var h House
	h = getHouse(i)

	var payload = newPayload()
	payload.Houses = append(payload.Houses, h)
	payload.Users = append(payload.Users, getUser("Seller"))
	payload.Users = append(payload.Users, getUser(u))
	t, err := template.ParseFiles("templates/listHouse.tmpl", "templates/head.tmpl", "templates/navbar.tmpl")
	check("Parse template", err)
	t.Execute(w, payload)
}

func bidsHandler(w http.ResponseWriter, r *http.Request) {
	var u string
	vars := mux.Vars(r)
	i := vars["houseID"]

	session, err := store.Get(r, "BlockPartySession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u = session.Values["user"].(string)

	payload := newPayload()
	payload.Bids = getHouseBids(i)
	payload.Users = append(payload.Users, getUser(u))
	payload.Users = append(payload.Users, getUser("Seller"))
	payload.Houses = append(payload.Houses, getHouse(i))
	t, err := template.ParseFiles("templates/bids.tmpl", "templates/head.tmpl", "templates/navbar.tmpl")
	check("Parse template", err)
	t.Execute(w, payload)

}

func enterBidHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "BlockPartySession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u := session.Values["user"].(string)

	vars := mux.Vars(r)
	i := vars["houseID"]
	a := r.PostFormValue("bidAmount")
	a = strings.Replace(a, ",", "", -1)
	a = strings.Replace(a, ".", "", -1)

	var h House
	h = getHouse(i)

	c := pool.Get()
	defer c.Close()

	key := "bid:" + i + ":" + u
	_, err = c.Do("HMSET", key, "userID", u, "userAddress", getUser(u).Address, "amount", a, "houseID", i, "houseAddress", h.Contract, "status", "Submitted")
	check("HMSET", err)

	http.Redirect(w, r, mainURL+"/house/"+i+"/myBid", http.StatusFound)
}

func getSigner(u string) (*bind.TransactOpts, error) {
	var uaddr = common.HexToAddress(getUser(u).Address)
	am := accounts.NewManager("keystore", accounts.StandardScryptN, accounts.StandardScryptP)
	ac, err := am.Find(accounts.Account{Address: uaddr})
	if err != nil {
		log.Fatalf("Account not found: %v", err)
	}
	accountJSON, err := am.Export(ac, "password01", "blockparty")
	if err != nil {
		log.Fatalf("Account JSON failed:  %v", err)
	}
	key, err := accounts.DecryptKey(accountJSON, "blockparty")
	if err != nil {
		log.Fatalf("Decrypt key failed: %v", err)
	}
	signer := bind.NewKeyedTransactor(key.PrivateKey)
	return signer, err
}

func enterListingHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "BlockPartySession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u := session.Values["user"].(string)
	vars := mux.Vars(r)
	i := vars["houseID"]
	a := r.PostFormValue("askingPrice")
	a = strings.Replace(a, ",", "", -1)
	a = strings.Replace(a, ".", "", -1)

	conn, err := ethclient.Dial("http://54.245.138.237:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	h := getHouse(i)

	var addr = common.HexToAddress(h.Contract)
	c, err := NewContract(addr, conn)
	if err != nil {
		log.Fatalf("Failed to bind to contract: %v", err)
	}

	askingPrice := new(big.Int)
	askingPrice.SetString(a, 10)

	signer, err := getSigner(u)
	if err != nil {
		log.Fatalf("getSigner failed: %v", err)
	}
	signer.GasLimit = maxGas
	_, err = c.ForSale(signer, askingPrice)
	if err != nil {
		log.Fatalf("ForSale failed: %v", err)
	}

	ap, err := strconv.ParseFloat(a, 64)
	err = changeHousePrice(i, ap)
	check("changeHousePrice", err)
	err = changeHouseStatus(i, "Listed")
	check("changeHouseStatus", err)
	err = deleteAllHouseBids(i)
	check("deleteAllHouseBids", err)
	err = deleteAllHouseMortgages(i)
	check("deleteAllHouseMortages", err)
	http.Redirect(w, r, mainURL+"/seller", http.StatusFound)
}

func delistHouseHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "BlockPartySession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u := session.Values["user"].(string)
	vars := mux.Vars(r)
	i := vars["houseID"]
	a := r.PostFormValue("askingPrice")
	a = strings.Replace(a, ",", "", -1)
	a = strings.Replace(a, ".", "", -1)

	conn, err := ethclient.Dial("http://54.245.138.237:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	h := getHouse(i)

	var addr = common.HexToAddress(h.Contract)
	c, err := NewContract(addr, conn)
	if err != nil {
		log.Fatalf("Failed to bind to contract: %v", err)
	}
	signer, err := getSigner(u)
	if err != nil {
		log.Fatalf("getSigner failed: %v", err)
	}
	log.Printf("maxGas is %v\n", maxGas)
	signer.GasLimit = maxGas
	_, err = c.NotForSale(signer)
	if err != nil {
		log.Fatalf("NotForSale failed: %v", err)
	}
	err = changeHouseStatus(i, "Sold")
	check("changeHouseStatus", err)
	http.Redirect(w, r, mainURL+"/seller", http.StatusFound)
}
func enterMortgageHandler(w http.ResponseWriter, r *http.Request) {
	var u string
	session, err := store.Get(r, "BlockPartySession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u = session.Values["user"].(string)

	vars := mux.Vars(r)
	i := vars["houseID"]
	a := r.PostFormValue("amount")
	a = strings.Replace(a, ",", "", -1)
	a = strings.Replace(a, ".", "", -1)
	l := r.PostFormValue("lender")

	c := pool.Get()
	defer c.Close()

	key := "mortgage:" + i + ":" + u

	_, err = c.Do("HMSET", key, "user", u, "amount", a, "houseID", i, "lender", l, "status", "Submitted")
	check("HMSET", err)
	http.Redirect(w, r, mainURL+"/house/"+i+"/mortgage/myMortgage", http.StatusFound)
}

func updateMortgageAmountHandler(w http.ResponseWriter, r *http.Request) {
	_, err := store.Get(r, "BlockPartySession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	i := vars["houseID"]
	a := r.PostFormValue("amount")
	a = strings.Replace(a, ",", "", -1)
	a = strings.Replace(a, ".", "", -1)
	um := r.PostFormValue("user")

	c := pool.Get()
	defer c.Close()

	key := "mortgage:" + i + ":" + um

	_, err = c.Do("HMSET", key, "amount", a, "override", "Yes")
	check("HSET", err)
	http.Redirect(w, r, mainURL+"/lender", http.StatusFound)
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
	i := vars["houseID"]

	var h House
	h = getHouse(i)
	t, err := template.ParseFiles("templates/apply.tmpl", "templates/head.tmpl", "templates/navbar.tmpl")
	check("Parse template", err)
	var payload = newPayload()
	payload.Houses = append(payload.Houses, h)
	payload.Users = append(payload.Users, getUser(u))
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
	i := vars["houseID"]

	var h House
	h = getHouse(i)
	t, err := template.ParseFiles("templates/mortgage.tmpl", "templates/head.tmpl", "templates/navbar.tmpl")
	check("Parse template", err)
	var payload = newPayload()
	payload.Houses = append(payload.Houses, h)
	payload.Users = append(payload.Users, getUser(u))
	t.Execute(w, payload)
}

func myBidHandler(w http.ResponseWriter, r *http.Request) {
	var u string

	session, err := store.Get(r, "BlockPartySession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u = session.Values["user"].(string)

	vars := mux.Vars(r)
	i := vars["houseID"]
	myBid := getMyBid(i, u)
	house := getHouse(i)

	t, err := template.ParseFiles("templates/myBid.tmpl", "templates/head.tmpl", "templates/navbar.tmpl")
	check("Parse template", err)
	var payload = newPayload()
	payload.Houses = append(payload.Houses, house)
	payload.Users = append(payload.Users, getUser(u))
	payload.Bids = append(payload.Bids, myBid)
	t.Execute(w, payload)
}

func myMortgageHandler(w http.ResponseWriter, r *http.Request) {
	var u string

	session, err := store.Get(r, "BlockPartySession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u = session.Values["user"].(string)

	vars := mux.Vars(r)
	i := vars["houseID"]
	myMortgage := getMyMortgage(i, u)
	house := getHouse(i)

	t, err := template.ParseFiles("templates/myMortgage.tmpl", "templates/head.tmpl", "templates/navbar.tmpl")
	check("Parse template", err)
	var payload = newPayload()
	payload.Houses = append(payload.Houses, house)
	payload.Users = append(payload.Users, getUser(u))
	payload.Mortgages = append(payload.Mortgages, myMortgage)
	t.Execute(w, payload)
}

func inspectHandler(w http.ResponseWriter, r *http.Request) {
	var u string
	session, err := store.Get(r, "BlockPartySession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u = session.Values["user"].(string)

	vars := mux.Vars(r)
	i := vars["houseID"]

	var h House
	h = getHouse(i)
	t, err := template.ParseFiles("templates/inspect.tmpl", "templates/head.tmpl", "templates/navbar.tmpl")
	check("Parse template", err)
	var payload = newPayload()
	payload.Houses = append(payload.Houses, h)
	payload.Users = append(payload.Users, getUser(u))
	t.Execute(w, payload)
}

func appraiseHandler(w http.ResponseWriter, r *http.Request) {
	var u string
	session, err := store.Get(r, "BlockPartySession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u = session.Values["user"].(string)

	vars := mux.Vars(r)
	i := vars["houseID"]
	um := vars["user"]

	var h House
	h = getHouse(i)
	t, err := template.ParseFiles("templates/appraise.tmpl", "templates/head.tmpl", "templates/navbar.tmpl")
	check("Parse template", err)
	var payload = newPayload()
	mortgages := getHouseMortgages(i)

	for _, m := range mortgages {
		if m.UserID == um {
			payload.Mortgages = append(payload.Mortgages, m)
		}
	}

	payload.Houses = append(payload.Houses, h)
	payload.Users = append(payload.Users, getUser(u))
	t.Execute(w, payload)
}

func lenderHandler(w http.ResponseWriter, r *http.Request) {
	var u string
	session, err := store.Get(r, "BlockPartySession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u = session.Values["user"].(string)

	t, err := template.ParseFiles("templates/lender.tmpl", "templates/head.tmpl", "templates/navbar.tmpl")
	check("Parse template", err)
	var payload = newPayload()
	payload.Houses = getHouses()
	payload.Mortgages = getMortgages("*")
	payload.Users = append(payload.Users, getUser(u))
	t.Execute(w, payload)
}

func scheduleAppraisalHandler(w http.ResponseWriter, r *http.Request) {
	var u string
	session, err := store.Get(r, "BlockPartySession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u = session.Values["user"].(string)

	vars := mux.Vars(r)
	i := vars["houseID"]
	mu := vars["user"]

	var h House
	h = getHouse(i)
	t, err := template.ParseFiles("templates/scheduleAppraisal.tmpl", "templates/head.tmpl", "templates/navbar.tmpl")
	check("Parse template", err)
	var payload = newPayload()
	payload.Houses = append(payload.Houses, h)
	payload.Users = append(payload.Users, getUser(u))
	payload.Mortgages = append(payload.Mortgages, getMortgage("mortgage:"+i+":"+mu))
	t.Execute(w, payload)
}

func scheduleInspectionHandler(w http.ResponseWriter, r *http.Request) {
	var u string
	session, err := store.Get(r, "BlockPartySession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u = session.Values["user"].(string)

	vars := mux.Vars(r)
	i := vars["houseID"]

	var h House
	h = getHouse(i)
	err = changeHouseQuality(i, 0)
	check("changeHouseQuality", err)
	t, err := template.ParseFiles("templates/scheduleInspection.tmpl", "templates/head.tmpl", "templates/navbar.tmpl")
	check("Parse template", err)
	var payload = newPayload()
	payload.Houses = append(payload.Houses, h)
	payload.Users = append(payload.Users, getUser(u))
	t.Execute(w, payload)
}

func enterInspectionAppointmentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i := vars["houseID"]
	d := r.PostFormValue("date")

	err := setInspectionDate(i, d)
	check("setInspectionDate", err)
	http.Redirect(w, r, mainURL+"/myMortgages", http.StatusFound)
}

func enterAppraisalAppointmentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i := vars["houseID"]
	d := r.PostFormValue("date")
	u := r.PostFormValue("user")

	err := setAppraisalDate(i, u, d)
	check("setAppraisalDate", err)
	http.Redirect(w, r, mainURL+"/lender", http.StatusFound)
}

func enterAppraisalHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i := vars["houseID"]
	a := r.PostFormValue("amount")
	a = strings.Replace(a, ",", "", -1)
	a = strings.Replace(a, ".", "", -1)
	log.Print(a)
	u := r.PostFormValue("user")

	err := setAppraisal(i, u, a)
	check("setAppraisal", err)
	http.Redirect(w, r, mainURL+"/appraiser", http.StatusFound)
}

func changeHouseStatusHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i := vars["houseID"]
	s := vars["status"]

	err := changeHouseStatus(i, s)
	check("changeHouseStatus", err)
	http.Redirect(w, r, mainURL+"/seller", http.StatusFound)
}

func changeHouseQualityHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i := vars["houseID"]
	q, err := strconv.Atoi(vars["quality"])

	err = changeHouseQuality(i, q)
	check("changeHouseStatus", err)
	http.Redirect(w, r, mainURL+"/inspector", http.StatusFound)
}

func changeBidStatusHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i := vars["houseID"]
	u := vars["user"]
	s := vars["status"]

	err := changeBidStatus(i, u, s)
	check("changeBidStatus", err)

	if s == "Accepted" {
		err := rejectOtherBids(i, u)
		check("rejectOtherBids", err)
		check("changeHouseStatus", err)
	}
	http.Redirect(w, r, mainURL+"/seller", http.StatusFound)
}

func changeMortgageStatusHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i := vars["houseID"]
	u := vars["user"]
	s := vars["status"]

	err := changeMortgageStatus(i, u, s)
	check("changeMortgageStatus", err)

	if s == "Accepted" {
		err := rejectOtherMortgages(i, u)
		check("rejectOtherMortgages", err)
		err = changeHouseStatus(i, "Sold")
		check("changeHouseStatus", err)
	}
	http.Redirect(w, r, mainURL+"/lender", http.StatusFound)
}

func checkBidStatusHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i := vars["houseID"]
	u := vars["user"]

	key := "bid:" + i + ":" + u
	bid := getBid(key)
	response, err := json.Marshal(bid)
	check("Marshal", err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func checkMortgageStatusHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i := vars["houseID"]
	u := vars["user"]

	key := "mortgage:" + i + ":" + u
	mortgage := getMortgage(key)
	response, err := json.Marshal(mortgage)
	check("Marshal", err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func appraiserHandler(w http.ResponseWriter, r *http.Request) {
	var u string
	session, err := store.Get(r, "BlockPartySession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u = session.Values["user"].(string)

	t, err := template.ParseFiles("templates/appraiser.tmpl", "templates/head.tmpl", "templates/navbar.tmpl")
	check("Parse template", err)
	var payload = newPayload()
	for _, m := range getMortgages("*") {
		if m.AppraisalDate != "" {
			payload.Mortgages = append(payload.Mortgages, m)
			payload.Houses = append(payload.Houses, getHouse(m.HouseID))
		}
	}
	payload.Users = append(payload.Users, getUser(u))
	t.Execute(w, payload)
}

func inspectorHandler(w http.ResponseWriter, r *http.Request) {
	var u string
	session, err := store.Get(r, "BlockPartySession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u = session.Values["user"].(string)

	t, err := template.ParseFiles("templates/inspector.tmpl", "templates/head.tmpl", "templates/navbar.tmpl")
	check("Parse template", err)
	var payload = newPayload()
	for _, h := range getHouses() {
		if h.Quality == 0 {
			payload.Houses = append(payload.Houses, h)
		}
	}
	payload.Users = append(payload.Users, getUser(u))
	t.Execute(w, payload)
}
func sellerHandler(w http.ResponseWriter, r *http.Request) {
	var u string
	session, err := store.Get(r, "BlockPartySession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u = session.Values["user"].(string)

	t, err := template.ParseFiles("templates/seller.tmpl", "templates/head.tmpl", "templates/navbar.tmpl")
	check("Parse template", err)
	var payload = newPayload()
	payload.Houses = getHouses()
	payload.Users = append(payload.Users, getUser("Seller"))
	payload.Users = append(payload.Users, getUser(u))
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

	myBids := getMyBids(u)
	var payload = newPayload()
	payload.Bids = myBids
	payload.Users = append(payload.Users, getUser(u))

	t, err := template.ParseFiles("templates/myBids.tmpl", "templates/head.tmpl", "templates/navbar.tmpl")
	check("Parse template", err)
	t.Execute(w, payload)
}

func myMortgagesHandler(w http.ResponseWriter, r *http.Request) {
	var u string
	session, err := store.Get(r, "BlockPartySession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u = session.Values["user"].(string)

	myMortgages := getMyMortgages(u)
	var payload = newPayload()
	payload.Mortgages = myMortgages
	payload.Users = append(payload.Users, getUser(u))

	for _, m := range myMortgages {
		payload.Houses = append(payload.Houses, getHouse(m.HouseID))
	}
	t, err := template.ParseFiles("templates/myMortgages.tmpl", "templates/head.tmpl", "templates/navbar.tmpl")
	check("Parse template", err)
	t.Execute(w, payload)
}

func missingRequirementsHandler(w http.ResponseWriter, r *http.Request) {
	payload := newPayload()
	t, err := template.ParseFiles("templates/missingRequirements.tmpl")
	check("Parse template", err)
	t.Execute(w, payload)

}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	var keys []string
	keys = append(keys, "house:", "bid:", "mortgage:", "address:", "user:")

	c := pool.Get()
	defer c.Close()

	for _, k := range keys {
		n, err := redis.Strings(c.Do("KEYS", k+"*"))
		check("Keys", err)
		for _, v := range n {
			_, err := c.Do("DEL", v)
			check("DEL", err)
			//log.Print("Deleted key: " + v)
		}
	}

	setDefaultHouses()
	setDefaultUsers()
	setDefaultAddresses()
	http.Redirect(w, r, mainURL, http.StatusFound)
}

func getHouseInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i := vars["houseID"]
	house := getHouse(i)
	log.Print(house)
	log.Print(house.Contract)
	response, err := json.Marshal(house)
	check("Marshal", err)
	log.Print(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func main() {
	initialize()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", listingsHandler)
	router.HandleFunc("/seller", sellerHandler)
	router.HandleFunc("/house/{houseID}", detailsHandler)
	router.HandleFunc("/house/{houseID}/listHouse", listHouseHandler)
	router.HandleFunc("/house/{houseID}/info", getHouseInfoHandler)
	router.HandleFunc("/house/{houseID}/enterListing", enterListingHandler)
	router.HandleFunc("/house/{houseID}/delistHouse", delistHouseHandler)
	router.HandleFunc("/house/{houseID}/enterBid", enterBidHandler)
	router.HandleFunc("/house/{houseID}/myBid", myBidHandler)
	router.HandleFunc("/myBids", myBidsHandler)
	router.HandleFunc("/house/{houseID}/bids", bidsHandler)
	router.HandleFunc("/lender", lenderHandler)
	router.HandleFunc("/house/{houseID}/applyForMortgage", applyHandler)
	router.HandleFunc("/house/{houseID}/enterMortgage", enterMortgageHandler)
	router.HandleFunc("/house/{houseID}/mortgage/myMortgage", myMortgageHandler)
	router.HandleFunc("/house/{houseID}/mortgage/{user}/changeStatus/{status}", changeMortgageStatusHandler)
	router.HandleFunc("/house/{houseID}/mortgage/{user}/checkStatus", checkMortgageStatusHandler)
	router.HandleFunc("/appraiser", appraiserHandler)
	router.HandleFunc("/house/{houseID}/mortgage/{user}/scheduleAppraisal", scheduleAppraisalHandler)
	router.HandleFunc("/house/{houseID}/mortgage/{user}/enterAppraisalAppointment", enterAppraisalAppointmentHandler)
	router.HandleFunc("/house/{houseID}/mortgage/{user}/appraise", appraiseHandler)
	router.HandleFunc("/house/{houseID}/mortgage/{user}/enterAppraisal", enterAppraisalHandler)
	router.HandleFunc("/house/{houseID}/mortgage/{user}/updateAmount", updateMortgageAmountHandler)
	router.HandleFunc("/house/{houseID}/inspect", inspectHandler)
	router.HandleFunc("/myMortgages", myMortgagesHandler)
	router.HandleFunc("/inspector", inspectorHandler)
	router.HandleFunc("/house/{houseID}/scheduleInspection", scheduleInspectionHandler)
	router.HandleFunc("/house/{houseID}/enterInspectionAppointment", enterInspectionAppointmentHandler)
	router.HandleFunc("/house/{houseID}/changeStatus/{status}", changeHouseStatusHandler)
	router.HandleFunc("/house/{houseID}/changeQuality/{quality}", changeHouseQualityHandler)
	router.HandleFunc("/house/{houseID}/bid/{user}/changeStatus/{status}", changeBidStatusHandler)
	router.HandleFunc("/house/{houseID}/bid/{user}/checkStatus", checkBidStatusHandler)
	router.HandleFunc("/missingRequirements", missingRequirementsHandler)
	router.HandleFunc("/reset", resetHandler)
	router.HandleFunc("/addressList", addressListHandler)

	http.Handle("/images/", http.FileServer(http.Dir("/app")))
	http.Handle("/css/", http.FileServer(http.Dir("/app")))
	http.Handle("/fonts/", http.FileServer(http.Dir("/app")))
	http.Handle("/js/", http.FileServer(http.Dir("/app")))
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
