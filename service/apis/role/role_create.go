package role

import (
	"fmt"
	"time"

	"prj-aiee/model"
	"prj-aiee/service/apis/common"
	"prj-aiee/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RoleCreateRequest struct {
	Name        string `json:"name"`
	Desc        string `json:"desc"`
	Key         string `json:"key"`
	Status      int    `json:"status"`
	Permissions []int  `json:"permissions"`
}

type CreateRoleApi struct {
	*common.ApiCommon
	Data RoleCreateRequest
}

// @Summary 创建角色并关联权限
// @Description 创建角色并关联权限
// @ID CreateRole
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token" default(token)
// @Param RequestBody body RoleCreateRequest true "新角色的详细信息"
// @Success 201 {object} common.CommonResponse "成功创建的用户信息"
// @Failure 400 {string} string "请求参数有误"
// @Router /v1/role/create [post]
// @Tags 角色管理
func CreateRole(c *gin.Context) {
	req := &CreateRoleApi{
		ApiCommon: common.NewRequest(c),
	}
	if err := req.BindRequest(&req.Data); err != nil {
		req.Reply.BindRequestFailed().Response(c)
		return
	}
	fmt.Println("req", req)
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

func (req *CreateRoleApi) checkParams() {
}

func (req *CreateRoleApi) createRecord() {
	now := time.Now().Unix()
	record := &model.Role{
		Name:      req.Data.Name,
		Desc:      req.Data.Desc,
		Status:    common.RecordStatusInit,
		CreatedAt: now,
		UpdatedAt: now,
	}
	// 事务
	err := mysql.DB.Model(&model.Role{}).Create(&record).Error
	if err != nil {
		req.Logger.Error("create record failed", zap.Any("Data", record), zap.Error(err))
		req.Reply.CreateFailed()
		return
	}
	var rolePermissions []*model.RolePermission
	for _, permissionID := range req.Data.Permissions {
		rolePermission := &model.RolePermission{
			RoleId:       record.Id,
			PermissionId: permissionID,
		}
		rolePermissions = append(rolePermissions, rolePermission)
	}
	errRole := mysql.DB.Model(&model.RolePermission{}).Create(&rolePermissions).Error
	if errRole != nil {
		req.Logger.Error("create record failed", zap.Any("Data", record), zap.Error(err))
		req.Reply.CreateFailed()
		return
	}
}
