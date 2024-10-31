package serviceworkerregistration

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/eventtarget"
	"github.com/realPy/hogosuru/base/initinterface"
	"github.com/realPy/hogosuru/base/navigationpreloadmanager"
	"github.com/realPy/hogosuru/base/promise"
	"github.com/realPy/hogosuru/base/pushmanager"
	"github.com/realPy/hogosuru/base/serviceworker"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var serviceworkerregistrationinterface js.Value

// GetInterface get the JS interface of serviceworkerregistration
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if serviceworkerregistrationinterface, err = baseobject.Get(js.Global(), "ServiceWorkerRegistration"); err != nil {
			serviceworkerregistrationinterface = js.Undefined()
		}
		baseobject.Register(serviceworkerregistrationinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

		promise.GetInterface()
		serviceworker.GetInterface()
		navigationpreloadmanager.GetInterface()
		pushmanager.GetInterface()
	})

	return serviceworkerregistrationinterface
}

type ServiceWorkerRegistration struct {
	eventtarget.EventTarget
}

type ServiceWorkerRegistrationFrom interface {
	ServiceWorkerRegistration_() ServiceWorkerRegistration
}

func (s ServiceWorkerRegistration) ServiceWorkerRegistration_() ServiceWorkerRegistration {
	return s
}

func NewFromJSObject(obj js.Value) (ServiceWorkerRegistration, error) {
	var s ServiceWorkerRegistration

	if si := GetInterface(); !si.IsUndefined() {
		if obj.InstanceOf(si) {
			s.BaseObject = s.SetObject(obj)
			return s, nil

		}
	}

	return s, ErrNotImplemented
}

func (s ServiceWorkerRegistration) getserviceworkerAttribute(attribute string) (serviceworker.ServiceWorker, error) {
	var err error
	var obj interface{}
	var sw serviceworker.ServiceWorker
	var ok bool

	if obj, err = s.GetAttributeGlobal(attribute); err == nil {

		if obj == nil {
			return sw, baseobject.ErrUndefinedValue
		} else {
			if sw, ok = obj.(serviceworker.ServiceWorker); !ok {
				err = serviceworker.ErrNotAServiceWorker

			}
		}

	}

	return sw, err
}

func (s ServiceWorkerRegistration) Active() (serviceworker.ServiceWorker, error) {

	return s.getserviceworkerAttribute("active")
}

func (s ServiceWorkerRegistration) Index() (int, error) {

	return s.GetAttributeInt("index")
}

func (s ServiceWorkerRegistration) Installing() (serviceworker.ServiceWorker, error) {

	return s.getserviceworkerAttribute("installing")
}

func (s ServiceWorkerRegistration) Scope() (string, error) {

	return s.GetAttributeString("scope")
}

func (s ServiceWorkerRegistration) Waiting() (serviceworker.ServiceWorker, error) {

	return s.getserviceworkerAttribute("waiting")
}

func (s ServiceWorkerRegistration) NavigationPreload() (navigationpreloadmanager.NavigationPreloadManager, error) {

	var err error
	var obj interface{}
	var n navigationpreloadmanager.NavigationPreloadManager
	var ok bool

	if obj, err = s.GetAttributeGlobal("navigationPreload"); err == nil {
		if n, ok = obj.(navigationpreloadmanager.NavigationPreloadManager); !ok {
			err = navigationpreloadmanager.ErrNotANavigationPreloadManager
		}
	}

	return n, err
}

func (s ServiceWorkerRegistration) PushManager() (pushmanager.PushManager, error) {

	var err error
	var obj interface{}
	var p pushmanager.PushManager
	var ok bool

	if obj, err = s.GetAttributeGlobal("pushManager"); err == nil {
		if p, ok = obj.(pushmanager.PushManager); !ok {
			err = pushmanager.ErrNotAPushManager
		}
	}

	return p, err
}

func (s ServiceWorkerRegistration) GetNotifications(title string, options ...map[string]interface{}) (promise.Promise, error) {

	var err error
	var obj js.Value
	var arrayJS []interface{}
	var p promise.Promise

	arrayJS = append(arrayJS, js.ValueOf(title))

	if len(options) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(options[0]))
	}

	if obj, err = s.Call("getNotifications", arrayJS...); err == nil {

		p, err = promise.NewFromJSObject(obj)

	}

	return p, err

}

func (s ServiceWorkerRegistration) ShowNotification(title string, options ...map[string]interface{}) (promise.Promise, error) {

	var err error
	var obj js.Value
	var arrayJS []interface{}
	var p promise.Promise

	arrayJS = append(arrayJS, js.ValueOf(title))

	if len(options) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(options[0]))
	}

	if obj, err = s.Call("shownotification", arrayJS...); err == nil {

		p, err = promise.NewFromJSObject(obj)

	}

	return p, err

}

func (s ServiceWorkerRegistration) Unregister() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = s.Call("getRegistration"); err == nil {

		p, err = promise.NewFromJSObject(obj)

	}
	return p, err

}

func (s ServiceWorkerRegistration) Update() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = s.Call("update"); err == nil {

		p, err = promise.NewFromJSObject(obj)

	}
	return p, err

}

func (s ServiceWorkerRegistration) UpdateViaCache() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = s.Call("updateViaCache"); err == nil {

		p, err = promise.NewFromJSObject(obj)

	}
	return p, err

}
