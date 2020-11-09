// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package cloudwatchstub

import (
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type DeleteAlarmsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteAlarmsFuture) Get(ctx workflow.Context) (*cloudwatch.DeleteAlarmsOutput, error) {
	var output cloudwatch.DeleteAlarmsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteAnomalyDetectorFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteAnomalyDetectorFuture) Get(ctx workflow.Context) (*cloudwatch.DeleteAnomalyDetectorOutput, error) {
	var output cloudwatch.DeleteAnomalyDetectorOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteDashboardsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteDashboardsFuture) Get(ctx workflow.Context) (*cloudwatch.DeleteDashboardsOutput, error) {
	var output cloudwatch.DeleteDashboardsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteInsightRulesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteInsightRulesFuture) Get(ctx workflow.Context) (*cloudwatch.DeleteInsightRulesOutput, error) {
	var output cloudwatch.DeleteInsightRulesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeAlarmHistoryFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeAlarmHistoryFuture) Get(ctx workflow.Context) (*cloudwatch.DescribeAlarmHistoryOutput, error) {
	var output cloudwatch.DescribeAlarmHistoryOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeAlarmsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeAlarmsFuture) Get(ctx workflow.Context) (*cloudwatch.DescribeAlarmsOutput, error) {
	var output cloudwatch.DescribeAlarmsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeAlarmsForMetricFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeAlarmsForMetricFuture) Get(ctx workflow.Context) (*cloudwatch.DescribeAlarmsForMetricOutput, error) {
	var output cloudwatch.DescribeAlarmsForMetricOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeAnomalyDetectorsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeAnomalyDetectorsFuture) Get(ctx workflow.Context) (*cloudwatch.DescribeAnomalyDetectorsOutput, error) {
	var output cloudwatch.DescribeAnomalyDetectorsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeInsightRulesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeInsightRulesFuture) Get(ctx workflow.Context) (*cloudwatch.DescribeInsightRulesOutput, error) {
	var output cloudwatch.DescribeInsightRulesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DisableAlarmActionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DisableAlarmActionsFuture) Get(ctx workflow.Context) (*cloudwatch.DisableAlarmActionsOutput, error) {
	var output cloudwatch.DisableAlarmActionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DisableInsightRulesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DisableInsightRulesFuture) Get(ctx workflow.Context) (*cloudwatch.DisableInsightRulesOutput, error) {
	var output cloudwatch.DisableInsightRulesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EnableAlarmActionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EnableAlarmActionsFuture) Get(ctx workflow.Context) (*cloudwatch.EnableAlarmActionsOutput, error) {
	var output cloudwatch.EnableAlarmActionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EnableInsightRulesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EnableInsightRulesFuture) Get(ctx workflow.Context) (*cloudwatch.EnableInsightRulesOutput, error) {
	var output cloudwatch.EnableInsightRulesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetDashboardFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetDashboardFuture) Get(ctx workflow.Context) (*cloudwatch.GetDashboardOutput, error) {
	var output cloudwatch.GetDashboardOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetInsightRuleReportFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetInsightRuleReportFuture) Get(ctx workflow.Context) (*cloudwatch.GetInsightRuleReportOutput, error) {
	var output cloudwatch.GetInsightRuleReportOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetMetricDataFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetMetricDataFuture) Get(ctx workflow.Context) (*cloudwatch.GetMetricDataOutput, error) {
	var output cloudwatch.GetMetricDataOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetMetricStatisticsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetMetricStatisticsFuture) Get(ctx workflow.Context) (*cloudwatch.GetMetricStatisticsOutput, error) {
	var output cloudwatch.GetMetricStatisticsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetMetricWidgetImageFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetMetricWidgetImageFuture) Get(ctx workflow.Context) (*cloudwatch.GetMetricWidgetImageOutput, error) {
	var output cloudwatch.GetMetricWidgetImageOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListDashboardsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListDashboardsFuture) Get(ctx workflow.Context) (*cloudwatch.ListDashboardsOutput, error) {
	var output cloudwatch.ListDashboardsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListMetricsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListMetricsFuture) Get(ctx workflow.Context) (*cloudwatch.ListMetricsOutput, error) {
	var output cloudwatch.ListMetricsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsForResourceFuture) Get(ctx workflow.Context) (*cloudwatch.ListTagsForResourceOutput, error) {
	var output cloudwatch.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutAnomalyDetectorFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutAnomalyDetectorFuture) Get(ctx workflow.Context) (*cloudwatch.PutAnomalyDetectorOutput, error) {
	var output cloudwatch.PutAnomalyDetectorOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutCompositeAlarmFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutCompositeAlarmFuture) Get(ctx workflow.Context) (*cloudwatch.PutCompositeAlarmOutput, error) {
	var output cloudwatch.PutCompositeAlarmOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutDashboardFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutDashboardFuture) Get(ctx workflow.Context) (*cloudwatch.PutDashboardOutput, error) {
	var output cloudwatch.PutDashboardOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutInsightRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutInsightRuleFuture) Get(ctx workflow.Context) (*cloudwatch.PutInsightRuleOutput, error) {
	var output cloudwatch.PutInsightRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutMetricAlarmFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutMetricAlarmFuture) Get(ctx workflow.Context) (*cloudwatch.PutMetricAlarmOutput, error) {
	var output cloudwatch.PutMetricAlarmOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutMetricDataFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutMetricDataFuture) Get(ctx workflow.Context) (*cloudwatch.PutMetricDataOutput, error) {
	var output cloudwatch.PutMetricDataOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SetAlarmStateFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SetAlarmStateFuture) Get(ctx workflow.Context) (*cloudwatch.SetAlarmStateOutput, error) {
	var output cloudwatch.SetAlarmStateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TagResourceFuture) Get(ctx workflow.Context) (*cloudwatch.TagResourceOutput, error) {
	var output cloudwatch.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UntagResourceFuture) Get(ctx workflow.Context) (*cloudwatch.UntagResourceOutput, error) {
	var output cloudwatch.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteAlarms(ctx workflow.Context, input *cloudwatch.DeleteAlarmsInput) (*cloudwatch.DeleteAlarmsOutput, error) {
	var output cloudwatch.DeleteAlarmsOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-DeleteAlarms", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteAlarmsAsync(ctx workflow.Context, input *cloudwatch.DeleteAlarmsInput) *DeleteAlarmsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-DeleteAlarms", input)
	return &DeleteAlarmsFuture{Future: future}
}

