package main

import (
	ng "github.com/gopherjs/go-angularjs"
	"net/http"
)

type Phone struct {
	Id, ImageUrl  string
	Name, Snippet string
	Age           int
}

func main() {
	app := ng.NewModule("gopherJsApp", []string{})

	app.NewController("PhoneListCtrl", func(scope *ng.Scope, httpService *ng.HttpService) {

		httpService.Get("phones/phones.json").Success(func(data []Phone, status int) {
			if status != http.StatusOK {
				println("request status:", status)
				return
			}

			scope.Set("phones", data)
		})

		scope.Set("orderProp", "Age")

	})

}
