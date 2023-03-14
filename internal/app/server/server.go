package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/AxulReich/kitchen/internal/config"
	"github.com/AxulReich/kitchen/internal/pkg/app/command/update_order_status"
	"github.com/AxulReich/kitchen/internal/pkg/app/query/get_orders"
	pb "github.com/AxulReich/kitchen/pkg/kitchen_api"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type KitchenServer struct {
	pb.UnimplementedKitchenServer
	grpcServer *grpc.Server
	config     *config.Config

	updateOrderStatus update_order_status.UpdateOrderStatusHandler
	getOrders         get_orders.GetOrdersHandler
}

func NewServer(
	config *config.Config,
	updateOrderStatus update_order_status.UpdateOrderStatusHandler,
	getOrders get_orders.GetOrdersHandler,
) *KitchenServer {
	return &KitchenServer{
		config:            config,
		updateOrderStatus: updateOrderStatus,
		getOrders:         getOrders,
	}
}

func (k *KitchenServer) Run(ctx context.Context) (err error) {
	go func() {
		mux := runtime.NewServeMux()
		err = pb.RegisterKitchenHandlerServer(ctx, mux, k)
		if err != nil {
			return
		}
		err = http.ListenAndServe(fmt.Sprintf("%s:%d", k.config.AppHost, k.config.AppHTTPPort), mux)
		if err != nil {
			return
		}
	}()

	k.grpcServer = grpc.NewServer()

	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", k.config.AppHost, k.config.AppGRPCPort))
	if err != nil {
		return err
	}

	pb.RegisterKitchenServer(k.grpcServer, k)
	// TODO: add reflection
	return k.grpcServer.Serve(listen)
}

func (k *KitchenServer) Stop() {
	k.grpcServer.Stop()
}
