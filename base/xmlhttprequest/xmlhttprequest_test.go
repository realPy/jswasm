package xmlhttprequest

import (
	"syscall/js"
	"testing"
	"time"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/formdata"
	"github.com/realPy/hogosuru/base/json"
	"github.com/realPy/hogosuru/base/progressevent"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}

func TestNew(t *testing.T) {

	if xhr, err := New(); testingutils.AssertErr(t, err) {

		testingutils.AssertExpect(t, "[object XMLHttpRequest]", xhr.ToString_())

	}
}

func TestNewFromJSObject(t *testing.T) {

	baseobject.Eval("xhr=new XMLHttpRequest()")

	if obj, err := baseobject.Get(js.Global(), "xhr"); testingutils.AssertErr(t, err) {
		if d, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object XMLHttpRequest]", d.ToString_())

		}
	}

}

func TestGetRequest(t *testing.T) {

	t.Run("get request with ResponseHeader", func(t *testing.T) {
		var io chan bool = make(chan bool)

		if xhr, err := New(); testingutils.AssertErr(t, err) {

			err := xhr.Open("GET", "http://localhost/get")
			testingutils.AssertErr(t, err)

			xhr.SetOnload(func(i interface{}) {

				if status, err := xhr.Status(); testingutils.AssertErr(t, err) {
					testingutils.AssertExpect(t, status, 200)
				}

				if header, err := xhr.GetResponseHeader("Content-Type"); testingutils.AssertErr(t, err) {

					testingutils.AssertExpect(t, "application/json", header)

				}
				if text, err := xhr.ResponseText(); testingutils.AssertErr(t, err) {

					if j, err := json.Parse(text); testingutils.AssertErr(t, err) {
						goValue := j.Map()

						url := goValue.(map[string]interface{})["url"]

						if url != nil {
							testingutils.AssertExpect(t, url, "http://localhost/get")
							io <- true
						} else {
							t.Error("No url present")
						}

					}

				}

			})

			xhr.Send()

			select {
			case <-io:
			case <-time.After(time.Duration(1000) * time.Millisecond):
				t.Errorf("No message channel receive")
			}
		}

	})

	t.Run("get request with GetAllResponseHeader", func(t *testing.T) {
		var io chan bool = make(chan bool)

		if xhr, err := New(); testingutils.AssertErr(t, err) {

			err := xhr.Open("GET", "http://localhost/get")
			testingutils.AssertErr(t, err)

			xhr.SetOnload(func(i interface{}) {

				if status, err := xhr.Status(); testingutils.AssertErr(t, err) {
					testingutils.AssertExpect(t, status, 200)
				}

				if headers, err := xhr.GetAllResponseHeader(); testingutils.AssertErr(t, err) {

					testingutils.AssertStringContains(t, "content-type: application/json\r\n", headers)
					testingutils.AssertStringContains(t, "content-length: ", headers)

				}
				if text, err := xhr.ResponseText(); testingutils.AssertErr(t, err) {

					if j, err := json.Parse(text); testingutils.AssertErr(t, err) {
						goValue := j.Map()

						url := goValue.(map[string]interface{})["url"]

						if url != nil {
							testingutils.AssertExpect(t, url, "http://localhost/get")
							io <- true
						} else {
							t.Error("No url present")
						}

					}

				}

			})

			xhr.Send()

			select {
			case <-io:
			case <-time.After(time.Duration(1000) * time.Millisecond):
				t.Errorf("No message channel receive")
			}
		}
	})

}

