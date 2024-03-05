package user

import (
	"prj-aiee/model"
	"prj-aiee/service/apis/common"
	"prj-aiee/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetUserApi struct {
	*common.ApiCommon
	Data GetUserRequest
}
type GetUserRequest struct {
	Id int `json:"id"`
}
type GetUserResponse struct {
	Data        *model.User         `json:"data"`
	Permissions []*model.Permission `json:"permissions"`
}

type GetUserInfo struct {
	Data        *model.User         `json:"data"`
	Permissions []*model.Permission `json:"permissions"`
}

// @Summary 获取用户
// @Description 获取用户和用户权限
// @ID getUser
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token" default(token)
// @Param RequestBody body GetUserRequest true "用户信息"
// @Success 201 {object} common.CommonResponse "用户信息"
// @Failure 400 {string} string "请求参数有误"
// @Router /v1/user/get [post]
// @Tags 用户管理
func GetUser(c *gin.Context) {
	req := &GetUserApi{
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
	req.Reply.DataSet(record)
	req.Reply.Response(c)
}

func (req *GetUserApi) checkParams() {
	if req.Data.Id <= 0 {
		req.Reply.MsgSet(common.ReplyStatusMissingParam, common.ReplyMessageMissingParam)
		return
	}
}

func (req *GetUserApi) getRecord() (record GetUserInfo) {
	err := mysql.DB.Model(&model.User{}).
		Where("status != ?", common.RecordStatusDeleted).
		Where("id = ?", req.Data.Id).
		First(&record.Data).Error
	if err != nil {
		req.Logger.Error("get record failed", zap.Error(err))
		req.Reply.ReadFailed()
		return
	}
	// 根据UserId，关联查询UserId关联的RoleId列表，再通过RoleId列表，关联查询PermissionId列表
	var roleIds []int
	err1 := mysql.DB.Model(&model.UserRole{}).Where("user_id = ?", record.Data.Id).Pluck("role_id", &roleIds).Error
	if err1 != nil {
		req.Logger.Error("get record failed", zap.Error(err))
		req.Reply.ReadFailed()
		return
	}

	var permissionIds []int
	err2 := mysql.DB.Model(&model.RolePermission{}).Where("role_id in (?)", roleIds).Pluck("permission_id", &permissionIds).Error
	if err2 != nil {
		req.Logger.Error("get record failed", zap.Error(err))
		req.Reply.ReadFailed()
		return
	}
	err3 := mysql.DB.Model(&model.Permission{}).Where("id in (?)", permissionIds).
		Find(&record.Permissions).Error
	if err3 != nil {
		req.Logger.Error("get record failed", zap.Error(err))
		req.Reply.ReadFailed()
		return
	}
	return record
}
