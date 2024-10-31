package baseobject

import (
	"errors"
	"strings"
	"syscall/js"
	"testing"
)

func TestEval(t *testing.T) {

	t.Run("Eval can execute js code", func(t *testing.T) {

		if _, err := Eval("mysum=2+2"); err != nil {
			t.Errorf("Must execute js code :%s", err.Error())
		} else {
			if intvalue, err := Get(js.Global(), "mysum"); err == nil {

				if intvalue.Type() == js.TypeNumber {

					if intvalue.Int() != 4 {
						t.Errorf("Must be equal 4 got %d", intvalue.Int())
					}

				} else {
					t.Errorf("Must be a number")

				}

			} else {
				t.Errorf("Can't get mysum :%s", err.Error())
			}
		}

	})

	t.Run("Syntax error", func(t *testing.T) {

		if _, err := Eval("mysum=2+"); err == nil {
			t.Errorf("Must return error Syntax")
		} else {
			if err.Error() != "Unexpected end of input" {
				t.Errorf("Must return Unexpected end of input %s", err.Error())
			}

		}

	})
}

func TestRegister(t *testing.T) {
	errorinterface := js.Global().Get("Error")
	Register(errorinterface, func(v js.Value) (interface{}, error) {
		return "Construct", nil
	})

	if _, ok := registry["Error"]; !ok {
		t.Errorf("Must contain the func constructor")
	}

}

func TestDiscover(t *testing.T) {
	var err error
	var obj js.Value

	errorinterface := js.Global().Get("Error")
	Register(errorinterface, func(v js.Value) (interface{}, error) {
		return BaseObject{}, nil
	})

	Eval("testerror=new Error('an error')")

	if obj, err = Get(js.Global(), "testerror"); err == nil {

		_, err = Discover(obj)

	}

	if err != nil {
		t.Errorf(err.Error())
	}

}

