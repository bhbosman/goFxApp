package FxServicesSlide

import (
	"context"
	FxAppManagerService2 "github.com/bhbosman/goFxApp/FxWrappers/Services/FxAppManagerService"
	"github.com/bhbosman/gocommon/messageRouter"
	"github.com/bhbosman/gocommon/messages"
	"github.com/bhbosman/gocommon/model"
	"sort"
)

type PublishInstanceDataFor struct {
	Name string
}

type SendActionsForService struct {
	Name    string
	Actions []string
}

type IdAndName struct {
	ServiceId         model.ServiceIdentifier
	ServiceDependency model.ServiceIdentifier
	Name              string
	Active            bool
}

type FxServicesManagerData struct {
	Name              string
	Active            bool
	ServiceId         model.ServiceIdentifier
	ServiceDependency model.ServiceIdentifier
	isDirty           bool
}

type fxServiceManagerSlideData struct {
	serviceListIsDirty         bool
	ss                         map[string]*FxServicesManagerData
	messageRouter              *messageRouter.MessageRouter
	onConnectionListChange     func(connectionList []IdAndName)
	onConnectionInstanceChange func(data *SendActionsForService)
	fxManagerService           FxAppManagerService2.IFxManagerService
}

func (self *fxServiceManagerSlideData) StartAllService() {
	_ = self.fxManagerService.StartAll(context.Background())
}

func (self *fxServiceManagerSlideData) StopAllService() {
	_ = self.fxManagerService.StopAll(context.Background())
}

func (self *fxServiceManagerSlideData) StartService(name string) {
	_ = self.fxManagerService.Start(context.Background(), name)
}

func (self *fxServiceManagerSlideData) StopService(name string) {
	_ = self.fxManagerService.Stop(context.Background(), name)
}

func NewData(fxManagerService FxAppManagerService2.IFxManagerService) *fxServiceManagerSlideData {
	result := &fxServiceManagerSlideData{
		ss:               make(map[string]*FxServicesManagerData),
		messageRouter:    messageRouter.NewMessageRouter(),
		fxManagerService: fxManagerService,
	}
	_ = result.messageRouter.Add(result.handleEmptyQueue)
	_ = result.messageRouter.Add(result.handleFxServiceStarted)
	_ = result.messageRouter.Add(result.handleFxServiceStopped)
	_ = result.messageRouter.Add(result.handleFxServiceAdded)
	_ = result.messageRouter.Add(result.handleFxServiceStatus)
	_ = result.messageRouter.Add(result.handlePublishInstanceDataFor)

	return result
}

func (self *fxServiceManagerSlideData) Send(data interface{}) error {
	_, err := self.messageRouter.Route(data)
	return err
}

func (self *fxServiceManagerSlideData) handlePublishInstanceDataFor(message *PublishInstanceDataFor) error {
	if instance, ok := self.ss[message.Name]; ok {
		instance.isDirty = true
	}
	return nil
}
func (self *fxServiceManagerSlideData) handleEmptyQueue(_ *messages.EmptyQueue) error {
	if self.serviceListIsDirty {
		self.DoServiceListChange()
		self.serviceListIsDirty = false
	}
	for _, serviceInformation := range self.ss {
		if serviceInformation.isDirty {
			self.DoServiceInstanceChange(serviceInformation)
			serviceInformation.isDirty = false
		}
	}
	return nil
}

func (self *fxServiceManagerSlideData) DoServiceListChange() {
	if self.onConnectionListChange != nil {
		ss := make([]string, 0, len(self.ss))

		for key := range self.ss {
			ss = append(ss, key)
		}
		sort.Strings(ss)
		cbData := make([]IdAndName, 0, len(self.ss))

		for _, s := range ss {
			if info, ok := self.ss[s]; ok {
				idAndName := IdAndName{
					ServiceId:         info.ServiceId,
					ServiceDependency: info.ServiceDependency,
					Name:              info.Name,
					Active:            info.Active,
				}
				cbData = append(cbData, idAndName)
			}
		}
		self.onConnectionListChange(cbData)
	}
}

const StopServiceString = "Stop Service"
const StartServiceString = "Start Service"
const StopAllServiceString = "Stop All Service"
const StartAllServiceString = "Start All Service"

func (self *fxServiceManagerSlideData) DoServiceInstanceChange(data *FxServicesManagerData) {
	if self.onConnectionInstanceChange != nil {
		var actions []string
		if data.Active {
			actions = append(actions, StopServiceString)
		} else {
			actions = append(actions, StartServiceString)
		}
		actions = append(actions, []string{"-", StartAllServiceString, StopAllServiceString}...)
		endActionsForService := &SendActionsForService{
			Name:    data.Name,
			Actions: actions,
		}
		self.onConnectionInstanceChange(endActionsForService)
	}
}

func (self *fxServiceManagerSlideData) SetConnectionInstanceChange(cb func(data *SendActionsForService)) {
	self.onConnectionInstanceChange = cb
}

func (self *fxServiceManagerSlideData) SetConnectionListChange(cb func(connectionList []IdAndName)) {
	self.onConnectionListChange = cb
}

func (self *fxServiceManagerSlideData) handleFxServiceStatus(message *FxAppManagerService2.FxServiceStatus) error {
	self.ss[message.Name] = &FxServicesManagerData{
		Name:              message.Name,
		Active:            message.Active,
		ServiceId:         message.ServiceId,
		ServiceDependency: message.ServiceDependency,
		isDirty:           true,
	}
	self.serviceListIsDirty = true
	return nil
}

func (self *fxServiceManagerSlideData) handleFxServiceStarted(message *FxAppManagerService2.FxServiceStarted) error {
	if instance, ok := self.ss[message.Name]; ok {
		instance.Active = true
		instance.isDirty = true
	}
	return nil
}

func (self *fxServiceManagerSlideData) handleFxServiceStopped(message *FxAppManagerService2.FxServiceStopped) error {
	if instance, ok := self.ss[message.Name]; ok {
		instance.Active = false
		instance.isDirty = true
	}
	return nil
}

func (self *fxServiceManagerSlideData) handleFxServiceAdded(message *FxAppManagerService2.FxServiceAdded) error {
	self.ss[message.Name] = &FxServicesManagerData{
		Name:    message.Name,
		Active:  false,
		isDirty: true,
	}
	self.serviceListIsDirty = true

	return nil
}
