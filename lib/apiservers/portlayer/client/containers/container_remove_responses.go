package containers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/vic/lib/apiservers/portlayer/models"
)

// ContainerRemoveReader is a Reader for the ContainerRemove structure.
type ContainerRemoveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerRemoveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewContainerRemoveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewContainerRemoveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewContainerRemoveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewContainerRemoveConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewContainerRemoveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewContainerRemoveDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewContainerRemoveOK creates a ContainerRemoveOK with default headers values
func NewContainerRemoveOK() *ContainerRemoveOK {
	return &ContainerRemoveOK{}
}

/*ContainerRemoveOK handles this case with default header values.

OK
*/
type ContainerRemoveOK struct {
}

func (o *ContainerRemoveOK) Error() string {
	return fmt.Sprintf("[DELETE /containers/{id}][%d] containerRemoveOK ", 200)
}

func (o *ContainerRemoveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerRemoveBadRequest creates a ContainerRemoveBadRequest with default headers values
func NewContainerRemoveBadRequest() *ContainerRemoveBadRequest {
	return &ContainerRemoveBadRequest{}
}

/*ContainerRemoveBadRequest handles this case with default header values.

bad parameter
*/
type ContainerRemoveBadRequest struct {
}

func (o *ContainerRemoveBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /containers/{id}][%d] containerRemoveBadRequest ", 400)
}

func (o *ContainerRemoveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerRemoveNotFound creates a ContainerRemoveNotFound with default headers values
func NewContainerRemoveNotFound() *ContainerRemoveNotFound {
	return &ContainerRemoveNotFound{}
}

/*ContainerRemoveNotFound handles this case with default header values.

no such container
*/
type ContainerRemoveNotFound struct {
}

func (o *ContainerRemoveNotFound) Error() string {
	return fmt.Sprintf("[DELETE /containers/{id}][%d] containerRemoveNotFound ", 404)
}

func (o *ContainerRemoveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerRemoveConflict creates a ContainerRemoveConflict with default headers values
func NewContainerRemoveConflict() *ContainerRemoveConflict {
	return &ContainerRemoveConflict{}
}

/*ContainerRemoveConflict handles this case with default header values.

conflict
*/
type ContainerRemoveConflict struct {
	Payload *models.Error
}

func (o *ContainerRemoveConflict) Error() string {
	return fmt.Sprintf("[DELETE /containers/{id}][%d] containerRemoveConflict  %+v", 409, o.Payload)
}

func (o *ContainerRemoveConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerRemoveInternalServerError creates a ContainerRemoveInternalServerError with default headers values
func NewContainerRemoveInternalServerError() *ContainerRemoveInternalServerError {
	return &ContainerRemoveInternalServerError{}
}

/*ContainerRemoveInternalServerError handles this case with default header values.

server error
*/
type ContainerRemoveInternalServerError struct {
}

func (o *ContainerRemoveInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /containers/{id}][%d] containerRemoveInternalServerError ", 500)
}

func (o *ContainerRemoveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerRemoveDefault creates a ContainerRemoveDefault with default headers values
func NewContainerRemoveDefault(code int) *ContainerRemoveDefault {
	return &ContainerRemoveDefault{
		_statusCode: code,
	}
}

/*ContainerRemoveDefault handles this case with default header values.

Error
*/
type ContainerRemoveDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the container remove default response
func (o *ContainerRemoveDefault) Code() int {
	return o._statusCode
}

func (o *ContainerRemoveDefault) Error() string {
	return fmt.Sprintf("[DELETE /containers/{id}][%d] ContainerRemove default  %+v", o._statusCode, o.Payload)
}

func (o *ContainerRemoveDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}