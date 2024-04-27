// Code generated by "make api"; DO NOT EDIT.
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hosts

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type StaticHostAttributes struct {
	Address string `json:"address,omitempty"`
}

func AttributesMapToStaticHostAttributes(in map[string]interface{}) (*StaticHostAttributes, error) {
	if in == nil {
		return nil, fmt.Errorf("nil input map")
	}
	var out StaticHostAttributes
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

func (pt *Host) GetStaticHostAttributes() (*StaticHostAttributes, error) {
	if pt.Type != "static" {
		return nil, fmt.Errorf("asked to fetch %s-type attributes but host is of type %s", "static", pt.Type)
	}
	return AttributesMapToStaticHostAttributes(pt.Attributes)
}
