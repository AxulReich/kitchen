package server

import (
	"github.com/AxulReich/kitchen/internal/pkg/domain"
	pb "github.com/AxulReich/kitchen/pkg/kitchen_api"
)

func convertItems(items []domain.Item) []*pb.OrderItem {
	result := make([]*pb.OrderItem, 0, len(items))
	for _, item := range items {
		result = append(result, &pb.OrderItem{
			Name:    item.Name,
			Comment: item.Comment,
		})
	}
	return result
}

func convertStatus(status domain.KitchenOrderStatus) pb.KitchenOrderStatus {
	switch status {
	case domain.KitchenOrderStatusNew:
		return pb.KitchenOrderStatus_kitchenOrderStatusNew
	case domain.KitchenOrderStatusCooking:
		return pb.KitchenOrderStatus_kitchenOrderStatusCooking
	case domain.KitchenOrderStatusCooked:
		return pb.KitchenOrderStatus_kitchenOrderStatusCooked
	}
	return pb.KitchenOrderStatus_KitchenOrderStatusInvalid
}
