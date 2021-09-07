package model

type API1Model struct {
	ContactAddress    string      `json:"contact_address"`
	WarehouseAddress  string      `json:"warehouse_address"`
	PackageDimensions []Dimention `json:"package_dimensions"`
}
