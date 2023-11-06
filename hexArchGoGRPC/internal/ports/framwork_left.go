package ports

import (
	"context"

	"github.com/alireza-frj4/app1/internal/adapters/framwork/left/grpc/pb"
)

type GEPCPort interface {
	Run()
	GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
	GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
	GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
}
