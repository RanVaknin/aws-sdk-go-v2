// Code generated by smithy-go-codegen DO NOT EDIT.

package verifiedpermissions

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/verifiedpermissions/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Modifies the validation setting for a policy store.
func (c *Client) UpdatePolicyStore(ctx context.Context, params *UpdatePolicyStoreInput, optFns ...func(*Options)) (*UpdatePolicyStoreOutput, error) {
	if params == nil {
		params = &UpdatePolicyStoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePolicyStore", params, optFns, c.addOperationUpdatePolicyStoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePolicyStoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePolicyStoreInput struct {

	// Specifies the ID of the policy store that you want to update
	//
	// This member is required.
	PolicyStoreId *string

	// A structure that defines the validation settings that want to enable for the
	// policy store.
	//
	// This member is required.
	ValidationSettings *types.ValidationSettings

	noSmithyDocumentSerde
}

type UpdatePolicyStoreOutput struct {

	// The Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the updated policy store.
	//
	// This member is required.
	Arn *string

	// The date and time that the policy store was originally created.
	//
	// This member is required.
	CreatedDate *time.Time

	// The date and time that the policy store was most recently updated.
	//
	// This member is required.
	LastUpdatedDate *time.Time

	// The ID of the updated policy store.
	//
	// This member is required.
	PolicyStoreId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePolicyStoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdatePolicyStore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdatePolicyStore{}, middleware.After)
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
	if err = addOpUpdatePolicyStoreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePolicyStore(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePolicyStore(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "verifiedpermissions",
		OperationName: "UpdatePolicyStore",
	}
}