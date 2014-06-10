package main

import (
	ng "github.com/gopherjs/go-angularjs"
)

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
			Otherwise(listRoute) // should be redirectTo
	})

	app.NewController("PhoneListCtrl", PhoneListCtrl)
	app.NewController("PhoneDetailCtrl", PhoneDetailCtrl)
}
