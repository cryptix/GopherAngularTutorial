GopherAngularTutorial
=====================

A port of the [AngularJS.org tutorial](https://docs.angularjs.org/tutorial) to [GopherJS](https://github.com/gopherjs/gopherjs) using the [go-angularjs](https://github.com/gopherjs/go-angularjs) wrapper.

This is a minkimal fork of the original tutorial and I advice you to work through it first if you don't know angular yet.
My aim was to see how far the angularjs wrapper for gopherjs is. I stripped out all the testing of the tutorial until jet. I also got rid of all the nodejs tools like bower to focus on the gopherjs stuff.


## Step 0 - [Bootstrapping](https://docs.angularjs.org/tutorial/step_00)

Run `go get github.com/gopherjs/go-angularjs` to get the angularjs wrapper code. It also fetches the gopherjs package.

I included `public/index.html` which is simular to the original `app/index.html` but I link to CDN hosted versions of angularjs and bootstrap (again, to not use bower).

Run `go run server.go` and open [localhost:3000](http://localhost:3000) with your browser to get started.

It is a minimal http server that hosts the `public` directory, much like the example in the [stdlib](http://golang.org/pkg/net/http/#example_FileServer) `http.FileServer`.


## Step 1 - [Static Template](https://docs.angularjs.org/tutorial/step_01)
I skipped this, nothing javascript related yet..

## Step 2 - [Angular Templates](https://docs.angularjs.org/tutorial/step_02)
Now the fun begins. The first controller in `app/controllers.go` is identical to the one in the original tutorial and is compiled to javascript like this: `gopherjs build app/controllers.go`. Which dumps a `contollers.js` and `controllers.js.map` in the directory you ran the `build` command in. There is a nicer way to do this:


`gopherjs build -w -o public/controllers.js app/controllers.go`

The `-o` flag specifies where to write the output to and the `-w` watches for changes to the input source and rebuilds automatically.


## Step 3 - [Filtering Repeaters](https://docs.angularjs.org/tutorial/step_03)
This makes no changes to the controllers but I included the changes to the HTML for completness.


## Step 4 - [Two-way Data Binding](https://docs.angularjs.org/tutorial/step_04)
The `Age` field is added to the type `Phone` of the `PhoneListCtrl` to demo the `orderBy` feature of angular.


## Step 5 - [XHRs & Dependency Injection](https://docs.angularjs.org/tutorial/step_05)
Here we inject the HttpService of AngularJS into our controller to request JSON data of the phones, instead of hardcoding them all in our controller.

The wrapped service in go ([*ng.HttpService](http://godoc.org/github.com/gopherjs/go-angularjs#HttpService)) is pretty cool, in that it does the JSON unmarshall for you, just supply the type of data you expect as the first parameter. `[]Phone` in this case and you are done.

~~One minor thing I tripped over at first is that you need to supply two arguments to the callback otherwise you get a panic like this: `runtime error: index out of range` and a callstack in js that is not really pretty...
It comes down to [this](https://github.com/gopherjs/go-angularjs/blob/master/http.go#L120) go code. The wrapped http service uses a hardcoded amount of arguments. I opend an [pull request](https://github.com/gopherjs/go-angularjs/pull/4) with a fix where it checks `len(in)` and gives you a nicer error if it is to short.~~ My PR landed, you now get a error message, explaining that you need two arguments in the callback.


## Step 6 - [Templating Links & Images](https://docs.angularjs.org/tutorial/step_06)
This just addes some more ng-html markup to display the images of the phones as thumbnails.


## Step 7 - [Routing & Multiple Views](https://docs.angularjs.org/tutorial/step_07)
This step introduces multiple views using the `ngRoute` module.
I took the chance to break out the controllers into named functions. If they grew more they could be moved out to other go files.

~~The only thing I need to complete this Step is the `$routeParams`, i need to investigate how the ngRoute module calls the controller.~~ [Pull Request #5](https://github.com/gopherjs/go-angularjs/pull/5) adds the missing `$routeParams` Wrapper.


## Step 8 - [More Templating](https://docs.angularjs.org/tutorial/step_08)
This adds the http request to the `PhoneDetailCtrl` and extends the template with detailed data.

There is an Issue regarding setting the phone data on `$scope`. It seems that string slices are somehoe mangled, they come out in go string representation in the templates, which in this case means that angular tries to set the `src` attribute of the `img` tags to something like `http://localhost:3000/[%22img/phones/dell-streak-7.0.jpg%22,%22img/phones/dell-streak-7.1.jpg%22,%22img/phones/dell-streak-7.2.jpg%22,%22img/phones/dell-streak-7.3.jpg%22,%22img/phones/dell-streak-7.4.jpg%22,%22%22]`. I added seperate scopes for these until it is fixed.


### TODO
- [ ] Step 9
- [x] The gopherjs wrapper seems to miss the angularjs `$routeParams`
- [ ] Something is wrong with `[]string` inside structs
- [ ] `RouteProvider` doesn't know about `redirectTo`
- [ ] Open Issue regarding the UTF8 encoding, but where?
- [ ] The testing part of the tutorial
