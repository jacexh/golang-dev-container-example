package trace

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jacexh/golang-dev-container-example/pkg/infection"
)

const (
	CtxKeyRequestIndex = "request-index"
)

func GenContextWithRequestIndex(c *gin.Context) context.Context {
	index := MustExtractRequestIndex(c)
	ctx := infection.GenContextWithDefaultTimeout()
	return infection.Store(ctx, CtxKeyRequestIndex, index)
}
