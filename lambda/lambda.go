package lambda

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
)

type LambdaFunc func(ctx context.Context, name interface{}) (interface{}, error)

func Execute(l *LambdaFunc) {
	lambda.Start(l)
}