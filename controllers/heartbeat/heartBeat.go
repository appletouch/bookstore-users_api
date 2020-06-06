package heartbeat

import (
	"github.com/appletouch/bookstore-users_api/utils/dates"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HeartBeat(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]string{"status": "200", "title": "Health OK", "detail": dates.GetDateString()})
	//ctx.String(http.StatusOK,"hello from server")
}
