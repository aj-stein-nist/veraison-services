// Copyright 2022-2023 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0
package cca_realm

import (
	"github.com/veraison/services/handler"
)

type EndorsementHandler struct{}

func (o EndorsementHandler) Init(params handler.EndorsementHandlerParams) error {
	return nil // no-op
}

func (o EndorsementHandler) Close() error {
	return nil // no-op
}

func (o EndorsementHandler) GetName() string {
	return "unsigned-corim (CCA realm profile)"
}

func (o EndorsementHandler) GetAttestationScheme() string {
	return SchemeName
}

func (o EndorsementHandler) GetSupportedMediaTypes() []string {
	return EndorsementMediaTypes
}

func (o EndorsementHandler) Decode(data []byte) (*handler.EndorsementHandlerResponse, error) {
	return nil, nil
}
