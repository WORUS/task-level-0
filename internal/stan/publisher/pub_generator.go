package publisher

import (
	"fmt"
	"task-level-0/internal/domain/model"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/sirupsen/logrus"
)

const (
	trackNumber = "WBILMTESTTRACK"
	entry       = "WBIL"
	provider    = "wbpay"
)

//TODO: make constans and items

func generateJSON(order *model.Order) *model.Order {
	currencies := [3]string{"USD", "RUB", "EUR"}
	banks := [3]string{"alpha", "sber", "tinkoff"}
	items := [5]int{1, 1, 1, 2, 3}
	itemsNumber := gofakeit.RandomInt(items[:])
	logrus.Info(itemsNumber)

	order.OrderUID = gofakeit.UUID()
	order.TrackNumber = trackNumber
	order.Entry = entry

	order.Delivery.Name = gofakeit.Name()
	order.Delivery.Phone = gofakeit.Phone()
	order.Delivery.Zip = fmt.Sprint(gofakeit.Number(1000000, 9999999))
	order.Delivery.City = gofakeit.City()
	order.Delivery.Address = gofakeit.Street()
	order.Delivery.Region = gofakeit.StreetName()
	order.Delivery.Email = gofakeit.Email()

	order.Payment.Transaction = order.OrderUID
	order.Payment.RequestID = ""
	order.Payment.Currency = gofakeit.RandomString(currencies[:])
	order.Payment.Provider = provider
	order.Payment.Amount = gofakeit.IntRange(1600, 9999)
	order.Payment.PaymentDt = gofakeit.Date().Unix()
	order.Payment.Bank = gofakeit.RandomString(banks[:])
	order.Payment.DeliveryCost = int(gofakeit.Uint8() * 3)
	order.Payment.GoodsTotal = order.Payment.Amount - order.Payment.DeliveryCost

	order.Payment.CustomFee = 0
	for i := 0; i < itemsNumber; i++ {
		addItem(order)
		order.Items[i].ChrtID = gofakeit.Number(1000000, 9999999)
		order.Items[i].TrackNumber = order.TrackNumber
		order.Items[i].Price = int(gofakeit.Uint16())
		order.Items[i].Rid = gofakeit.UUID()
		order.Items[i].Name = gofakeit.ProductName()
		order.Items[i].Sale = gofakeit.IntRange(0, 100)
		order.Items[i].Size = "0"
		order.Items[i].TotalPrice = order.Items[i].Price * (order.Items[i].Sale / 100)
		order.Items[i].NmID = int(gofakeit.Uint32())
		order.Items[i].Brand = gofakeit.AppName()
		order.Items[i].Status = 202
	}

	order.Locale = gofakeit.RandomString([]string{"ru", "en"})
	order.InternalSignature = ""
	order.CustomerID = gofakeit.Username()
	order.DeliveryService = gofakeit.ProductMaterial()
	order.Shardkey = fmt.Sprint(gofakeit.Uint8())
	order.SmID = int(gofakeit.Uint8())
	order.DateCreated = time.Now().UTC().Format(time.RFC3339)
	order.OofShard = "1"

	return order
}

func addItem(order *model.Order) {
	order.Items = append(order.Items, struct {
		ChrtID      int    "json:\"chrt_id\""
		TrackNumber string "json:\"track_number\""
		Price       int    "json:\"price\""
		Rid         string "json:\"rid\""
		Name        string "json:\"name\""
		Sale        int    "json:\"sale\""
		Size        string "json:\"size\""
		TotalPrice  int    "json:\"total_price\""
		NmID        int    "json:\"nm_id\""
		Brand       string "json:\"brand\""
		Status      int    "json:\"status\""
	}{})
}
