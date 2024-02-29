/*
service
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

var validPath = regexp.MustCompile(
^/(add|bid|lookup)$`)

// port of service, this case http
var PORT = 8080	

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}

type item struct{
var name string
var descr string
var minBid int
var currBid int

}

// Where items are stored
// In a real service, this should be a database
// so that data is saved in case service shuts down
var itemTable map[string]item

func addHandler(w http.ResponseWriter, r *http.Request) {
//get the info for a new item
	name := r.URL.Query().Get("name")
	//pass a struct here
	i := r.URL.Query().Get("item")

	itemTable[name] = i
}

// Supports looking up item
func lookupHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	
	n := itemTable[name].name

	fmt.Fprint(w, n)
}

// Supports bidding
func bidHandler(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("name")
	item := r.URL.Query().Get("item")
	bid := r.URL.Query().Get("num")
	 
	
	i := itemTable[item]
	if bid > i.minBid && bid > i.currBid {
	i.currBid = bid
	// try regex to change the descr within the {}, announce who bid how much
	re = regexp.MustCompile(`{.*}`)
    fmt.Println(re.ReplaceAllString(i.descr, "{" + user + " " + bid + " }"))
	fmt.Fprint(w, i)
}

func main() {
	itemTable = make(map[string]item)
	http.HandleFunc("/lookup", makeHandler(lookupHandler))
	http.HandleFunc("/add", makeHandler(addHandler))
	http.HandleFunc("/bit", makeHandler(bidHandler))
	
	log.Fatal(http.ListenAndServe(PORT_STR, nil))
}
