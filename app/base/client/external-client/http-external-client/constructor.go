package http_external_client

import "github.com/ic2hrmk/skeleton/app/base/client/external-client"

type externalServiceClient struct {

}

func NewHttpExternalClient() external_client.ExternalServiceClient {
	return &externalServiceClient{}
}
