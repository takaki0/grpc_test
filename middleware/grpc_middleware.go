package middleware

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"path"
	"time"
)

//参考：https://qiita.com/lightstaff/items/28724d9dd8a6b30b236d
func UnaryServerInterCeptor(logger log.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		fmt.Println("start middleware")
		var err error
		defer func(begin time.Time){
			//メソッド名取得
			method := path.Base(info.FullMethod)
			took := time.Since(begin)
			if err != nil {
				logger.Printf("method: %v, took: %v", method, took)
			}
			fmt.Printf("method: %v, took: %v", method, took)
			fmt.Println("end middleware")
		}(time.Now())

		res, hErr := handler(ctx, req)
		if hErr != nil {
			err = hErr
		}
		return res, err
	}
}