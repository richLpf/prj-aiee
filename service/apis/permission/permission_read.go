package permission

import (
	"prj-aiee/model"
	"prj-aiee/service/apis/common"
	"prj-aiee/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetPermissionApi struct {
	*common.ApiCommon
	Data GetPermissionRequest
}
type GetPermissionRequest struct {
	Id int `json:"id"`
}
type GetPermissionResponse struct {
	Data *model.Permission `json:"data"`
}

func GetPermission(c *gin.Context) {
	req := &GetPermissionApi{
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
	res := GetPermissionResponse{
		Data: record,
	}
	req.Reply.DataSet(res)
	req.Reply.Response(c)
}

func (req *GetPermissionApi) checkParams() {
	if req.Data.Id <= 0 {
		req.Reply.MsgSet(common.ReplyStatusMissingParam, common.ReplyMessageMissingParam)
		return
	}
}

func (req *GetPermissionApi) getRecord() (record *model.Permission) {
	err := mysql.DB.Model(&model.Permission{}).
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
