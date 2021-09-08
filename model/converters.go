package model

func ConvertAPI1ToStandard(payload API1Model) []DataModel {

	data := make([]DataModel, 0)

	for _, item := range payload.PackageDimensions {

		data = append(data, DataModel{
			ShippingSourceAddress:      payload.WarehouseAddress,
			ShippingDestinationAddress: payload.ContactAddress,
			Dimensions:                 item,
		})
	}

	return data
}

func ConvertAPI2ToStandard(payload API2Model) []DataModel {

	data := make([]DataModel, 0)

	for _, item := range payload.Cartons {

		data = append(data, DataModel{
			ShippingSourceAddress:      payload.Consignee,
			ShippingDestinationAddress: payload.Consignor,
			Dimensions:                 item,
		})
	}

	return data
}
