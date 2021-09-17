package models

import "time"

type Supplier struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	Deleted bool      `json:"deleted"`
	ImgURL  string    `json:"img_url"`
	Type    string    `json:"type"`
	Opening string    `json:"opening"`
	Closing string    `json:"closing"`
}

type ResponseBodyRestaurants struct {
	Suppliers []ParsedSupplier `json:"suppliers"`
}

type WorkingHours struct {
	Opening string `json:"opening"`
	Closing string `json:"closing"`
}

type ParsedSupplier struct {
	ID           int          `json:"id"`
	Image        string       `json:"image"`
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	WorkingHours WorkingHours `json:"workingHours"`
}

type SupplierResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	ImgURL  string `json:"img_url"`
	Type    string `json:"type"`
	Opening string `json:"opening"`
	Closing string `json:"closing"`
}

type SupplierRequestID struct {
	ID int `json:"id"`
}

type SupplierRequestType struct {
	Type string `json:"type"`
}

type SupplierRequestTime struct {
	Time string `json:"time"`
}

func TransformSupplierForResponse(supplier *Supplier) *SupplierResponse {
	SupplierResponse := SupplierResponse{
		ID:      supplier.ID,
		Name:    supplier.Name,
		ImgURL:  supplier.ImgURL,
		Type:    supplier.Type,
		Opening: supplier.Opening,
		Closing: supplier.Closing,
	}
	return &SupplierResponse
}
