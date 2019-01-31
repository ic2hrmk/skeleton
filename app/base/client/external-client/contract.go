package external_client

import "errors"

var (
	ErrUnreachable = errors.New("ErrUnreachable")
)

type ExternalServiceClient interface {
	Ping() error
}
