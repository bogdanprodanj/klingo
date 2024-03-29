// Code generated by go-swagger; DO NOT EDIT.

package character

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new character API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for character API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetCharacter Retrival of a single character
*/
func (a *Client) GetCharacter(params *GetCharacterParams) (*GetCharacterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharacterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCharacter",
		Method:             "GET",
		PathPattern:        "/character",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCharacterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCharacterOK), nil

}

/*
PostCharacterSearch Searching characters
*/
func (a *Client) PostCharacterSearch(params *PostCharacterSearchParams) (*PostCharacterSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCharacterSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostCharacterSearch",
		Method:             "POST",
		PathPattern:        "/character/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostCharacterSearchReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostCharacterSearchOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
