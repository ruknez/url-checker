// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	"sync"
	entity "url-checker/internal/domain"
)

// GetUrlStatuserMock is a mock implementation of checker.getUrlStatuser.
//
//	func TestSomethingThatUsesgetUrlStatuser(t *testing.T) {
//
//		// make and configure a mocked checker.getUrlStatuser
//		mockedgetUrlStatuser := &GetUrlStatuserMock{
//			GetUrlStatusFunc: func(ctx context.Context, url string) (entity.Status, error) {
//				panic("mock out the GetUrlStatus method")
//			},
//		}
//
//		// use mockedgetUrlStatuser in code that requires checker.getUrlStatuser
//		// and then make assertions.
//
//	}
type GetUrlStatuserMock struct {
	// GetUrlStatusFunc mocks the GetUrlStatus method.
	GetUrlStatusFunc func(ctx context.Context, url string) (entity.Status, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetUrlStatus holds details about calls to the GetUrlStatus method.
		GetUrlStatus []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// URL is the url argument value.
			URL string
		}
	}
	lockGetUrlStatus sync.RWMutex
}

// GetUrlStatus calls GetUrlStatusFunc.
func (mock *GetUrlStatuserMock) GetUrlStatus(ctx context.Context, url string) (entity.Status, error) {
	callInfo := struct {
		Ctx context.Context
		URL string
	}{
		Ctx: ctx,
		URL: url,
	}
	mock.lockGetUrlStatus.Lock()
	mock.calls.GetUrlStatus = append(mock.calls.GetUrlStatus, callInfo)
	mock.lockGetUrlStatus.Unlock()
	if mock.GetUrlStatusFunc == nil {
		var (
			statusOut entity.Status
			errOut    error
		)
		return statusOut, errOut
	}
	return mock.GetUrlStatusFunc(ctx, url)
}

// GetUrlStatusCalls gets all the calls that were made to GetUrlStatus.
// Check the length with:
//
//	len(mockedgetUrlStatuser.GetUrlStatusCalls())
func (mock *GetUrlStatuserMock) GetUrlStatusCalls() []struct {
	Ctx context.Context
	URL string
} {
	var calls []struct {
		Ctx context.Context
		URL string
	}
	mock.lockGetUrlStatus.RLock()
	calls = mock.calls.GetUrlStatus
	mock.lockGetUrlStatus.RUnlock()
	return calls
}