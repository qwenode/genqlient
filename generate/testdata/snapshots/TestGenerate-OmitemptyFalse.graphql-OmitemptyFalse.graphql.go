// Code generated by github.com/qwenode/genqlient, DO NOT EDIT.

package test

import (
	"github.com/qwenode/genqlient/graphql"
)

// OmitemptyFalseResponse is returned by OmitemptyFalse on success.
type OmitemptyFalseResponse struct {
	Omitempty bool `json:"omitempty"`
}

// GetOmitempty returns OmitemptyFalseResponse.Omitempty, and is useful for accessing the field via an interface.
func (v *OmitemptyFalseResponse) GetOmitempty() bool { return v.Omitempty }

type OmitemptyInput struct {
	Field         string `json:"field"`
	NullableField string `json:"nullableField,omitempty"`
}

// GetField returns OmitemptyInput.Field, and is useful for accessing the field via an interface.
func (v *OmitemptyInput) GetField() string { return v.Field }

// GetNullableField returns OmitemptyInput.NullableField, and is useful for accessing the field via an interface.
func (v *OmitemptyInput) GetNullableField() string { return v.NullableField }

// __OmitemptyFalseInput is used internally by genqlient
type __OmitemptyFalseInput struct {
	Input OmitemptyInput `json:"input,omitempty"`
}

// GetInput returns __OmitemptyFalseInput.Input, and is useful for accessing the field via an interface.
func (v *__OmitemptyFalseInput) GetInput() OmitemptyInput { return v.Input }

// The query or mutation executed by OmitemptyFalse.
const OmitemptyFalse_Operation = `
query OmitemptyFalse ($input: OmitemptyInput) {
	omitempty(input: $input)
}
`

func OmitemptyFalse(
	client_ graphql.Client,
	input OmitemptyInput,
) (*OmitemptyFalseResponse, error) {
	req_ := &graphql.Request{
		OpName: "OmitemptyFalse",
		Query:  OmitemptyFalse_Operation,
		Variables: &__OmitemptyFalseInput{
			Input: input,
		},
	}
	var err_ error

	var data_ OmitemptyFalseResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		nil,
		req_,
		resp_,
	)

	return &data_, err_
}