func TestPostRequest(t *testing.T) {

	var io chan bool = make(chan bool)

	if xhr, err := New(); testingutils.AssertErr(t, err) {

		err := xhr.Open("POST", "http://localhost/post")
		testingutils.AssertErr(t, err)

		xhr.SetOnload(func(i interface{}) {
			progressevent.GetInterface()
			testingutils.AssertExpect(t, "[object ProgressEvent]", i.(baseobject.BaseObject).ToString_())

			if status, err := xhr.Status(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, status, 200)
			}

			if header, err := xhr.GetResponseHeader("Content-Type"); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, "application/json", header)

			}
			if text, err := xhr.ResponseText(); testingutils.AssertErr(t, err) {

				if j, err := json.Parse(text); testingutils.AssertErr(t, err) {
					goValue := j.Map()

					form := goValue.(map[string]interface{})["form"]

					if form != nil {
						customValue := form.(map[string]interface{})["data"]

						testingutils.AssertExpect(t, customValue, "testing")
						io <- true

					} else {
						t.Error("No form present")
					}

				}

			}

		})
		f, _ := formdata.New()
		f.Append("data", "testing")

		xhr.Send(f)

		select {
		case <-io:
		case <-time.After(time.Duration(1000) * time.Millisecond):
			t.Errorf("No message channel receive")
		}
	}
}

func TestSetRequestHeader(t *testing.T) {

	var io chan bool = make(chan bool)

	if xhr, err := New(); testingutils.AssertErr(t, err) {

		err := xhr.Open("GET", "http://localhost/get")
		testingutils.AssertErr(t, err)

		xhr.SetOnload(func(i interface{}) {
			progressevent.GetInterface()
			testingutils.AssertExpect(t, "[object ProgressEvent]", i.(progressevent.ProgressEvent).ToString_())

			if status, err := xhr.Status(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, status, 200)
			}

			if header, err := xhr.GetResponseHeader("Content-Type"); testingutils.AssertErr(t, err) {

				testingutils.AssertExpect(t, "application/json", header)

			}
			if text, err := xhr.ResponseText(); testingutils.AssertErr(t, err) {

				if j, err := json.Parse(text); testingutils.AssertErr(t, err) {
					goValue := j.Map()

					headers := goValue.(map[string]interface{})["headers"]

					if headers != nil {
						customValue := headers.(map[string]interface{})["Xcustomvalue"]

						if customValue != nil {
							if customValue.(string) == "Test" {
								io <- true
							} else {
								t.Errorf("Xcustomvalue not match %s", customValue.(string))
							}
						} else {
							t.Error("No Xcustomvalue headers present")
						}

					} else {
						t.Error("No headers present")
					}

				}

			}

		})

		xhr.SetRequestHeader("Xcustomvalue", "Test")
		xhr.Send()

		select {
		case <-io:
		case <-time.After(time.Duration(2000) * time.Millisecond):
			t.Errorf("No message channel receive")
		}
	}
}

func TestOnError(t *testing.T) {

	var io chan bool = make(chan bool)

	if xhr, err := New(); testingutils.AssertErr(t, err) {
		progressevent.GetInterface()
		err := xhr.Open("GET", "m://httpbin.org/get")
		testingutils.AssertErr(t, err)
		xhr.SetOnError(func(i interface{}) {
			testingutils.AssertExpect(t, "[object ProgressEvent]", i.(progressevent.ProgressEvent).ToString_())

			io <- true
		})

		xhr.Send()

		select {
		case <-io:
		case <-time.After(time.Duration(1000) * time.Millisecond):
			t.Errorf("No message channel receive")
		}
	}
}

func TestOnAbort(t *testing.T) {

	var io chan bool = make(chan bool)

	if xhr, err := New(); testingutils.AssertErr(t, err) {
		progressevent.GetInterface()
		err := xhr.Open("GET", "http://localhost/get")
		testingutils.AssertErr(t, err)

		xhr.SetOnAbort(func(i interface{}) {
			testingutils.AssertExpect(t, "[object ProgressEvent]", i.(progressevent.ProgressEvent).ToString_())

			go func() {
				io <- true
			}()

		})

		xhr.Send()
		xhr.Abort() //call SetOnAbort in the same current go routine (be carefull on deadlock channel)

		select {
		case <-io:
		case <-time.After(time.Duration(1000) * time.Millisecond):
			t.Errorf("No message channel receive")
		}
	}
}
