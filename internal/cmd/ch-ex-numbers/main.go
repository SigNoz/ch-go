package main

import (
	"context"
	"fmt"

	"github.com/go-faster/ch"
	"github.com/go-faster/ch/proto"
)

func main() {
	ctx := context.Background()
	c, err := ch.Dial(ctx, "localhost:9000", ch.Options{})
	if err != nil {
		panic(err)
	}
	var (
		numbers int
		data    proto.ColUInt64
	)
	if err := c.Query(ctx, ch.Query{
		Body: "SELECT number FROM system.numbers LIMIT 500000000",
		OnData: func(ctx context.Context) error {
			numbers += len(data)
			return nil
		},
		Result: []proto.ResultColumn{
			{Name: "number", Data: &data},
		},
	}); err != nil {
		panic(err)
	}
	fmt.Println("numbers:", numbers)
}