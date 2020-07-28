// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package http

import "github.com/mainflux/mainflux/things"

var _ apiReq = (*identifyReq)(nil)

type apiReq interface {
	validate() error
}

type identifyReq struct {
	Token string `json:"token"`
}

func (req identifyReq) validate() error {
	if req.Token == "" {
		return things.ErrUnauthenticated
	}

	return nil
}

type canAccessByKeyReq struct {
	chanID string
	Token  string `json:"token"`
}

func (req canAccessByKeyReq) validate() error {
	if req.Token == "" {
		return things.ErrUnauthenticated
	}

	if req.chanID == "" {
		return things.ErrUnauthorized
	}

	return nil
}

type canAccessByIDReq struct {
	chanID  string
	ThingID string `json:"thing_id"`
}

func (req canAccessByIDReq) validate() error {
	if req.ThingID == "" {
		return things.ErrUnauthenticated
	}

	if req.chanID == "" {
		return things.ErrUnauthorized
	}

	return nil
}
