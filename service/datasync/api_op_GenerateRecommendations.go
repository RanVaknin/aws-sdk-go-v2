// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates recommendations about where to migrate your data to in Amazon Web
// Services. Recommendations are generated based on information that DataSync
// Discovery collects about your on-premises storage system's resources. For more
// information, see Recommendations provided by DataSync Discovery (https://docs.aws.amazon.com/datasync/latest/userguide/discovery-understand-recommendations.html)
// . Once generated, you can view your recommendations by using the
// DescribeStorageSystemResources (https://docs.aws.amazon.com/datasync/latest/userguide/API_DescribeStorageSystemResources.html)
// operation. If your discovery job completes successfully (https://docs.aws.amazon.com/datasync/latest/userguide/discovery-job-statuses.html#discovery-job-statuses-table)
// , you don't need to use this operation. DataSync Discovery generates the
// recommendations for you automatically.
func (c *Client) GenerateRecommendations(ctx context.Context, params *GenerateRecommendationsInput, optFns ...func(*Options)) (*GenerateRecommendationsOutput, error) {
	if params == nil {
		params = &GenerateRecommendationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GenerateRecommendations", params, optFns, c.addOperationGenerateRecommendationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GenerateRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GenerateRecommendationsInput struct {

	// Specifies the Amazon Resource Name (ARN) of the discovery job that collects
	// information about your on-premises storage system.
	//
	// This member is required.
	DiscoveryJobArn *string

	// Specifies the universally unique identifiers (UUIDs) of the resources in your
	// storage system that you want recommendations on.
	//
	// This member is required.
	ResourceIds []string

	// Specifies the type of resource in your storage system that you want
	// recommendations on.
	//
	// This member is required.
	ResourceType types.DiscoveryResourceType

	noSmithyDocumentSerde
}

type GenerateRecommendationsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGenerateRecommendationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGenerateRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGenerateRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opGenerateRecommendationsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGenerateRecommendationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGenerateRecommendations(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opGenerateRecommendationsMiddleware struct {
}

func (*endpointPrefix_opGenerateRecommendationsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGenerateRecommendationsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "discovery-" + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opGenerateRecommendationsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opGenerateRecommendationsMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opGenerateRecommendations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "GenerateRecommendations",
	}
}