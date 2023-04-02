package domain

type KitchenOrder struct {
	ID          int64                  `json:"kitchen_order_id,omitempty"`
	ShopOrderID int64                  `json:"id"`
	Status      KitchenOrderStatusEnum `json:"status"`
	Items       []Item                 `json:"items,omitempty"`
}

type KitchenOrderStatusEnum interface {
	isKitchenOrderStatusEnumValue()
}

type invalid struct{}

func (i *invalid) String() string {
	return "invalid"
}

func (_ *invalid) isKitchenOrderStatusEnumValue() {}

type novel struct{}

func (n *novel) String() string {
	return "new"
}

func (_ *novel) isKitchenOrderStatusEnumValue() {}

type cookingStart struct{}

func (s *cookingStart) String() string {
	return "cooking_start"
}

func (_ *cookingStart) isKitchenOrderStatusEnumValue() {}

type cookingEnd struct{}

func (c *cookingEnd) String() string {
	return "cooking_end"
}

func (_ *cookingEnd) isKitchenOrderStatusEnumValue() {}

var (
	// Values ...
	Values = struct {
		Invalid      *invalid
		New          *novel
		CookingStart *cookingStart
		CookingEnd   *cookingEnd
	}{
		Invalid:      &invalid{},
		New:          &novel{},
		CookingStart: &cookingStart{},
		CookingEnd:   &cookingEnd{},
	}
)
