package storage

// https://developer.mozilla.org/fr/docs/Mozilla/Add-ons/WebExtensions/API/storage

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var storageinterface js.Value

type StorageType string

const (
	StorageTypeSession StorageType = "sessionStorage"
	StorageTypeLocal   StorageType = "localStorage"
)

// GetInterface get the Storage interface
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if storageinterface, err = baseobject.Get(js.Global(), "Storage"); err != nil {
			storageinterface = js.Undefined()
		}
		baseobject.Register(storageinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return storageinterface
}

type Storage struct {
	baseobject.BaseObject
}

type StorageFrom interface {
	Storage_() Storage
}

func (s Storage) Storage_() Storage {
	return s
}

func NewStorage(storageType StorageType) (*Storage, error) {

	winobj, err := baseobject.Get(js.Global(), "window")
	if err != nil {
		return nil, err
	}

	storageobj, err := baseobject.Get(winobj, string(storageType))
	if err != nil {
		return nil, err
	}

	store, err := NewFromJSObject(storageobj)
	return &store, err

}

func NewFromJSObject(obj js.Value) (Storage, error) {
	var s Storage
	var err error
	if si := GetInterface(); !si.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(si) {
				s.BaseObject = s.SetObject(obj)

			} else {
				err = ErrNotAStorage
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return s, err
}

func (l Storage) SetItem(key, value string) error {
	var err error
	_, err = l.Call("setItem", js.ValueOf(key), js.ValueOf(value))
	return err
}

func (l Storage) GetItem(key string) (interface{}, error) {
	var err error
	var itemObject js.Value
	var ret interface{}

	if itemObject, err = l.Call("getItem", js.ValueOf(key)); err == nil {
		if !itemObject.IsUndefined() && !itemObject.IsNull() {
			if itemObject.Type() == js.TypeString {
				ret = itemObject.String()
			}
		}

	}
	return ret, err
}

func (l Storage) RemoveItem(key string) error {
	var err error
	_, err = l.Call("removeItem", js.ValueOf(key))
	return err
}

func (l Storage) Clear() error {
	var err error
	_, err = l.Call("clear")
	return err
}
func (l Storage) Key(index int) (interface{}, error) {
	var err error
	var itemObject js.Value
	var ret interface{}

	if itemObject, err = l.Call("key", js.ValueOf(index)); err == nil {
		if !itemObject.IsUndefined() && !itemObject.IsNull() {
			if itemObject.Type() == js.TypeString {
				ret = itemObject.String()
			}
		}
	}
	return ret, err
}
