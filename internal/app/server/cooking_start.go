package server

import (
	"context"

	"github.com/AxulReich/kitchen/internal/pkg/app/command/update_order_status"
	"github.com/AxulReich/kitchen/internal/pkg/domain"
	pb "github.com/AxulReich/kitchen/pkg/kitchen_api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k *KitchenServer) CookingStart(ctx context.Context, request *pb.CookingStartRequest) (*pb.CookingStartResponse, error) {
	err := k.updateOrderStatus.Handle(ctx, update_order_status.Command{
		ShopOrderID: request.OrderId,
		Status:      domain.Values.CookingStart,
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CookingStartResponse{}, nil
}
