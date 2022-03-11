package navigator

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/initinterface"
	"github.com/realPy/hogosuru/permissions"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var navigatorinterface js.Value

//GetInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if navigatorinterface, err = baseobject.Get(js.Global(), "Navigator"); err != nil {
			navigatorinterface = js.Undefined()
		}
		baseobject.Register(navigatorinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

		permissions.GetInterface()
	})

	return navigatorinterface
}

type Navigator struct {
	baseobject.BaseObject
}

type NavigatorFrom interface {
	Navigator_() Navigator
}

func (n Navigator) Navigator_() Navigator {
	return n
}

func NewFromJSObject(obj js.Value) (Navigator, error) {
	var n Navigator

	if ni := GetInterface(); !ni.IsUndefined() {
		if obj.InstanceOf(ni) {
			n.BaseObject = n.SetObject(obj)
			return n, nil

		}
	}

	return n, ErrNotImplemented
}

func (n Navigator) CookieEnabled() (bool, error) {

	return n.GetAttributeBool("cookieEnabled")
}

func (n Navigator) Permissions() (permissions.Permissions, error) {
	var err error
	var obj interface{}
	var p permissions.Permissions
	var ok bool

	if obj, err = n.GetAttributeGlobal("permissions"); err == nil {
		if p, ok = obj.(permissions.Permissions); !ok {
			err = permissions.ErrNotAPermissions
		}
	}

	return p, err
}

/*
func (n Navigator) Credentials() (Credentials, error) {

	return n.GetAttributeBool("credentials")
}
*/

func (n Navigator) DeviceMemory() (float64, error) {

	return n.GetAttributeDouble("deviceMemory")
}

func (n Navigator) HardwareConcurrency() (int, error) {

	return n.GetAttributeInt("hardwareConcurrency")
}

func (n Navigator) UserAgent() (string, error) {

	return n.GetAttributeString("userAgent")
}

func (n Navigator) Language() (string, error) {

	return n.GetAttributeString("language")
}

func (n Navigator) Vendor() (string, error) {

	return n.GetAttributeString("vendor")
}

func (n Navigator) JavaEnabled() (bool, error) {
	return n.CallBool("javaEnabled")
}
