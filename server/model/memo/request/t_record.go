package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type TRecordSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	RecordName     string     `json:"recordName" form:"recordName" `
	CreatedBy      *int       `json:"createdBy" form:"createdBy" `
	request.PageInfo
}
