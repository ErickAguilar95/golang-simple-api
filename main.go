package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ifaces, err := net.Interfaces()
		if err != nil {
			panic(err)
		}

		var response Response
		var ip string

		for _, i := range ifaces {
			addrs, err := i.Addrs()
			if err != nil {
				panic(err)
			}

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						ip = fmt.Sprintf("Hello response by %s", ipnet.IP.String())
					}
				}
			}
		}

		response.Message = ip
		fmt.Println(response)
		json.NewEncoder(w).Encode(response)
	})

	http.ListenAndServe(":8000", nil)
}
