import service from '@/utils/request'

// @Tags TRecord
// @Summary 创建备忘录记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TRecord true "创建备忘录记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tRecord/createTRecord [post]
export const createTRecord = (data) => {
  return service({
    url: '/tRecord/createTRecord',
    method: 'post',
    data
  })
}

// @Tags TRecord
// @Summary 删除备忘录记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TRecord true "删除备忘录记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tRecord/deleteTRecord [delete]
export const deleteTRecord = (params) => {
  return service({
    url: '/tRecord/deleteTRecord',
    method: 'delete',
    params
  })
}

// @Tags TRecord
// @Summary 批量删除备忘录记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除备忘录记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tRecord/deleteTRecord [delete]
export const deleteTRecordByIds = (params) => {
  return service({
    url: '/tRecord/deleteTRecordByIds',
    method: 'delete',
    params
  })
}

// @Tags TRecord
// @Summary 更新备忘录记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TRecord true "更新备忘录记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tRecord/updateTRecord [put]
export const updateTRecord = (data) => {
  return service({
    url: '/tRecord/updateTRecord',
    method: 'put',
    data
  })
}

// @Tags TRecord
// @Summary 用id查询备忘录记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TRecord true "用id查询备忘录记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tRecord/findTRecord [get]
export const findTRecord = (params) => {
  return service({
    url: '/tRecord/findTRecord',
    method: 'get',
    params
  })
}

// @Tags TRecord
// @Summary 分页获取备忘录记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取备忘录记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tRecord/getTRecordList [get]
export const getTRecordList = (params) => {
  return service({
    url: '/tRecord/getTRecordList',
    method: 'get',
    params
  })
}
