package user

import (
	"time"

	"prj-aiee/model"
	"prj-aiee/service/apis/common"
	"prj-aiee/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UpdateUserApi struct {
	*common.ApiCommon
	UpdateUserRequest
}
type UpdateUserRequest struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Username string `json:"username"`
	RoleIds  []int  `json:"role_ids"`
}

// @Summary 更新用户
// @Description 更新用户
// @ID UpdateUser
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token" default(token)
// @Param RequestBody body UpdateUserApi true "用户信息"
// @Success 201 {object} common.CommonResponse "用户信息"
// @Failure 400 {string} string "请求参数有误"
// @Router /v1/user/update [post]
// @Tags 用户管理
func UpdateUser(c *gin.Context) {
	req := &UpdateUserApi{
		ApiCommon: common.NewRequest(c),
	}
	if err := req.BindRequest(&req); err != nil {
		req.Reply.BindRequestFailed().Response(c)
		return
	}
	req.checkParams()
	if req.Reply.IsStatusFailed() {
		req.Reply.Response(c)
		return
	}
	req.updateRecord()
	if req.Reply.IsStatusFailed() {
		req.Reply.Response(c)
		return
	}

	req.Reply.Response(c)
}

func (req *UpdateUserApi) checkParams() {
	if req.Id <= 0 {
		req.Reply.MsgSet(common.ReplyStatusMissingParam, common.ReplyMessageMissingParam)
		return
	}
	var count int64
	err := mysql.DB.Model(&model.User{}).
		Where("status != ?", common.RecordStatusDeleted).
		Where("id = ?", req.Id).
		Count(&count).Error
	if err != nil {
		req.Logger.Error("check record failed", zap.Error(err))
		req.Reply.ReadFailed()
		return
	}
	if count <= 0 {
		req.Reply.MsgSet(common.ReplyStatusBindRequestFailed, common.ReplyMessageBindRequestFailed)
		return
	}
}

func (req *UpdateUserApi) updateRecord() {
	now := time.Now().Unix()
	record := &model.User{
		Username:  req.Username,
		Password:  req.Password,
		Name:      req.Name,
		UpdatedAt: now,
	}
	tx := mysql.DB
	err := tx.Model(&model.User{}).Where("id = ?", req.Id).Updates(&record).Error
	if err != nil {
		req.Logger.Error("update record failed", zap.Any("Data", record), zap.Error(err))
		req.Reply.UpdateFailed()
		return
	}
	err = tx.Where("user_id = ?", req.Id).Delete(&model.UserRole{}).Error
	if err != nil {
		req.Logger.Error("delete record failed", zap.Error(err))
		req.Reply.DeleteFailed()
		return
	}
	roleRecords := make([]*model.Role, 0)
	errRole := tx.Model(&model.Role{}).Where("id in (?)", req.RoleIds).Find(&roleRecords).Error
	if errRole != nil {
		req.Logger.Error("get records failed", zap.Error(errRole))
		req.Reply.ReadFailed()
		return
	}
	var userRoles []*model.UserRole
	for _, roleId := range req.RoleIds {
		userRoles = append(userRoles, &model.UserRole{
			RoleId:    roleId,
			UserId:    req.Id,
			RoleName:  GetRoleNameByRoleId(roleRecords, roleId),
			CreatedAt: now,
			UpdatedAt: now,
		})
	}
	errUser := mysql.DB.Model(&model.UserRole{}).Create(&userRoles).Error
	if errUser != nil {
		req.Logger.Error("create record failed", zap.Any("Data", userRoles), zap.Error(err))
		req.Reply.CreateFailed()
		return
	}
}
