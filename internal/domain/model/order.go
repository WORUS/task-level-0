package model

type Order struct {
	OrderUID    string `json:"order_uid" validate:"required,lowercase"`
	TrackNumber string `json:"track_number" validate:"required,uppercase"`
	Entry       string `json:"entry" validate:"required,uppercase"`
	Delivery    struct {
		Name    string `json:"name" validate:"required,min=3,max=100"`
		Phone   string `json:"phone" validate:"required,e164"`
		Zip     string `json:"zip" validate:"required,number"`
		City    string `json:"city" validate:"required"`
		Address string `json:"address" validate:"required"`
		Region  string `json:"region" validate:"required"`
		Email   string `json:"email" validate:"required,email"`
	} `json:"delivery" validate:"required"`
	Payment struct {
		Transaction  string `json:"transaction" validate:"required,lowercase"`
		RequestID    string `json:"request_id"`
		Currency     string `json:"currency" validate:"required,oneof=USD RUB EUR"`
		Provider     string `json:"provider" validate:"required"`
		Amount       int    `json:"amount" validate:"required,gte=0"`
		PaymentDt    int64  `json:"payment_dt" validate:"required"`
		Bank         string `json:"bank" validate:"required,oneof=sber alpha tinkoff"`
		DeliveryCost int    `json:"delivery_cost" validate:"required,gte=0"`
		GoodsTotal   int    `json:"goods_total" validate:"required,gte=0"`
		CustomFee    int    `json:"custom_fee" validate:"required,gte=0"`
	} `json:"payment" validate:"required"`
	Items []struct {
		ChrtID      int    `json:"chrt_id" validate:"required,number"`
		TrackNumber string `json:"track_number" validate:"required,uppercase"`
		Price       int    `json:"price" validate:"required"`
		Rid         string `json:"rid" validate:"required,lowercase"`
		Name        string `json:"name" validate:"required,min=1"`
		Sale        int    `json:"sale" validate:"required,gte=0,lte=100"`
		Size        string `json:"size" validate:"required,min=1"`
		TotalPrice  int    `json:"total_price" validate:"required,ltefield=Price"`
		NmID        int    `json:"nm_id" validate:"required,gte=0,len=7"`
		Brand       string `json:"brand" validate:"required,min=1,max=140"`
		Status      int    `json:"status" validate:"required,gte=0"`
	} `json:"items" validate:"required"`
	Locale            string `json:"locale" validate:"required,oneof=ru en kz"`
	InternalSignature string `json:"internal_signature"`
	CustomerID        string `json:"customer_id" validate:"required,min=1,max=100"`
	DeliveryService   string `json:"delivery_service" validate:"required,min=1,max=100"`
	Shardkey          string `json:"shardkey" validate:"required,number,gte=0"`
	SmID              int    `json:"sm_id" validate:"required,gte=0"`
	DateCreated       string `json:"date_created" validate:"required"`
	OofShard          string `json:"oof_shard" validate:"required,number,gte=0"`
}
