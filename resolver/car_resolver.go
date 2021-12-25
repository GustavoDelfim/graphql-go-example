package resolver

import (
	"GustavoDelfim/graphql-go-example/middleware"
	"GustavoDelfim/graphql-go-example/schema/scalar"
	"context"
	"fmt"
	"time"
)

// This is the Car Model
type Car struct {
	Name       string `json:"name"`
	Price      int32  `json:"price"`
	Created_at string `json:"created_at"`
}

// Request input arguments
type AddCarArgs struct {
	Name  string       `json:"name"`
	Price int32        `json:"price"`
	Data  []scalar.Map `json:"data"`
}

func (*RootResolver) AddCar(ctx context.Context, args AddCarArgs) (Car, error) {
	user, _ := ctx.Value(middleware.UserCtxKey).(*middleware.User)
	fmt.Println(user)

	// Here you can make your interactions with Repositories, Databases.
	newCar := Car{
		Name:       args.Name,
		Price:      args.Price,
		Created_at: time.Now().String(),
	}

	return newCar, nil
}

func (*RootResolver) Cars(ctx context.Context) ([]Car, error) {
	user, _ := ctx.Value(middleware.UserCtxKey).(*middleware.User)
	fmt.Println(user)

	// Here you can make your interactions with Repositories, Databases.
	cars := []Car{
		{
			Name:       "Tesla Model 3",
			Price:      110,
			Created_at: time.Now().String(),
		},
		{
			Name:       "Tesla Y",
			Price:      90,
			Created_at: time.Now().String(),
		},
	}

	return cars, nil
}
