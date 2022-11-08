package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	u, err := url.Parse(s)

	if err != nil {
		panic(err)
	}

	fmt.Println(u)
	fmt.Println("Scheme: ", u.Scheme)
	fmt.Println("User: ", u.User)
	fmt.Println("Username: ", u.User.Username())
	up, set := u.User.Password()
	if set {
		fmt.Println("Password: ", up)
	}
	fmt.Println("Host: ", u.Host)
	fmt.Println("Hostname: ", u.Hostname())
	_, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("Port: ", port)
	fmt.Println("Escaped path: ", u.EscapedPath())
	fmt.Println("Path: ", u.Path)
	fmt.Println("Raw path: ", u.RawPath)
	fmt.Println("Raw query: ", u.RawQuery)
	fmt.Println("Query map: ", u.Query())
	fmt.Println("Fragment: ", u.Fragment)
	fmt.Println("Raw fragment: ", u.RawFragment)
	fmt.Println("Escaped fragment: ", u.EscapedFragment())
}