func (a *stub) DeleteAnomalyDetector(ctx workflow.Context, input *cloudwatch.DeleteAnomalyDetectorInput) (*cloudwatch.DeleteAnomalyDetectorOutput, error) {
	var output cloudwatch.DeleteAnomalyDetectorOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-DeleteAnomalyDetector", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteAnomalyDetectorAsync(ctx workflow.Context, input *cloudwatch.DeleteAnomalyDetectorInput) *DeleteAnomalyDetectorFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-DeleteAnomalyDetector", input)
	return &DeleteAnomalyDetectorFuture{Future: future}
}

func (a *stub) DeleteDashboards(ctx workflow.Context, input *cloudwatch.DeleteDashboardsInput) (*cloudwatch.DeleteDashboardsOutput, error) {
	var output cloudwatch.DeleteDashboardsOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-DeleteDashboards", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteDashboardsAsync(ctx workflow.Context, input *cloudwatch.DeleteDashboardsInput) *DeleteDashboardsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-DeleteDashboards", input)
	return &DeleteDashboardsFuture{Future: future}
}

func (a *stub) DeleteInsightRules(ctx workflow.Context, input *cloudwatch.DeleteInsightRulesInput) (*cloudwatch.DeleteInsightRulesOutput, error) {
	var output cloudwatch.DeleteInsightRulesOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-DeleteInsightRules", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteInsightRulesAsync(ctx workflow.Context, input *cloudwatch.DeleteInsightRulesInput) *DeleteInsightRulesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-DeleteInsightRules", input)
	return &DeleteInsightRulesFuture{Future: future}
}

func (a *stub) DescribeAlarmHistory(ctx workflow.Context, input *cloudwatch.DescribeAlarmHistoryInput) (*cloudwatch.DescribeAlarmHistoryOutput, error) {
	var output cloudwatch.DescribeAlarmHistoryOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-DescribeAlarmHistory", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAlarmHistoryAsync(ctx workflow.Context, input *cloudwatch.DescribeAlarmHistoryInput) *DescribeAlarmHistoryFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-DescribeAlarmHistory", input)
	return &DescribeAlarmHistoryFuture{Future: future}
}

