package routing

import (
	"errors"
	"sync"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/base/event"
	"github.com/realPy/hogosuru/base/node"
	"github.com/realPy/hogosuru/base/promise"
	"github.com/realPy/hogosuru/base/window"
)

var (
	//ErrRoutingAlreadyPresent ErrRoutingAlreadyPresent err
	ErrRoutingAlreadyPresent = errors.New("This route is already present")
)

const (
	//STDROUTE
	STDROUTE = iota
	//HASHROUTE
	HASHROUTE
)

var singletonRoute sync.Once
var route RouteMap

// Rendering interface
type Rendering interface {
	OnLoad(d document.Document, n node.Node, route string) (*promise.Promise, []Rendering)
	OnEndChildsRendering()
	OnEndChildRendering(r Rendering)
	//Node attach childs to this node
	Node(r Rendering) node.Node
	OnUnload()
}

// Router struct
type RouteMap struct {
	mode             int
	defaultRendering Rendering
	currentHashRoute string
	currentStdRoute  string
	currentRendering Rendering
	nextRendering    Rendering
	routing          map[string]Rendering
	cancel           *bool
}

// Router
func Router() *RouteMap {

	singletonRoute.Do(func() {
		route.routing = make(map[string]Rendering)
		if w, err := window.New(); err == nil {

			w.OnPopState(func(e event.Event) {

				route.onurlchange()

			})

		} else {
			println("Router " + err.Error())
		}
	})

	return &route
}

func (r *RouteMap) DefaultRendering(obj Rendering) {
	r.defaultRendering = obj
}

func (r *RouteMap) loadDefaultRendering() {
	if d, err := document.New(); hogosuru.AssertErr(err) {
		if body, err := d.Body(); hogosuru.AssertErr(err) {
			r.loadChilds(d, r.defaultRendering, body.Node)
		}
	}
}

func (r *RouteMap) Route() string {
	if r.mode == STDROUTE {
		return r.currentStdRoute
	}
	return r.currentHashRoute
}

func (r *RouteMap) SetRoute(route string) {
	if r.mode == STDROUTE {
		r.currentStdRoute = route
	}
	r.currentHashRoute = route
}

func (r *RouteMap) loadChilds(d document.Document, obj Rendering, node node.Node) promise.Promise {
	p, arrayRendering := obj.OnLoad(d, node, r.Route())

	var allpromise []interface{}
	if p != nil {
		allpromise = append(allpromise, *p)
	}

	if arrayRendering != nil {
		for _, render := range arrayRendering {
			var rthis Rendering
			rthis = render
			attachChilds := obj.Node(rthis)
			childpromise := r.loadChilds(d, rthis, attachChilds)
			if childpromise.Empty() {

				obj.OnEndChildRendering(rthis)

			} else {
				childpromise.Finally(func() {

					obj.OnEndChildRendering(rthis)
				})
				allpromise = append(allpromise, childpromise)
			}

		}
	}

	var promisewaitAll promise.Promise
	var err error
	var cancel *bool = r.cancel

	if p != nil {
		if promisewaitAll, err = promise.All(allpromise...); hogosuru.AssertErr(err) {
			promisewaitAll.Finally(func() {

				if cancel != nil {
					if *cancel == true {
						return
					}
				}

				if r.nextRendering == obj {
					if r.currentRendering != nil {
						r.currentRendering.OnUnload()
					}
					r.currentRendering = r.nextRendering

				}

				obj.OnEndChildsRendering()

			})
		}
	} else {

		if obj == r.defaultRendering {
			if promisewaitAll, err = promise.All(allpromise...); hogosuru.AssertErr(err) {
				promisewaitAll.Finally(func() {

					r.onurlchange()

				})
			}
		}

		if r.nextRendering == obj {
			if r.currentRendering != nil {
				r.currentRendering.OnUnload()
			}
			r.currentRendering = r.nextRendering

		}

		obj.OnEndChildsRendering()
	}

	return promisewaitAll
}

func (r *RouteMap) Go(newroute string) {

	if w, err := window.New(); err == nil {

		if historyObj, err := w.History(); err == nil {
			historyObj.PushState(nil, newroute, newroute)
			if r.currentRendering != r.nextRendering {
				*r.cancel = true
			}
			r.onurlchange()

		}
	}

}

func (r *RouteMap) onChangeRoute(newroute string) {
	if len(r.routing) == 0 {
		r.SetRoute(newroute)

	}

	for route, render := range r.routing {
		if newroute == route {

			r.SetRoute(newroute)
			r.LoadRendering(render)

		}
	}

}
func (r *RouteMap) LoadRendering(obj Rendering) {

	r.nextRendering = obj
	var cancel bool = false
	r.cancel = &cancel
	if d, err := document.New(); err == nil {

		if r.defaultRendering != nil {
			r.loadChilds(d, obj, r.defaultRendering.Node(r.defaultRendering))
		} else {
			if body, err := d.Body(); err == nil {
				r.loadChilds(d, obj, body.Node)

			}
		}
	}

}

func (r *RouteMap) Start(mode int) {
	r.mode = mode
	r.loadDefaultRendering()

}

func (r *RouteMap) Add(route string, obj Rendering) error {
	var err error
	if _, ok := r.routing[route]; ok {
		err = ErrRoutingAlreadyPresent
	} else {

		r.routing[route] = obj
	}
	return err
}

func (r *RouteMap) onhashechange() {
	if w, err := window.New(); err == nil {

		if l, err := w.Location(); err == nil {

			if h, err := l.Hash(); err == nil {

				if len(h) > 1 {
					r.onChangeRoute(h[1:])
				} else {
					r.onChangeRoute("")
				}

			} else {
				println("Router " + err.Error())
			}

		} else {
			println("Router " + err.Error())
		}

	} else {
		println("Router " + err.Error())
	}

}

func (r *RouteMap) onurlchange() {
	if w, err := window.New(); err == nil {

		if l, err := w.Location(); err == nil {
			var route string = ""
			var err error
			if r.mode == STDROUTE {
				route, err = l.Pathname()

			} else {
				route, err = l.Hash()

				if len(route) > 1 {
					route = route[1:]

				}
			}
			if err == nil {

				r.onChangeRoute(route)

			} else {
				println("Router " + err.Error())
			}

		} else {
			println("Router " + err.Error())
		}

	} else {
		println("Router " + err.Error())
	}

}
