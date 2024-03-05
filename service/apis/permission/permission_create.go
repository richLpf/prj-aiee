package permission

import (
	"fmt"
	"time"

	"prj-aiee/model"
	"prj-aiee/service/apis/common"
	"prj-aiee/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CreatePermissionApi struct {
	*common.ApiCommon
	Data *model.Permission
}

// @Summary 创建权限
// @Description 创建用户权限
// @ID CreatePermission
// @Accept json
// @Produce json
// @Param RequestBody body model.Permission true "新用户的详细信息"
// @Success 201 {object} common.CommonResponse "成功创建的用户信息"
// @Failure 400 {string} string "请求参数有误"
// @Router /v1/permission/create [post]
// @Tags 权限管理
func CreatePermission(c *gin.Context) {
	req := &CreatePermissionApi{
		ApiCommon: common.NewRequest(c),
	}
	fmt.Println("req", req)
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

func (req *CreatePermissionApi) checkParams() {
}

func (req *CreatePermissionApi) createRecord() {
	now := time.Now().Unix()
	record := &model.Permission{
		Key:       req.Data.Key,
		Type:      req.Data.Type,
		Desc:      req.Data.Desc,
		Attribute: req.Data.Attribute,
		CreatedAt: now,
		UpdatedAt: now,
	}
	err := mysql.DB.Model(&model.Permission{}).Create(&record).Error
	if err != nil {
		req.Logger.Error("create record failed", zap.Any("Data", record), zap.Error(err))
		req.Reply.CreateFailed()
		return
	}
}
