// Code generated by go-swagger; DO NOT EDIT.

package cluster_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// UpgradeClusterReader is a Reader for the UpgradeCluster structure.
type UpgradeClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpgradeClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpgradeClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpgradeClusterOK creates a UpgradeClusterOK with default headers values
func NewUpgradeClusterOK() *UpgradeClusterOK {
	return &UpgradeClusterOK{}
}

/*UpgradeClusterOK handles this case with default header values.

UpgradeClusterOK upgrade cluster o k
*/
type UpgradeClusterOK struct {
	Payload *models.OpenpitrixUpgradeClusterResponse
}

func (o *UpgradeClusterOK) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/upgrade][%d] upgradeClusterOK  %+v", 200, o.Payload)
}

func (o *UpgradeClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixUpgradeClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}