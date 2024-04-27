// Code generated by "make api"; DO NOT EDIT.
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accounts

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type PasswordAccountAttributes struct {
	LoginName string `json:"login_name,omitempty"`
	Password  string `json:"password,omitempty"`
}

func AttributesMapToPasswordAccountAttributes(in map[string]interface{}) (*PasswordAccountAttributes, error) {
	if in == nil {
		return nil, fmt.Errorf("nil input map")
	}
	var out PasswordAccountAttributes
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:  &out,
		TagName: "json",
	})
	if err != nil {
		return nil, fmt.Errorf("error creating mapstructure decoder: %w", err)
	}
	if err := dec.Decode(in); err != nil {
		return nil, fmt.Errorf("error decoding: %w", err)
	}
	return &out, nil
}

func (pt *Account) GetPasswordAccountAttributes() (*PasswordAccountAttributes, error) {
	if pt.Type != "password" {
		return nil, fmt.Errorf("asked to fetch %s-type attributes but account is of type %s", "password", pt.Type)
	}
	return AttributesMapToPasswordAccountAttributes(pt.Attributes)
}