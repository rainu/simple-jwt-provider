// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mailer

import (
	"gopkg.in/mail.v2"
	"sync"
)

var (
	locktemplateMockRender sync.RWMutex
)

// Ensure, that templateMock does implement template.
// If this is not the case, regenerate this file with moq.
var _ template = &templateMock{}

// templateMock is a mock implementation of template.
//
//     func TestSomethingThatUsestemplate(t *testing.T) {
//
//         // make and configure a mocked template
//         mockedtemplate := &templateMock{
//             RenderFunc: func(args interface{}) (*mail.Message, error) {
// 	               panic("mock out the Render method")
//             },
//         }
//
//         // use mockedtemplate in code that requires template
//         // and then make assertions.
//
//     }
type templateMock struct {
	// RenderFunc mocks the Render method.
	RenderFunc func(args interface{}) (*mail.Message, error)

	// calls tracks calls to the methods.
	calls struct {
		// Render holds details about calls to the Render method.
		Render []struct {
			// Args is the args argument value.
			Args interface{}
		}
	}
}

// Render calls RenderFunc.
func (mock *templateMock) Render(args interface{}) (*mail.Message, error) {
	if mock.RenderFunc == nil {
		panic("templateMock.RenderFunc: method is nil but template.Render was just called")
	}
	callInfo := struct {
		Args interface{}
	}{
		Args: args,
	}
	locktemplateMockRender.Lock()
	mock.calls.Render = append(mock.calls.Render, callInfo)
	locktemplateMockRender.Unlock()
	return mock.RenderFunc(args)
}

// RenderCalls gets all the calls that were made to Render.
// Check the length with:
//     len(mockedtemplate.RenderCalls())
func (mock *templateMock) RenderCalls() []struct {
	Args interface{}
} {
	var calls []struct {
		Args interface{}
	}
	locktemplateMockRender.RLock()
	calls = mock.calls.Render
	locktemplateMockRender.RUnlock()
	return calls
}
