package server

import (
	"context"

	"github.com/AxulReich/kitchen/internal/pkg/app/command/update_order_status"
	"github.com/AxulReich/kitchen/internal/pkg/domain"
	pb "github.com/AxulReich/kitchen/pkg/kitchen_api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k *KitchenServer) CookingEnd(ctx context.Context, request *pb.CookingEndRequest) (*pb.CookingEndResponse, error) {
	err := k.updateOrderStatus.Handle(ctx, update_order_status.Command{
		KitchenOrderID: request.OrderId,
		Status:         domain.KitchenOrderStatusCooked,
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CookingEndResponse{}, nil
}
