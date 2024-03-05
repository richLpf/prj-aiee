package user

import (
	"prj-aiee/model"
	"prj-aiee/service/apis/common"
	"prj-aiee/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetUserListApi struct {
	*common.ApiCommon
	Data GetUserListRequest
}
type GetUserListRequest struct {
	common.Page
}
type UserResponse struct {
	Id        int              `json:"id"`
	CreatedAt int              `json:"created_at"`
	Username  string           `json:"username"`
	Name      string           `json:"name"`
	Status    int              `json:"status"`
	Roles     []model.UserRole `json:"roles" gorm:"foreignKey:UserId"`
}
type GetUserListResponse struct {
	Data  []*UserResponse `json:"data"`
	Count int64           `json:"count"`
}

// @Summary 获取用户列表
// @Description 用户列表
// @ID GetAllUser
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token" default(token)
// @Param RequestBody body GetUserListRequest true "新用户的详细信息"
// @Success 201 {object} common.CommonResponse "用户信息"
// @Failure 400 {string} string "请求参数有误"
// @Router /v1/user/getList [post]
// @Tags 用户管理
func GetUserList(c *gin.Context) {
	req := &GetUserListApi{
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

	res := GetUserListResponse{
		Data:  records,
		Count: count,
	}
	req.Reply.DataSet(res)
	req.Reply.Response(c)
}

func (req *GetUserListApi) getRecords() (records []*UserResponse, count int64) {
	records = make([]*UserResponse, 0)
	tx := mysql.DB.Model(&model.User{}).Where("status != ?", common.RecordStatusDeleted)

	tx = tx.Count(&count)
	if count == 0 {
		return
	}

	if req.Data.Limit > 0 && req.Data.Offset > 0 {
		tx = tx.Limit(req.Data.Limit).Offset(req.Data.Limit * (req.Data.Offset - 1))
	}
	err := tx.Preload("Roles").Order("created_at desc").Find(&records).Error
	if err != nil {
		req.Logger.Error("get records failed", zap.Error(err))
		req.Reply.ReadFailed()
		return
	}

	return records, count
}
