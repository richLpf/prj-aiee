package role

import (
	"time"

	"prj-aiee/model"
	"prj-aiee/service/apis/common"
	"prj-aiee/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DeleteRoleApi struct {
	*common.ApiCommon
	Data DeleteRoleRequest
}
type DeleteRoleRequest struct {
	Id	int	`json:"id"`
}

func DeleteRole(c *gin.Context) {
	req := &DeleteRoleApi{
		ApiCommon: common.NewRequest(c),
	}
	if err := req.BindRequest(&req.Data); err != nil {
		req.Reply.BindRequestFailed()
		req.Reply.Response(c)
		return
	}
	req.checkParams()
	if req.Reply.IsStatusFailed() {
		req.Reply.Response(c)
		return
	}
	req.deleteRecord()
	if req.Reply.IsStatusFailed() {
		req.Reply.Response(c)
		return
	}

	req.Reply.Response(c)
}

func (req *DeleteRoleApi) checkParams() {
	if req.Data.Id <= 0 {
		req.Reply.MsgSet(common.ReplyStatusMissingParam, common.ReplyMessageMissingParam)
		return
	}
	var count int64
	err := mysql.DB.Model(&model.Role{}).
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

func (req *DeleteRoleApi) deleteRecord() {
	now := time.Now().Unix()
	record := &model.Role{
		Status:    common.RecordStatusDeleted,
		UpdatedAt: now,
	}
	err := mysql.DB.Model(&model.Role{}).
		Where("id = ?", req.Data.Id).
		Updates(&record).Error
	if err != nil {
		req.Logger.Error("delete record failed", zap.Any("Data", record), zap.Error(err))
		req.Reply.DeleteFailed()
		return
	}
}
