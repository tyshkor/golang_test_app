package model

type DataModel struct {
	ShippingSourceAddress      string
	ShippingDestinationAddress string
	Dimensions                 Dimention
	Cost                       uint
}
