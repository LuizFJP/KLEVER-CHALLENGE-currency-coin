package main

import (
	"context"
	"fmt"

	pb "github.com/LuizFJP/currency-coin-grpc-go/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type test struct {
	Id string `bson:"_id"`
	Name string `bson:"name"`
	Price float64 `bson:"price"`
	Vote int64 `bson:"vote"`
}

func UpdateByName(name string) (error) {
	filter := bson.D{primitive.E{Key: "Name", Value: name}}
	update := bson.D{primitive.E{Key: "$inc", Value: bson.D{primitive.E{Key: "Vote", Value: 1}}}}
	
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error: %v", err),
		)
	}
	return nil

}

func (s *Server) UpvoteCoin(ctx context.Context, in *pb.CoinRequest) (*pb.CoinResponse, error) {
	result := &test{}
	data := test {
		Name: in.Name,
	}

	err := UpdateByName(data.Name)

	if err != nil {
		return nil, err
	}
	
	err = collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "Name", Value: data.Name}}).Decode(&result)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error: %v", err),
		)
	}

	return &pb.CoinResponse{
		Name: result.Name,
		Price: result.Price,
		Vote: result.Vote,
		
	}, nil
}
