package main

import (
	"log"

	"github.com/HMasataka/go-plugin/shared"
)

func main() {
	greeters := shared.NewManager("*", "./bin", &shared.GreeterPlugin{})
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
			log.Fatal(err.Error())
		}
		log.Printf("\n\n%s plugin gives me: %s\n\n", pluginName, p.(shared.Greeter).Greet())
	}
}
