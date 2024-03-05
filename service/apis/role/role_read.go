package role

import (
	"prj-aiee/model"
	"prj-aiee/service/apis/common"
	"prj-aiee/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetRoleApi struct {
	*common.ApiCommon
	Data GetRoleRequest
}
type GetRoleRequest struct {
	Id	int	`json:"id"`
}
type GetRoleResponse struct {
	Data	*model.Role	`json:"data"`
}

func GetRole(c *gin.Context) {
	req := &GetRoleApi{
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
	res := GetRoleResponse{
		Data:  record,
	}
	req.Reply.DataSet(res)
	req.Reply.Response(c)
}

func (req *GetRoleApi) checkParams() {
	if req.Data.Id <= 0 {
		req.Reply.MsgSet(common.ReplyStatusMissingParam, common.ReplyMessageMissingParam)
		return
	}
}

func (req *GetRoleApi) getRecord() (record *model.Role) {
	err := mysql.DB.Model(&model.Role{}).
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
