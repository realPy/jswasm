package keyobservable

import (
	"errors"
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/objectmap"
)

var singletonKeyObservable sync.Once
var ko Observable

var (
	//ErrKeyNotFound ErrKeyNotFound error
	ErrKeyNotFound = errors.New("Key not found")
)

type KeyObservableFunc func(value interface{})

// KeyObservable struct
type Observable struct {
	register    map[string]map[*KeyObservableFunc]bool
	persistData objectmap.ObjectMap
}

func KeyObservable() *Observable {

	singletonKeyObservable.Do(func() {
		ko.register = make(map[string]map[*KeyObservableFunc]bool)

		if observablePersist, err := baseobject.Get(js.Global(), "persistObservable"); err == nil && !observablePersist.IsUndefined() {
			ko.persistData, _ = objectmap.NewFromJSObject(observablePersist)
		} else {
			if observablePersist, err := objectmap.New(); err == nil {
				ko.persistData = observablePersist
				//	observablePersist.Export("persistObservable")
			} else {
				hogosuru.AssertErr(err)
			}
		}

	})

	return &ko
}

func (ko *Observable) RegisterFunc(key string, f KeyObservableFunc) *KeyObservableFunc {
	if ko.register[key] == nil {
		ko.register[key] = make(map[*KeyObservableFunc]bool)
	}
	ko.register[key][&f] = true
	return &f
}

func (ko *Observable) Put(key string, value interface{}) {

	ko.persistData.Set(key, value)

}

func (ko *Observable) Set(key string, value interface{}, persist bool) {

	if persist {
		ko.persistData.Set(key, value)
	}

	if callFuncs, ok := ko.register[key]; ok {

		if callFuncs != nil {
			for keyf, _ := range callFuncs {
				f := *keyf
				f(value)
			}
		}

	}
}

func (ko *Observable) UnRegisterFunc(key string, f *KeyObservableFunc) {

	if funcs, ok := ko.register[key]; ok {
		if _, ok := funcs[f]; ok {
			delete(ko.register[key], f)
			if len(ko.register[key]) == 0 {
				delete(ko.register, key)
			}
		}

	}
}

// Get Get key in persist array . return error if key is not found
func (ko *Observable) Get(key string) (interface{}, error) {
	if haskey, err := ko.persistData.Has(key); err == nil && haskey {
		return ko.persistData.Get(key)
	} else {
		return nil, ErrKeyNotFound
	}

}
