// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cognitosync

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/private/protocol/restjson"
)

// Amazon Cognito Sync provides an AWS service and client library that enable
// cross-device syncing of application-related user data. High-level client
// libraries are available for both iOS and Android. You can use these libraries
// to persist data locally so that it's available even if the device is offline.
// Developer credentials don't need to be stored on the mobile device to access
// the service. You can use Amazon Cognito to obtain a normalized user ID and
// credentials. User data is persisted in a dataset that can store up to 1 MB
// of key-value pairs, and you can have up to 20 datasets per user identity.
//
// With Amazon Cognito Sync, the data stored for each identity is accessible
// only to credentials assigned to that identity. In order to use the Cognito
// Sync service, you need to make API calls using credentials retrieved with
// Amazon Cognito Identity service (http://docs.aws.amazon.com/cognitoidentity/latest/APIReference/Welcome.html).
//
// If you want to use Cognito Sync in an Android or iOS application, you will
// probably want to make API calls via the AWS Mobile SDK. To learn more, see
// the Developer Guide for Android (http://docs.aws.amazon.com/mobile/sdkforandroid/developerguide/cognito-sync.html)
// and the Developer Guide for iOS (http://docs.aws.amazon.com/mobile/sdkforios/developerguide/cognito-sync.html).
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type CognitoSync struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "cognito-sync"

// New creates a new instance of the CognitoSync client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a CognitoSync client from just a session.
//     svc := cognitosync.New(mySession)
//
//     // Create a CognitoSync client with additional configuration
//     svc := cognitosync.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *CognitoSync {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *CognitoSync {
	svc := &CognitoSync{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2014-06-30",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restjson.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restjson.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restjson.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restjson.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a CognitoSync operation and runs any
// custom request initialization.
func (c *CognitoSync) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}