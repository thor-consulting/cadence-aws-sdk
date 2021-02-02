// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehend"
	"github.com/aws/aws-sdk-go/service/comprehend/comprehendiface"

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
	client comprehendiface.ComprehendAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := comprehend.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (comprehendiface.ComprehendAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return comprehend.New(sess), nil
}

func (a *Activities) BatchDetectDominantLanguage(ctx context.Context, input *comprehend.BatchDetectDominantLanguageInput) (*comprehend.BatchDetectDominantLanguageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchDetectDominantLanguageWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) BatchDetectEntities(ctx context.Context, input *comprehend.BatchDetectEntitiesInput) (*comprehend.BatchDetectEntitiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchDetectEntitiesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) BatchDetectKeyPhrases(ctx context.Context, input *comprehend.BatchDetectKeyPhrasesInput) (*comprehend.BatchDetectKeyPhrasesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchDetectKeyPhrasesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) BatchDetectSentiment(ctx context.Context, input *comprehend.BatchDetectSentimentInput) (*comprehend.BatchDetectSentimentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchDetectSentimentWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) BatchDetectSyntax(ctx context.Context, input *comprehend.BatchDetectSyntaxInput) (*comprehend.BatchDetectSyntaxOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchDetectSyntaxWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ClassifyDocument(ctx context.Context, input *comprehend.ClassifyDocumentInput) (*comprehend.ClassifyDocumentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ClassifyDocumentWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateDocumentClassifier(ctx context.Context, input *comprehend.CreateDocumentClassifierInput) (*comprehend.CreateDocumentClassifierOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateDocumentClassifierWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateEndpoint(ctx context.Context, input *comprehend.CreateEndpointInput) (*comprehend.CreateEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateEndpointWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateEntityRecognizer(ctx context.Context, input *comprehend.CreateEntityRecognizerInput) (*comprehend.CreateEntityRecognizerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateEntityRecognizerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteDocumentClassifier(ctx context.Context, input *comprehend.DeleteDocumentClassifierInput) (*comprehend.DeleteDocumentClassifierOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteDocumentClassifierWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteEndpoint(ctx context.Context, input *comprehend.DeleteEndpointInput) (*comprehend.DeleteEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteEndpointWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteEntityRecognizer(ctx context.Context, input *comprehend.DeleteEntityRecognizerInput) (*comprehend.DeleteEntityRecognizerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteEntityRecognizerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeDocumentClassificationJob(ctx context.Context, input *comprehend.DescribeDocumentClassificationJobInput) (*comprehend.DescribeDocumentClassificationJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeDocumentClassificationJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeDocumentClassifier(ctx context.Context, input *comprehend.DescribeDocumentClassifierInput) (*comprehend.DescribeDocumentClassifierOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeDocumentClassifierWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeDominantLanguageDetectionJob(ctx context.Context, input *comprehend.DescribeDominantLanguageDetectionJobInput) (*comprehend.DescribeDominantLanguageDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeDominantLanguageDetectionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeEndpoint(ctx context.Context, input *comprehend.DescribeEndpointInput) (*comprehend.DescribeEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeEndpointWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeEntitiesDetectionJob(ctx context.Context, input *comprehend.DescribeEntitiesDetectionJobInput) (*comprehend.DescribeEntitiesDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeEntitiesDetectionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeEntityRecognizer(ctx context.Context, input *comprehend.DescribeEntityRecognizerInput) (*comprehend.DescribeEntityRecognizerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeEntityRecognizerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeEventsDetectionJob(ctx context.Context, input *comprehend.DescribeEventsDetectionJobInput) (*comprehend.DescribeEventsDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeEventsDetectionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeKeyPhrasesDetectionJob(ctx context.Context, input *comprehend.DescribeKeyPhrasesDetectionJobInput) (*comprehend.DescribeKeyPhrasesDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeKeyPhrasesDetectionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribePiiEntitiesDetectionJob(ctx context.Context, input *comprehend.DescribePiiEntitiesDetectionJobInput) (*comprehend.DescribePiiEntitiesDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribePiiEntitiesDetectionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeSentimentDetectionJob(ctx context.Context, input *comprehend.DescribeSentimentDetectionJobInput) (*comprehend.DescribeSentimentDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeSentimentDetectionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeTopicsDetectionJob(ctx context.Context, input *comprehend.DescribeTopicsDetectionJobInput) (*comprehend.DescribeTopicsDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeTopicsDetectionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DetectDominantLanguage(ctx context.Context, input *comprehend.DetectDominantLanguageInput) (*comprehend.DetectDominantLanguageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DetectDominantLanguageWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DetectEntities(ctx context.Context, input *comprehend.DetectEntitiesInput) (*comprehend.DetectEntitiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DetectEntitiesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DetectKeyPhrases(ctx context.Context, input *comprehend.DetectKeyPhrasesInput) (*comprehend.DetectKeyPhrasesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DetectKeyPhrasesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DetectPiiEntities(ctx context.Context, input *comprehend.DetectPiiEntitiesInput) (*comprehend.DetectPiiEntitiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DetectPiiEntitiesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DetectSentiment(ctx context.Context, input *comprehend.DetectSentimentInput) (*comprehend.DetectSentimentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DetectSentimentWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DetectSyntax(ctx context.Context, input *comprehend.DetectSyntaxInput) (*comprehend.DetectSyntaxOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DetectSyntaxWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListDocumentClassificationJobs(ctx context.Context, input *comprehend.ListDocumentClassificationJobsInput) (*comprehend.ListDocumentClassificationJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListDocumentClassificationJobsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListDocumentClassifiers(ctx context.Context, input *comprehend.ListDocumentClassifiersInput) (*comprehend.ListDocumentClassifiersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListDocumentClassifiersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListDominantLanguageDetectionJobs(ctx context.Context, input *comprehend.ListDominantLanguageDetectionJobsInput) (*comprehend.ListDominantLanguageDetectionJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListDominantLanguageDetectionJobsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListEndpoints(ctx context.Context, input *comprehend.ListEndpointsInput) (*comprehend.ListEndpointsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListEndpointsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListEntitiesDetectionJobs(ctx context.Context, input *comprehend.ListEntitiesDetectionJobsInput) (*comprehend.ListEntitiesDetectionJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListEntitiesDetectionJobsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListEntityRecognizers(ctx context.Context, input *comprehend.ListEntityRecognizersInput) (*comprehend.ListEntityRecognizersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListEntityRecognizersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListEventsDetectionJobs(ctx context.Context, input *comprehend.ListEventsDetectionJobsInput) (*comprehend.ListEventsDetectionJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListEventsDetectionJobsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListKeyPhrasesDetectionJobs(ctx context.Context, input *comprehend.ListKeyPhrasesDetectionJobsInput) (*comprehend.ListKeyPhrasesDetectionJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListKeyPhrasesDetectionJobsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListPiiEntitiesDetectionJobs(ctx context.Context, input *comprehend.ListPiiEntitiesDetectionJobsInput) (*comprehend.ListPiiEntitiesDetectionJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListPiiEntitiesDetectionJobsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListSentimentDetectionJobs(ctx context.Context, input *comprehend.ListSentimentDetectionJobsInput) (*comprehend.ListSentimentDetectionJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListSentimentDetectionJobsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *comprehend.ListTagsForResourceInput) (*comprehend.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTopicsDetectionJobs(ctx context.Context, input *comprehend.ListTopicsDetectionJobsInput) (*comprehend.ListTopicsDetectionJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTopicsDetectionJobsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartDocumentClassificationJob(ctx context.Context, input *comprehend.StartDocumentClassificationJobInput) (*comprehend.StartDocumentClassificationJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartDocumentClassificationJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartDominantLanguageDetectionJob(ctx context.Context, input *comprehend.StartDominantLanguageDetectionJobInput) (*comprehend.StartDominantLanguageDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartDominantLanguageDetectionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartEntitiesDetectionJob(ctx context.Context, input *comprehend.StartEntitiesDetectionJobInput) (*comprehend.StartEntitiesDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartEntitiesDetectionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartEventsDetectionJob(ctx context.Context, input *comprehend.StartEventsDetectionJobInput) (*comprehend.StartEventsDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartEventsDetectionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartKeyPhrasesDetectionJob(ctx context.Context, input *comprehend.StartKeyPhrasesDetectionJobInput) (*comprehend.StartKeyPhrasesDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartKeyPhrasesDetectionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartPiiEntitiesDetectionJob(ctx context.Context, input *comprehend.StartPiiEntitiesDetectionJobInput) (*comprehend.StartPiiEntitiesDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartPiiEntitiesDetectionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartSentimentDetectionJob(ctx context.Context, input *comprehend.StartSentimentDetectionJobInput) (*comprehend.StartSentimentDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartSentimentDetectionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartTopicsDetectionJob(ctx context.Context, input *comprehend.StartTopicsDetectionJobInput) (*comprehend.StartTopicsDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartTopicsDetectionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopDominantLanguageDetectionJob(ctx context.Context, input *comprehend.StopDominantLanguageDetectionJobInput) (*comprehend.StopDominantLanguageDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopDominantLanguageDetectionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopEntitiesDetectionJob(ctx context.Context, input *comprehend.StopEntitiesDetectionJobInput) (*comprehend.StopEntitiesDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopEntitiesDetectionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopEventsDetectionJob(ctx context.Context, input *comprehend.StopEventsDetectionJobInput) (*comprehend.StopEventsDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopEventsDetectionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopKeyPhrasesDetectionJob(ctx context.Context, input *comprehend.StopKeyPhrasesDetectionJobInput) (*comprehend.StopKeyPhrasesDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopKeyPhrasesDetectionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopPiiEntitiesDetectionJob(ctx context.Context, input *comprehend.StopPiiEntitiesDetectionJobInput) (*comprehend.StopPiiEntitiesDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopPiiEntitiesDetectionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopSentimentDetectionJob(ctx context.Context, input *comprehend.StopSentimentDetectionJobInput) (*comprehend.StopSentimentDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopSentimentDetectionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopTrainingDocumentClassifier(ctx context.Context, input *comprehend.StopTrainingDocumentClassifierInput) (*comprehend.StopTrainingDocumentClassifierOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopTrainingDocumentClassifierWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopTrainingEntityRecognizer(ctx context.Context, input *comprehend.StopTrainingEntityRecognizerInput) (*comprehend.StopTrainingEntityRecognizerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopTrainingEntityRecognizerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *comprehend.TagResourceInput) (*comprehend.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *comprehend.UntagResourceInput) (*comprehend.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateEndpoint(ctx context.Context, input *comprehend.UpdateEndpointInput) (*comprehend.UpdateEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateEndpointWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
