package role_permission

import (
	"prj-aiee/model"
	"prj-aiee/service/apis/common"
	"prj-aiee/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetRolePermissionApi struct {
	*common.ApiCommon
	Data GetRolePermissionRequest
}
type GetRolePermissionRequest struct {
	Id	int	`json:"id"`
}
type GetRolePermissionResponse struct {
	Data	*model.RolePermission	`json:"data"`
}

func GetRolePermission(c *gin.Context) {
	req := &GetRolePermissionApi{
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
	record := req.getRecord()
	if req.Reply.IsStatusFailed() {
		req.Reply.Response(c)
		return
	}
	res := GetRolePermissionResponse{
		Data:  record,
	}
	req.Reply.DataSet(res)
	req.Reply.Response(c)
}

func (req *GetRolePermissionApi) checkParams() {
	if req.Data.Id <= 0 {
		req.Reply.MsgSet(common.ReplyStatusMissingParam, common.ReplyMessageMissingParam)
		return
	}
}

func (req *GetRolePermissionApi) getRecord() (record *model.RolePermission) {
	err := mysql.DB.Model(&model.RolePermission{}).
		Where("status != ?", common.RecordStatusDeleted).
		Where("id = ?", req.Data.Id).
		First(&record).Error
	if err != nil {
		req.Logger.Error("get record failed", zap.Error(err))
		req.Reply.ReadFailed()
		return
	}
	return record
}
