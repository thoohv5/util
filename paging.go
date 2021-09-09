package util

import (
	"reflect"
)

const (
	// DefaultStart 默认起始页
	DefaultStart uint32 = 0
	// DefaultLimit 默认分页长度
	DefaultLimit uint32 = 20
	// MaxLimit 最大分页长度
	MaxLimit uint32 = 50
)

type paging struct {
	DefaultStart uint32
	DefaultLimit uint32
	MaxLimit     uint32

	Start uint32
	Limit uint32
}

type IPage interface {
	ParsePage(page *BasePage) IPage
	BuildResponse(list interface{}) *BaseResp
}

func New(opts ...Option) IPage {
	p := new(paging)
	p.DefaultStart = DefaultStart
	p.DefaultLimit = DefaultLimit
	p.MaxLimit = MaxLimit
	for _, o := range opts {
		o.apply(p)
	}
	return p
}

type Option interface {
	apply(*paging)
}
type optionFunc func(*paging)

func (f optionFunc) apply(o *paging) {
	f(o)
}

func WithMaxLimit(maxLimit uint32) Option {
	return optionFunc(func(p *paging) {
		p.MaxLimit = maxLimit
	})
}

func WithDefaultLimit(defaultLimit uint32) Option {
	return optionFunc(func(p *paging) {
		p.DefaultLimit = defaultLimit
	})
}

type BasePage struct {
	Start uint32 `json:"start" form:"start,default=0"`
	Limit uint32 `json:"limit" form:"limit,default=15"`
}

func (p *paging) ParsePage(page *BasePage) IPage {

	start := page.Start
	limit := page.Limit

	if start < 1 {
		start = p.DefaultStart
	}

	if limit < 1 || limit > p.MaxLimit {
		limit = p.DefaultLimit
	}

	limit++

	p.Start = start
	p.Limit = limit
	page.Start = p.Start
	page.Limit = p.Start

	return p
}

type BaseResp struct {
	HasMore bool        `json:"hasMore"`
	Start   uint32      `json:"start"`
	List    interface{} `json:"list"`
}

func (p *paging) BuildResponse(list interface{}) *BaseResp {
	var (
		recordNum uint32
		start     uint32
		hasMore   bool
		result    *BaseResp
	)
	v := reflect.ValueOf(list)
	if v.Kind() != reflect.Slice {
		panic("build response fail")
	}

	if 0 == v.Len() {
		list = []interface{}{}
	}

	recordNum = uint32(v.Len())
	if recordNum >= p.Limit {
		start = p.Start + recordNum - 1
		hasMore = true
	} else {
		start = p.Start + recordNum
	}

	if hasMore && p.Limit > 0 {
		list = v.Slice(0, int(p.Limit-1)).Interface()
	}
	result = new(BaseResp)

	result.List = list
	result.Start = start
	result.HasMore = hasMore

	return result
}

func ParsePage(page *BasePage, opts ...Option) IPage {
	p := New(opts...)
	p.ParsePage(page)
	return p
}
