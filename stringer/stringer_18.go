// https://tour.golang.org/methods/18
package main

import "fmt"

type IPAddr [4]byte

func (i IPAddr) String() string {
	var a string
	for _, b := range i {
		a += fmt.Sprintf("%d.", b)
	}
	return a
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
