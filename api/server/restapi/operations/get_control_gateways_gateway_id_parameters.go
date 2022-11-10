// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetControlGatewaysGatewayIDParams creates a new GetControlGatewaysGatewayIDParams object
//
// There are no default values defined in the spec.
func NewGetControlGatewaysGatewayIDParams() GetControlGatewaysGatewayIDParams {

	return GetControlGatewaysGatewayIDParams{}
}

// GetControlGatewaysGatewayIDParams contains all the bound params for the get control gateways gateway ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetControlGatewaysGatewayID
type GetControlGatewaysGatewayIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Gateway ID
	  Required: true
	  In: path
	*/
	GatewayID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetControlGatewaysGatewayIDParams() beforehand.
func (o *GetControlGatewaysGatewayIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rGatewayID, rhkGatewayID, _ := route.Params.GetOK("gatewayId")
	if err := o.bindGatewayID(rGatewayID, rhkGatewayID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindGatewayID binds and validates parameter GatewayID from path.
func (o *GetControlGatewaysGatewayIDParams) bindGatewayID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("gatewayId", "path", "int64", raw)
	}
	o.GatewayID = value

	return nil
}
