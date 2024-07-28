package memo

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
)

type RouterGroup struct{ TRecordRouter }

var (
	tRecordApi = api.ApiGroupApp.MemoApiGroup.TRecordApi
)
