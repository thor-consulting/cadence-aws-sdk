// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package forecastqueryservicestub

import (
	"github.com/aws/aws-sdk-go/service/forecastqueryservice"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type QueryForecastFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *QueryForecastFuture) Get(ctx workflow.Context) (*forecastqueryservice.QueryForecastOutput, error) {
	var output forecastqueryservice.QueryForecastOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) QueryForecast(ctx workflow.Context, input *forecastqueryservice.QueryForecastInput) (*forecastqueryservice.QueryForecastOutput, error) {
	var output forecastqueryservice.QueryForecastOutput
	err := workflow.ExecuteActivity(ctx, "aws-forecastqueryservice-QueryForecast", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) QueryForecastAsync(ctx workflow.Context, input *forecastqueryservice.QueryForecastInput) *QueryForecastFuture {
	future := workflow.ExecuteActivity(ctx, "aws-forecastqueryservice-QueryForecast", input)
	return &QueryForecastFuture{Future: future}
}
