package usecase_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/tyshkor/golang_test_app/model"
	"github.com/tyshkor/golang_test_app/usecase"
)

var TestMockDBAPI1 = []model.DataModel{
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

var TestMockDBAPI2 = []model.DataModel{
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

var TestMockDBAPI3 = []model.DataModel{
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

func TestPayloadHandler(t *testing.T) {

	data := []model.DataModel{
		model.DataModel{
			ShippingSourceAddress:      "a",
			ShippingDestinationAddress: "a",
			Dimensions: model.Dimention{
				Height: 1,
				Width:  2,
				Depth:  3,
			},
		},
	}

	testMockDBList := [][]model.DataModel{
		TestMockDBAPI1,
		TestMockDBAPI2,
		TestMockDBAPI3,
	}

	got := usecase.PayloadHandler(data, testMockDBList)

	var want uint
	want = 2

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCallAPI(t *testing.T) {

	var syncMap sync.Map

	var wg sync.WaitGroup

	wg.Add(1)

	go usecase.CallAPI(&syncMap, TestMockDBAPI1, &wg)

	wg.Wait()

	item := TestMockDBAPI1[0]

	item.Cost = 0

	str := fmt.Sprint(item)

	res, _ := syncMap.Load(str)

	got, _ := res.(uint)

	var want uint

	want = 2

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
