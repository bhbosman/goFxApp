// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bhbosman/gocommon/internal/Services/FxAppManagerService (interfaces: IFxManager)

// Package FxAppManagerService is a generated GoMock package.
package FxAppManagerService

import (
	context "context"
	fmt "fmt"

	errors "github.com/bhbosman/gocommon/errors"
)

// Interface A Comment
// Interface github.com/bhbosman/gocommon/internal/Services/FxAppManagerService
// Interface IFxManager
// Interface IFxManager, Method: Start
type IFxManagerStartIn struct {
	arg0 context.Context
	arg1 []string
}

type IFxManagerStartOut struct {
	Args0 error
}
type IFxManagerStartError struct {
	InterfaceName string
	MethodName    string
	Reason        string
}

func (self *IFxManagerStartError) Error() string {
	return fmt.Sprintf("error in data coming back from %v::%v. Reason: %v", self.InterfaceName, self.MethodName, self.Reason)
}

type IFxManagerStart struct {
	inData         IFxManagerStartIn
	outDataChannel chan IFxManagerStartOut
}

func NewIFxManagerStart(waitToComplete bool, arg0 context.Context, arg1 ...string) *IFxManagerStart {
	var outDataChannel chan IFxManagerStartOut
	if waitToComplete {
		outDataChannel = make(chan IFxManagerStartOut)
	} else {
		outDataChannel = nil
	}
	return &IFxManagerStart{
		inData: IFxManagerStartIn{
			arg0: arg0,
			arg1: arg1,
		},
		outDataChannel: outDataChannel,
	}
}

func (self *IFxManagerStart) Wait(onError func(interfaceName string, methodName string, err error) error) (IFxManagerStartOut, error) {
	data, ok := <-self.outDataChannel
	if !ok {
		generatedError := &IFxManagerStartError{
			InterfaceName: "IFxManager",
			MethodName:    "Start",
			Reason:        "Channel for IFxManager::Start returned false",
		}
		if onError != nil {
			err := onError("IFxManager", "Start", generatedError)
			return IFxManagerStartOut{}, err
		} else {
			return IFxManagerStartOut{}, generatedError
		}
	}
	return data, nil
}

func (self *IFxManagerStart) Close() error {
	close(self.outDataChannel)
	return nil
}
func CallIFxManagerStart(context context.Context, channel chan<- interface{}, waitToComplete bool, arg0 context.Context, arg1 ...string) (IFxManagerStartOut, error) {
	if context != nil && context.Err() != nil {
		return IFxManagerStartOut{}, context.Err()
	}
	data := NewIFxManagerStart(waitToComplete, arg0, arg1...)
	if waitToComplete {
		defer func(data *IFxManagerStart) {
			err := data.Close()
			if err != nil {
			}
		}(data)
	}
	if context != nil && context.Err() != nil {
		return IFxManagerStartOut{}, context.Err()
	}
	channel <- data
	var err error
	var v IFxManagerStartOut
	if waitToComplete {
		v, err = data.Wait(func(interfaceName string, methodName string, err error) error {
			return err
		})
	} else {
		err = errors.NoWaitOperationError
	}
	if err != nil {
		return IFxManagerStartOut{}, err
	}
	return v, nil
}

// Interface IFxManager, Method: StartAll
type IFxManagerStartAllIn struct {
	arg0 context.Context
}

type IFxManagerStartAllOut struct {
	Args0 error
}
type IFxManagerStartAllError struct {
	InterfaceName string
	MethodName    string
	Reason        string
}

func (self *IFxManagerStartAllError) Error() string {
	return fmt.Sprintf("error in data coming back from %v::%v. Reason: %v", self.InterfaceName, self.MethodName, self.Reason)
}

type IFxManagerStartAll struct {
	inData         IFxManagerStartAllIn
	outDataChannel chan IFxManagerStartAllOut
}

func NewIFxManagerStartAll(waitToComplete bool, arg0 context.Context) *IFxManagerStartAll {
	var outDataChannel chan IFxManagerStartAllOut
	if waitToComplete {
		outDataChannel = make(chan IFxManagerStartAllOut)
	} else {
		outDataChannel = nil
	}
	return &IFxManagerStartAll{
		inData: IFxManagerStartAllIn{
			arg0: arg0,
		},
		outDataChannel: outDataChannel,
	}
}

