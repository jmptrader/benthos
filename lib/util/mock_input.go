/*
Copyright (c) 2014 Ashley Jeffs

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package util

import (
	"time"

	"github.com/jeffail/benthos/lib/types"
)

//--------------------------------------------------------------------------------------------------

// MockInputType - Implements the input.Type interface.
type MockInputType struct {
	MsgChan chan types.Message
	ResChan <-chan types.Response
}

// StartListening - Sets the channel used for reading responses.
func (m *MockInputType) StartListening(resChan <-chan types.Response) error {
	m.ResChan = resChan
	return nil
}

// MessageChan - Returns the messages channel.
func (m *MockInputType) MessageChan() <-chan types.Message {
	return m.MsgChan
}

// CloseAsync - Does nothing.
func (m MockInputType) CloseAsync() {
	// Do nothing
}

// WaitForClose - Does nothing.
func (m MockInputType) WaitForClose(time.Duration) error {
	// Do nothing
	return nil
}

//--------------------------------------------------------------------------------------------------
