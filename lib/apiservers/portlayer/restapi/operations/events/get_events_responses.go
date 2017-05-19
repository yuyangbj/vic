package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"
)

/*GetEventsOK OK

swagger:response getEventsOK
*/
type GetEventsOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewGetEventsOK creates GetEventsOK with default headers values
func NewGetEventsOK() *GetEventsOK {
	return &GetEventsOK{}
}

// WithPayload adds the payload to the get events o k response
func (o *GetEventsOK) WithPayload(payload io.ReadCloser) *GetEventsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get events o k response
func (o *GetEventsOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetEventsInternalServerError Failed to get events

swagger:response getEventsInternalServerError
*/
type GetEventsInternalServerError struct {
}

// NewGetEventsInternalServerError creates GetEventsInternalServerError with default headers values
func NewGetEventsInternalServerError() *GetEventsInternalServerError {
	return &GetEventsInternalServerError{}
}

// WriteResponse to the client
func (o *GetEventsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
}