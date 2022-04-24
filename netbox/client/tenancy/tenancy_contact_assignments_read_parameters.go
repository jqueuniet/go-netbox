// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewTenancyContactAssignmentsReadParams creates a new TenancyContactAssignmentsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenancyContactAssignmentsReadParams() *TenancyContactAssignmentsReadParams {
	return &TenancyContactAssignmentsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyContactAssignmentsReadParamsWithTimeout creates a new TenancyContactAssignmentsReadParams object
// with the ability to set a timeout on a request.
func NewTenancyContactAssignmentsReadParamsWithTimeout(timeout time.Duration) *TenancyContactAssignmentsReadParams {
	return &TenancyContactAssignmentsReadParams{
		timeout: timeout,
	}
}

// NewTenancyContactAssignmentsReadParamsWithContext creates a new TenancyContactAssignmentsReadParams object
// with the ability to set a context for a request.
func NewTenancyContactAssignmentsReadParamsWithContext(ctx context.Context) *TenancyContactAssignmentsReadParams {
	return &TenancyContactAssignmentsReadParams{
		Context: ctx,
	}
}

// NewTenancyContactAssignmentsReadParamsWithHTTPClient creates a new TenancyContactAssignmentsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenancyContactAssignmentsReadParamsWithHTTPClient(client *http.Client) *TenancyContactAssignmentsReadParams {
	return &TenancyContactAssignmentsReadParams{
		HTTPClient: client,
	}
}

/* TenancyContactAssignmentsReadParams contains all the parameters to send to the API endpoint
   for the tenancy contact assignments read operation.

   Typically these are written to a http.Request.
*/
type TenancyContactAssignmentsReadParams struct {

	/* ID.

	   A unique integer value identifying this contact assignment.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenancy contact assignments read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyContactAssignmentsReadParams) WithDefaults() *TenancyContactAssignmentsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenancy contact assignments read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyContactAssignmentsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenancy contact assignments read params
func (o *TenancyContactAssignmentsReadParams) WithTimeout(timeout time.Duration) *TenancyContactAssignmentsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy contact assignments read params
func (o *TenancyContactAssignmentsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy contact assignments read params
func (o *TenancyContactAssignmentsReadParams) WithContext(ctx context.Context) *TenancyContactAssignmentsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy contact assignments read params
func (o *TenancyContactAssignmentsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy contact assignments read params
func (o *TenancyContactAssignmentsReadParams) WithHTTPClient(client *http.Client) *TenancyContactAssignmentsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy contact assignments read params
func (o *TenancyContactAssignmentsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the tenancy contact assignments read params
func (o *TenancyContactAssignmentsReadParams) WithID(id int64) *TenancyContactAssignmentsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the tenancy contact assignments read params
func (o *TenancyContactAssignmentsReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyContactAssignmentsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}