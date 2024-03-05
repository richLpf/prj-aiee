package user_role

import (
	"time"

	"prj-aiee/model"
	"prj-aiee/service/apis/common"
	"prj-aiee/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CreateUserRoleApi struct {
	*common.ApiCommon
	Data *model.UserRole
}

// @Summary 用户绑定角色
// @Description 用户关联角色
// @ID createUserRole
// @Accept json
// @Produce json
// @Param RequestBody body model.UserRole true "用户绑定角色"
// @Success 201 {object} common.CommonResponse "用户信息"
// @Failure 400 {string} string "请求参数有误"
// @Router /v1/user_role/create [post]
// @Tags 用户管理
func CreateUserRole(c *gin.Context) {
	req := &CreateUserRoleApi{
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
	req.createRecord()
	if req.Reply.IsStatusFailed() {
		req.Reply.Response(c)
		return
	}

	req.Reply.Response(c)
}

func (req *CreateUserRoleApi) checkParams() {
}

func (req *CreateUserRoleApi) createRecord() {
	now := time.Now().Unix()
	record := &model.UserRole{
		RoleId:    req.Data.RoleId,
		CreatedAt: now,
		UpdatedAt: now,
	}
	err := mysql.DB.Model(&model.UserRole{}).Create(&record).Error
	if err != nil {
		req.Logger.Error("create record failed", zap.Any("Data", record), zap.Error(err))
		req.Reply.CreateFailed()
		return
	}
}
