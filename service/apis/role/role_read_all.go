package role

import (
	"prj-aiee/model"
	"prj-aiee/service/apis/common"
	"prj-aiee/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetRoleListApi struct {
	*common.ApiCommon
	Data GetRoleListRequest
}
type GetRoleListRequest struct {
	common.Page
}
type GetRoleListResponse struct {
	Data  []*model.Role `json:"data"`
	Count int64         `json:"count"`
}

// @Summary 获取角色列表
// @Description 角色列表
// @ID GetAllRole
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token" default(token)
// @Param RequestBody body GetRoleListRequest true "新角色的详细信息"
// @Success 201 {object} common.CommonResponse "角色信息"
// @Failure 400 {string} string "请求参数有误"
// @Router /v1/role/getList [post]
// @Tags 角色管理
func GetRoleList(c *gin.Context) {
	req := &GetRoleListApi{
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

	res := GetRoleListResponse{
		Data:  records,
		Count: count,
	}
	req.Reply.DataSet(res)
	req.Reply.Response(c)
}

func (req *GetRoleListApi) getRecords() (records []*model.Role, count int64) {
	records = make([]*model.Role, 0)
	tx := mysql.DB.Model(&model.Role{}).Where("status != ?", common.RecordStatusDeleted)

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
