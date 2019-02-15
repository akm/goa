// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// chatter service
//
// Command:
// $ goa gen goa.design/goa/examples/streaming/design -o
// $(GOPATH)/src/goa.design/goa/examples/streaming

package chattersvc

import (
	"context"

	"goa.design/goa"
	chattersvcviews "goa.design/goa/examples/streaming/gen/chatter/views"
	"goa.design/goa/security"
)

// The chatter service implements a simple client and server chat.
type Service interface {
	// Creates a valid JWT token for auth to chat.
	Login(context.Context, *LoginPayload) (res string, err error)
	// Echoes the message sent by the client.
	Echoer(*EchoerPayload, EchoerServerStream) (err error)
	// Listens to the messages sent by the client.
	Listener(*ListenerPayload, ListenerServerStream) (err error)
	// Summarizes the chat messages sent by the client.
	Summary(*SummaryPayload, SummaryServerStream) (err error)
	// Returns the chat messages sent to the server.
	// The "view" return value must have one of the following views
	//	- "tiny"
	//	- "default"
	History(*HistoryPayload, HistoryServerStream) (err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// BasicAuth implements the authorization logic for the Basic security scheme.
	BasicAuth(ctx context.Context, user, pass string, schema *security.BasicScheme) (context.Context, error)
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "chatter"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"login", "echoer", "listener", "summary", "history"}

// EchoerServerStream is the interface a "echoer" endpoint server stream must
// satisfy.
type EchoerServerStream interface {
	// Send streams instances of "string".
	Send(string) error
	// Recv reads instances of "string" from the stream.
	Recv() (string, error)
	// Close closes the stream.
	Close() error
	goa.Contextualizer
}

// EchoerClientStream is the interface a "echoer" endpoint client stream must
// satisfy.
type EchoerClientStream interface {
	// Send streams instances of "string".
	Send(string) error
	// Recv reads instances of "string" from the stream.
	Recv() (string, error)
	// Close closes the stream.
	Close() error
	goa.Contextualizer
}

// ListenerServerStream is the interface a "listener" endpoint server stream
// must satisfy.
type ListenerServerStream interface {
	// Recv reads instances of "string" from the stream.
	Recv() (string, error)
	// Close closes the stream.
	Close() error
	goa.Contextualizer
}

// ListenerClientStream is the interface a "listener" endpoint client stream
// must satisfy.
type ListenerClientStream interface {
	// Send streams instances of "string".
	Send(string) error
	// Close closes the stream.
	Close() error
	goa.Contextualizer
}

// SummaryServerStream is the interface a "summary" endpoint server stream must
// satisfy.
type SummaryServerStream interface {
	// SendAndClose streams instances of "ChatSummaryCollection" and closes the
	// stream.
	SendAndClose(ChatSummaryCollection) error
	// Recv reads instances of "string" from the stream.
	Recv() (string, error)
	goa.Contextualizer
}

// SummaryClientStream is the interface a "summary" endpoint client stream must
// satisfy.
type SummaryClientStream interface {
	// Send streams instances of "string".
	Send(string) error
	// CloseAndRecv stops sending messages to the stream and reads instances of
	// "ChatSummaryCollection" from the stream.
	CloseAndRecv() (ChatSummaryCollection, error)
	goa.Contextualizer
}

// HistoryServerStream is the interface a "history" endpoint server stream must
// satisfy.
type HistoryServerStream interface {
	// Send streams instances of "ChatSummary".
	Send(*ChatSummary) error
	// Close closes the stream.
	Close() error
	// SetView sets the view used to render the result before streaming.
	SetView(view string)
	goa.Contextualizer
}

// HistoryClientStream is the interface a "history" endpoint client stream must
// satisfy.
type HistoryClientStream interface {
	// Recv reads instances of "ChatSummary" from the stream.
	Recv() (*ChatSummary, error)
	goa.Contextualizer
}

// Credentials used to authenticate to retrieve JWT token
type LoginPayload struct {
	User     string
	Password string
}

// EchoerPayload is the payload type of the chatter service echoer method.
type EchoerPayload struct {
	// JWT used for authentication
	Token string
}

// ListenerPayload is the payload type of the chatter service listener method.
type ListenerPayload struct {
	// JWT used for authentication
	Token string
}

// SummaryPayload is the payload type of the chatter service summary method.
type SummaryPayload struct {
	// JWT used for authentication
	Token string
}

// ChatSummaryCollection is the result type of the chatter service summary
// method.
type ChatSummaryCollection []*ChatSummary

// HistoryPayload is the payload type of the chatter service history method.
type HistoryPayload struct {
	// JWT used for authentication
	Token string
	// View to use to render the result
	View *string
}

// ChatSummary is the result type of the chatter service history method.
type ChatSummary struct {
	// Message sent to the server
	Message string
	// Length of the message sent
	Length *int
	// Time at which the message was sent
	SentAt *string
}

// Credentials are invalid
type Unauthorized string

type InvalidScopes string

// Error returns an error description.
func (e Unauthorized) Error() string {
	return "Credentials are invalid"
}

// ErrorName returns "unauthorized".
func (e Unauthorized) ErrorName() string {
	return "unauthorized"
}

// Error returns an error description.
func (e InvalidScopes) Error() string {
	return ""
}

// ErrorName returns "invalid-scopes".
func (e InvalidScopes) ErrorName() string {
	return "invalid-scopes"
}

// NewChatSummaryCollection initializes result type ChatSummaryCollection from
// viewed result type ChatSummaryCollection.
func NewChatSummaryCollection(vres chattersvcviews.ChatSummaryCollection) ChatSummaryCollection {
	var res ChatSummaryCollection
	switch vres.View {
	case "tiny":
		res = newChatSummaryCollectionTiny(vres.Projected)
	case "default", "":
		res = newChatSummaryCollection(vres.Projected)
	}
	return res
}

// NewViewedChatSummaryCollection initializes viewed result type
// ChatSummaryCollection from result type ChatSummaryCollection using the given
// view.
func NewViewedChatSummaryCollection(res ChatSummaryCollection, view string) chattersvcviews.ChatSummaryCollection {
	var vres chattersvcviews.ChatSummaryCollection
	switch view {
	case "tiny":
		p := newChatSummaryCollectionViewTiny(res)
		vres = chattersvcviews.ChatSummaryCollection{p, "tiny"}
	case "default", "":
		p := newChatSummaryCollectionView(res)
		vres = chattersvcviews.ChatSummaryCollection{p, "default"}
	}
	return vres
}

// NewChatSummary initializes result type ChatSummary from viewed result type
// ChatSummary.
func NewChatSummary(vres *chattersvcviews.ChatSummary) *ChatSummary {
	var res *ChatSummary
	switch vres.View {
	case "tiny":
		res = newChatSummaryTiny(vres.Projected)
	case "default", "":
		res = newChatSummary(vres.Projected)
	}
	return res
}

// NewViewedChatSummary initializes viewed result type ChatSummary from result
// type ChatSummary using the given view.
func NewViewedChatSummary(res *ChatSummary, view string) *chattersvcviews.ChatSummary {
	var vres *chattersvcviews.ChatSummary
	switch view {
	case "tiny":
		p := newChatSummaryViewTiny(res)
		vres = &chattersvcviews.ChatSummary{p, "tiny"}
	case "default", "":
		p := newChatSummaryView(res)
		vres = &chattersvcviews.ChatSummary{p, "default"}
	}
	return vres
}

// newChatSummaryCollectionTiny converts projected type ChatSummaryCollection
// to service type ChatSummaryCollection.
func newChatSummaryCollectionTiny(vres chattersvcviews.ChatSummaryCollectionView) ChatSummaryCollection {
	res := make(ChatSummaryCollection, len(vres))
	for i, n := range vres {
		res[i] = newChatSummaryTiny(n)
	}
	return res
}

// newChatSummaryCollection converts projected type ChatSummaryCollection to
// service type ChatSummaryCollection.
func newChatSummaryCollection(vres chattersvcviews.ChatSummaryCollectionView) ChatSummaryCollection {
	res := make(ChatSummaryCollection, len(vres))
	for i, n := range vres {
		res[i] = newChatSummary(n)
	}
	return res
}

// newChatSummaryCollectionViewTiny projects result type ChatSummaryCollection
// into projected type ChatSummaryCollectionView using the "tiny" view.
func newChatSummaryCollectionViewTiny(res ChatSummaryCollection) chattersvcviews.ChatSummaryCollectionView {
	vres := make(chattersvcviews.ChatSummaryCollectionView, len(res))
	for i, n := range res {
		vres[i] = newChatSummaryViewTiny(n)
	}
	return vres
}

// newChatSummaryCollectionView projects result type ChatSummaryCollection into
// projected type ChatSummaryCollectionView using the "default" view.
func newChatSummaryCollectionView(res ChatSummaryCollection) chattersvcviews.ChatSummaryCollectionView {
	vres := make(chattersvcviews.ChatSummaryCollectionView, len(res))
	for i, n := range res {
		vres[i] = newChatSummaryView(n)
	}
	return vres
}

// newChatSummaryTiny converts projected type ChatSummary to service type
// ChatSummary.
func newChatSummaryTiny(vres *chattersvcviews.ChatSummaryView) *ChatSummary {
	res := &ChatSummary{}
	if vres.Message != nil {
		res.Message = *vres.Message
	}
	return res
}

// newChatSummary converts projected type ChatSummary to service type
// ChatSummary.
func newChatSummary(vres *chattersvcviews.ChatSummaryView) *ChatSummary {
	res := &ChatSummary{
		Length: vres.Length,
		SentAt: vres.SentAt,
	}
	if vres.Message != nil {
		res.Message = *vres.Message
	}
	return res
}

// newChatSummaryViewTiny projects result type ChatSummary into projected type
// ChatSummaryView using the "tiny" view.
func newChatSummaryViewTiny(res *ChatSummary) *chattersvcviews.ChatSummaryView {
	vres := &chattersvcviews.ChatSummaryView{
		Message: &res.Message,
	}
	return vres
}

// newChatSummaryView projects result type ChatSummary into projected type
// ChatSummaryView using the "default" view.
func newChatSummaryView(res *ChatSummary) *chattersvcviews.ChatSummaryView {
	vres := &chattersvcviews.ChatSummaryView{
		Message: &res.Message,
		Length:  res.Length,
		SentAt:  res.SentAt,
	}
	return vres
}
