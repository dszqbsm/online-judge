package middleware

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/iox"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/syncx"
	"github.com/zeromicro/go-zero/core/utils"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

const (
	limitDetailedBodyBytes = 4096
	defaultSlowThreshold   = time.Millisecond * 500
)

const (
	defaultClientIPFieldKey = "client_ip"
	defaultRequestFieldKey  = "request"
	defaultResponseFieldKey = "response"
)

type detailLoggedRequestFields struct {
	Method      string            `json:"method"`
	Uri         string            `json:"uri,omitempty"`
	Headers     map[string]string `json:"headers,omitempty"`
	QueryString map[string]string `json:"querystring,omitempty"`
	Body        string            `json:"body"`
}

type detailLoggedResponseFields struct {
	Headers map[string]string `json:"headers,omitempty"`
	Body    string            `json:"body"`
	Status  int               `json:"status"`
}

type detailLoggedResponseWriter struct {
	writer http.ResponseWriter
	code   int
	buf    *bytes.Buffer
}

func newDetailLoggedResponseWriter(writer http.ResponseWriter, buf *bytes.Buffer) *detailLoggedResponseWriter {
	return &detailLoggedResponseWriter{
		writer: writer,
		code:   http.StatusOK,
		buf:    buf,
	}
}

func (w *detailLoggedResponseWriter) Flush() {
	if flusher, ok := w.writer.(http.Flusher); ok {
		flusher.Flush()
	}
}

func (w *detailLoggedResponseWriter) Header() http.Header {
	return w.writer.Header()
}

// Hijack implements the http.Hijacker interface.
// This expands the Response to fulfill http.Hijacker if the underlying http.ResponseWriter supports it.
func (w *detailLoggedResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if hijacked, ok := w.writer.(http.Hijacker); ok {
		return hijacked.Hijack()
	}

	return nil, nil, errors.New("server doesn't support hijacking")
}

func (w *detailLoggedResponseWriter) Write(bs []byte) (int, error) {
	w.buf.Write(bs)
	return w.writer.Write(bs)
}

func (w *detailLoggedResponseWriter) WriteHeader(code int) {
	w.writer.WriteHeader(code)
	w.code = code
}

var slowThreshold = syncx.ForAtomicDuration(defaultSlowThreshold)

// SetSlowThreshold sets the slow threshold.
func SetSlowThreshold(threshold time.Duration) {
	slowThreshold.Set(threshold)
}

func NewInterceptorMiddleware() *InterceptorMiddleware {
	return &InterceptorMiddleware{}
}

type InterceptorMiddleware struct {
}

func (m *InterceptorMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		timer := utils.NewElapsedTimer()
		var buf bytes.Buffer
		lrw := newDetailLoggedResponseWriter(w, &buf)

		var dup io.ReadCloser
		r.Body, dup = iox.LimitDupReadCloser(r.Body, limitDetailedBodyBytes)
		next(lrw, r)
		r.Body = dup
		logDetails(r, lrw, timer)
	}
}

func logDetails(r *http.Request, response *detailLoggedResponseWriter, timer *utils.ElapsedTimer) {
	duration := timer.Duration()
	logger := logx.WithContext(r.Context()).WithDuration(duration)

	// Set client IP.
	clientIP, _, _ := net.SplitHostPort(r.RemoteAddr)
	logger = logger.WithFields(logx.Field(defaultClientIPFieldKey, clientIP))

	// Set request fields.
	requestFields := detailLoggedRequestFields{
		Method:      r.Method,
		Uri:         r.RequestURI,
		QueryString: make(map[string]string),
		Headers:     make(map[string]string),
	}
	queryValues := r.URL.Query()
	for key := range queryValues {
		requestFields.QueryString[strings.ToLower(key)] = queryValues.Get(key)
	}
	headers := r.Header
	for key := range headers {
		requestFields.Headers[strings.ToLower(key)] = headers.Get(key)
	}
	var bodyBuf bytes.Buffer
	_, _ = bodyBuf.ReadFrom(r.Body)
	requestFields.Body = bodyBuf.String()
	logger = logger.WithFields(logx.Field(defaultRequestFieldKey, requestFields))

	// Set response fields.
	responseFields := detailLoggedResponseFields{
		Headers: make(map[string]string),
	}
	headers = response.Header()
	for key := range headers {
		responseFields.Headers[strings.ToLower(key)] = headers.Get(key)
	}
	responseFields.Body = response.buf.String()
	responseFields.Status = response.code
	logger = logger.WithFields(logx.Field(defaultResponseFieldKey, responseFields))

	// Write log.
	if response.code < http.StatusInternalServerError {
		if duration > slowThreshold.Load() {
			logger.Slow()
		} else {
			logger.Info()
		}
	} else {
		logger.Error()
	}
}

func RpcLogDetail(ctx context.Context, method, in, out string, err error, timer *utils.ElapsedTimer) {
	duration := timer.Duration()
	logger := logx.WithContext(ctx).WithDuration(duration)

	// Set client IP.
	var clientIP string
	if p, ok := peer.FromContext(ctx); ok && p.Addr != net.Addr(nil) {
		clientIP, _, _ = net.SplitHostPort(p.Addr.String())
	}
	logger = logger.WithFields(logx.Field(defaultClientIPFieldKey, clientIP))

	// Set request fields.
	requestFields := detailLoggedRequestFields{
		Method: method,
		Body:   in,
	}
	logger = logger.WithFields(logx.Field(defaultRequestFieldKey, requestFields))

	// Set response fields.
	responseFields := detailLoggedResponseFields{
		Body:   out,
		Status: int(status.Code(err)),
	}
	logger = logger.WithFields(logx.Field(defaultResponseFieldKey, responseFields))

	// Write log.
	if err == nil {
		if duration > slowThreshold.Load() {
			logger.Slow()
		} else {
			logger.Info()
		}
	} else {
		logger.Error()
	}
}
