// Code generated by go-swagger; DO NOT EDIT.

package app_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// DescribeActiveAppsReader is a Reader for the DescribeActiveApps structure.
type DescribeActiveAppsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeActiveAppsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDescribeActiveAppsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDescribeActiveAppsOK creates a DescribeActiveAppsOK with default headers values
func NewDescribeActiveAppsOK() *DescribeActiveAppsOK {
	return &DescribeActiveAppsOK{}
}

/*DescribeActiveAppsOK handles this case with default header values.

A successful response.
*/
type DescribeActiveAppsOK struct {
	Payload *models.OpenpitrixDescribeAppsResponse
}

func (o *DescribeActiveAppsOK) Error() string {
	return fmt.Sprintf("[GET /api/AppManager.DescribeActiveApps][%d] describeActiveAppsOK  %+v", 200, o.Payload)
}

func (o *DescribeActiveAppsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixDescribeAppsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}