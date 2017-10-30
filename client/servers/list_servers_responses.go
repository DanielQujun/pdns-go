// Code generated by go-swagger; DO NOT EDIT.

package servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ffledgling/pdns-go/models"
)

// ListServersReader is a Reader for the ListServers structure.
type ListServersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListServersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListServersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListServersOK creates a ListServersOK with default headers values
func NewListServersOK() *ListServersOK {
	return &ListServersOK{}
}

/*ListServersOK handles this case with default header values.

An array of servers
*/
type ListServersOK struct {
	Payload models.Servers
}

func (o *ListServersOK) Error() string {
	return fmt.Sprintf("[GET /servers][%d] listServersOK  %+v", 200, o.Payload)
}

func (o *ListServersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
