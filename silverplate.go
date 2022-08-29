package main

import (
	"fmt"
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
)

func main() {
	fmt.Print(
		"_________________________________________________________________\n" +
			"      __                              ____                       \n" +
			"    /    )   ,   /                    /    )   /                 \n" +
			"----\\-----------/---------__---)__---/____/---/----__--_/_----__-\n" +
			"     \\     /   /   | /  /___) /   ) /        /   /   ) /    /___)\n" +
			"_(____/___/___/____|/__(___ _/_____/________/___(___(_(_ __(___ _\n" +
			"                                                                 \n" +
			"                                                                 \n\n")

	config := ParseConfig()
	ruleConfig := ParseRules()

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = config.VerboseMode
	proxy.KeepDestinationHeaders = config.KeepDestinationHeaders
	proxy.KeepHeader = config.KeepHeader

	fmt.Println("verbose mode set to " + config.VerboseModeToString())

	ruleConfig.ApplyRules(proxy)
	if RuleCount() > 0 {
		fmt.Println(RuleCountToString() + " rejecting rules parsed and applied")
	} else {
		fmt.Println("No rejecting rules parsed, accepting all requests")
	}

	if config.HandleNonProxyRequests {
		proxy.NonproxyHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			if req.Host == "" {
				_, err := fmt.Fprintln(w, "silverplate: cannot handle requests without host header")
				if err != nil {
					panic(err)
				}
				return
			}
			req.URL.Scheme = "http"
			req.URL.Host = req.Host
			proxy.ServeHTTP(w, req)
		})
	}

	fmt.Println("HTTP/S proxy listening on :" + config.PortToString())
	log.Fatalln(http.ListenAndServe(":"+config.PortToString(), proxy))
}