func (a *stub) DescribeAlarms(ctx workflow.Context, input *cloudwatch.DescribeAlarmsInput) (*cloudwatch.DescribeAlarmsOutput, error) {
	var output cloudwatch.DescribeAlarmsOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-DescribeAlarms", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAlarmsAsync(ctx workflow.Context, input *cloudwatch.DescribeAlarmsInput) *DescribeAlarmsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-DescribeAlarms", input)
	return &DescribeAlarmsFuture{Future: future}
}

func (a *stub) DescribeAlarmsForMetric(ctx workflow.Context, input *cloudwatch.DescribeAlarmsForMetricInput) (*cloudwatch.DescribeAlarmsForMetricOutput, error) {
	var output cloudwatch.DescribeAlarmsForMetricOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-DescribeAlarmsForMetric", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAlarmsForMetricAsync(ctx workflow.Context, input *cloudwatch.DescribeAlarmsForMetricInput) *DescribeAlarmsForMetricFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-DescribeAlarmsForMetric", input)
	return &DescribeAlarmsForMetricFuture{Future: future}
}

func (a *stub) DescribeAnomalyDetectors(ctx workflow.Context, input *cloudwatch.DescribeAnomalyDetectorsInput) (*cloudwatch.DescribeAnomalyDetectorsOutput, error) {
	var output cloudwatch.DescribeAnomalyDetectorsOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-DescribeAnomalyDetectors", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAnomalyDetectorsAsync(ctx workflow.Context, input *cloudwatch.DescribeAnomalyDetectorsInput) *DescribeAnomalyDetectorsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-DescribeAnomalyDetectors", input)
	return &DescribeAnomalyDetectorsFuture{Future: future}
}

func (a *stub) DescribeInsightRules(ctx workflow.Context, input *cloudwatch.DescribeInsightRulesInput) (*cloudwatch.DescribeInsightRulesOutput, error) {
	var output cloudwatch.DescribeInsightRulesOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-DescribeInsightRules", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeInsightRulesAsync(ctx workflow.Context, input *cloudwatch.DescribeInsightRulesInput) *DescribeInsightRulesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-DescribeInsightRules", input)
	return &DescribeInsightRulesFuture{Future: future}
}

func (a *stub) DisableAlarmActions(ctx workflow.Context, input *cloudwatch.DisableAlarmActionsInput) (*cloudwatch.DisableAlarmActionsOutput, error) {
	var output cloudwatch.DisableAlarmActionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-DisableAlarmActions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisableAlarmActionsAsync(ctx workflow.Context, input *cloudwatch.DisableAlarmActionsInput) *DisableAlarmActionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-DisableAlarmActions", input)
	return &DisableAlarmActionsFuture{Future: future}
}

func (a *stub) DisableInsightRules(ctx workflow.Context, input *cloudwatch.DisableInsightRulesInput) (*cloudwatch.DisableInsightRulesOutput, error) {
	var output cloudwatch.DisableInsightRulesOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-DisableInsightRules", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisableInsightRulesAsync(ctx workflow.Context, input *cloudwatch.DisableInsightRulesInput) *DisableInsightRulesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-DisableInsightRules", input)
	return &DisableInsightRulesFuture{Future: future}
}

func (a *stub) EnableAlarmActions(ctx workflow.Context, input *cloudwatch.EnableAlarmActionsInput) (*cloudwatch.EnableAlarmActionsOutput, error) {
	var output cloudwatch.EnableAlarmActionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-EnableAlarmActions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) EnableAlarmActionsAsync(ctx workflow.Context, input *cloudwatch.EnableAlarmActionsInput) *EnableAlarmActionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-EnableAlarmActions", input)
	return &EnableAlarmActionsFuture{Future: future}
}

func (a *stub) EnableInsightRules(ctx workflow.Context, input *cloudwatch.EnableInsightRulesInput) (*cloudwatch.EnableInsightRulesOutput, error) {
	var output cloudwatch.EnableInsightRulesOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-EnableInsightRules", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) EnableInsightRulesAsync(ctx workflow.Context, input *cloudwatch.EnableInsightRulesInput) *EnableInsightRulesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-EnableInsightRules", input)
	return &EnableInsightRulesFuture{Future: future}
}

