// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package medialive

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/medialive"
	"github.com/aws/aws-sdk-go/service/medialive/medialiveiface"

	"github.com/banzaicloud/cadence-aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client medialiveiface.MediaLiveAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := medialive.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (medialiveiface.MediaLiveAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return medialive.New(sess), nil
}

func (a *Activities) AcceptInputDeviceTransfer(ctx context.Context, input *medialive.AcceptInputDeviceTransferInput) (*medialive.AcceptInputDeviceTransferOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AcceptInputDeviceTransferWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) BatchDelete(ctx context.Context, input *medialive.BatchDeleteInput) (*medialive.BatchDeleteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchDeleteWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) BatchStart(ctx context.Context, input *medialive.BatchStartInput) (*medialive.BatchStartOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchStartWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) BatchStop(ctx context.Context, input *medialive.BatchStopInput) (*medialive.BatchStopOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchStopWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) BatchUpdateSchedule(ctx context.Context, input *medialive.BatchUpdateScheduleInput) (*medialive.BatchUpdateScheduleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchUpdateScheduleWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CancelInputDeviceTransfer(ctx context.Context, input *medialive.CancelInputDeviceTransferInput) (*medialive.CancelInputDeviceTransferOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CancelInputDeviceTransferWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateChannel(ctx context.Context, input *medialive.CreateChannelInput) (*medialive.CreateChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateChannelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateInput(ctx context.Context, input *medialive.CreateInputInput) (*medialive.CreateInputOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateInputWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateInputSecurityGroup(ctx context.Context, input *medialive.CreateInputSecurityGroupInput) (*medialive.CreateInputSecurityGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateInputSecurityGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateMultiplex(ctx context.Context, input *medialive.CreateMultiplexInput) (*medialive.CreateMultiplexOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateMultiplexWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateMultiplexProgram(ctx context.Context, input *medialive.CreateMultiplexProgramInput) (*medialive.CreateMultiplexProgramOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateMultiplexProgramWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateTags(ctx context.Context, input *medialive.CreateTagsInput) (*medialive.CreateTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateTagsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteChannel(ctx context.Context, input *medialive.DeleteChannelInput) (*medialive.DeleteChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteChannelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteInput(ctx context.Context, input *medialive.DeleteInputInput) (*medialive.DeleteInputOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteInputWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteInputSecurityGroup(ctx context.Context, input *medialive.DeleteInputSecurityGroupInput) (*medialive.DeleteInputSecurityGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteInputSecurityGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteMultiplex(ctx context.Context, input *medialive.DeleteMultiplexInput) (*medialive.DeleteMultiplexOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteMultiplexWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteMultiplexProgram(ctx context.Context, input *medialive.DeleteMultiplexProgramInput) (*medialive.DeleteMultiplexProgramOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteMultiplexProgramWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteReservation(ctx context.Context, input *medialive.DeleteReservationInput) (*medialive.DeleteReservationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteReservationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteSchedule(ctx context.Context, input *medialive.DeleteScheduleInput) (*medialive.DeleteScheduleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteScheduleWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteTags(ctx context.Context, input *medialive.DeleteTagsInput) (*medialive.DeleteTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteTagsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeChannel(ctx context.Context, input *medialive.DescribeChannelInput) (*medialive.DescribeChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeChannelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeInput(ctx context.Context, input *medialive.DescribeInputInput) (*medialive.DescribeInputOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeInputWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeInputDevice(ctx context.Context, input *medialive.DescribeInputDeviceInput) (*medialive.DescribeInputDeviceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeInputDeviceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeInputDeviceThumbnail(ctx context.Context, input *medialive.DescribeInputDeviceThumbnailInput) (*medialive.DescribeInputDeviceThumbnailOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeInputDeviceThumbnailWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeInputSecurityGroup(ctx context.Context, input *medialive.DescribeInputSecurityGroupInput) (*medialive.DescribeInputSecurityGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeInputSecurityGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeMultiplex(ctx context.Context, input *medialive.DescribeMultiplexInput) (*medialive.DescribeMultiplexOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeMultiplexWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeMultiplexProgram(ctx context.Context, input *medialive.DescribeMultiplexProgramInput) (*medialive.DescribeMultiplexProgramOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeMultiplexProgramWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeOffering(ctx context.Context, input *medialive.DescribeOfferingInput) (*medialive.DescribeOfferingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeOfferingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeReservation(ctx context.Context, input *medialive.DescribeReservationInput) (*medialive.DescribeReservationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeReservationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeSchedule(ctx context.Context, input *medialive.DescribeScheduleInput) (*medialive.DescribeScheduleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeScheduleWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListChannels(ctx context.Context, input *medialive.ListChannelsInput) (*medialive.ListChannelsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListChannelsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListInputDeviceTransfers(ctx context.Context, input *medialive.ListInputDeviceTransfersInput) (*medialive.ListInputDeviceTransfersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListInputDeviceTransfersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListInputDevices(ctx context.Context, input *medialive.ListInputDevicesInput) (*medialive.ListInputDevicesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListInputDevicesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListInputSecurityGroups(ctx context.Context, input *medialive.ListInputSecurityGroupsInput) (*medialive.ListInputSecurityGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListInputSecurityGroupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListInputs(ctx context.Context, input *medialive.ListInputsInput) (*medialive.ListInputsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListInputsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListMultiplexPrograms(ctx context.Context, input *medialive.ListMultiplexProgramsInput) (*medialive.ListMultiplexProgramsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListMultiplexProgramsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListMultiplexes(ctx context.Context, input *medialive.ListMultiplexesInput) (*medialive.ListMultiplexesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListMultiplexesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListOfferings(ctx context.Context, input *medialive.ListOfferingsInput) (*medialive.ListOfferingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListOfferingsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListReservations(ctx context.Context, input *medialive.ListReservationsInput) (*medialive.ListReservationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListReservationsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *medialive.ListTagsForResourceInput) (*medialive.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PurchaseOffering(ctx context.Context, input *medialive.PurchaseOfferingInput) (*medialive.PurchaseOfferingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PurchaseOfferingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RejectInputDeviceTransfer(ctx context.Context, input *medialive.RejectInputDeviceTransferInput) (*medialive.RejectInputDeviceTransferOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RejectInputDeviceTransferWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartChannel(ctx context.Context, input *medialive.StartChannelInput) (*medialive.StartChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartChannelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartMultiplex(ctx context.Context, input *medialive.StartMultiplexInput) (*medialive.StartMultiplexOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartMultiplexWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopChannel(ctx context.Context, input *medialive.StopChannelInput) (*medialive.StopChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopChannelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopMultiplex(ctx context.Context, input *medialive.StopMultiplexInput) (*medialive.StopMultiplexOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopMultiplexWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TransferInputDevice(ctx context.Context, input *medialive.TransferInputDeviceInput) (*medialive.TransferInputDeviceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TransferInputDeviceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateChannel(ctx context.Context, input *medialive.UpdateChannelInput) (*medialive.UpdateChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateChannelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateChannelClass(ctx context.Context, input *medialive.UpdateChannelClassInput) (*medialive.UpdateChannelClassOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateChannelClassWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateInput(ctx context.Context, input *medialive.UpdateInputInput) (*medialive.UpdateInputOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateInputWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateInputDevice(ctx context.Context, input *medialive.UpdateInputDeviceInput) (*medialive.UpdateInputDeviceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateInputDeviceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateInputSecurityGroup(ctx context.Context, input *medialive.UpdateInputSecurityGroupInput) (*medialive.UpdateInputSecurityGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateInputSecurityGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateMultiplex(ctx context.Context, input *medialive.UpdateMultiplexInput) (*medialive.UpdateMultiplexOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateMultiplexWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateMultiplexProgram(ctx context.Context, input *medialive.UpdateMultiplexProgramInput) (*medialive.UpdateMultiplexProgramOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateMultiplexProgramWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateReservation(ctx context.Context, input *medialive.UpdateReservationInput) (*medialive.UpdateReservationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateReservationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) WaitUntilChannelCreated(ctx context.Context, input *medialive.DescribeChannelInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilChannelCreatedWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilChannelDeleted(ctx context.Context, input *medialive.DescribeChannelInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilChannelDeletedWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilChannelRunning(ctx context.Context, input *medialive.DescribeChannelInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilChannelRunningWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilChannelStopped(ctx context.Context, input *medialive.DescribeChannelInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilChannelStoppedWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilInputAttached(ctx context.Context, input *medialive.DescribeInputInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilInputAttachedWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilInputDeleted(ctx context.Context, input *medialive.DescribeInputInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilInputDeletedWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilInputDetached(ctx context.Context, input *medialive.DescribeInputInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilInputDetachedWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilMultiplexCreated(ctx context.Context, input *medialive.DescribeMultiplexInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilMultiplexCreatedWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilMultiplexDeleted(ctx context.Context, input *medialive.DescribeMultiplexInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilMultiplexDeletedWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilMultiplexRunning(ctx context.Context, input *medialive.DescribeMultiplexInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilMultiplexRunningWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilMultiplexStopped(ctx context.Context, input *medialive.DescribeMultiplexInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilMultiplexStoppedWithContext(ctx, input, options...))
	})
}
