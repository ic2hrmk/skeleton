package helpers

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/ic2hrmk/skeleton/shared/gateway/representation"
)

func logError(err error) {
	var message = "error is empty"

	if err != nil {
		message = err.Error()
	}

	log.Printf("   ERROR  | error during request proccessing, %s", message)
}

func ResponseWithBadRequest(response *restful.Response, err error, message string) {
	logError(err)

	if err := response.WriteHeaderAndEntity(http.StatusBadRequest, &representation.ErrorResponse{
		Message: message,
	}); err != nil {
		logError(err)
	}
}

func ResponseWithInternalError(response *restful.Response, err error) {
	logError(err)

	if err := response.WriteHeaderAndEntity(http.StatusInternalServerError, &representation.ErrorResponse{
		Message: http.StatusText(http.StatusInternalServerError),
	}); err != nil {
		logError(err)
	}
}

func ResponseWithOK(response *restful.Response, data interface{}) {
	if err := response.WriteHeaderAndEntity(http.StatusOK, data); err != nil {
		logError(err)
	}
}

func ResponseWithNoContent(response *restful.Response) {
	response.WriteHeader(http.StatusNoContent)
}
