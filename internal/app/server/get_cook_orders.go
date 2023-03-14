package server

import (
	"context"

	"github.com/AxulReich/kitchen/internal/pkg/app/query/get_orders"
	pb "github.com/AxulReich/kitchen/pkg/kitchen_api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k *KitchenServer) GetCookOrders(ctx context.Context, request *pb.GetCookOrdersRequest) (*pb.GetCookOrdersResponse, error) {
	result, err := k.getOrders.Handle(ctx, get_orders.Command{
		Offset: request.Offset,
		Limit:  request.Limit,
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	ordersPb := make([]*pb.Order, 0, len(result.Orders))

	for _, order := range result.Orders {
		ordersPb = append(ordersPb, &pb.Order{
			Id:     order.ShopOrderID,
			Status: convertStatus(order.Status),
			Items:  convertItems(order.Items),
		})
	}

	return &pb.GetCookOrdersResponse{
		Orders: ordersPb,
		Total:  result.Total,
	}, nil
}
