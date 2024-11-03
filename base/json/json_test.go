package json

import (
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}

func TestParse(t *testing.T) {

	var str = `{
		"name":"John",
		"age":30,
		"cars":[ "Ford", "BMW", "Fiat" ]
		}`

	var badstr = `{
			"name":"John",
			"age":30,
			"cars:[ "Ford", "BMW", "Fiat" ]
			}`

	if j, err := Parse(str); testingutils.AssertErr(t, err) {

		k, _ := j.Keys()
		l, _ := k.Length()
		testingutils.AssertExpect(t, 3, l)
		goValue := j.Map()

		testingutils.AssertExpect(t, "John", goValue.(map[string]interface{})["name"])

	}

	if _, err := Parse(badstr); err == nil {
		t.Error("Must give an error")
	}
}
func TestStringify(t *testing.T) {

	if str, err := Stringify(1, "hello", true); testingutils.AssertErr(t, err) {

		testingutils.AssertExpect(t, "[1,\"hello\",true]", str)

	}

}

func TestStringifyObject(t *testing.T) {

	if str, err := StringifyObject(map[string]interface{}{"hello": "world"}); testingutils.AssertErr(t, err) {

		testingutils.AssertExpect(t, "{\"hello\":\"world\"}", str)

	}
}
