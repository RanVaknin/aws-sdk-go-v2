// Code generated by smithy-go-codegen DO NOT EDIT.

// Package ivschat provides the API client, operations, and parameter types for
// Amazon Interactive Video Service Chat.
//
//	Introduction
//
// The Amazon IVS Chat control-plane API enables you to create and manage Amazon
// IVS Chat resources. You also need to integrate with the [Amazon IVS Chat Messaging API], to enable users to
// interact with chat rooms in real time.
//
// The API is an AWS regional service. For a list of supported regions and Amazon
// IVS Chat HTTPS service endpoints, see the Amazon IVS Chat information on the [Amazon IVS page]in
// the AWS General Reference.
//
// This document describes HTTP operations. There is a separate messaging API for
// managing Chat resources; see the [Amazon IVS Chat Messaging API Reference].
//
// Notes on terminology:
//
//   - You create service applications using the Amazon IVS Chat API. We refer to
//     these as applications.
//
//   - You create front-end client applications (browser and Android/iOS apps)
//     using the Amazon IVS Chat Messaging API. We refer to these as clients.
//
// # Resources
//
// The following resources are part of Amazon IVS Chat:
//
//   - LoggingConfiguration — A configuration that allows customers to store and
//     record sent messages in a chat room. See the Logging Configuration endpoints for
//     more information.
//
//   - Room — The central Amazon IVS Chat resource through which clients connect
//     to and exchange chat messages. See the Room endpoints for more information.
//
// # Tagging
//
// A tag is a metadata label that you assign to an AWS resource. A tag comprises a
// key and a value, both set by you. For example, you might set a tag as
// topic:nature to label a particular video category. See [Best practices and strategies] in Tagging Amazon Web
// Services Resources and Tag Editor for details, including restrictions that apply
// to tags and "Tag naming limits and requirements"; Amazon IVS Chat has no
// service-specific constraints beyond what is documented there.
//
// Tags can help you identify and organize your AWS resources. For example, you
// can use the same tag for different resources to indicate that they are related.
// You can also use tags to manage access (see [Access Tags]).
//
// The Amazon IVS Chat API has these tag-related operations: TagResource, UntagResource, and ListTagsForResource. The
// following resource supports tagging: Room.
//
// At most 50 tags can be applied to a resource.
//
// # API Access Security
//
// Your Amazon IVS Chat applications (service applications and clients) must be
// authenticated and authorized to access Amazon IVS Chat resources. Note the
// differences between these concepts:
//
//   - Authentication is about verifying identity. Requests to the Amazon IVS Chat
//     API must be signed to verify your identity.
//
//   - Authorization is about granting permissions. Your IAM roles need to have
//     permissions for Amazon IVS Chat API requests.
//
// Users (viewers) connect to a room using secure access tokens that you create
// using the CreateChatTokenoperation through the AWS SDK. You call CreateChatToken for every
// user’s chat session, passing identity and authorization information about the
// user.
//
// # Signing API Requests
//
// HTTP API requests must be signed with an AWS SigV4 signature using your AWS
// security credentials. The AWS Command Line Interface (CLI) and the AWS SDKs take
// care of signing the underlying API calls for you. However, if your application
// calls the Amazon IVS Chat HTTP API directly, it’s your responsibility to sign
// the requests.
//
// You generate a signature using valid AWS credentials for an IAM role that has
// permission to perform the requested action. For example, DeleteMessage requests
// must be made using an IAM role that has the ivschat:DeleteMessage permission.
//
// For more information:
//
//   - Authentication and generating signatures — See [Authenticating Requests (Amazon Web Services Signature Version 4)]in the Amazon Web Services
//     General Reference.
//
//   - Managing Amazon IVS permissions — See [Identity and Access Management]on the Security page of the Amazon
//     IVS User Guide.
//
// Amazon Resource Names (ARNs)
//
// ARNs uniquely identify AWS resources. An ARN is required when you need to
// specify a resource unambiguously across all of AWS, such as in IAM policies and
// API calls. For more information, see [Amazon Resource Names]in the AWS General Reference.
//
// [Amazon Resource Names]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
// [Amazon IVS page]: https://docs.aws.amazon.com/general/latest/gr/ivs.html
// [Identity and Access Management]: https://docs.aws.amazon.com/ivs/latest/userguide/security-iam.html
// [Amazon IVS Chat Messaging API]: https://docs.aws.amazon.com/ivs/latest/chatmsgapireference/chat-messaging-api.html
// [Best practices and strategies]: https://docs.aws.amazon.com/tag-editor/latest/userguide/best-practices-and-strats.html
// [Amazon IVS Chat Messaging API Reference]: https://docs.aws.amazon.com/ivs/latest/chatmsgapireference/chat-messaging-api.html
// [Access Tags]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html
// [Authenticating Requests (Amazon Web Services Signature Version 4)]: https://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-authenticating-requests.html
package ivschat
