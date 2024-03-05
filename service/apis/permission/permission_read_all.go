package permission

import (
	"prj-aiee/model"
	"prj-aiee/service/apis/common"
	"prj-aiee/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetPermissionListApi struct {
	*common.ApiCommon
	Data GetPermissionListRequest
}
type GetPermissionListRequest struct {
	common.Page
}
type GetPermissionListResponse struct {
	Data  []*model.Permission `json:"data"`
	Count int64               `json:"count"`
}

// @Summary 获取权限列表
// @Description 权限列表
// @ID GetAllPermission
// @Accept json
// @Produce json
// @Param RequestBody body GetPermissionListRequest true "新权限的详细信息"
// @Success 201 {object} common.CommonResponse "成功创建的用户信息"
// @Failure 400 {string} string "请求参数有误"
// @Router /v1/permission/getList [post]
// @Tags 权限管理
func GetPermissionList(c *gin.Context) {
	req := &GetPermissionListApi{
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

	res := GetPermissionListResponse{
		Data:  records,
		Count: count,
	}
	req.Reply.DataSet(res)
	req.Reply.Response(c)
}

func (req *GetPermissionListApi) getRecords() (records []*model.Permission, count int64) {
	records = make([]*model.Permission, 0)
	tx := mysql.DB.Model(&model.Permission{})

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
