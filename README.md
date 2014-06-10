GopherAngularTutorial
=====================

A port of the [AngularJS.org tutorial](https://docs.angularjs.org/tutorial) to [GopherJS](https://github.com/gopherjs/gopherjs) using the [go-angularjs](https://github.com/gopherjs/go-angularjs) wrapper.

This is a minkimal fork of the original tutorial and I advice you to work through it first if you don't know angular yet.
My aim was to see how far the angularjs wrapper for gopherjs is. I stripped out all the testing of the tutorial until jet. I also got rid of all the nodejs tools like bower to focus on the gopherjs stuff.

## Step 0 - [Bootstrapping](https://docs.angularjs.org/tutorial/step_00)

Run `go get github.com/gopherjs/go-angularjs` to get the angularjs wrapper code. It also fetches the gopherjs package.

I included `public/index.html` which is simular to the original `app/index.html` but I link to CDN hosted versions of angularjs and bootstrap (again, to not use bower).

Run `go run server.go` and open [localhost:3000](http://localhost:3000) with your browser to get started.

It is a minimal http server that hosts the `public` directory using [negroni](https://github.com/codegangsta/negroni) (you might need to `go get` that, too). The [pure stdlib](http://golang.org/pkg/net/http/#example_FileServer) `http.FileServer` would work, too but you don't get request logging which is kind of nice to have.


## Step 1 - [Static Template](https://docs.angularjs.org/tutorial/step_01)
I skipped this, nothing javascript related yet..


### TODO
- [ ] Step 2
- [ ] The testing part of the tutorial

### Issues found
* [utf8 weirdness]()