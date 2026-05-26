package sse

import (
	"context"
	"time"

	"go.uber.org/zap"
)

// HubOption hub option
type HubOption func(*hubOptions)

type hubOptions struct {
	store Store
	// enable resend events from store specified by event ID, default is false, dependency store
	enableResendEvents bool

	ctx                context.Context
	cancel             context.CancelFunc
	logger             *zap.Logger
	pushBufferSize     int
	pushFailedHandleFn func(uid string, event *Event)
	workerNum          int
}

func defaultHubOptions() *hubOptions { _ = "STUB: not implemented"; return nil }

func (h *hubOptions) apply(opts ...HubOption) { _ = "STUB: not implemented"; return }

// WithContext set context and cancel
func WithContext(ctx context.Context, cancel context.CancelFunc) HubOption {
	_ = "STUB: not implemented"
	return *new(HubOption)
}

// WithStore set store
func WithStore(store Store) HubOption { _ = "STUB: not implemented"; return *new(HubOption) }

// WithEnableResendEvents enable resend events
func WithEnableResendEvents() HubOption { _ = "STUB: not implemented"; return *new(HubOption) }

// WithLogger set logger
func WithLogger(logger *zap.Logger) HubOption { _ = "STUB: not implemented"; return *new(HubOption) }

// WithPushBufferSize set push events buffer size
func WithPushBufferSize(size int) HubOption { _ = "STUB: not implemented"; return *new(HubOption) }

// WithPushFailedHandleFn set push failed handle function
func WithPushFailedHandleFn(fn func(uid string, event *Event)) HubOption {
	_ = "STUB: not implemented"
	return *new(HubOption)
}

// WithWorkerNum set worker num
func WithWorkerNum(num int) HubOption { _ = "STUB: not implemented"; return *new(HubOption) }

// ------------------------------------------------------------------------------------------

// UserEvent user event
type UserEvent struct {
	UID   string `json:"uid"`
	Event *Event `json:"event"`
}

// Hub event center, manage client connections, receive user events, and broadcast them to online users
type Hub struct {
	store Store
	// default is false, if it is enabled and store is not nil, send event messages starting
	// from the specified event ID after disconnecting and reconnecting.
	enableResendEvents bool

	clients            *SafeMap // userID -> *Client
	register           chan *UserClient
	unregister         chan *UserClient
	broadcast          chan *UserEvent
	asyncTaskPool      *AsyncTaskPool
	PushStats          *PushStats
	maxRetry           int
	pushBufferSize     int
	pushFailedHandleFn func(uid string, event *Event)

	ctx       context.Context
	cancel    context.CancelFunc
	zapLogger *zap.Logger
}

// NewHub create a new event center
func NewHub(opts ...HubOption) *Hub { _ = "STUB: not implemented"; return nil }

// default buffer size is 1000
// default worker num is 10

func (h *Hub) run() { _ = "STUB: not implemented"; return }

//h.zapLogger.Info("[sse] pushed event to client", zap.String("uid", ue.UID), zap.String("event_id", ue.Event.ID))

// Push to specified users or all online users
func (h *Hub) Push(uids []string, events ...*Event) error { _ = "STUB: not implemented"; return nil }

// push to specified users

// push to all online users

func pushOne(h *Hub, ue *UserEvent) { _ = "STUB: not implemented"; return }

// try to send without blocking, fast path

// use async task pool to submit push task with timeout retry

func pushAll(h *Hub, e *Event) { _ = "STUB: not implemented"; return }

// Asynchronous retry push with timeout logic
func tryPushWithTimeout(h *Hub, ue *UserEvent, timeout time.Duration) {
	_ = "STUB: not implemented"
	return
}

// all retry failed

// Asynchronous retry send with timeout logic
func trySendWithTimeout(h *Hub, cli *UserClient, ue *UserEvent, timeout time.Duration) {
	_ = "STUB: not implemented"
	return
}

// all retry failed

// resend events to specified client after disconnecting and reconnecting
func (h *Hub) resendEvents(client *UserClient, eventType string, lastEventID string) {
	_ = "STUB: not implemented"
	return
}

// update lastID to nextID

// if nextID is empty, it means all events have been sent, exit loop

// PrintPushStats print push stats
func (h *Hub) PrintPushStats() { _ = "STUB: not implemented"; return }

// PushHeartBeat push heart beat event to specified user
func (h *Hub) PushHeartBeat(uid string) { _ = "STUB: not implemented"; return }

// OnlineClientsNum get online clients num
func (h *Hub) OnlineClientsNum() int { _ = "STUB: not implemented"; return 0 }

// Close event center and stop all worker,
// By default, send a shutdown event to the client. If you want the client
// to automatically reconnect, please set the tryToReconnect parameter to false.
func (h *Hub) Close(tryToReconnect ...bool) { _ = "STUB: not implemented"; return }

// ------------------------------------------------------------------------------------------

// PushStats push stats
type PushStats struct {
	total   int64 // total push count
	success int64 // success push count
	failed  int64 // failed push count
	timeout int64 // timeout push count
}

// IncTotal increment total push count
func (s *PushStats) IncTotal() { _ = "STUB: not implemented"; return }

// IncSuccess increment success push count
func (s *PushStats) IncSuccess() { _ = "STUB: not implemented"; return }

// IncFailed increment failed push count
func (s *PushStats) IncFailed() { _ = "STUB: not implemented"; return }

// IncTimeout increment timeout push count
func (s *PushStats) IncTimeout() { _ = "STUB: not implemented"; return }

// Snapshot get push stats snapshot
func (s *PushStats) Snapshot() (total, success, failed, timeout int64) {
	_ = "STUB: not implemented" //nolint
	return 0, 0, 0, 0
}

func newStringID() string { _ = "STUB: not implemented"; return "" }
