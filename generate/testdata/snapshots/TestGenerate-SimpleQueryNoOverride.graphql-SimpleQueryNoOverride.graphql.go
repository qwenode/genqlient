// Code generated by github.com/qwenode/genqlient, DO NOT EDIT.

package test

import (
	"github.com/qwenode/genqlient/graphql"
	"github.com/qwenode/genqlient/internal/testutil"
)

// SimpleQueryNoOverrideResponse is returned by SimpleQueryNoOverride on success.
type SimpleQueryNoOverrideResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User SimpleQueryNoOverrideUser `json:"user"`
}

// GetUser returns SimpleQueryNoOverrideResponse.User, and is useful for accessing the field via an interface.
func (v *SimpleQueryNoOverrideResponse) GetUser() SimpleQueryNoOverrideUser { return v.User }

// SimpleQueryNoOverrideUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type SimpleQueryNoOverrideUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// GetId returns SimpleQueryNoOverrideUser.Id, and is useful for accessing the field via an interface.
func (v *SimpleQueryNoOverrideUser) GetId() testutil.ID { return v.Id }

// GetName returns SimpleQueryNoOverrideUser.Name, and is useful for accessing the field via an interface.
func (v *SimpleQueryNoOverrideUser) GetName() string { return v.Name }

// The query or mutation executed by SimpleQueryNoOverride.
const SimpleQueryNoOverride_Operation = `
query SimpleQueryNoOverride {
	user {
		id
		name
	}
}
`

func SimpleQueryNoOverride(
	client_ graphql.Client,
) (*SimpleQueryNoOverrideResponse, error) {
	req_ := &graphql.Request{
		OpName: "SimpleQueryNoOverride",
		Query:  SimpleQueryNoOverride_Operation,
	}
	var err_ error

	var data_ SimpleQueryNoOverrideResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		nil,
		req_,
		resp_,
	)

	return &data_, err_
}

