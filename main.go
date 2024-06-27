package main

import (
	"log"

	"github.com/HMasataka/go-plugin/shared"
)

func main() {
	greeters := shared.NewManager("*", "./bin", shared.NewGreeterPluginClient())
	defer greeters.Dispose()

	if err := greeters.Init(); err != nil {
		log.Fatal(err)
	}

	if err := greeters.Launch(); err != nil {
		log.Fatal(err)
	}

	for _, pluginName := range []string{"foo", "hello"} {
		p, err := greeters.Get(pluginName)
		if err != nil {
			log.Fatal(err)
		}

		resp, err := p.(shared.Plugin).Greet(pluginName)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("\n\n%s plugin gives me: %s\n\n", pluginName, resp)
	}
}
