package main

import (
	"context"
	"fmt"

	"github.com/swethabhageerath/storage/pkg/storage/aws/secrets"
)

func main() {
	s := secrets.Secrets{}
	connectionString, err := s.GetValueString(context.Background(), "PgConnection")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(connectionString)
	}
}
