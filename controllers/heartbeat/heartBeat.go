package heartbeat

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HeartBeat(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}
