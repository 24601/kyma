// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import (
	apperrors "github.com/kyma-project/kyma/components/application-registry/internal/apperrors"
	applications "github.com/kyma-project/kyma/components/application-registry/internal/metadata/applications"

	mock "github.com/stretchr/testify/mock"

	model "github.com/kyma-project/kyma/components/application-registry/internal/metadata/model"

	strategy "github.com/kyma-project/kyma/components/application-registry/internal/metadata/secrets/strategy"
)

// Factory is an autogenerated mock type for the Factory type
type Factory struct {
	mock.Mock
}

// NewSecretAccessStrategy provides a mock function with given fields: credentials
func (_m *Factory) NewSecretAccessStrategy(credentials *applications.Credentials) (strategy.AccessStrategy, apperrors.AppError) {
	ret := _m.Called(credentials)

	var r0 strategy.AccessStrategy
	if rf, ok := ret.Get(0).(func(*applications.Credentials) strategy.AccessStrategy); ok {
		r0 = rf(credentials)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(strategy.AccessStrategy)
		}
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(*applications.Credentials) apperrors.AppError); ok {
		r1 = rf(credentials)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}

// NewSecretModificationStrategy provides a mock function with given fields: credentials
func (_m *Factory) NewSecretModificationStrategy(credentials *model.CredentialsWithCSRF) (strategy.ModificationStrategy, apperrors.AppError) {
	ret := _m.Called(credentials)

	var r0 strategy.ModificationStrategy
	if rf, ok := ret.Get(0).(func(*model.CredentialsWithCSRF) strategy.ModificationStrategy); ok {
		r0 = rf(credentials)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(strategy.ModificationStrategy)
		}
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(*model.CredentialsWithCSRF) apperrors.AppError); ok {
		r1 = rf(credentials)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}
