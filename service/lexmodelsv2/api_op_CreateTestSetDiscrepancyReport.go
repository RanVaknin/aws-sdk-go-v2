// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Create a report that describes the differences between the bot and the test set.
func (c *Client) CreateTestSetDiscrepancyReport(ctx context.Context, params *CreateTestSetDiscrepancyReportInput, optFns ...func(*Options)) (*CreateTestSetDiscrepancyReportOutput, error) {
	if params == nil {
		params = &CreateTestSetDiscrepancyReportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTestSetDiscrepancyReport", params, optFns, c.addOperationCreateTestSetDiscrepancyReportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTestSetDiscrepancyReportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTestSetDiscrepancyReportInput struct {

	// The target bot for the test set discrepancy report.
	//
	// This member is required.
	Target *types.TestSetDiscrepancyReportResourceTarget

	// The test set Id for the test set discrepancy report.
	//
	// This member is required.
	TestSetId *string

	noSmithyDocumentSerde
}

type CreateTestSetDiscrepancyReportOutput struct {

	// The creation date and time for the test set discrepancy report.
	CreationDateTime *time.Time

	// The target bot for the test set discrepancy report.
	Target *types.TestSetDiscrepancyReportResourceTarget

	// The unique identifier of the test set discrepancy report to describe.
	TestSetDiscrepancyReportId *string

	// The test set Id for the test set discrepancy report.
	TestSetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTestSetDiscrepancyReportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateTestSetDiscrepancyReport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateTestSetDiscrepancyReport{}, middleware.After)
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
	if err = addOpCreateTestSetDiscrepancyReportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTestSetDiscrepancyReport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTestSetDiscrepancyReport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "CreateTestSetDiscrepancyReport",
	}
}