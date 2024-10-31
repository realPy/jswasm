package xmlhttprequest

// https://developer.mozilla.org/fr/docs/Web/API/XMLHttpRequest/XMLHttpRequest

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/initinterface"
	"github.com/realPy/hogosuru/base/progressevent"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var xhrinterface js.Value

// XMLHTTPRequest XMLHTTPRequest struct
type XMLHTTPRequest struct {
	baseobject.BaseObject
}

type XMLHTTPRequestFrom interface {
	XMLHTTPRequest_() XMLHTTPRequest
}

func (x XMLHTTPRequest) XMLHTTPRequest_() XMLHTTPRequest {
	return x
}

// GetInterface Get the JS XMLHTTPRequest Interface If nil browser doesn't implement it
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if xhrinterface, err = baseobject.Get(js.Global(), "XMLHttpRequest"); err != nil {
			xhrinterface = js.Undefined()
		}
		baseobject.Register(xhrinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return xhrinterface
}

func NewFromJSObject(obj js.Value) (XMLHTTPRequest, error) {
	var x XMLHTTPRequest
	var err error
	if si := GetInterface(); !si.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(si) {
				x.BaseObject = x.SetObject(obj)

			} else {
				err = ErrNotAXMLHTTPRequest
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return x, err
}

// New Get an XML HTTP Request
func New() (XMLHTTPRequest, error) {
	var request XMLHTTPRequest
	var objnew js.Value
	var err error
	if xhri := GetInterface(); !xhri.IsUndefined() {

		if objnew, err = baseobject.New(xhri); err == nil {
			request.BaseObject = request.SetObject(objnew)
		}

	} else {
		err = ErrNotImplemented
	}
	return request, err
}

func (x XMLHTTPRequest) Open(method string, url string) error {
	var err error
	_, err = x.Call("open", js.ValueOf(method), js.ValueOf(url))
	return err
}

func (x XMLHTTPRequest) SetRequestHeader(header string, value string) error {
	var err error
	_, err = x.Call("setRequestHeader", js.ValueOf(header), js.ValueOf(value))
	return err
}

// Send the form. Can accept a form data in args
func (x XMLHTTPRequest) Send(value ...interface{}) error {
	var err error
	var arrayJS []interface{}
	if len(value) > 0 {
		arrayJS = append(arrayJS, baseobject.GetJsValueOf(value[0]))
	}
	_, err = x.Call("send", arrayJS...)
	return err
}

func (x XMLHTTPRequest) Abort() error {
	var err error
	_, err = x.Call("abort")
	return err
}

func (x XMLHTTPRequest) setHandler(jshandlername string, handler func(i interface{})) {

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var i interface{}
		if len(args) > 0 {
			i, _ = baseobject.Discover(args[0])
		}
		handler(i)

		return nil
	})

	x.JSObject().Set(jshandlername, jsfunc)
}

// SetOnload Set OnLoad
func (x XMLHTTPRequest) SetOnload(handler func(i interface{})) {
	x.setHandler("onload", handler)
}

// SetOnAbort Set SetOnAbort
func (x XMLHTTPRequest) SetOnAbort(handler func(i interface{})) {
	x.setHandler("onabort", handler)
}

// SetOnError Set SetOnError
func (x XMLHTTPRequest) SetOnError(handler func(i interface{})) {
	x.setHandler("onerror", handler)
}

// SetOnReadyStateChange Set SetOnReadyStateChange
func (x XMLHTTPRequest) SetOnReadyStateChange(handler func(i interface{})) {
	x.setHandler("onreadystatechange", handler)
}

// SetOnProgress Set  OnProgress
func (x XMLHTTPRequest) SetOnProgress(handler func(progressevent.ProgressEvent)) {
	onprogress := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		if pe, err := progressevent.NewFromJSObject(args[0]); err == nil {
			handler(pe)
		}

		return nil
	})

	x.JSObject().Set("onprogress", onprogress)

}

func (x XMLHTTPRequest) ReadyState() (int, error) {

	return x.GetAttributeInt("readyState")
}

func (x XMLHTTPRequest) ResponseText() (string, error) {

	return x.GetAttributeString("responseText")
}

// GetResponseHeader https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/getResponseHeader
func (x XMLHTTPRequest) GetResponseHeader(header string) (string, error) {
	var responseHeader js.Value
	var err error
	if responseHeader, err = x.Call("getResponseHeader", js.ValueOf(header)); err == nil {

		if responseHeader.Type() == js.TypeString {
			return responseHeader.String(), nil
		} else {
			return "", baseobject.ErrObjectNotString
		}

	}
	return "", err
}

