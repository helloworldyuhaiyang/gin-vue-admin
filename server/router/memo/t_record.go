package memo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TRecordRouter struct{}

// InitTRecordRouter 初始化 备忘录记录 路由信息
func (s *TRecordRouter) InitTRecordRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	tRecordRouter := Router.Group("tRecord").Use(middleware.OperationRecord())
	tRecordRouterWithoutRecord := Router.Group("tRecord")
	tRecordRouterWithoutAuth := PublicRouter.Group("tRecord")
	{
		tRecordRouter.POST("createTRecord", tRecordApi.CreateTRecord)             // 新建备忘录记录
		tRecordRouter.DELETE("deleteTRecord", tRecordApi.DeleteTRecord)           // 删除备忘录记录
		tRecordRouter.DELETE("deleteTRecordByIds", tRecordApi.DeleteTRecordByIds) // 批量删除备忘录记录
		tRecordRouter.PUT("updateTRecord", tRecordApi.UpdateTRecord)              // 更新备忘录记录
	}
	{
		tRecordRouterWithoutRecord.GET("findTRecord", tRecordApi.FindTRecord)       // 根据ID获取备忘录记录
		tRecordRouterWithoutRecord.GET("getTRecordList", tRecordApi.GetTRecordList) // 获取备忘录记录列表
	}
	{
		tRecordRouterWithoutAuth.GET("getTRecordPublic", tRecordApi.GetTRecordPublic) // 获取备忘录记录列表
	}
}