func TestString(t *testing.T) {
	var err error
	var obj js.Value

	Eval("testerror=new Error('an error')")

	if obj, err = Get(js.Global(), "testerror"); err == nil {
		if String(obj) != "<object>" {
			t.Errorf("must Be <object>")
		}

	}
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestToStringWithErr(t *testing.T) {
	var err error
	var obj js.Value

	Eval("testerror=new Error('an error')")

	if obj, err = Get(js.Global(), "testerror"); err == nil {
		var str string
		if str, err = ToStringWithErr(obj); err == nil {
			if str != "Error: an error" {
				t.Errorf("Must be Error: an error")
			}
		}

	}

	if err != nil {
		t.Errorf(err.Error())
	}

}

func TestNewFromJSObject(t *testing.T) {
	var err error
	var obj js.Value

	Eval("testerror=new Error('an error')")

	if obj, err = Get(js.Global(), "testerror"); err == nil {
		var b BaseObject

		if b, err = NewFromJSObject(obj); err == nil {
			if b.object == nil {
				t.Errorf("JS Object must be attached")
			}
		}

	}

	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestSelf(t *testing.T) {
	if obj, err := Self(); err == nil {
		if b, ok := obj.(ObjectFrom); ok {

			if !strings.EqualFold(b.BaseObject_().ConstructName_(), "Window") {
				t.Errorf("Self must be a Window")
			}

		} else {
			t.Errorf("Self must be a object")
		}
	}

}

func TestSetObject(t *testing.T) {

	var err error
	var obj js.Value

	Eval("testerror=new Error('an error')")

	if obj, err = Get(js.Global(), "testerror"); err == nil {
		var b BaseObject = BaseObject{}

		b = b.SetObject(obj)
		if b.object == nil {
			t.Errorf("JS Object must be attached")
		}

	}

	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestEmpty(t *testing.T) {

	var err error
	var obj js.Value
	var b BaseObject = BaseObject{}

	if !b.Empty() {
		t.Errorf("Must be Empty")
	}

	Eval("testerror=new Error('an error')")

	if obj, err = Get(js.Global(), "testerror"); err == nil {

		b = b.SetObject(obj)
		if b.Empty() {
			t.Errorf("Base Object must not be empty")
		}

	}

	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestBaseObjectDiscover(t *testing.T) {

	var err error
	var obj js.Value
	var b BaseObject = BaseObject{}
	errorinterface := js.Global().Get("Error")

	Register(errorinterface, func(v js.Value) (interface{}, error) {
		return BaseObject{}, nil
	})

	Eval("testerror=new Error('an error')")

	if obj, err = Get(js.Global(), "testerror"); err == nil {

		if b, err = NewFromJSObject(obj); err == nil {
			_, err = b.Discover()

		}

	}

	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestJSObject(t *testing.T) {

	var obj js.Value
	var err error
	var b BaseObject = BaseObject{}

	obj = b.JSObject()
	if obj.Type() != js.TypeUndefined {
		t.Errorf("JSObject Must be undefined")
	}

	if err != nil {
		t.Errorf(err.Error())
	}

	Eval("testerror=new Error('an error')")

	if obj, err = Get(js.Global(), "testerror"); err == nil {

		if b, err = NewFromJSObject(obj); err == nil {

			obj = b.JSObject()
			if obj.Type() != js.TypeObject {
				t.Errorf("JSObject Must be object")
			}

		}

	}

	if err != nil {
		t.Errorf(err.Error())
	}

	Eval("teststring=\"hello\"")

	if obj, err = Get(js.Global(), "teststring"); err == nil {

		if b, err = NewFromJSObject(obj); err == nil {

			obj = b.JSObject()
			if obj.Type() != js.TypeString {
				t.Errorf("JSObject Must be string")
			}

		}

	}

	if err != nil {
		t.Errorf(err.Error())
	}

}

func TestBaseObjectString(t *testing.T) {

	var obj js.Value
	var err error
	var b BaseObject

	Eval("testerror=new Error('an error')")

	if obj, err = Get(js.Global(), "testerror"); err == nil {

		if b, err = NewFromJSObject(obj); err == nil {

			if b.String() != "<object>" {
				t.Errorf("Must be <object>")
			}

		}

	}
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestBaseObjectToString(t *testing.T) {

	var obj js.Value
	var err error
	var b BaseObject

	Eval("testerror=new Error('an error')")

	if obj, err = Get(js.Global(), "testerror"); err == nil {

		if b, err = NewFromJSObject(obj); err == nil {
			var str string

			if str, err = b.ToString(); err == nil {
				if str != "Error: an error" {
					t.Errorf("Must be Error: an error")
				}

			}
		}

	}
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestValue(t *testing.T) {

	var obj js.Value
	var err error
	var b BaseObject

	Eval("testerror=new Error('an error')")

	if obj, err = Get(js.Global(), "testerror"); err == nil {

		if b, err = NewFromJSObject(obj); err == nil {
			var str string
			str = b.Value()

			if str != "<object>" {
				t.Errorf("Must be <object>")
			}

		}

	}
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestLength(t *testing.T) {

	var obj js.Value
	var err error
	var b BaseObject

	Eval("testerror=new Error('an error')")

	if obj, err = Get(js.Global(), "testerror"); err == nil {

		if b, err = NewFromJSObject(obj); err == nil {
			var i int
			i = b.Length()

			if i != 0 {
				t.Errorf("Must be 0")
			}

		}

	}

	if err != nil {
		t.Errorf(err.Error())
	}

	Eval("testarray=new Array(2)")

	if obj, err = Get(js.Global(), "testarray"); err == nil {

		if b, err = NewFromJSObject(obj); err == nil {
			var i int
			i = b.Length()

			if i != 2 {
				t.Errorf("Must be 2")
			}

		}

	}
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestImplement(t *testing.T) {
	var obj js.Value
	var err error
	var b BaseObject

	if obj, err = Get(js.Global(), "window"); err == nil {

		if b, err = NewFromJSObject(obj); err == nil {
			var ok bool
			if ok, err = b.Implement("alert"); err == nil {
				if !ok {
					t.Errorf("window must implement alert")
				}
			}

			if err != nil {
				t.Errorf(err.Error())
			}

			if ok, err = b.Implement("alerte"); err == nil {
				if ok {
					t.Errorf("window must not implement alerte")
				}
			}

			if err != nil {
				t.Errorf(err.Error())
			}

		}

	}
	if err != nil {
		t.Errorf(err.Error())
	}

}
func TestClass(t *testing.T) {
	var obj js.Value
	var err error
	var b BaseObject

	if obj, err = Get(js.Global(), "window"); err == nil {

		if b, err = NewFromJSObject(obj); err == nil {

			var classname string
			if classname, err = b.Class(); err == nil {
				if classname != "Window" {
					t.Errorf("window must be Windows class")
				}
			}

			if err != nil {
				t.Errorf(err.Error())
			}

		}

	}
	if err != nil {
		t.Errorf(err.Error())
	}

}

func TestSetFunc(t *testing.T) {
	var obj js.Value
	var err error
	var b BaseObject

	if obj, err = Get(js.Global(), "window"); err == nil {

		if b, err = NewFromJSObject(obj); err == nil {

			b.SetFunc("helloworld", func(this js.Value, args []js.Value) interface{} {

				return "Hello World"
			})
			var ok bool
			if ok, err = b.Implement("helloworld"); err == nil {
				if !ok {
					t.Errorf("window must be implement helloworld")
				}

			}

			if err != nil {
				t.Errorf(err.Error())
			}

			var objstr js.Value

			if objstr, err = b.Call("helloworld"); err == nil {
				if objstr.Type() == js.TypeString {
					if objstr.String() != "Hello World" {
						t.Errorf("Must return Hello World")
					}
				} else {
					err = errors.New("Must be string")
				}

			}

			if err != nil {
				t.Errorf(err.Error())
			}

		}

	}
	if err != nil {
		t.Errorf(err.Error())
	}

}

func TestExport(t *testing.T) {
	var obj js.Value
	var err error
	var b BaseObject

	Eval("testarray=new Array(2)")

	if obj, err = Get(js.Global(), "testarray"); err == nil {

		if b, err = NewFromJSObject(obj); err == nil {

			b.Export("arrayGoExported")

			var objExported js.Value
			if objExported, err = Get(js.Global(), "arrayGoExported"); err == nil {
				if objExported.IsUndefined() {
					t.Errorf("Base Object is not exported")
				}

			} else {

				t.Errorf(err.Error())

			}

		} else {

			t.Errorf(err.Error())

		}

	} else {

		t.Errorf(err.Error())

	}

}

func TestDelete(t *testing.T) {

	var err error
	var obj js.Value
	var b BaseObject

	Eval("customobject=new Object()")
	Eval("customobject.Value=\"HelloWorld\"")

	if obj, err = Get(js.Global(), "customobject"); err == nil {

		if b, err = NewFromJSObject(obj); err == nil {
			Delete(b.JSObject(), "Value")

			v, _ := b.Get("Value")
			if !v.IsUndefined() {
				t.Error("delete should set undefined")
			}

		}
	}

}
func TestGetAttributeString(t *testing.T) {

	var err error
	var obj js.Value
	var b BaseObject
	var v string

	Eval("customobject=new Object()")
	Eval("customobject.Value=\"HelloWorld\"")
	Eval("customobject.Event=New Event()")

	if obj, err = Get(js.Global(), "customobject"); err == nil {

		if b, err = NewFromJSObject(obj); err == nil {
			if v, err = b.GetAttributeString("Value"); err == nil {
				if v != "HelloWorld" {
					t.Errorf("Must be equal to HelloWorld")
				}
			} else {

				t.Errorf(err.Error())

			}

			if _, err := b.GetAttributeString("Event"); !errors.Is(err, ErrUndefinedValue) {
				t.Errorf("Must Return %s", ErrUndefinedValue.Error())

			}

			if _, err := b.GetAttributeString("Eventp"); !errors.Is(err, ErrUndefinedValue) {
				t.Errorf("Must Return %s", ErrUndefinedValue.Error())
			}
		} else {

			t.Errorf(err.Error())

		}

	} else {

		t.Errorf(err.Error())

	}

}

func TestGetAttributeGlobal(t *testing.T) {
	var w js.Value
	var err error
	var l interface{}
	var b BaseObject
	var bf ObjectFrom
	var ok bool

	if w, err = Get(js.Global(), "window"); err == nil {
		if b, err = NewFromJSObject(w); err == nil {
			if l, err = b.GetAttributeGlobal("location"); err == nil {
				if bf, ok = l.(ObjectFrom); ok {

					if bf.BaseObject_().Class_() != "Location" {
						t.Errorf("Must an Location Object")
					}
				} else {
					t.Errorf("Must an ObjectFrom")
				}

			} else {

				t.Errorf(err.Error())

			}
		}

	} else {

		t.Errorf(err.Error())

	}

}
func TestSetAttributeString(t *testing.T) {
	var obj js.Value
	var err error
	var b BaseObject

	Eval("customobject=new Object()")

	if obj, err = Get(js.Global(), "customobject"); err == nil {

		if b, err = NewFromJSObject(obj); err == nil {
			if err = b.SetAttributeString("Value", "custom"); err == nil {
				var v string
				if v, err = b.GetAttributeString("Value"); err == nil {
					if v != "custom" {
						t.Errorf("Must be equal to custom have %s", v)
					}
				} else {

					t.Errorf(err.Error())

				}

			} else {

				t.Errorf(err.Error())

			}

		} else {

			t.Errorf(err.Error())

		}

	} else {

		t.Errorf(err.Error())

	}

}

func TestGetAttributeBool(t *testing.T) {

	var err error
	var obj js.Value
	var b BaseObject
	var v bool

	Eval("customobject=new Object()")
	Eval("customobject.Value=true")
	Eval("customobject.Event=New Event()")

	if obj, err = Get(js.Global(), "customobject"); err == nil {

		if b, err = NewFromJSObject(obj); err == nil {
			if v, err = b.GetAttributeBool("Value"); err == nil {
				if v != true {
					t.Errorf("Must be equal true")
				}
			} else {

				t.Errorf(err.Error())

			}

			if _, err := b.GetAttributeBool("Event"); !errors.Is(err, ErrObjectNotBool) {
				t.Errorf("Must Return %s", ErrObjectNotBool.Error())

			}

			if _, err := b.GetAttributeBool("Eventp"); !errors.Is(err, ErrObjectNotBool) {
				t.Errorf("Must Return %s", ErrObjectNotBool.Error())
			}

		} else {

			t.Errorf(err.Error())

		}

	} else {

		t.Errorf(err.Error())

	}

}

func TestSetAttributeBool(t *testing.T) {
	var obj js.Value
	var err error
	var b BaseObject

	Eval("customobject=new Object()")

	if obj, err = Get(js.Global(), "customobject"); err == nil {

		if b, err = NewFromJSObject(obj); err == nil {
			if err = b.SetAttributeBool("Value", true); err == nil {
				var v bool
				if v, err = b.GetAttributeBool("Value"); err == nil {
					if v != true {
						t.Errorf("Must be equal to true")
					}
				} else {

					t.Errorf(err.Error())

				}

			} else {

				t.Errorf(err.Error())

			}

		} else {

			t.Errorf(err.Error())

		}

	} else {

		t.Errorf(err.Error())

	}

}

func TestGetAttributeInt(t *testing.T) {

	var err error
	var obj js.Value
	var b BaseObject
	var v int

	Eval("customobject=new Object()")
	Eval("customobject.Value=69")
	Eval("customobject.Event=New Event()")

	if obj, err = Get(js.Global(), "customobject"); err == nil {

		if b, err = NewFromJSObject(obj); err == nil {
			if v, err = b.GetAttributeInt("Value"); err == nil {
				if v != 69 {
					t.Errorf("Must be equal to 69")
				}
			} else {

				t.Errorf(err.Error())

			}

			if _, err := b.GetAttributeInt("Event"); !errors.Is(err, ErrUndefinedValue) {
				t.Errorf("Must Return %s", ErrUndefinedValue.Error())

			}

			if _, err := b.GetAttributeInt("Eventp"); !errors.Is(err, ErrUndefinedValue) {
				t.Errorf("Must Return %s", ErrUndefinedValue.Error())
			}
		} else {

			t.Errorf(err.Error())

		}

	} else {

		t.Errorf(err.Error())

	}

}

func TestSetAttributeInt64(t *testing.T) {
	var obj js.Value
	var err error
	var b BaseObject

	Eval("customobject=new Object()")

	if obj, err = Get(js.Global(), "customobject"); err == nil {

		if b, err = NewFromJSObject(obj); err == nil {
			if err = b.SetAttributeInt64("Value", 69); err == nil {
				var v int64
				if v, err = b.GetAttributeInt64("Value"); err == nil {
					if v != 69 {
						t.Errorf("Must be equal to 69")
					}
				} else {

					t.Errorf(err.Error())

				}

			} else {

				t.Errorf(err.Error())

			}

		} else {

			t.Errorf(err.Error())

		}

	} else {

		t.Errorf(err.Error())

	}

}

func TestSetAttributeInt(t *testing.T) {
	var obj js.Value
	var err error
	var b BaseObject

	Eval("customobject=new Object()")

	if obj, err = Get(js.Global(), "customobject"); err == nil {

		if b, err = NewFromJSObject(obj); err == nil {
			if err = b.SetAttributeInt("Value", 69); err == nil {
				var v int
				if v, err = b.GetAttributeInt("Value"); err == nil {
					if v != 69 {
						t.Errorf("Must be equal to 69")
					}
				} else {

					t.Errorf(err.Error())

				}

			} else {

				t.Errorf(err.Error())

			}

		} else {

			t.Errorf(err.Error())

		}

	} else {

		t.Errorf(err.Error())

	}

}

func TestGetAttributeDouble(t *testing.T) {

	var err error
	var obj js.Value
	var b BaseObject
	var v float64

	Eval("customobject=new Object()")
	Eval("customobject.Value=3.8")
	Eval("customobject.Event=New Event()")
	if obj, err = Get(js.Global(), "customobject"); err == nil {

		if b, err = NewFromJSObject(obj); err == nil {
			if v, err = b.GetAttributeDouble("Value"); err == nil {
				if v != 3.8 {
					t.Errorf("Must be equal to 3.8")
				}
			} else {

				t.Errorf(err.Error())

			}

			if _, err := b.GetAttributeDouble("Event"); !errors.Is(err, ErrUndefinedValue) {
				t.Errorf("Must Return %s", ErrUndefinedValue.Error())

			}

			if _, err := b.GetAttributeDouble("Eventp"); !errors.Is(err, ErrUndefinedValue) {
				t.Errorf("Must Return %s", ErrUndefinedValue.Error())
			}

		} else {

			t.Errorf(err.Error())

		}

	} else {

		t.Errorf(err.Error())

	}

}
func TestCallInt64(t *testing.T) {

	var err error
	var obj js.Value
	var b BaseObject

	Eval("customobject=new Object()")

	if obj, err = Get(js.Global(), "customobject"); err == nil {

		if b, err = NewFromJSObject(obj); err == nil {

			b.SetFunc("get64", func(this js.Value, args []js.Value) interface{} {
				return 1234
			})

			if v, err := b.CallInt64("get64"); err == nil {

				if v != 1234 {
					t.Errorf("Must Return 1234")
				}
			} else {

				t.Errorf(err.Error())

			}

			if v, err := b.CallInt64("get64"); err == nil {

				if v != 1234 {
					t.Errorf("Must Return 1234")
				}
			} else {

				t.Errorf(err.Error())

			}
		} else {

			t.Errorf(err.Error())

		}
	} else {

		t.Errorf(err.Error())

	}

}

func TestCallBool(t *testing.T) {

	var err error
	var obj js.Value
	var b BaseObject
	var v bool

	Eval("customobject=new Object()")

	if obj, err = Get(js.Global(), "customobject"); err == nil {

		if b, err = NewFromJSObject(obj); err == nil {

			b.SetFunc("getbool", func(this js.Value, args []js.Value) interface{} {
				return true
			})

			if v, err = b.CallBool("getbool"); err == nil {

				if v != true {
					t.Errorf("Must Return true")
				}
			}
		}
	}
	if err != nil {
		t.Errorf(err.Error())
	}

}

func TestGoValue(t *testing.T) {
	var err error
	var obj js.Value

	Eval("customobject=1")

	if obj, err = Get(js.Global(), "customobject"); err == nil {

		v := GoValue_(obj)
		if vi, ok := v.(int64); ok {
			if vi != 1 {
				t.Errorf("Value not match")
			}
		} else {
			t.Errorf("Must be int")
		}

	}

	Eval("customobject=1.6")

	if obj, err = Get(js.Global(), "customobject"); err == nil {

		v := GoValue_(obj)
		if vi, ok := v.(float64); ok {
			if vi != 1.6 {
				t.Errorf("Value not match")
			}
		} else {
			t.Errorf("Must be float64")
		}

	}

	Eval("customobject=\"string\"")

	if obj, err = Get(js.Global(), "customobject"); err == nil {

		v := GoValue_(obj)
		if vi, ok := v.(string); ok {
			if vi != "string" {
				t.Errorf("Value not match")
			}
		} else {
			t.Errorf("Must be string")
		}

	}

	Eval("customobject=true")

	if obj, err = Get(js.Global(), "customobject"); err == nil {

		v := GoValue_(obj)
		if vi, ok := v.(bool); ok {
			if vi != true {
				t.Errorf("Value not match")
			}
		} else {
			t.Errorf("Must be bool")
		}

	}

	Eval("customobject=new Array(1)")

	if obj, err = Get(js.Global(), "customobject"); err == nil {

		v := GoValue_(obj)
		if _, ok := v.(ObjectFrom); !ok {
			t.Errorf("Must be Object")
		}
	}

}

func TestMain(m *testing.M) {
	SetSyscall()
	m.Run()
}
