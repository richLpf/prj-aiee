package role_permission

import (
	"time"

	"prj-aiee/model"
	"prj-aiee/service/apis/common"
	"prj-aiee/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UpdateRolePermissionApi struct {
	*common.ApiCommon
	Data UpdateRolePermissionRequest
}
type UpdateRolePermissionRequest struct {
	model.RolePermission
}

func UpdateRolePermission(c *gin.Context) {
	req := &UpdateRolePermissionApi{
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

func (req *UpdateRolePermissionApi) checkParams() {
	if req.Data.Id <= 0 {
		req.Reply.MsgSet(common.ReplyStatusMissingParam, common.ReplyMessageMissingParam)
		return
	}
	var count int64
	err := mysql.DB.Model(&model.RolePermission{}).
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

func (req *UpdateRolePermissionApi) updateRecord() {
	now := time.Now().Unix()
	record := &model.RolePermission{
		PermissionId: req.Data.PermissionId,
		UpdatedAt:    now,
	}
	err := mysql.DB.Model(&model.RolePermission{}).
		Where("id = ?", req.Data.Id).
		Updates(&record).Error
	if err != nil {
		req.Logger.Error("update record failed", zap.Any("Data", record), zap.Error(err))
		req.Reply.UpdateFailed()
		return
	}
}
