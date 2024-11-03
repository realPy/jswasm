package document

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/element"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}

func TestDomain(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		if str, err := d.Domain(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "localhost", str)
		}

		if err := d.SetDomain("testing.com"); err == nil {
			t.Error("Must return error")
		} else {

			testingutils.AssertExpect(t, "Failed to set the 'domain' property on 'Document': 'testing.com' is not a suffix of 'localhost'.", err.Error())

		}

	}

}

func TestTitle(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {
		if str, err := d.Title(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "Go wasm", str)

		}

		testingutils.AssertErr(t, d.SetTitle("Hello"))

		if str, err := d.Title(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "Hello", str)

		}

	}

}

func TestBody(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {
		if b, err := d.Body(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object HTMLBodyElement]", b.ToString_())
		}

	}

}

func TestActiveElement(t *testing.T) {
	if d, err := New(); testingutils.AssertErr(t, err) {
		if el, err := d.ActiveElement(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object HTMLBodyElement]", el.ToString_())
		}
	}
}

func TestCharacterSet(t *testing.T) {
	if d, err := New(); testingutils.AssertErr(t, err) {
		if s, err := d.CharacterSet(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "UTF-8", s)
		}

	}
}

func TestChildElementCount(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {
		if c, err := d.ChildElementCount(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 1, c)
		}

	}
}

func TestChildren(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {
		if collections, err := d.Children(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 1, collections.Length())
			if item, err := collections.Item(0); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object HTMLHtmlElement]", item.(baseobject.ObjectFrom).BaseObject_().ToString_())
			}
		}

	}
}

func TestCompatMode(t *testing.T) {

	//var expectElementString []string = []string{"", ""}
	if d, err := New(); testingutils.AssertErr(t, err) {
		if compatmode, err := d.CompatMode(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "CSS1Compat", compatmode)
		}

	}
}

func TestContentType(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		if str, err := d.ContentType(); testingutils.AssertErr(t, err) {
			if str != "text/html" {
				t.Errorf("Content Type must be text/html, have %s", str)
			}
		}

	}

}
func TestDocumentElement(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		if elem, err := d.DocumentElement(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object HTMLHtmlElement]", elem.ToString_())
		}

	}

}

func TestDocumentURI(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		if uri, err := d.DocumentURI(); testingutils.AssertErr(t, err) {
			var expect string = "http://localhost"
			if !strings.Contains(uri, expect) {
				t.Errorf("Must contain %s have %s", expect, uri)
			}

		}

	}
}

func TestEmbeds(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {
		baseobject.Eval("emb=document.createElement(\"embed\");document.body.appendChild(emb)")

		if collections, err := d.Embeds(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 1, collections.Length())
			if item, err := collections.Item(0); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object HTMLEmbedElement]", item.(baseobject.ObjectFrom).BaseObject_().ToString_())
			}
		}
		baseobject.Eval("emb.parentNode.removeChild(emb)")
	}
}

func TestFirstElementChild(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		if elem, err := d.FirstElementChild(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object HTMLHtmlElement]", elem.ToString_())
		}

	}

}
func TestForms(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {
		baseobject.Eval("form=document.createElement(\"form\");document.body.appendChild(form)")

		if collections, err := d.Forms(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 1, collections.Length())
			if item, err := collections.Item(0); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object HTMLFormElement]", item.(baseobject.ObjectFrom).BaseObject_().ToString_())
			}
		}

		baseobject.Eval("form.parentNode.removeChild(form)")

	}
}

func TestFullscreenElement(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		if elem, err := d.FullscreenElement(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "", elem.ToString_())
		}

	}

}

func TestHead(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {
		if h, err := d.Head(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object HTMLHeadElement]", h.ToString_())
		}

	}

}

func TestHidden(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		if h, err := d.Hidden(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, false, h)
		}

	}

}

func TestImages(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {
		baseobject.Eval("img=document.createElement(\"img\");document.body.appendChild(img)")

		if collections, err := d.Images(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 1, collections.Length())
			if item, err := collections.Item(0); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object HTMLImageElement]", item.(baseobject.ObjectFrom).BaseObject_().ToString_())
			}
		}
		baseobject.Eval("img.parentNode.removeChild(img)")
	}
}

func TestLastElementChild(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		if elem, err := d.LastElementChild(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object HTMLHtmlElement]", elem.ToString_())
		}

	}

}

