package indexeddb

import (
	"errors"
	"fmt"
	"sync"

	"github.com/realPy/jswasm/indexeddb/store"
	"github.com/realPy/jswasm/js"
	"github.com/realPy/jswasm/object"
)

func getEventTargetError(ev js.Value) (js.Value, error) {
	if target, err := ev.GetWithErr("target"); err == nil {
		if errorResult, err := target.GetWithErr("error"); err == nil {
			return errorResult, nil
		} else {
			return js.Value{}, fmt.Errorf("error not found")
		}
	} else {
		return js.Value{}, fmt.Errorf("target not found")
	}
}

func stringFromTargetError(ev js.Value) (string, error) {
	var err error
	var jserr js.Value

	if jserr, err = getEventTargetError(ev); err == nil {

		return object.StringWithErr(jserr)

	}

	return "", err

}

var singleton sync.Once

var indexeddbinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//GetJSInterface get teh JS interface of broadcast channel
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var indexeddbinstance JSInterface
		var window js.Value
		var err error

		if window, err = js.Global().GetWithErr("window"); err == nil {
			if indexeddbinstance.objectInterface, err = window.GetWithErr("indexedDB"); err == nil {
				indexeddbinterface = &indexeddbinstance
			}

		}
	})

	return indexeddbinterface
}

type IDBOpenDBRequest struct {
	object.Object
}

func Open(name string, version int,
	automigratehandler func(IDBOpenDBRequest) error,
	onsuccesshandler func(IDBOpenDBRequest) error,
	onerrorhandler func(IDBOpenDBRequest, error)) (IDBOpenDBRequest, error) {
	var i IDBOpenDBRequest
	var waitableOpen js.Value
	var err error

	if dbi := GetJSInterface(); dbi != nil {
		if version == 0 {
			waitableOpen, err = dbi.objectInterface.CallWithErr("open", js.ValueOf(name))
		} else {
			waitableOpen, err = dbi.objectInterface.CallWithErr("open", js.ValueOf(name), js.ValueOf(version))
		}

		if err == nil {
			migrateDBFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

				if db, err := waitableOpen.GetWithErr("result"); err == nil {
					i.Object = i.SetObject(db)
					return automigratehandler(i)
				} else {
					fmt.Printf("result not found..\n")
				}

				return nil
			})

			waitableOpen.Set("onupgradeneeded", migrateDBFunc)

			onsuccess := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

				if db, err := waitableOpen.GetWithErr("result"); err == nil {
					i.Object = i.SetObject(db)
					onsuccesshandler(i)
				}

				return nil
			})

			onerror := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
				err = fmt.Errorf("Unable to open indexeddb")
				var errorString string

				if len(args) > 0 {
					event := args[0]
					if errorString, err = stringFromTargetError(event); err == nil {
						err = errors.New(errorString)
					}
				}

				onerrorhandler(i, err)
				return nil
			})

			waitableOpen.Set("onsuccess", onsuccess)
			waitableOpen.Set("onerror", onerror)

		}

	}

	return i, err
}

func (i IDBOpenDBRequest) CreateStore(name string, schema map[string]interface{}) (store.Store, error) {

	if storeObject, err := i.JSObject().CallWithErr("createObjectStore", js.ValueOf(name), schema); err == nil {

		return store.NewFromJSObject(storeObject)
	} else {
		return store.Store{}, err
	}

}

func (i IDBOpenDBRequest) GetObjectStore(table string, permission string) (store.Store, error) {
	if transaction, err := i.JSObject().CallWithErr("transaction", js.ValueOf(table), js.ValueOf(permission)); err == nil {

		if objectstore, err := transaction.CallWithErr("objectStore", js.ValueOf(table)); err == nil {
			return store.NewFromJSObject(objectstore)
		} else {
			return store.Store{}, err
		}
	} else {
		return store.Store{}, err
	}
}
