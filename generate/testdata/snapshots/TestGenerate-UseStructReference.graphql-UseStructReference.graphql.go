// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package test

import (
	"github.com/Khan/genqlient/graphql"
)

type StructInput struct {
	Field string `json:"field"`
}

// GetField returns StructInput.Field, and is useful for accessing the field via an interface.
func (v *StructInput) GetField() string { return v.Field }

// UseStructReferenceResponse is returned by UseStructReference on success.
type UseStructReferenceResponse struct {
	UseStructReferencesInput bool `json:"useStructReferencesInput"`
}

// GetUseStructReferencesInput returns UseStructReferenceResponse.UseStructReferencesInput, and is useful for accessing the field via an interface.
func (v *UseStructReferenceResponse) GetUseStructReferencesInput() bool {
	return v.UseStructReferencesInput
}

type UseStructReferencesInput struct {
	Struct         StructInput   `json:"struct"`
	NullableStruct StructInput   `json:"nullableStruct"`
	List           []StructInput `json:"list"`
	ListOfNullable []StructInput `json:"listOfNullable"`
	NullableList   []StructInput `json:"nullableList"`
}

// GetStruct returns UseStructReferencesInput.Struct, and is useful for accessing the field via an interface.
func (v *UseStructReferencesInput) GetStruct() StructInput { return v.Struct }

// GetNullableStruct returns UseStructReferencesInput.NullableStruct, and is useful for accessing the field via an interface.
func (v *UseStructReferencesInput) GetNullableStruct() StructInput { return v.NullableStruct }

// GetList returns UseStructReferencesInput.List, and is useful for accessing the field via an interface.
func (v *UseStructReferencesInput) GetList() []StructInput { return v.List }

// GetListOfNullable returns UseStructReferencesInput.ListOfNullable, and is useful for accessing the field via an interface.
func (v *UseStructReferencesInput) GetListOfNullable() []StructInput { return v.ListOfNullable }

// GetNullableList returns UseStructReferencesInput.NullableList, and is useful for accessing the field via an interface.
func (v *UseStructReferencesInput) GetNullableList() []StructInput { return v.NullableList }

// __UseStructReferenceInput is used internally by genqlient
type __UseStructReferenceInput struct {
	Input UseStructReferencesInput `json:"input"`
}

// GetInput returns __UseStructReferenceInput.Input, and is useful for accessing the field via an interface.
func (v *__UseStructReferenceInput) GetInput() UseStructReferencesInput { return v.Input }

// The query or mutation executed by UseStructReference.
const UseStructReference_Operation = `
query UseStructReference ($input: UseStructReferencesInput!) {
	useStructReferencesInput(input: $input)
}
`

// https://github.com/Khan/genqlient/issues/342
func UseStructReference(
	client_ graphql.Client,
	input UseStructReferencesInput,
) (*UseStructReferenceResponse, error) {
	req_ := &graphql.Request{
		OpName: "UseStructReference",
		Query:  UseStructReference_Operation,
		Variables: &__UseStructReferenceInput{
			Input: input,
		},
	}
	var err_ error

	var data_ UseStructReferenceResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		nil,
		req_,
		resp_,
	)

	return &data_, err_
}
