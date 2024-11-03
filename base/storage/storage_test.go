package storage

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()
	m.Run()
}

func TestNewStorage(t *testing.T) {

	t.Run("session", func(t *testing.T) {
		if obj, err := baseobject.Get(js.Global(), "window"); testingutils.AssertErr(t, err) {

			if storageobj, err := baseobject.Get(obj, "sessionStorage"); testingutils.AssertErr(t, err) {
				if stor, err := NewFromJSObject(storageobj); testingutils.AssertErr(t, err) {
					testingutils.AssertExpect(t, "[object Storage]", stor.ToString_())
				}

			}
		}

	})

	t.Run("local", func(t *testing.T) {
		if obj, err := baseobject.Get(js.Global(), "window"); testingutils.AssertErr(t, err) {

			if storageobj, err := baseobject.Get(obj, "localStorage"); testingutils.AssertErr(t, err) {
				if stor, err := NewFromJSObject(storageobj); testingutils.AssertErr(t, err) {
					testingutils.AssertExpect(t, "[object Storage]", stor.ToString_())
				}

			}
		}

	})

	t.Run("new", func(t *testing.T) {

		if stor, err := NewStorage(StorageTypeLocal); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "[object Storage]", stor.ToString_())
		}

	})

}

func TestGetItem(t *testing.T) {

	t.Run("sessionStorage", func(t *testing.T) {
		baseobject.Eval("window.sessionStorage.setItem(\"hello\",\"world\")")

		if obj, err := baseobject.Get(js.Global(), "window"); testingutils.AssertErr(t, err) {

			if storageobj, err := baseobject.Get(obj, "sessionStorage"); testingutils.AssertErr(t, err) {
				if stor, err := NewFromJSObject(storageobj); testingutils.AssertErr(t, err) {

					if str, err := stor.GetItem("hello"); testingutils.AssertErr(t, err) {
						testingutils.AssertExpect(t, "world", str)
					}
					if str, err := stor.GetItem("hello2"); testingutils.AssertErr(t, err) {

						testingutils.AssertExpect(t, nil, str)
					}
				}

			}
		}

	})

	t.Run("localStorage", func(t *testing.T) {
		baseobject.Eval("window.localStorage.setItem(\"hello\",\"world\")")
		if obj, err := baseobject.Get(js.Global(), "window"); testingutils.AssertErr(t, err) {

			if storageobj, err := baseobject.Get(obj, "localStorage"); testingutils.AssertErr(t, err) {
				if stor, err := NewFromJSObject(storageobj); testingutils.AssertErr(t, err) {

					if str, err := stor.GetItem("hello"); testingutils.AssertErr(t, err) {

						testingutils.AssertExpect(t, "world", str)
					}

					if str, err := stor.GetItem("hello2"); testingutils.AssertErr(t, err) {

						testingutils.AssertExpect(t, nil, str)
					}
				}

			}
		}

	})

}

func TestSetItem(t *testing.T) {

	t.Run("sessionStorage", func(t *testing.T) {
		baseobject.Eval("window.sessionStorage.setItem(\"hello\",\"world\")")

		if obj, err := baseobject.Get(js.Global(), "window"); testingutils.AssertErr(t, err) {

			if storageobj, err := baseobject.Get(obj, "sessionStorage"); testingutils.AssertErr(t, err) {
				if stor, err := NewFromJSObject(storageobj); testingutils.AssertErr(t, err) {

					if err := stor.SetItem("hello", "you"); testingutils.AssertErr(t, err) {
						if str, err := stor.GetItem("hello"); testingutils.AssertErr(t, err) {

							testingutils.AssertExpect(t, "you", str)
						}
					}
				}

			}
		}

	})

	t.Run("localStorage", func(t *testing.T) {
		baseobject.Eval("window.localStorage.setItem(\"hello\",\"world\")")
		if obj, err := baseobject.Get(js.Global(), "window"); testingutils.AssertErr(t, err) {

			if storageobj, err := baseobject.Get(obj, "localStorage"); testingutils.AssertErr(t, err) {
				if stor, err := NewFromJSObject(storageobj); testingutils.AssertErr(t, err) {

					if err := stor.SetItem("hello", "you"); testingutils.AssertErr(t, err) {
						if str, err := stor.GetItem("hello"); testingutils.AssertErr(t, err) {

							testingutils.AssertExpect(t, "you", str)
						}
					}
				}

			}
		}

	})

}

