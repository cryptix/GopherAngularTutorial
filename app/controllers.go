package main

import (
	ng "github.com/gopherjs/go-angularjs"
)

type Phone struct {
	Name, Snippet string
	Age           int
}

func main() {
	app := ng.NewModule("gopherJsApp", []string{})

	app.NewController("PhoneListCtrl", func(scope *ng.Scope) {

		scope.Set("phones", []Phone{
			Phone{"Nexus S", "Fast just got faster with Nexus S.", 1},
			Phone{"Motorola XOOM™ with Wi-Fi", "The Next, Next Generation tablet.", 2},
			Phone{"MOTOROLA XOOM™", "The Next, Next Generation tablet.", 3},
		})

		scope.Set("orderProp", "Age")

	})

}
