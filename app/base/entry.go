package main

import (
	"log"

	"github.com/ic2hrmk/skeleton/app/base/internal"
	"github.com/ic2hrmk/skeleton/shared/app"
)

func FactoryMethod() (app.ApplicationService, error) {
	//
	// TODO: Do initialization here
	//
	return internal.NewBaseService(), nil
}

func main() {
	//
	// Read cli & env. here
	//
	// TODO:
	//  	cli.Load()
	//		env.Load()

	//
	// Receive application entity
	//
	application, err := FactoryMethod()
	if err != nil {
		log.Fatal(err)
	}

	//
	// Run application
	//     TODO: Add system exit signals
	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}