func TestRemoveItem(t *testing.T) {

	t.Run("sessionStorage", func(t *testing.T) {
		baseobject.Eval("window.sessionStorage.setItem(\"hello\",\"world\")")

		if obj, err := baseobject.Get(js.Global(), "window"); testingutils.AssertErr(t, err) {

			if storageobj, err := baseobject.Get(obj, "sessionStorage"); testingutils.AssertErr(t, err) {
				if stor, err := NewFromJSObject(storageobj); testingutils.AssertErr(t, err) {

					if err := stor.SetItem("objrmv", "yes"); testingutils.AssertErr(t, err) {
						if err := stor.RemoveItem("objrmv"); testingutils.AssertErr(t, err) {
							if str, err := stor.GetItem("objrmv"); testingutils.AssertErr(t, err) {

								testingutils.AssertExpect(t, nil, str)
							}

						}
					}

				}

			}
		}

	})

	t.Run("localStorage", func(t *testing.T) {
		baseobject.Eval("window.sessionStorage.setItem(\"hello\",\"world\")")

		if obj, err := baseobject.Get(js.Global(), "window"); testingutils.AssertErr(t, err) {

			if storageobj, err := baseobject.Get(obj, "localStorage"); testingutils.AssertErr(t, err) {
				if stor, err := NewFromJSObject(storageobj); testingutils.AssertErr(t, err) {

					if err := stor.SetItem("objrmv", "yes"); testingutils.AssertErr(t, err) {
						if err := stor.RemoveItem("objrmv"); testingutils.AssertErr(t, err) {
							if str, err := stor.GetItem("objrmv"); testingutils.AssertErr(t, err) {

								testingutils.AssertExpect(t, nil, str)
							}

						}
					}

				}

			}
		}

	})

}

func TestClear(t *testing.T) {

	t.Run("sessionStorage", func(t *testing.T) {

		if obj, err := baseobject.Get(js.Global(), "window"); testingutils.AssertErr(t, err) {

			if storageobj, err := baseobject.Get(obj, "sessionStorage"); testingutils.AssertErr(t, err) {
				if stor, err := NewFromJSObject(storageobj); testingutils.AssertErr(t, err) {

					if err := stor.SetItem("objclear", "yes"); testingutils.AssertErr(t, err) {
						if err := stor.Clear(); testingutils.AssertErr(t, err) {
							if str, err := stor.GetItem("objclear"); testingutils.AssertErr(t, err) {

								testingutils.AssertExpect(t, nil, str)
							}

						}
					}

				}

			}
		}

	})

	t.Run("localStorage", func(t *testing.T) {

		if obj, err := baseobject.Get(js.Global(), "window"); testingutils.AssertErr(t, err) {

			if storageobj, err := baseobject.Get(obj, "localStorage"); testingutils.AssertErr(t, err) {
				if stor, err := NewFromJSObject(storageobj); testingutils.AssertErr(t, err) {

					if err := stor.SetItem("objclear", "yes"); testingutils.AssertErr(t, err) {
						if err := stor.Clear(); testingutils.AssertErr(t, err) {
							if str, err := stor.GetItem("objclear"); testingutils.AssertErr(t, err) {

								testingutils.AssertExpect(t, nil, str)
							}

						}
					}

				}

			}
		}

	})

}

func TestKey(t *testing.T) {

	t.Run("sessionStorage", func(t *testing.T) {

		if obj, err := baseobject.Get(js.Global(), "window"); testingutils.AssertErr(t, err) {

			if storageobj, err := baseobject.Get(obj, "sessionStorage"); testingutils.AssertErr(t, err) {
				if stor, err := NewFromJSObject(storageobj); testingutils.AssertErr(t, err) {
					stor.Clear()

					if err := stor.SetItem("objkey", "yes"); testingutils.AssertErr(t, err) {

						if str, err := stor.Key(0); testingutils.AssertErr(t, err) {
							testingutils.AssertExpect(t, "objkey", str)
						}

					}

				}

			}
		}

	})

	t.Run("localStorage", func(t *testing.T) {

		if obj, err := baseobject.Get(js.Global(), "window"); testingutils.AssertErr(t, err) {

			if storageobj, err := baseobject.Get(obj, "localStorage"); testingutils.AssertErr(t, err) {
				if stor, err := NewFromJSObject(storageobj); testingutils.AssertErr(t, err) {
					stor.Clear()

					if err := stor.SetItem("objkey", "yes"); testingutils.AssertErr(t, err) {

						if str, err := stor.Key(0); testingutils.AssertErr(t, err) {
							testingutils.AssertExpect(t, "objkey", str)
						}

					}

				}

			}
		}

	})

}
