package main

import (
	ng "github.com/gopherjs/go-angularjs"
)

type Phone struct {
	Id, ImageUrl  string
	Name, Snippet string
	Age           int
}

func main() {
	app := ng.NewModule("gopherJsApp", []string{"ngRoute"})

	app.Config(func(rp *ng.RouteProvider) {
		listRoute := ng.RouteOptions(
			ng.RouteController{"PhoneListCtrl"},
			ng.RouteTemplate{"partials/phone-list.html"},
		)

		detailRoute := ng.RouteOptions(
			ng.RouteController{"PhoneDetailCtrl"},
			ng.RouteTemplate{"partials/phone-detail.html"},
		)

		rp.
			When("/phones", listRoute).
			When("/phones/:phoneId", detailRoute).
			Otherwise(listRoute)
	})

	app.NewController("PhoneListCtrl", PhoneListCtrl)
	app.NewController("PhoneDetailCtrl", PhoneDetailCtrl)
}

func PhoneListCtrl(scope *ng.Scope, httpService *ng.HttpService) {

	httpService.Get("phones/phones.json").Success(func(data []Phone, status int) {
		if status != 200 {
			println("request status:", status)
			return
		}

		scope.Set("phones", data)
	})

	scope.Set("orderProp", "Age")
}

// TODO find out where to get routeParams from
func PhoneDetailCtrl(scope *ng.Scope) {
	scope.Set("phoneId", "42")
}
