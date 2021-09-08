package mockapi

import (
	"encoding/json"
	"log"

	"github.com/tyshkor/golang_test_app/model"
)

var MockDBAPI1 = []model.DataModel{
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

var MockDBAPI2 = []model.DataModel{
	model.DataModel{
		ShippingSourceAddress:      "a",
		ShippingDestinationAddress: "a",
		Dimensions: model.Dimention{
			Height: 1,
			Width:  2,
			Depth:  3,
		},
		Cost: 1,
	},
}

var MockDBAPI3 = []model.DataModel{
	model.DataModel{
		ShippingSourceAddress:      "a",
		ShippingDestinationAddress: "a",
		Dimensions: model.Dimention{
			Height: 1,
			Width:  2,
			Depth:  3,
		},
		Cost: 3,
	},
}

func GetMockAPI(mockDB []model.DataModel) []byte {

	byteList, err := json.Marshal(mockDB)
	if err != nil {
		log.Printf("GetMockAPI, Err: %v", err)
		return nil
	}

	return byteList
}
