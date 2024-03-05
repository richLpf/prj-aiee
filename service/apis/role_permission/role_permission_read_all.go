package role_permission

import (
	"prj-aiee/model"
	"prj-aiee/service/apis/common"
	"prj-aiee/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetRolePermissionListApi struct {
	*common.ApiCommon
	Data GetRolePermissionListRequest
}
type GetRolePermissionListRequest struct {
	common.Page
}
type GetRolePermissionListResponse struct {
	Data	[]*model.RolePermission	`json:"data"`
	Count int64	`json:"count"`
}

func GetRolePermissionList(c *gin.Context) {
	req := &GetRolePermissionListApi{
		ApiCommon: common.NewRequest(c),
	}
	if err := req.BindRequest(&req.Data); err != nil {
		req.Reply.BindRequestFailed()
		req.Reply.Response(c)
		return
	}

	records, count := req.getRecords()
	if req.Reply.IsStatusFailed() {
		req.Reply.Response(c)
		return
	}

	res := GetRolePermissionListResponse{
		Data:  records,
		Count: count,
	}
	req.Reply.DataSet(res)
	req.Reply.Response(c)
}

func (req *GetRolePermissionListApi) getRecords() (records []*model.RolePermission, count int64) {
	records = make([]*model.RolePermission, 0)
	tx := mysql.DB.Model(&model.RolePermission{}).Where("status != ?", common.RecordStatusDeleted)

	tx = tx.Count(&count)
	if count == 0 {
		return
	}

	if req.Data.Limit > 0 && req.Data.Offset > 0 {
		tx = tx.Limit(req.Data.Limit).Offset(req.Data.Limit * (req.Data.Offset - 1))
	}
	err := tx.Order("created_at desc").Find(&records).Error
	if err != nil {
		req.Logger.Error("get records failed", zap.Error(err))
		req.Reply.ReadFailed()
		return
	}

	return records, count
}
