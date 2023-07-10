package main

import (
	"context"
	"fmt"
	"time"

	//polarpb "git.leyantech.com/leyan/leyan-proto-golang/cmdb/polardb"
	rdspb "git.leyantech.com/leyan/leyan-proto-golang/cmdb/rds"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

const token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpc3MiOiJkZXBsb3lhcHAiLCJpYXQiOjE2NjQ0NjExNzIsImV4cCI6MjYxMDU0MTE3MiwiYXVkIjoiY2xvdWRmYXJtIiwib3BlcmF0b3IiOiJmaXJlcGlwZSIsInBlcm1pc3Npb25zIjpbMl19.Mu9e9swhdgZabaxW8d3aeyMom6oH1IwHOXiLRbqYlYD2UDbkShUrZleT7-KUxr5a0BurD0Z8I0fzFUNmzRqVt3BbFQvJ_AbuGKiAzk8C9Fk18dxvquJip4YUzpB9QSA00oHXfjmDX_rZ-wm8hYYSNhAvpZn1jU86ybDjynN2kj23KUG7F_siSSSjvweQbplXcLfEsI_lw940NU9HtCzZCgvv8JpVkdvvgBehtLeKaBpsiAn199tKvydc5O26rgLmk5pmEb5_hxNLqCRMK_1ujx86h2e0CP67VceRAYIKmX-LyqD_TdXTZiu0dpoLErOEza1jhtOVfuknREscZ9Du4cL0SVorGuk0ofipSwqpbvKvGuk1eKZXwwADhyUQ_Q9rZ14FSNgtf7brX76RYXMFm3Kr0b5WZQN37vcVAt6h7YEdH8SXyBebadhYs3v9JDjNHmHL8dTZJ9VS-aBKLNKOc4YkaLnMpOXpZzcAUzqxhY42fMoRP-kfxQIt0gLxj7v4"

func main() {
	conn, err := grpc.Dial(
		"localhost:3333",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx, "Authorization", "Bearer "+token)

	c := rdspb.NewRdsServiceClient(conn)
	req := rdspb.GetRdsRequest{
		InstanceId: "rm-k2j75oip05ielyr13",
	}
	resp, err := c.GetRds(ctx, &req)
	if err != nil {
		fmt.Println("GetRds error: ", err.Error())
		return
	}
	fmt.Println(resp)

	//reqp := polarpb.GetPolardbRequest{
	//	ClusterId: "pc-k2jiv0g7j0u6ywu66",
	//}
	//respp, err := cp.GetPolardb(ctx, &reqp)
	//if err != nil {
	//	fmt.Println("GetPolardb error: ", err.Error())
	//	return
	//}
	//fmt.Println(respp)
	//fmt.Println(respp.Cluster.Tags)	cp := polarpb.NewPolardbServiceClient(conn)

}
