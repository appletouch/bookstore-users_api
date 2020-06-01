package heartbeat

import (
	"github.com/appletouch/bookstore-users_api/utils/dates"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HeartBeat(ctx *gin.Context) {
	ctx.String(http.StatusOK, dates.GetDateString())
	//ctx.String(http.StatusOK,"hello from server")
}
