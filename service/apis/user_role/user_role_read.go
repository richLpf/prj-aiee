package user_role

import (
	"prj-aiee/model"
	"prj-aiee/service/apis/common"
	"prj-aiee/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetUserRoleApi struct {
	*common.ApiCommon
	Data GetUserRoleRequest
}
type GetUserRoleRequest struct {
	Id	int	`json:"id"`
}
type GetUserRoleResponse struct {
	Data	*model.UserRole	`json:"data"`
}

func GetUserRole(c *gin.Context) {
	req := &GetUserRoleApi{
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
	res := GetUserRoleResponse{
		Data:  record,
	}
	req.Reply.DataSet(res)
	req.Reply.Response(c)
}

func (req *GetUserRoleApi) checkParams() {
	if req.Data.Id <= 0 {
		req.Reply.MsgSet(common.ReplyStatusMissingParam, common.ReplyMessageMissingParam)
		return
	}
}

func (req *GetUserRoleApi) getRecord() (record *model.UserRole) {
	err := mysql.DB.Model(&model.UserRole{}).
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
