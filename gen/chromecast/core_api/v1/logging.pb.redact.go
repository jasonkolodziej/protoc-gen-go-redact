// Code generated by protoc-gen-redact. DO NOT EDIT.
// source: chromecast/core_api/v1/logging.proto

package core_apiv1

import (
	context "context"
	redact "github.com/jasonkolodziej/protoc-gen-go-redact/gen/redact"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ grpc.Server
	_ context.Context
	_ redact.Redactor
	_ codes.Code
	_ status.Status
)

// Redact method implementation for SocketEvent
func (x *SocketEvent) Redact() {
	if x == nil {
		return
	}

	// Safe field: Type

	// Safe field: TimestampMicros

	// Safe field: Details

	// Safe field: NetReturnValue

	// Safe field: MessageNamespace

	// Safe field: ReadyState

	// Safe field: ConnectionState

	// Safe field: ReadState

	// Safe field: WriteState

	// Safe field: ErrorState

	// Safe field: ChallengeReplyErrorType

	// Safe field: NssErrorCode
}

// Redact method implementation for AggregatedSocketEvent
func (x *AggregatedSocketEvent) Redact() {
	if x == nil {
		return
	}

	// Safe field: Id

	// Safe field: EndpointId

	// Safe field: ChannelAuthType

	// Safe field: SocketEvent

	// Safe field: BytesRead

	// Safe field: BytesWritten
}

// Redact method implementation for Log
func (x *Log) Redact() {
	if x == nil {
		return
	}

	// Safe field: AggregatedSocketEvent

	// Safe field: NumEvictedAggregatedSocketEvents

	// Safe field: NumEvictedSocketEvents
}