func (a *stub) GetDashboard(ctx workflow.Context, input *cloudwatch.GetDashboardInput) (*cloudwatch.GetDashboardOutput, error) {
	var output cloudwatch.GetDashboardOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-GetDashboard", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetDashboardAsync(ctx workflow.Context, input *cloudwatch.GetDashboardInput) *GetDashboardFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-GetDashboard", input)
	return &GetDashboardFuture{Future: future}
}

func (a *stub) GetInsightRuleReport(ctx workflow.Context, input *cloudwatch.GetInsightRuleReportInput) (*cloudwatch.GetInsightRuleReportOutput, error) {
	var output cloudwatch.GetInsightRuleReportOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-GetInsightRuleReport", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetInsightRuleReportAsync(ctx workflow.Context, input *cloudwatch.GetInsightRuleReportInput) *GetInsightRuleReportFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-GetInsightRuleReport", input)
	return &GetInsightRuleReportFuture{Future: future}
}

func (a *stub) GetMetricData(ctx workflow.Context, input *cloudwatch.GetMetricDataInput) (*cloudwatch.GetMetricDataOutput, error) {
	var output cloudwatch.GetMetricDataOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-GetMetricData", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetMetricDataAsync(ctx workflow.Context, input *cloudwatch.GetMetricDataInput) *GetMetricDataFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-GetMetricData", input)
	return &GetMetricDataFuture{Future: future}
}

func (a *stub) GetMetricStatistics(ctx workflow.Context, input *cloudwatch.GetMetricStatisticsInput) (*cloudwatch.GetMetricStatisticsOutput, error) {
	var output cloudwatch.GetMetricStatisticsOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-GetMetricStatistics", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetMetricStatisticsAsync(ctx workflow.Context, input *cloudwatch.GetMetricStatisticsInput) *GetMetricStatisticsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-GetMetricStatistics", input)
	return &GetMetricStatisticsFuture{Future: future}
}

func (a *stub) GetMetricWidgetImage(ctx workflow.Context, input *cloudwatch.GetMetricWidgetImageInput) (*cloudwatch.GetMetricWidgetImageOutput, error) {
	var output cloudwatch.GetMetricWidgetImageOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-GetMetricWidgetImage", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetMetricWidgetImageAsync(ctx workflow.Context, input *cloudwatch.GetMetricWidgetImageInput) *GetMetricWidgetImageFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-GetMetricWidgetImage", input)
	return &GetMetricWidgetImageFuture{Future: future}
}

func (a *stub) ListDashboards(ctx workflow.Context, input *cloudwatch.ListDashboardsInput) (*cloudwatch.ListDashboardsOutput, error) {
	var output cloudwatch.ListDashboardsOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-ListDashboards", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDashboardsAsync(ctx workflow.Context, input *cloudwatch.ListDashboardsInput) *ListDashboardsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-ListDashboards", input)
	return &ListDashboardsFuture{Future: future}
}

func (a *stub) ListMetrics(ctx workflow.Context, input *cloudwatch.ListMetricsInput) (*cloudwatch.ListMetricsOutput, error) {
	var output cloudwatch.ListMetricsOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-ListMetrics", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListMetricsAsync(ctx workflow.Context, input *cloudwatch.ListMetricsInput) *ListMetricsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-ListMetrics", input)
	return &ListMetricsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *cloudwatch.ListTagsForResourceInput) (*cloudwatch.ListTagsForResourceOutput, error) {
	var output cloudwatch.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *cloudwatch.ListTagsForResourceInput) *ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-ListTagsForResource", input)
	return &ListTagsForResourceFuture{Future: future}
}

func (a *stub) PutAnomalyDetector(ctx workflow.Context, input *cloudwatch.PutAnomalyDetectorInput) (*cloudwatch.PutAnomalyDetectorOutput, error) {
	var output cloudwatch.PutAnomalyDetectorOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-PutAnomalyDetector", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutAnomalyDetectorAsync(ctx workflow.Context, input *cloudwatch.PutAnomalyDetectorInput) *PutAnomalyDetectorFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-PutAnomalyDetector", input)
	return &PutAnomalyDetectorFuture{Future: future}
}

