// Client 

package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

// port of service where we send requests, this case its 880 for http
var PORT = 8080	

type item struct{
var name string
var descr string
var minBid int
var currBid int

}

// Add an item address to our email list
func add(scanner *bufio.Scanner) {
	fmt.Print("Please enter the your name: ")
	scanner.Scan()
	user := scanner.Text()
	fmt.Print("Please enter the item name: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Print("Please enter your bid: ")
	scanner.Scan()
	minB := scanner.Text()
	currB := minB
	fmt.Print("Please enter description: ")
	scanner.Scan()
	descr := scanner.Text() + "{" + name + minB + "}"
	
	item := item{ name, descr, minB, currB}
	
	//constructing url, url.PathEscape("string") to make it url safe
	request := fmt.Sprintf("http://localhost:%v/add?name=%v&item=%v", PORT, url.PathEscape(name), url.PathEscape(item))
	_, err := http.Get(request)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
}

// Lookup an item in our item list -- names are assumed to be unique
func lookup(scanner *bufio.Scanner) {
	fmt.Print("Please enter the item: ")
	scanner.Scan()
	item := scanner.Text()
	request := fmt.Sprintf("http://localhost:8080/lookup?item=%v", url.PathEscape(item))
	resp, err := http.Get(request)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	name := string(body)
	if "" == name {
		fmt.Printf("%v is not in the item list\n", name)
	} else {
		fmt.Printf("The item name is: %v\n", string(body))
	}
}

// bid an item in our item list -- names are assumed to be unique
func bid(scanner *bufio.Scanner) {
	fmt.Print("Please enter your name: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Print("Please enter the item's name: ")
	scanner.Scan()
	item := scanner.Text()
	fmt.Print("Please enter the bid: ")
	scanner.Scan()
	num := scanner.Text()
	request := fmt.Sprintf("http://localhost:%v/bid?name=%v&item=%vnum=%v", PORT, url.PathEscape(name), url.PathEscape(item), url.PathEscape(num))
	resp, err := http.Get(request)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	item := struct(body)
	if "" == item.name {
		fmt.Printf("%v is not in the item list\n", name)
	}
	//the actual bidding has already been decided on the host side
	if num < item.currBid {
		fmt.Printf("%v is lower than current bid of %v\n", num, item.currBid)
	}
	if num < item.minBid {
		fmt.Printf("%v is lower than current bid of %v\n", num, item.minBid)
	}
	else {
		fmt.Printf(item.descr)
	}
}

func main() {
	//client := new(http.Client)

	fmt.Println("Biddings\n")

	var cmd string
	scanner := bufio.NewScanner(os.Stdin)

	quit := "no"
	for quit == "no" {
		fmt.Printf("Add, Bid, or Lookup (a/b/l)? ")
		fmt.Scanf("%s ", &cmd)
		if "a" == cmd {
			add(scanner)
		} else "b" == cmd{
			bid(scanner)
		} else if "l" == cmd {
			lookup(scanner)
		}

		fmt.Print("Do you want to quit? (yes/no): ")
		fmt.Scanf("%s ", &quit)
	}

}
