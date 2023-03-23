package domain

type ShopOrderStatus string

const (
	ShopOrderStatusInvalid   ShopOrderStatus = "invalid"
	ShopOrderStatusConfirmed ShopOrderStatus = "confirmed"
)

type ShopOrder struct {
	ID     int64
	Status ShopOrderStatus
	Items  []ShopOrderItem
}

type ShopOrderItem struct {
	ShopOrderID int64
	Name        string
	Comment     string
}