func (a *stub) PutCompositeAlarm(ctx workflow.Context, input *cloudwatch.PutCompositeAlarmInput) (*cloudwatch.PutCompositeAlarmOutput, error) {
	var output cloudwatch.PutCompositeAlarmOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-PutCompositeAlarm", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutCompositeAlarmAsync(ctx workflow.Context, input *cloudwatch.PutCompositeAlarmInput) *PutCompositeAlarmFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-PutCompositeAlarm", input)
	return &PutCompositeAlarmFuture{Future: future}
}

func (a *stub) PutDashboard(ctx workflow.Context, input *cloudwatch.PutDashboardInput) (*cloudwatch.PutDashboardOutput, error) {
	var output cloudwatch.PutDashboardOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-PutDashboard", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutDashboardAsync(ctx workflow.Context, input *cloudwatch.PutDashboardInput) *PutDashboardFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-PutDashboard", input)
	return &PutDashboardFuture{Future: future}
}

func (a *stub) PutInsightRule(ctx workflow.Context, input *cloudwatch.PutInsightRuleInput) (*cloudwatch.PutInsightRuleOutput, error) {
	var output cloudwatch.PutInsightRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-PutInsightRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutInsightRuleAsync(ctx workflow.Context, input *cloudwatch.PutInsightRuleInput) *PutInsightRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-PutInsightRule", input)
	return &PutInsightRuleFuture{Future: future}
}

func (a *stub) PutMetricAlarm(ctx workflow.Context, input *cloudwatch.PutMetricAlarmInput) (*cloudwatch.PutMetricAlarmOutput, error) {
	var output cloudwatch.PutMetricAlarmOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-PutMetricAlarm", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutMetricAlarmAsync(ctx workflow.Context, input *cloudwatch.PutMetricAlarmInput) *PutMetricAlarmFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-PutMetricAlarm", input)
	return &PutMetricAlarmFuture{Future: future}
}

func (a *stub) PutMetricData(ctx workflow.Context, input *cloudwatch.PutMetricDataInput) (*cloudwatch.PutMetricDataOutput, error) {
	var output cloudwatch.PutMetricDataOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-PutMetricData", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutMetricDataAsync(ctx workflow.Context, input *cloudwatch.PutMetricDataInput) *PutMetricDataFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-PutMetricData", input)
	return &PutMetricDataFuture{Future: future}
}

func (a *stub) SetAlarmState(ctx workflow.Context, input *cloudwatch.SetAlarmStateInput) (*cloudwatch.SetAlarmStateOutput, error) {
	var output cloudwatch.SetAlarmStateOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-SetAlarmState", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetAlarmStateAsync(ctx workflow.Context, input *cloudwatch.SetAlarmStateInput) *SetAlarmStateFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-SetAlarmState", input)
	return &SetAlarmStateFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *cloudwatch.TagResourceInput) (*cloudwatch.TagResourceOutput, error) {
	var output cloudwatch.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *cloudwatch.TagResourceInput) *TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-TagResource", input)
	return &TagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *cloudwatch.UntagResourceInput) (*cloudwatch.UntagResourceOutput, error) {
	var output cloudwatch.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudwatch-UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *cloudwatch.UntagResourceInput) *UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-UntagResource", input)
	return &UntagResourceFuture{Future: future}
}

func (a *stub) WaitUntilAlarmExists(ctx workflow.Context, input *cloudwatch.DescribeAlarmsInput) error {
	return workflow.ExecuteActivity(ctx, "aws-cloudwatch-WaitUntilAlarmExists", input).Get(ctx, nil)
}

func (a *stub) WaitUntilAlarmExistsAsync(ctx workflow.Context, input *cloudwatch.DescribeAlarmsInput) *clients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-WaitUntilAlarmExists", input)
	return clients.NewVoidFuture(future)
}

func (a *stub) WaitUntilCompositeAlarmExists(ctx workflow.Context, input *cloudwatch.DescribeAlarmsInput) error {
	return workflow.ExecuteActivity(ctx, "aws-cloudwatch-WaitUntilCompositeAlarmExists", input).Get(ctx, nil)
}

func (a *stub) WaitUntilCompositeAlarmExistsAsync(ctx workflow.Context, input *cloudwatch.DescribeAlarmsInput) *clients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudwatch-WaitUntilCompositeAlarmExists", input)
	return clients.NewVoidFuture(future)
}
