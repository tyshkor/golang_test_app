package model

type API2Model struct {
	Consignee string      `json:"consignee" xml:"consignee"`
	Consignor string      `json:"consignor" xml:"consignor"`
	Cartons   []Dimention `json:"cartons" xml:"cartons"`
}
