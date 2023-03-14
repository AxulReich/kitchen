package domain

type KitchenOrderStatus string

const (
	KitchenOrderStatusInvalid KitchenOrderStatus = "invalid"
	KitchenOrderStatusNew     KitchenOrderStatus = "new"
	KitchenOrderStatusCooking KitchenOrderStatus = "cooking_start"
	KitchenOrderStatusCooked  KitchenOrderStatus = "kooking_end"
)

type KitchenOrder struct {
	ID          int64
	ShopOrderID int64
	Status      KitchenOrderStatus
	Items       []Item
}