func TestLinks(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {
		baseobject.Eval("a=document.createElement(\"a\");a.href=\"testing://localhost\";document.body.appendChild(a)")

		if collections, err := d.Links(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 1, collections.Length())
			if item, err := collections.Item(0); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "testing://localhost", item.(baseobject.ObjectFrom).BaseObject_().ToString_())
			}
		}
		baseobject.Eval("a.parentNode.removeChild(a)")
	}
}

func TestPictureInPictureElement(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		if elem, err := d.PictureInPictureElement(); testingutils.AssertErr(t, err) {
			if !elem.Empty() {
				t.Error("PictureInPicture must be empty")
			}

		}

	}
}

func TestPictureInPictureEnabled(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		if ok, err := d.PictureInPictureEnabled(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, true, ok)
		}

	}
}

func TestPlugins(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {
		baseobject.Eval("emb=document.createElement(\"embed\");document.body.appendChild(emb)")

		if collections, err := d.Plugins(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 1, collections.Length())
			if item, err := collections.Item(0); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object HTMLEmbedElement]", item.(baseobject.ObjectFrom).BaseObject_().ToString_())
			}
		}
		baseobject.Eval("emb.parentNode.removeChild(emb)")
	}

}

func TestScripts(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		if collections, err := d.Scripts(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 2, collections.Length())
			if item, err := collections.Item(0); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object HTMLScriptElement]", item.(baseobject.ObjectFrom).BaseObject_().ToString_())
			}
		}

	}

}
func TestScrollingElement(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		if elem, err := d.ScrollingElement(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object HTMLHtmlElement]", elem.ToString_())
		}

	}

}
func TestVisibilityState(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		if state, err := d.VisibilityState(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "visible", state)
		}

	}

}

func TestLastModified(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		if state, err := d.LastModified(); testingutils.AssertErr(t, err) {
			if len(state) == 0 {
				t.Errorf("Must have value")
			}
		}

	}

}

func TestReadyState(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {
		time.Sleep(100 * time.Millisecond)
		if state, err := d.ReadyState(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "complete", state)
		}

	}

}

func TestReferrer(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		if state, err := d.Referrer(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "", state)
		}

	}

}
func TestURL(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		if state, err := d.URL(); testingutils.AssertErr(t, err) {
			var expect string = "http://localhost"
			if !strings.Contains(state, expect) {
				t.Errorf("Must contain %s have %s", expect, state)
			}
		}

	}

}

func TestCookie(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		if str, err := d.Cookie(); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "", str)

		}
		testingutils.AssertErr(t, d.SetCookie("hello world"))
		if str, err := d.Cookie(); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "hello world", str)
		}

	}

}

func TestAppend(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		if err := d.Append("hello"); err == nil {

			t.Error("Must return an error")

		}
	}

}

func TestCreateAttribute(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		if attr, err := d.CreateAttribute("myattr"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object Attr]", attr.ToString_())
			if n, err := attr.Name(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "myattr", n)
			}

			if n, err := attr.LocalName(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "myattr", n)
			}

			if n, err := attr.Value(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "", n)
			}

		}
	}

}

func TestCreateComment(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		if comment, err := d.CreateComment("com"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object Comment]", comment.ToString_())

		}
	}
}

func TestCreateDocumentFragment(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		if fragment, err := d.CreateDocumentFragment(); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object DocumentFragment]", fragment.ToString_())

		}
	}

}

func TestCreateHTMLElement(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		if elem, err := d.CreateHTMLElement("test"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object HTMLUnknownElement]", elem.ToString_())

		}

		if elem, err := d.CreateHTMLElement("input"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object HTMLInputElement]", elem.ToString_())

		}

		if elem, err := d.CreateHTMLElement("button"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object HTMLButtonElement]", elem.ToString_())

		}
	}

}

func TestCreateElement(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		if elem, err := d.CreateElement("test"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object HTMLUnknownElement]", elem.ToString_())

		}
		if elem, err := d.CreateElement("input"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object HTMLInputElement]", elem.ToString_())

		}

		if elem, err := d.CreateElement("button"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object HTMLButtonElement]", elem.ToString_())

		}
	}

}

func TestCreateElementNS(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		if elem, err := d.CreateElementNS("http://www.w3.org/1999/xhtml", "test"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object HTMLUnknownElement]", elem.ToString_())

		}
		if elem, err := d.CreateElementNS("http://www.w3.org/1999/xhtml", "input"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object HTMLInputElement]", elem.ToString_())

		}

		if elem, err := d.CreateElementNS("http://www.w3.org/1999/xhtml", "button"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object HTMLButtonElement]", elem.ToString_())

		}
	}

}

