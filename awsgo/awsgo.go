package awsgo

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Cfg aws.Config
var err error

func InicializoAWS() {
	fmt.Println("Entramos a AWSGO")
	Ctx = context.TODO()
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))

	fmt.Println("vamos a la mitad de AWSGO")

	if err != nil {
		panic("Error al cargar la configutations de .aws/config " + err.Error())
	}

	fmt.Println("Salimos de AWSGO")
}
