package main

//custom data models
type customEvent struct {
	Type       string `json:"type"`
	Decription string `json:"description"`
}

type ProductType struct {
	productTypeID string `json:"productTypeID"`
	productType   string `json:"productType"`
	category      string `json:"category"`
	description   string `json:"description"`
	feature1      string `json:"feature1"`
	UOM1          string `json:"UOM1"`
	minValue1     string `json:"minValue1"`
	maxValue1     string `json:"maxValue1"`
	feature2      string `json:"feature2"`
	UOM2          string `json:"UOM2"`
	minValue2     string `json:"minValue2"`
	maxValue2     string `json:"maxValue2"`
	feature3      string `json:"feature3"`
	UOM3          string `json:"UOM3"`
	minValue3     string `json:"minValue3"`
	maxValue3     string `json:"maxValue3"`
}

type PurchaseOrder struct {
	poID                 int    `json:"poID"`
	buyer                Org    `json:"buyer"`
	seller               Org    `json:"seller"`
	poDate               string `json:"poDate"`
	expectedDeliveryDate string `json:"expectedDeliveryDate"`
	actualDeliveryDate   string `json:"actualDeliveryDate"`
	shipToAddress        string `json:"shipToAddress"`
	Status               string `json:"Status"`
}

type Org struct {
	orgID       string `json:"orgID"`
	orgType     string `json:"orgType"`
	name        string `json:"name"`
	description string `json:"description"`
	address     string `json:"address"`
	rating      string `json:"rating"`
}

type LineItem struct {
	itemID             string        `json:"itemID"`
	productType        ProductType   `json:"productType"`
	po                 PurchaseOrder `json:"po"`
	qty                float32       `json:"qty"`
	price              float32       `json:"price"`
	inspectionRequired bool          `json:"inspectionRequired"`
}