func TestCreateEvent(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		if event, err := d.CreateEvent("KeyboardEvent"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object KeyboardEvent]", event.ToString_())

		}

		if _, err := d.CreateEvent("testEvent"); err == nil {

			t.Error("Must return an err")

		}
	}

}

func TestCreateTextNode(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		if textnode, err := d.CreateTextNode("test"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object Text]", textnode.ToString_())

		}

	}

}

func TestGetElementsByClassName(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {
		baseobject.Eval("div=document.createElement(\"div\");document.body.appendChild(div);div.className=\"test\"")

		if collections, err := d.GetElementsByClassName("test"); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 1, collections.Length())
			if item, err := collections.Item(0); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object HTMLDivElement]", item.(baseobject.ObjectFrom).BaseObject_().ToString_())
			}
		}
		baseobject.Eval("div.parentNode.removeChild(div)")

		if collections, err := d.GetElementsByClassName("test"); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 0, collections.Length())

		}
	}

}

func TestGetElementsByTagName(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {
		baseobject.Eval("div=document.createElement(\"div\");document.body.appendChild(div);div.className=\"test\"")

		if collection, err := d.GetElementsByTagName("div"); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 1, collection.Length())

			if item, err := collection.Item(0); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object HTMLDivElement]", item.(baseobject.ObjectFrom).BaseObject_().ToString_())
			}
		}
		baseobject.Eval("div.parentNode.removeChild(div)")

		if collection, err := d.GetElementsByTagName("div"); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 0, collection.Length())

		}
	}

}

func TestGetElementsByTagNameNS(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {
		baseobject.Eval("div=document.createElement(\"div\");document.body.appendChild(div)")

		if collection, err := d.GetElementsByTagNameNS("http://www.w3.org/1999/xhtml", "div"); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 1, collection.Length())

			if item, err := collection.Item(0); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object HTMLDivElement]", item.(baseobject.ObjectFrom).BaseObject_().ToString_())
			}
		}
		baseobject.Eval("div.parentNode.removeChild(div)")

		if collection, err := d.GetElementsByTagNameNS("http://www.w3.org/1999/xhtml", "div"); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 0, collection.Length())

		}

	}

}

func TestGetElementById(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {
		baseobject.Eval("div=document.createElement(\"div\");document.body.appendChild(div);div.id=\"testid\"")

		if item, err := d.GetElementById("testid"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object HTMLDivElement]", item.ToString_())

		}

		baseobject.Eval("div.parentNode.removeChild(div)")
		if _, err := d.GetElementById("testid"); !errors.Is(err, ErrElementNotFound) {
			t.Errorf("Must return err %s", ErrElementNotFound.Error())
		}

	}

}

func TestQuerySelector(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {
		baseobject.Eval("div=document.createElement(\"div\");document.body.appendChild(div);div.id=\"testid\"")

		if item, err := d.QuerySelector("#testid"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "[object HTMLDivElement]", item.ToString_())

		}

		baseobject.Eval("div.parentNode.removeChild(div)")
		if _, err := d.QuerySelector("#testid"); !errors.Is(err, ErrElementNotFound) {
			t.Errorf("Must return err %s", ErrElementNotFound.Error())
		}

	}

}

func TestQuerySelectorAll(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {
		baseobject.Eval("div=document.createElement(\"div\");document.body.appendChild(div);div.id=\"testid\"")

		if nodeslist, err := d.QuerySelectorAll("div"); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, 1, nodeslist.Length())
			if item, err := nodeslist.Item(0); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "[object HTMLDivElement]", item.ToString_())
			}
		}

		baseobject.Eval("div.parentNode.removeChild(div)")
		if nodeslist, err := d.QuerySelectorAll("divv"); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 0, nodeslist.Length())
		}

	}

}

func TestImportNode(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {
		baseobject.Eval("div=document.createElement(\"div\");document.body.appendChild(div);div.id=\"testid\"")

		if div, err := d.GetElementById("testid"); testingutils.AssertErr(t, err) {

			if clone, err := d.ImportNode(div.Node, true); testingutils.AssertErr(t, err) {
				div.SetID("te")
				if clondelem, err := element.NewFromJSObject(clone.(baseobject.ObjectFrom).BaseObject_().JSObject()); testingutils.AssertErr(t, err) {

					testingutils.AssertExpect(t, "testid", clondelem.ID_())

				}

			}

		}

	}

}
