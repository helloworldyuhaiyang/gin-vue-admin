package memo

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ TRecordApi }

var (
	tRecordService = service.ServiceGroupApp.MemoServiceGroup.TRecordService
)
