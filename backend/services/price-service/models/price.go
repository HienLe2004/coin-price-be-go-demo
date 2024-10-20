package models

type Price struct {
	Symbol string  `json:"symbol,omitempty" bson:"symbol"`
	Name   string  `json:"name,omitempty" bson:"name,omitempty"`
	Price  float32 `json:"price,omitempty" bson:"price,omitempty"`
}
