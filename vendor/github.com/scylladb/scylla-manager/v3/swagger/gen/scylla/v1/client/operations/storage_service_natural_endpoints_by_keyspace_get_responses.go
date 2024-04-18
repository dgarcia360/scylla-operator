// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/v3/swagger/gen/scylla/v1/models"
)

// StorageServiceNaturalEndpointsByKeyspaceGetReader is a Reader for the StorageServiceNaturalEndpointsByKeyspaceGet structure.
type StorageServiceNaturalEndpointsByKeyspaceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceNaturalEndpointsByKeyspaceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceNaturalEndpointsByKeyspaceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceNaturalEndpointsByKeyspaceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceNaturalEndpointsByKeyspaceGetOK creates a StorageServiceNaturalEndpointsByKeyspaceGetOK with default headers values
func NewStorageServiceNaturalEndpointsByKeyspaceGetOK() *StorageServiceNaturalEndpointsByKeyspaceGetOK {
	return &StorageServiceNaturalEndpointsByKeyspaceGetOK{}
}

/*
StorageServiceNaturalEndpointsByKeyspaceGetOK handles this case with default header values.

Success
*/
type StorageServiceNaturalEndpointsByKeyspaceGetOK struct {
	Payload []string
}

func (o *StorageServiceNaturalEndpointsByKeyspaceGetOK) GetPayload() []string {
	return o.Payload
}

func (o *StorageServiceNaturalEndpointsByKeyspaceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceNaturalEndpointsByKeyspaceGetDefault creates a StorageServiceNaturalEndpointsByKeyspaceGetDefault with default headers values
func NewStorageServiceNaturalEndpointsByKeyspaceGetDefault(code int) *StorageServiceNaturalEndpointsByKeyspaceGetDefault {
	return &StorageServiceNaturalEndpointsByKeyspaceGetDefault{
		_statusCode: code,
	}
}

/*
StorageServiceNaturalEndpointsByKeyspaceGetDefault handles this case with default header values.

internal server error
*/
type StorageServiceNaturalEndpointsByKeyspaceGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service natural endpoints by keyspace get default response
func (o *StorageServiceNaturalEndpointsByKeyspaceGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceNaturalEndpointsByKeyspaceGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceNaturalEndpointsByKeyspaceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceNaturalEndpointsByKeyspaceGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}