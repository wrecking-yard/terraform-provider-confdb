package main

import (
	"context"
	"log"

	"codeberg.org/wrecking-yard/terraform-provider-confdb/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	opts := providerserver.ServeOpts{
		Address: "localhost/dev/confdb",
	}

	version := "dev"
	err := providerserver.Serve(context.Background(), provider.New(version), opts)
	if err != nil {
		log.Fatal(err.Error())
	}
}
