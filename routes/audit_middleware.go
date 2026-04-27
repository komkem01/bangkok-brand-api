package routes

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/url"
	"strings"
	"time"

	"bangkok-brand/app/modules"
	"bangkok-brand/app/modules/auth"
	"bangkok-brand/app/modules/entities/ent"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const maxAuditPayloadBytes = 1024 * 1024

type responseBodyRecorder struct {
	gin.ResponseWriter
	body bytes.Buffer
}

func (w *responseBodyRecorder) Write(b []byte) (int, error) {
	appendCapped(&w.body, b, maxAuditPayloadBytes)
	return w.ResponseWriter.Write(b)
}

func (w *responseBodyRecorder) WriteString(s string) (int, error) {
	appendCapped(&w.body, []byte(s), maxAuditPayloadBytes)
	return w.ResponseWriter.WriteString(s)
}

func auditLogMiddleware(mod *modules.Modules) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		requestPayload := captureRequestPayload(ctx)

		recorder := &responseBodyRecorder{ResponseWriter: ctx.Writer}
		ctx.Writer = recorder

		ctx.Next()

		if mod == nil || mod.ENT == nil || mod.ENT.Svc == nil {
			return
		}

		fullPath := ctx.FullPath()
		if fullPath == "" {
			fullPath = ctx.Request.URL.Path
		}

		actorID := actorIDFromContext(ctx)
		actorType := actorTypeFromContext(ctx)
		ip := emptyToNil(ctx.ClientIP())
		userAgent := emptyToNil(ctx.Request.UserAgent())
		requestID := emptyToNil(ctx.GetHeader("X-Request-ID"))
		responsePayload := captureResponsePayload(ctx, recorder.body.Bytes())

		oldValues := map[string]any{
			"request_payload": requestPayload,
		}

		newValues := map[string]any{
			"response_payload": responsePayload,
			"method":           ctx.Request.Method,
			"path":             fullPath,
			"status_code":      ctx.Writer.Status(),
			"latency_ms":       time.Since(start).Milliseconds(),
		}

		item := &ent.AuditLog{
			TableName:     tableNameFromPath(fullPath),
			Action:        ctx.Request.Method,
			ActorID:       actorID,
			ActorType:     actorType,
			OldValues:     oldValues,
			NewValues:     newValues,
			ChangedFields: []string{"request_payload", "response_payload", "method", "path", "status_code", "latency_ms"},
			IPAddress:     ip,
			UserAgent:     userAgent,
			RequestID:     requestID,
			CreatedAt:     time.Now(),
		}

		if _, err := mod.ENT.Svc.CreateAuditLog(ctx.Request.Context(), item); err != nil {
			slog.Error("audit middleware write failed", "error", err, "path", fullPath, "status", ctx.Writer.Status())
		}
	}
}

func captureRequestPayload(ctx *gin.Context) any {
	if ctx.Request == nil || ctx.Request.Body == nil {
		return nil
	}

	body, err := io.ReadAll(io.LimitReader(ctx.Request.Body, maxAuditPayloadBytes+1))
	if err != nil {
		return map[string]any{"_error": "read_request_failed"}
	}
	ctx.Request.Body = io.NopCloser(bytes.NewReader(body))

	if len(body) == 0 {
		if len(ctx.Request.URL.Query()) == 0 {
			return nil
		}
		return maskSensitive(convertStringSliceMap(ctx.Request.URL.Query()))
	}

	if len(body) > maxAuditPayloadBytes {
		body = body[:maxAuditPayloadBytes]
	}

	parsed := parsePayload(body)
	return maskSensitive(parsed)
}

func captureResponsePayload(ctx *gin.Context, body []byte) any {
	if len(body) == 0 {
		return nil
	}

	if len(body) > maxAuditPayloadBytes {
		body = body[:maxAuditPayloadBytes]
	}

	parsed := parsePayload(body)
	return maskSensitive(parsed)
}

func parsePayload(body []byte) any {
	var parsed any
	if err := json.Unmarshal(body, &parsed); err == nil {
		return parsed
	}

	text := strings.TrimSpace(string(body))
	if text == "" {
		return nil
	}

	if q, err := url.ParseQuery(text); err == nil && len(q) > 0 {
		return convertStringSliceMap(q)
	}

	return map[string]any{"_raw": text}
}

func appendCapped(buf *bytes.Buffer, b []byte, max int) {
	if max <= 0 || len(b) == 0 {
		return
	}

	remaining := max - buf.Len()
	if remaining <= 0 {
		return
	}

	if len(b) > remaining {
		_, _ = buf.Write(b[:remaining])
		return
	}

	_, _ = buf.Write(b)
}

func convertStringSliceMap(in map[string][]string) map[string]any {
	out := make(map[string]any, len(in))
	for k, v := range in {
		if len(v) == 1 {
			out[k] = v[0]
			continue
		}
		arr := make([]any, 0, len(v))
		for _, item := range v {
			arr = append(arr, item)
		}
		out[k] = arr
	}
	return out
}

func maskSensitive(v any) any {
	switch val := v.(type) {
	case map[string]any:
		out := make(map[string]any, len(val))
		for k, child := range val {
			if isSensitiveKey(k) {
				out[k] = "***"
				continue
			}
			out[k] = maskSensitive(child)
		}
		return out
	case []any:
		out := make([]any, 0, len(val))
		for _, item := range val {
			out = append(out, maskSensitive(item))
		}
		return out
	default:
		return v
	}
}

func isSensitiveKey(k string) bool {
	key := strings.ToLower(strings.TrimSpace(k))
	if key == "" {
		return false
	}

	sensitiveParts := []string{
		"password",
		"passwd",
		"pwd",
		"token",
		"authorization",
		"secret",
		"api_key",
		"apikey",
		"client_secret",
		"jwt",
	}

	for _, part := range sensitiveParts {
		if strings.Contains(key, part) {
			return true
		}
	}

	return false
}

func tableNameFromPath(path string) string {
	trimmed := strings.Trim(path, "/")
	if trimmed == "" {
		return "http_requests"
	}

	parts := strings.Split(trimmed, "/")
	if len(parts) >= 4 {
		return parts[3]
	}

	return strings.ReplaceAll(trimmed, "/", "_")
}

func actorIDFromContext(ctx *gin.Context) *uuid.UUID {
	v, ok := ctx.Get(auth.ContextKeyMemberID)
	if !ok {
		return nil
	}

	switch id := v.(type) {
	case uuid.UUID:
		idCopy := id
		return &idCopy
	case string:
		parsed, err := uuid.Parse(id)
		if err != nil {
			return nil
		}
		return &parsed
	default:
		return nil
	}
}

func actorTypeFromContext(ctx *gin.Context) *string {
	v, ok := ctx.Get(auth.ContextKeyMember)
	if !ok {
		return nil
	}

	if member, ok := v.(auth.AuthMember); ok {
		role := string(member.Role)
		if strings.TrimSpace(role) != "" {
			return &role
		}
		t := "member"
		return &t
	}

	t := "member"
	return &t
}

func emptyToNil(v string) *string {
	v = strings.TrimSpace(v)
	if v == "" {
		return nil
	}
	return &v
}
