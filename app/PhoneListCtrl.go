package main

import (
	ng "github.com/gopherjs/go-angularjs"
)

type Phone struct {
	Id, ImageUrl  string
	Name, Snippet string
	Age           int
}

func (p *Phone) GetId() string {
	return p.Id
}

func (p *Phone) GetSnippet() string {
	return p.Snippet
}

func (p *Phone) GetImageUrl() string {
	return p.ImageUrl
}

func (p *Phone) GetName() string {
	return p.Name
}

func PhoneListCtrl(scope *ng.Scope, httpService *ng.HttpService) {

	httpService.Get("phones/phones.json").Success(func(data []*Phone, status int) {
		if status != 200 {
			println("request status:", status)
			return
		}

		scope.Set("phones", data)
	})

	scope.Set("orderProp", "Age")
}
