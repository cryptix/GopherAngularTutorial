package main

import (
	ng "github.com/gopherjs/go-angularjs"
)

// generated using https://mholt.github.io/json-to-go/
type PhoneDatails struct {
	AdditionalFeatures string `json:"additionalFeatures"`
	Android            struct {
		Os string `json:"os"`
		Ui string `json:"ui"`
	} `json:"android"`
	Availability []string `json:"availability"`
	Battery      struct {
		StandbyTime string `json:"standbyTime"`
		TalkTime    string `json:"talkTime"`
		Type        string `json:"type"`
	} `json:"battery"`
	Camera struct {
		Features []string `json:"features"`
		Primary  string   `json:"primary"`
	} `json:"camera"`
	Connectivity struct {
		Bluetooth string `json:"bluetooth"`
		Cell      string `json:"cell"`
		Gps       bool   `json:"gps"`
		Infrared  bool   `json:"infrared"`
		Wifi      string `json:"wifi"`
	} `json:"connectivity"`
	Description string `json:"description"`
	Display     struct {
		ScreenResolution string `json:"screenResolution"`
		ScreenSize       string `json:"screenSize"`
		TouchScreen      bool   `json:"touchScreen"`
	} `json:"display"`
	Hardware struct {
		Accelerometer    bool   `json:"accelerometer"`
		AudioJack        string `json:"audioJack"`
		Cpu              string `json:"cpu"`
		FmRadio          bool   `json:"fmRadio"`
		PhysicalKeyboard bool   `json:"physicalKeyboard"`
		Usb              string `json:"usb"`
	} `json:"hardware"`
	Id            string   `json:"id"`
	Images        []string `json:"images"`
	Name          string   `json:"name"`
	SizeAndWeight struct {
		Dimensions []string `json:"dimensions"`
		Weight     string   `json:"weight"`
	} `json:"sizeAndWeight"`
	Storage struct {
		Flash string `json:"flash"`
		Ram   string `json:"ram"`
	} `json:"storage"`
}

func PhoneDetailCtrl(scope *ng.Scope, httpService *ng.HttpService, params *ng.RouteParams) {
	httpService.Get("phones/" + params.Get("phoneId") + ".json").Success(func(data PhoneDatails, status int) {
		if status != 200 {
			println("request status:", status)
			return
		}
		// println(data)

		scope.Set("phone", data)
		scope.Set("mainImageUrl", data.Images[0])

		// hack until $externalize works properly
		scope.Set("phoneImages", data.Images)
		scope.Set("phoneAvailability", data.Availability)
		scope.Set("phoneDimensions", data.SizeAndWeight.Dimensions)
		scope.Set("phoneCamFeatures", data.Camera.Features)
	})

	scope.Set("setImage", func(url string) {
		scope.Set("mainImageUrl", url)
	})
}
