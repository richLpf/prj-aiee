package permission

import (
	"time"

	"prj-aiee/model"
	"prj-aiee/service/apis/common"
	"prj-aiee/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UpdatePermissionApi struct {
	*common.ApiCommon
	Data UpdatePermissionRequest
}
type UpdatePermissionRequest struct {
	model.Permission
}

func UpdatePermission(c *gin.Context) {
	req := &UpdatePermissionApi{
		ApiCommon: common.NewRequest(c),
	}
	if err := req.BindRequest(&req.Data); err != nil {
		req.Reply.BindRequestFailed().Response(c)
		return
	}
	req.checkParams()
	if req.Reply.IsStatusFailed() {
		req.Reply.Response(c)
		return
	}
	req.updateRecord()
	if req.Reply.IsStatusFailed() {
		req.Reply.Response(c)
		return
	}

	req.Reply.Response(c)
}

func (req *UpdatePermissionApi) checkParams() {
	if req.Data.Id <= 0 {
		req.Reply.MsgSet(common.ReplyStatusMissingParam, common.ReplyMessageMissingParam)
		return
	}
	var count int64
	err := mysql.DB.Model(&model.Permission{}).
		Where("status != ?", common.RecordStatusDeleted).
		Where("id = ?", req.Data.Id).
		Count(&count).Error
	if err != nil {
		req.Logger.Error("check record failed", zap.Error(err))
		req.Reply.ReadFailed()
		return
	}
	if count <= 0 {
		req.Reply.MsgSet(common.ReplyStatusBindRequestFailed, common.ReplyMessageBindRequestFailed)
		return
	}
}

func (req *UpdatePermissionApi) updateRecord() {
	now := time.Now().Unix()
	record := &model.Permission{
		Key:       req.Data.Key,
		Type:      req.Data.Type,
		Desc:      req.Data.Desc,
		Attribute: req.Data.Attribute,
		CreatedAt: req.Data.CreatedAt,
		UpdatedAt: now,
	}
	err := mysql.DB.Model(&model.Permission{}).
		Where("id = ?", req.Data.Id).
		Updates(&record).Error
	if err != nil {
		req.Logger.Error("update record failed", zap.Any("Data", record), zap.Error(err))
		req.Reply.UpdateFailed()
		return
	}
}
