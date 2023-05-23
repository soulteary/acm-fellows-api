package humanname

import (
	"context"
	"time"

	pb "github.com/soulteary/acm-fellows-api/model/humanname/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	TestName = "Dr. Juan Q. Xavier de la Vega III (Doc Vega)"
	GrpcAddr = "localhost:8081"
)

func ParseNameByRpc(name string) (string, error) {
	conn, err := grpc.Dial(GrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", err
	}
	defer conn.Close()
	c := pb.NewConverterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.HumanName(ctx, &pb.ConvertRequest{Name: name})
	if err != nil {
		return "", err
	}
	return r.GetMessage(), nil
}
