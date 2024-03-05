package role_permission

import (
	"time"

	"prj-aiee/model"
	"prj-aiee/service/apis/common"
	"prj-aiee/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CreateRolePermissionApi struct {
	*common.ApiCommon
	Data *model.RolePermission
}

// @Summary 角色关联权限
// @Description 关联权限
// @ID CreateRolePermission
// @Accept json
// @Produce json
// @Param RequestBody body model.RolePermission true "角色关联权限"
// @Success 201 {object} common.CommonResponse "成功创建的用户信息"
// @Failure 400 {string} string "请求参数有误"
// @Router /v1/role_permission/create [post]
// @Tags 角色管理
func CreateRolePermission(c *gin.Context) {
	req := &CreateRolePermissionApi{
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

func (req *CreateRolePermissionApi) checkParams() {
}

func (req *CreateRolePermissionApi) createRecord() {
	now := time.Now().Unix()
	record := &model.RolePermission{
		PermissionId: req.Data.PermissionId,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	err := mysql.DB.Model(&model.RolePermission{}).Create(&record).Error
	if err != nil {
		req.Logger.Error("create record failed", zap.Any("Data", record), zap.Error(err))
		req.Reply.CreateFailed()
		return
	}
}
