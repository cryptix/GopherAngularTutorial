package main

import (
	ng "github.com/gopherjs/go-angularjs"
)

type Phone struct {
	Name, Snippet string
}

func main() {
	app := ng.NewModule("gopherJsApp", []string{})

	app.NewController("PhoneListCtrl", func(scope *ng.Scope) {

		scope.Set("phones", []Phone{
			Phone{"Nexus S", "Fast just got faster with Nexus S."},
			Phone{"Motorola XOOM™ with Wi-Fi", "The Next, Next Generation tablet."},
			Phone{"MOTOROLA XOOM™", "The Next, Next Generation tablet."},
		})

	})

}
