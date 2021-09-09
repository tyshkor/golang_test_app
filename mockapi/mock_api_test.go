package mockapi_test

import "github.com/tyshkor/golang_test_app/model"

var TestMockDBAPI = []model.DataModel{
	model.DataModel{
		ShippingSourceAddress:      "a",
		ShippingDestinationAddress: "a",
		Dimensions: model.Dimention{
			Height: 1,
			Width:  2,
			Depth:  3,
		},
		Cost: 2,
	},
}