func (self *IFxManagerStartAll) Wait(onError func(interfaceName string, methodName string, err error) error) (IFxManagerStartAllOut, error) {
	data, ok := <-self.outDataChannel
	if !ok {
		generatedError := &IFxManagerStartAllError{
			InterfaceName: "IFxManager",
			MethodName:    "StartAll",
			Reason:        "Channel for IFxManager::StartAll returned false",
		}
		if onError != nil {
			err := onError("IFxManager", "StartAll", generatedError)
			return IFxManagerStartAllOut{}, err
		} else {
			return IFxManagerStartAllOut{}, generatedError
		}
	}
	return data, nil
}

func (self *IFxManagerStartAll) Close() error {
	close(self.outDataChannel)
	return nil
}
func CallIFxManagerStartAll(context context.Context, channel chan<- interface{}, waitToComplete bool, arg0 context.Context) (IFxManagerStartAllOut, error) {
	if context != nil && context.Err() != nil {
		return IFxManagerStartAllOut{}, context.Err()
	}
	data := NewIFxManagerStartAll(waitToComplete, arg0)
	if waitToComplete {
		defer func(data *IFxManagerStartAll) {
			err := data.Close()
			if err != nil {
			}
		}(data)
	}
	if context != nil && context.Err() != nil {
		return IFxManagerStartAllOut{}, context.Err()
	}
	channel <- data
	var err error
	var v IFxManagerStartAllOut
	if waitToComplete {
		v, err = data.Wait(func(interfaceName string, methodName string, err error) error {
			return err
		})
	} else {
		err = errors.NoWaitOperationError
	}
	if err != nil {
		return IFxManagerStartAllOut{}, err
	}
	return v, nil
}

// Interface IFxManager, Method: Stop
type IFxManagerStopIn struct {
	arg0 context.Context
	arg1 []string
}

type IFxManagerStopOut struct {
	Args0 error
}
type IFxManagerStopError struct {
	InterfaceName string
	MethodName    string
	Reason        string
}

func (self *IFxManagerStopError) Error() string {
	return fmt.Sprintf("error in data coming back from %v::%v. Reason: %v", self.InterfaceName, self.MethodName, self.Reason)
}

type IFxManagerStop struct {
	inData         IFxManagerStopIn
	outDataChannel chan IFxManagerStopOut
}

func NewIFxManagerStop(waitToComplete bool, arg0 context.Context, arg1 ...string) *IFxManagerStop {
	var outDataChannel chan IFxManagerStopOut
	if waitToComplete {
		outDataChannel = make(chan IFxManagerStopOut)
	} else {
		outDataChannel = nil
	}
	return &IFxManagerStop{
		inData: IFxManagerStopIn{
			arg0: arg0,
			arg1: arg1,
		},
		outDataChannel: outDataChannel,
	}
}

func (self *IFxManagerStop) Wait(onError func(interfaceName string, methodName string, err error) error) (IFxManagerStopOut, error) {
	data, ok := <-self.outDataChannel
	if !ok {
		generatedError := &IFxManagerStopError{
			InterfaceName: "IFxManager",
			MethodName:    "Stop",
			Reason:        "Channel for IFxManager::Stop returned false",
		}
		if onError != nil {
			err := onError("IFxManager", "Stop", generatedError)
			return IFxManagerStopOut{}, err
		} else {
			return IFxManagerStopOut{}, generatedError
		}
	}
	return data, nil
}

func (self *IFxManagerStop) Close() error {
	close(self.outDataChannel)
	return nil
}
func CallIFxManagerStop(context context.Context, channel chan<- interface{}, waitToComplete bool, arg0 context.Context, arg1 ...string) (IFxManagerStopOut, error) {
	if context != nil && context.Err() != nil {
		return IFxManagerStopOut{}, context.Err()
	}
	data := NewIFxManagerStop(waitToComplete, arg0, arg1...)
	if waitToComplete {
		defer func(data *IFxManagerStop) {
			err := data.Close()
			if err != nil {
			}
		}(data)
	}
	if context != nil && context.Err() != nil {
		return IFxManagerStopOut{}, context.Err()
	}
	channel <- data
	var err error
	var v IFxManagerStopOut
	if waitToComplete {
		v, err = data.Wait(func(interfaceName string, methodName string, err error) error {
			return err
		})
	} else {
		err = errors.NoWaitOperationError
	}
	if err != nil {
		return IFxManagerStopOut{}, err
	}
	return v, nil
}

