package internal

import (
	"io"
	"net/http"
)

type requestF func(url string) *Response

type request struct {
	url    string
	cancel <-chan struct{}
	res    chan *Response
}

type Response struct {
	Body      interface{}
	Err       error
	Cancelled bool
}

type entity struct {
	res    *Response
	ready  chan struct{}
	cancel <-chan struct{}
}

func (e *entity) call(f requestF, url string) {
	select {
	case <-e.cancel:
		e.res.Cancelled = true
		close(e.ready)
	default:
		e.res = f(url)
		close(e.ready)
	}
}

func (e *entity) deliver(ch chan<- *Response) {
	<-e.ready
	ch <- e.res
}

type Cache struct {
	f          requestF
	requests   map[string]*entity
	getRequest chan request
}

func NewCache(f requestF) *Cache {
	return &Cache{f: f, requests: make(map[string]*entity), getRequest: make(chan request)}
}

func (c *Cache) Get(url string, cancel <-chan struct{}) *Response {
	res := make(chan *Response)
	c.getRequest <- request{url, cancel, res}
	return <-res
}

func (c *Cache) Serve() {
	for req := range c.getRequest {
		ent, ok := c.requests[req.url]
		if !ok {
			ent = &entity{res: nil, ready: make(chan struct{}), cancel: req.cancel}
			c.requests[req.url] = ent
			go ent.call(c.f, req.url)
			go func() {
				select {
				case <-ent.cancel:
					delete(c.requests, req.url)
				case <-ent.ready:
				}
			}()
		}
		go ent.deliver(req.res)
	}
}

func BodyCacheFunc(url string) *Response {
	res, err := http.Get(url)
	if err != nil {
		return &Response{Body: nil, Err: err}
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return &Response{Body: nil, Err: err}
	}
	return &Response{Body: string(body), Err: nil}
}
