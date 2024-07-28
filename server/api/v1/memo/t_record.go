package memo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/memo"
	memoReq "github.com/flipped-aurora/gin-vue-admin/server/model/memo/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TRecordApi struct{}

// CreateTRecord 创建备忘录记录
// @Tags TRecord
// @Summary 创建备忘录记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body memo.TRecord true "创建备忘录记录"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /tRecord/createTRecord [post]
func (tRecordApi *TRecordApi) CreateTRecord(c *gin.Context) {
	var tRecord memo.TRecord
	err := c.ShouldBindJSON(&tRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	tRecord.CreatedBy = utils.GetUserID(c)
	err = tRecordService.CreateTRecord(&tRecord)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteTRecord 删除备忘录记录
// @Tags TRecord
// @Summary 删除备忘录记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body memo.TRecord true "删除备忘录记录"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /tRecord/deleteTRecord [delete]
func (tRecordApi *TRecordApi) DeleteTRecord(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := tRecordService.DeleteTRecord(ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteTRecordByIds 批量删除备忘录记录
// @Tags TRecord
// @Summary 批量删除备忘录记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /tRecord/deleteTRecordByIds [delete]
func (tRecordApi *TRecordApi) DeleteTRecordByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := tRecordService.DeleteTRecordByIds(IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateTRecord 更新备忘录记录
// @Tags TRecord
// @Summary 更新备忘录记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body memo.TRecord true "更新备忘录记录"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /tRecord/updateTRecord [put]
func (tRecordApi *TRecordApi) UpdateTRecord(c *gin.Context) {
	var tRecord memo.TRecord
	err := c.ShouldBindJSON(&tRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	tRecord.UpdatedBy = utils.GetUserID(c)
	err = tRecordService.UpdateTRecord(tRecord)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindTRecord 用id查询备忘录记录
// @Tags TRecord
// @Summary 用id查询备忘录记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query memo.TRecord true "用id查询备忘录记录"
// @Success 200 {object} response.Response{data=object{retRecord=memo.TRecord},msg=string} "查询成功"
// @Router /tRecord/findTRecord [get]
func (tRecordApi *TRecordApi) FindTRecord(c *gin.Context) {
	ID := c.Query("ID")
	retRecord, err := tRecordService.GetTRecord(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(retRecord, c)
}

// GetTRecordList 分页获取备忘录记录列表
// @Tags TRecord
// @Summary 分页获取备忘录记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query memoReq.TRecordSearch true "分页获取备忘录记录列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /tRecord/getTRecordList [get]
func (tRecordApi *TRecordApi) GetTRecordList(c *gin.Context) {
	var pageInfo memoReq.TRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := tRecordService.GetTRecordInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetTRecordPublic 不需要鉴权的备忘录记录接口
// @Tags TRecord
// @Summary 不需要鉴权的备忘录记录接口
// @accept application/json
// @Produce application/json
// @Param data query memoReq.TRecordSearch true "分页获取备忘录记录列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /tRecord/getTRecordPublic [get]
func (tRecordApi *TRecordApi) GetTRecordPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的备忘录记录接口信息",
	}, "获取成功", c)
}
