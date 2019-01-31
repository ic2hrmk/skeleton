package internal

import (
	"sync"

	"github.com/ic2hrmk/skeleton/shared/app"
)

//
// Main application service
//
type BaseService struct {
	// Services
	//
	// - rest.APIService
	// - grpc.APIService

	// Clients
	// ...

	// Persistence
	// ...

	// Logger ?
	// ...
}

//
// Service constructor
//
func NewBaseService(
	//
	// Initialized sub-services
	//
) app.ApplicationService {
	return &BaseService{

	}
}

//
// Application main run method
//
func (rcv *BaseService) Run() error {
	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		wg.Done()
	}()

	wg.Wait()

	return nil
}