// Interface IFxManager, Method: StopAll
type IFxManagerStopAllIn struct {
	arg0 context.Context
}

type IFxManagerStopAllOut struct {
	Args0 error
}
type IFxManagerStopAllError struct {
	InterfaceName string
	MethodName    string
	Reason        string
}

func (self *IFxManagerStopAllError) Error() string {
	return fmt.Sprintf("error in data coming back from %v::%v. Reason: %v", self.InterfaceName, self.MethodName, self.Reason)
}

type IFxManagerStopAll struct {
	inData         IFxManagerStopAllIn
	outDataChannel chan IFxManagerStopAllOut
}

func NewIFxManagerStopAll(waitToComplete bool, arg0 context.Context) *IFxManagerStopAll {
	var outDataChannel chan IFxManagerStopAllOut
	if waitToComplete {
		outDataChannel = make(chan IFxManagerStopAllOut)
	} else {
		outDataChannel = nil
	}
	return &IFxManagerStopAll{
		inData: IFxManagerStopAllIn{
			arg0: arg0,
		},
		outDataChannel: outDataChannel,
	}
}

func (self *IFxManagerStopAll) Wait(onError func(interfaceName string, methodName string, err error) error) (IFxManagerStopAllOut, error) {
	data, ok := <-self.outDataChannel
	if !ok {
		generatedError := &IFxManagerStopAllError{
			InterfaceName: "IFxManager",
			MethodName:    "StopAll",
			Reason:        "Channel for IFxManager::StopAll returned false",
		}
		if onError != nil {
			err := onError("IFxManager", "StopAll", generatedError)
			return IFxManagerStopAllOut{}, err
		} else {
			return IFxManagerStopAllOut{}, generatedError
		}
	}
	return data, nil
}

func (self *IFxManagerStopAll) Close() error {
	close(self.outDataChannel)
	return nil
}
func CallIFxManagerStopAll(context context.Context, channel chan<- interface{}, waitToComplete bool, arg0 context.Context) (IFxManagerStopAllOut, error) {
	if context != nil && context.Err() != nil {
		return IFxManagerStopAllOut{}, context.Err()
	}
	data := NewIFxManagerStopAll(waitToComplete, arg0)
	if waitToComplete {
		defer func(data *IFxManagerStopAll) {
			err := data.Close()
			if err != nil {
			}
		}(data)
	}
	if context != nil && context.Err() != nil {
		return IFxManagerStopAllOut{}, context.Err()
	}
	channel <- data
	var err error
	var v IFxManagerStopAllOut
	if waitToComplete {
		v, err = data.Wait(func(interfaceName string, methodName string, err error) error {
			return err
		})
	} else {
		err = errors.NoWaitOperationError
	}
	if err != nil {
		return IFxManagerStopAllOut{}, err
	}
	return v, nil
}

func ChannelEventsForIFxManager(next IFxManager, event interface{}) (bool, error) {
	switch v := event.(type) {
	case *IFxManagerStart:
		data := IFxManagerStartOut{}
		data.Args0 = next.Start(v.inData.arg0, v.inData.arg1...)
		if v.outDataChannel != nil {
			v.outDataChannel <- data
		}
		return true, nil
	case *IFxManagerStartAll:
		data := IFxManagerStartAllOut{}
		data.Args0 = next.StartAll(v.inData.arg0)
		if v.outDataChannel != nil {
			v.outDataChannel <- data
		}
		return true, nil
	case *IFxManagerStop:
		data := IFxManagerStopOut{}
		data.Args0 = next.Stop(v.inData.arg0, v.inData.arg1...)
		if v.outDataChannel != nil {
			v.outDataChannel <- data
		}
		return true, nil
	case *IFxManagerStopAll:
		data := IFxManagerStopAllOut{}
		data.Args0 = next.StopAll(v.inData.arg0)
		if v.outDataChannel != nil {
			v.outDataChannel <- data
		}
		return true, nil
	default:
		return false, nil
	}
}
