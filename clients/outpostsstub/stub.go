// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package outpostsstub

import (
	"github.com/aws/aws-sdk-go/service/outposts"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type OutpostsCreateOutpostFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *OutpostsCreateOutpostFuture) Get(ctx workflow.Context) (*outposts.CreateOutpostOutput, error) {
	var output outposts.CreateOutpostOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type OutpostsDeleteOutpostFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *OutpostsDeleteOutpostFuture) Get(ctx workflow.Context) (*outposts.DeleteOutpostOutput, error) {
	var output outposts.DeleteOutpostOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type OutpostsDeleteSiteFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *OutpostsDeleteSiteFuture) Get(ctx workflow.Context) (*outposts.DeleteSiteOutput, error) {
	var output outposts.DeleteSiteOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type OutpostsGetOutpostFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *OutpostsGetOutpostFuture) Get(ctx workflow.Context) (*outposts.GetOutpostOutput, error) {
	var output outposts.GetOutpostOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type OutpostsGetOutpostInstanceTypesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *OutpostsGetOutpostInstanceTypesFuture) Get(ctx workflow.Context) (*outposts.GetOutpostInstanceTypesOutput, error) {
	var output outposts.GetOutpostInstanceTypesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type OutpostsListOutpostsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *OutpostsListOutpostsFuture) Get(ctx workflow.Context) (*outposts.ListOutpostsOutput, error) {
	var output outposts.ListOutpostsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type OutpostsListSitesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *OutpostsListSitesFuture) Get(ctx workflow.Context) (*outposts.ListSitesOutput, error) {
	var output outposts.ListSitesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateOutpost(ctx workflow.Context, input *outposts.CreateOutpostInput) (*outposts.CreateOutpostOutput, error) {
	var output outposts.CreateOutpostOutput
	err := workflow.ExecuteActivity(ctx, "aws-outposts-CreateOutpost", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateOutpostAsync(ctx workflow.Context, input *outposts.CreateOutpostInput) *OutpostsCreateOutpostFuture {
	future := workflow.ExecuteActivity(ctx, "aws-outposts-CreateOutpost", input)
	return &OutpostsCreateOutpostFuture{Future: future}
}

func (a *stub) DeleteOutpost(ctx workflow.Context, input *outposts.DeleteOutpostInput) (*outposts.DeleteOutpostOutput, error) {
	var output outposts.DeleteOutpostOutput
	err := workflow.ExecuteActivity(ctx, "aws-outposts-DeleteOutpost", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteOutpostAsync(ctx workflow.Context, input *outposts.DeleteOutpostInput) *OutpostsDeleteOutpostFuture {
	future := workflow.ExecuteActivity(ctx, "aws-outposts-DeleteOutpost", input)
	return &OutpostsDeleteOutpostFuture{Future: future}
}

func (a *stub) DeleteSite(ctx workflow.Context, input *outposts.DeleteSiteInput) (*outposts.DeleteSiteOutput, error) {
	var output outposts.DeleteSiteOutput
	err := workflow.ExecuteActivity(ctx, "aws-outposts-DeleteSite", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteSiteAsync(ctx workflow.Context, input *outposts.DeleteSiteInput) *OutpostsDeleteSiteFuture {
	future := workflow.ExecuteActivity(ctx, "aws-outposts-DeleteSite", input)
	return &OutpostsDeleteSiteFuture{Future: future}
}

func (a *stub) GetOutpost(ctx workflow.Context, input *outposts.GetOutpostInput) (*outposts.GetOutpostOutput, error) {
	var output outposts.GetOutpostOutput
	err := workflow.ExecuteActivity(ctx, "aws-outposts-GetOutpost", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetOutpostAsync(ctx workflow.Context, input *outposts.GetOutpostInput) *OutpostsGetOutpostFuture {
	future := workflow.ExecuteActivity(ctx, "aws-outposts-GetOutpost", input)
	return &OutpostsGetOutpostFuture{Future: future}
}

func (a *stub) GetOutpostInstanceTypes(ctx workflow.Context, input *outposts.GetOutpostInstanceTypesInput) (*outposts.GetOutpostInstanceTypesOutput, error) {
	var output outposts.GetOutpostInstanceTypesOutput
	err := workflow.ExecuteActivity(ctx, "aws-outposts-GetOutpostInstanceTypes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetOutpostInstanceTypesAsync(ctx workflow.Context, input *outposts.GetOutpostInstanceTypesInput) *OutpostsGetOutpostInstanceTypesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-outposts-GetOutpostInstanceTypes", input)
	return &OutpostsGetOutpostInstanceTypesFuture{Future: future}
}

func (a *stub) ListOutposts(ctx workflow.Context, input *outposts.ListOutpostsInput) (*outposts.ListOutpostsOutput, error) {
	var output outposts.ListOutpostsOutput
	err := workflow.ExecuteActivity(ctx, "aws-outposts-ListOutposts", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListOutpostsAsync(ctx workflow.Context, input *outposts.ListOutpostsInput) *OutpostsListOutpostsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-outposts-ListOutposts", input)
	return &OutpostsListOutpostsFuture{Future: future}
}

func (a *stub) ListSites(ctx workflow.Context, input *outposts.ListSitesInput) (*outposts.ListSitesOutput, error) {
	var output outposts.ListSitesOutput
	err := workflow.ExecuteActivity(ctx, "aws-outposts-ListSites", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListSitesAsync(ctx workflow.Context, input *outposts.ListSitesInput) *OutpostsListSitesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-outposts-ListSites", input)
	return &OutpostsListSitesFuture{Future: future}
}