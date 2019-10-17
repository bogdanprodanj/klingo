// Code generated by go-swagger; DO NOT EDIT.

package character

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostCharacterSearchParams creates a new PostCharacterSearchParams object
// with the default values initialized.
func NewPostCharacterSearchParams() *PostCharacterSearchParams {
	var ()
	return &PostCharacterSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCharacterSearchParamsWithTimeout creates a new PostCharacterSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCharacterSearchParamsWithTimeout(timeout time.Duration) *PostCharacterSearchParams {
	var ()
	return &PostCharacterSearchParams{

		timeout: timeout,
	}
}

// NewPostCharacterSearchParamsWithContext creates a new PostCharacterSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCharacterSearchParamsWithContext(ctx context.Context) *PostCharacterSearchParams {
	var ()
	return &PostCharacterSearchParams{

		Context: ctx,
	}
}

// NewPostCharacterSearchParamsWithHTTPClient creates a new PostCharacterSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCharacterSearchParamsWithHTTPClient(client *http.Client) *PostCharacterSearchParams {
	var ()
	return &PostCharacterSearchParams{
		HTTPClient: client,
	}
}

/*PostCharacterSearchParams contains all the parameters to send to the API endpoint
for the post character search operation typically these are written to a http.Request
*/
type PostCharacterSearchParams struct {

	/*AlternateReality
	  Whether it should be a alternate reality character

	*/
	AlternateReality *bool
	/*APIKey
	  API key

	*/
	APIKey *string
	/*Deceased
	  Whether it should be a deceased character

	*/
	Deceased *bool
	/*FictionalCharacter
	  Whether it should be a fictional character (from universe point of view)

	*/
	FictionalCharacter *bool
	/*Gender
	  Character gender

	*/
	Gender *string
	/*Hologram
	  Whether it should be a hologram

	*/
	Hologram *bool
	/*Mirror
	  Whether it should be a mirror universe character

	*/
	Mirror *bool
	/*Name
	  Character name

	*/
	Name *string
	/*PageNumber
	  Zero-based page number

	*/
	PageNumber *int32
	/*PageSize
	  Page size

	*/
	PageSize *int32
	/*Sort
	  Sorting, serialized like this: fieldName,ASC;anotherFieldName,DESC

	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post character search params
func (o *PostCharacterSearchParams) WithTimeout(timeout time.Duration) *PostCharacterSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post character search params
func (o *PostCharacterSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post character search params
func (o *PostCharacterSearchParams) WithContext(ctx context.Context) *PostCharacterSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post character search params
func (o *PostCharacterSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post character search params
func (o *PostCharacterSearchParams) WithHTTPClient(client *http.Client) *PostCharacterSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post character search params
func (o *PostCharacterSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlternateReality adds the alternateReality to the post character search params
func (o *PostCharacterSearchParams) WithAlternateReality(alternateReality *bool) *PostCharacterSearchParams {
	o.SetAlternateReality(alternateReality)
	return o
}

// SetAlternateReality adds the alternateReality to the post character search params
func (o *PostCharacterSearchParams) SetAlternateReality(alternateReality *bool) {
	o.AlternateReality = alternateReality
}

// WithAPIKey adds the aPIKey to the post character search params
func (o *PostCharacterSearchParams) WithAPIKey(aPIKey *string) *PostCharacterSearchParams {
	o.SetAPIKey(aPIKey)
	return o
}

// SetAPIKey adds the apiKey to the post character search params
func (o *PostCharacterSearchParams) SetAPIKey(aPIKey *string) {
	o.APIKey = aPIKey
}

// WithDeceased adds the deceased to the post character search params
func (o *PostCharacterSearchParams) WithDeceased(deceased *bool) *PostCharacterSearchParams {
	o.SetDeceased(deceased)
	return o
}

// SetDeceased adds the deceased to the post character search params
func (o *PostCharacterSearchParams) SetDeceased(deceased *bool) {
	o.Deceased = deceased
}

// WithFictionalCharacter adds the fictionalCharacter to the post character search params
func (o *PostCharacterSearchParams) WithFictionalCharacter(fictionalCharacter *bool) *PostCharacterSearchParams {
	o.SetFictionalCharacter(fictionalCharacter)
	return o
}

// SetFictionalCharacter adds the fictionalCharacter to the post character search params
func (o *PostCharacterSearchParams) SetFictionalCharacter(fictionalCharacter *bool) {
	o.FictionalCharacter = fictionalCharacter
}

// WithGender adds the gender to the post character search params
func (o *PostCharacterSearchParams) WithGender(gender *string) *PostCharacterSearchParams {
	o.SetGender(gender)
	return o
}

// SetGender adds the gender to the post character search params
func (o *PostCharacterSearchParams) SetGender(gender *string) {
	o.Gender = gender
}

// WithHologram adds the hologram to the post character search params
func (o *PostCharacterSearchParams) WithHologram(hologram *bool) *PostCharacterSearchParams {
	o.SetHologram(hologram)
	return o
}

// SetHologram adds the hologram to the post character search params
func (o *PostCharacterSearchParams) SetHologram(hologram *bool) {
	o.Hologram = hologram
}

// WithMirror adds the mirror to the post character search params
func (o *PostCharacterSearchParams) WithMirror(mirror *bool) *PostCharacterSearchParams {
	o.SetMirror(mirror)
	return o
}

// SetMirror adds the mirror to the post character search params
func (o *PostCharacterSearchParams) SetMirror(mirror *bool) {
	o.Mirror = mirror
}

// WithName adds the name to the post character search params
func (o *PostCharacterSearchParams) WithName(name *string) *PostCharacterSearchParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the post character search params
func (o *PostCharacterSearchParams) SetName(name *string) {
	o.Name = name
}

// WithPageNumber adds the pageNumber to the post character search params
func (o *PostCharacterSearchParams) WithPageNumber(pageNumber *int32) *PostCharacterSearchParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the post character search params
func (o *PostCharacterSearchParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the post character search params
func (o *PostCharacterSearchParams) WithPageSize(pageSize *int32) *PostCharacterSearchParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the post character search params
func (o *PostCharacterSearchParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSort adds the sort to the post character search params
func (o *PostCharacterSearchParams) WithSort(sort *string) *PostCharacterSearchParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the post character search params
func (o *PostCharacterSearchParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *PostCharacterSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AlternateReality != nil {

		// form param alternateReality
		var frAlternateReality bool
		if o.AlternateReality != nil {
			frAlternateReality = *o.AlternateReality
		}
		fAlternateReality := swag.FormatBool(frAlternateReality)
		if fAlternateReality != "" {
			if err := r.SetFormParam("alternateReality", fAlternateReality); err != nil {
				return err
			}
		}

	}

	if o.APIKey != nil {

		// query param apiKey
		var qrAPIKey string
		if o.APIKey != nil {
			qrAPIKey = *o.APIKey
		}
		qAPIKey := qrAPIKey
		if qAPIKey != "" {
			if err := r.SetQueryParam("apiKey", qAPIKey); err != nil {
				return err
			}
		}

	}

	if o.Deceased != nil {

		// form param deceased
		var frDeceased bool
		if o.Deceased != nil {
			frDeceased = *o.Deceased
		}
		fDeceased := swag.FormatBool(frDeceased)
		if fDeceased != "" {
			if err := r.SetFormParam("deceased", fDeceased); err != nil {
				return err
			}
		}

	}

	if o.FictionalCharacter != nil {

		// form param fictionalCharacter
		var frFictionalCharacter bool
		if o.FictionalCharacter != nil {
			frFictionalCharacter = *o.FictionalCharacter
		}
		fFictionalCharacter := swag.FormatBool(frFictionalCharacter)
		if fFictionalCharacter != "" {
			if err := r.SetFormParam("fictionalCharacter", fFictionalCharacter); err != nil {
				return err
			}
		}

	}

	if o.Gender != nil {

		// form param gender
		var frGender string
		if o.Gender != nil {
			frGender = *o.Gender
		}
		fGender := frGender
		if fGender != "" {
			if err := r.SetFormParam("gender", fGender); err != nil {
				return err
			}
		}

	}

	if o.Hologram != nil {

		// form param hologram
		var frHologram bool
		if o.Hologram != nil {
			frHologram = *o.Hologram
		}
		fHologram := swag.FormatBool(frHologram)
		if fHologram != "" {
			if err := r.SetFormParam("hologram", fHologram); err != nil {
				return err
			}
		}

	}

	if o.Mirror != nil {

		// form param mirror
		var frMirror bool
		if o.Mirror != nil {
			frMirror = *o.Mirror
		}
		fMirror := swag.FormatBool(frMirror)
		if fMirror != "" {
			if err := r.SetFormParam("mirror", fMirror); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// form param name
		var frName string
		if o.Name != nil {
			frName = *o.Name
		}
		fName := frName
		if fName != "" {
			if err := r.SetFormParam("name", fName); err != nil {
				return err
			}
		}

	}

	if o.PageNumber != nil {

		// query param pageNumber
		var qrPageNumber int32
		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := swag.FormatInt32(qrPageNumber)
		if qPageNumber != "" {
			if err := r.SetQueryParam("pageNumber", qPageNumber); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int32
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}

	}

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
