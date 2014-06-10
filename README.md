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

One minor thing I tripped over at first is that you need to supply two arguments to the callback otherwise you get a panic like this: `runtime error: index out of range` and a callstack in js that is not really pretty...
It comes down to [this](https://github.com/gopherjs/go-angularjs/blob/master/http.go#L120) go code. The wrapped http service uses a hardcoded amount of arguments where it could check the `len(in)` but that can easly be fixed.


### TODO
- [ ] Step 6
- [ ] Open Issue regarding the UTF8 encoding, but where?
- [ ] The testing part of the tutorial

### Issues found
* utf8 weirdness