// GetAllResponseHeader https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/getAllResponseHeaders
func (x XMLHTTPRequest) GetAllResponseHeader() (string, error) {
	var responseHeader js.Value
	var err error
	if responseHeader, err = x.Call("getAllResponseHeaders"); err == nil {

		if responseHeader.Type() == js.TypeString {
			return responseHeader.String(), nil
		} else {
			return "", baseobject.ErrObjectNotString
		}

	}
	return "", err
}

// overrideMimeType https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/overrideMimeType
func (x XMLHTTPRequest) OverrideMimeType(mimetype string) error {
	var err error
	_, err = x.Call("overrideMimeType", js.ValueOf(mimetype))

	return err
}

// Response
func (x XMLHTTPRequest) Response() (js.Value, error) {
	return x.Get("response")
}

func (x XMLHTTPRequest) SetResponseType(typeResponse string) {

	x.JSObject().Set("responseType", js.ValueOf(typeResponse))

}

func (x XMLHTTPRequest) SetWithCredentials(withcredentials bool) {

	x.JSObject().Set("withCredentials", js.ValueOf(withcredentials))

}

func (x XMLHTTPRequest) ResponseURL() (string, error) {

	return x.GetAttributeString("responseURL")
}

func (x XMLHTTPRequest) ResponseXML() (js.Value, error) {
	var responseXML js.Value
	var err error
	if responseXML, err = x.Get("responseXML"); err == nil {
		//return a document object : TO DO IMPLEMENTATION
		return responseXML, nil

	}
	return js.Value{}, err
}

func (x XMLHTTPRequest) Status() (int, error) {
	var readystate js.Value
	var err error
	if readystate, err = x.Get("status"); err == nil {
		if readystate.Type() == js.TypeNumber {
			return readystate.Int(), nil
		} else {
			return 0, baseobject.ErrObjectNotNumber
		}

	}
	return 0, err
}

func (x XMLHTTPRequest) StatusText() (string, error) {
	var responseUrl js.Value
	var err error
	if responseUrl, err = x.Get("statusText"); err == nil {

		if responseUrl.Type() == js.TypeString {
			return responseUrl.String(), nil
		} else {
			return "", baseobject.ErrObjectNotString
		}

	}
	return "", err
}

func (x XMLHTTPRequest) uploadSetHandler(jshandlername string, handler func(XMLHTTPRequest)) {
	var uploadAbstractObject js.Value
	var err error

	if uploadAbstractObject, err = x.Get("upload"); err == nil {

		jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			handler(x)

			return nil
		})

		uploadAbstractObject.Set(jshandlername, jsfunc)
	}

}

// UploadSetOnloadstart
func (x XMLHTTPRequest) UploadSetOnloadstart(handler func(XMLHTTPRequest)) {

	x.uploadSetHandler("onloadstart", handler)

}

// UploadSetOnabort
func (x XMLHTTPRequest) UploadSetOnabort(handler func(XMLHTTPRequest)) {

	x.uploadSetHandler("onabort", handler)

}

// UploadSetOnerror
func (x XMLHTTPRequest) UploadSetOnerror(handler func(XMLHTTPRequest)) {

	x.uploadSetHandler("onerror", handler)

}

// UploadSetOnload
func (x XMLHTTPRequest) UploadSetOnload(handler func(XMLHTTPRequest)) {

	x.uploadSetHandler("onload", handler)

}

// UploadSetOntimeout
func (x XMLHTTPRequest) UploadSetOntimeout(handler func(XMLHTTPRequest)) {

	x.uploadSetHandler("ontimeout", handler)

}

// UploadSetOnloadend
func (x XMLHTTPRequest) UploadSetOnloadend(handler func(XMLHTTPRequest)) {

	x.uploadSetHandler("onloadend", handler)

}

// UploadSetOnprogress
func (x XMLHTTPRequest) UploadSetOnprogress(handler func(XMLHTTPRequest, progressevent.ProgressEvent)) {

	var uploadAbstractObject js.Value
	var err error

	if uploadAbstractObject, err = x.Get("upload"); err == nil {

		jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			if pe, err := progressevent.NewFromJSObject(args[0]); err == nil {
				handler(x, pe)
			}

			return nil
		})

		uploadAbstractObject.Set("onprogress", jsfunc)
	}

}
