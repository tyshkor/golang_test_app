package usecase

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/tyshkor/golang_test_app/mockapi"
	"github.com/tyshkor/golang_test_app/model"
)

func PayloadHandler(data []model.DataModel, mockDBList [][]model.DataModel) uint {

	var syncMap sync.Map

	var wg sync.WaitGroup

	wg.Add(len(mockDBList))

	for _, item := range mockDBList {

		go CallAPI(&syncMap, item, &wg)
	}

	wg.Wait()

	var res uint

	for _, item := range data {

		str := fmt.Sprint(item)

		itemFromMap, ok := syncMap.Load(str)

		if ok {

			itemParsed := itemFromMap.(uint)

			res += itemParsed
		}
	}

	return res
}

func CallAPI(syncMap *sync.Map, item []model.DataModel, wg *sync.WaitGroup) {

	defer wg.Done()

	byteList := mockapi.GetMockAPI(item)

	var list []model.DataModel

	err := json.Unmarshal(byteList, &list)

	if err != nil {

		// we don't want to interrupt the flow - we can ignore API failure
		log.Printf("json.Unmarshal, Err: %v", err)
	}

	for _, listItem := range list {

		cost := listItem.Cost

		listItem.Cost = 0

		str := fmt.Sprint(listItem)

		val, ok := syncMap.Load(str)

		if ok {

			valParsed, okVal := val.(uint)

			if okVal && valParsed > cost {

				syncMap.Store(str, cost)
			}

		} else {

			syncMap.Store(str, cost)
		}

	}
}
