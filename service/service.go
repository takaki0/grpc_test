package service

import (
	"context"
	"errors"
	"fmt"
	"grpc_test/pb"
)

//参考：https://medium.com/eureka-engineering/golang%E3%81%A7grpc%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%9Fapiserver%E3%82%92%E4%BD%9C%E3%81%A3%E3%81%A6%E3%81%BF%E3%82%8B-9e608be200b1
type MyCatService struct {
}

func (c *MyCatService) GetMyCat(ctx context.Context, message *pb.GetMyCatMessage) (*pb.MyCatResponse, error) {
	fmt.Println("start GetMyCat")
	switch message.TargetCat {
	case "tama":
		fmt.Println("end GetMyCat")
		return &pb.MyCatResponse{
			Name: "tama",
			Kind: "mainecoon",
		}, nil
	case "mike":
		return &pb.MyCatResponse{
			Name: "mike",
			Kind: "Norwegian Forest Cat",
		}, nil
	}
	return nil, errors.New("Not Found Your Cat")
}

func (c *MyCatService) GetMyCats(ctx context.Context, message *pb.GetMyCatsMessage) (*pb.MyCatsResponse, error) {
	fmt.Println("start GetMyCats")

	var myCats []*pb.MyCatResponse
	for _, cat := range message.TargetCats {
		switch cat {
		case "tama":
			myCats = append(myCats, &pb.MyCatResponse{
				Name: "tama",
				Kind: "mainecoon",
			})
		case "mike":
			myCats = append(myCats, &pb.MyCatResponse{
				Name: "mike",
				Kind: "Norwegian Forest Cat",
			})
		default:
			myCats = append(myCats, &pb.MyCatResponse{
				Name: "other",
				Kind: "various mixed cat",
			})
		}
	}
	return &pb.MyCatsResponse{MyCat: myCats}, nil
}
