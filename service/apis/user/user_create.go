package user

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"prj-aiee/model"
	"prj-aiee/service/apis/common"
	"prj-aiee/service/mysql"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CreateUserRole struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Age      int64  `json:"age"`
	Username string `json:"username"`
	RoleIds  []int  `json:"role_ids"`
}

type CreateUserApi struct {
	*common.ApiCommon
	Data CreateUserRole
}

type UserLogin struct {
	*common.ApiCommon
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserInfo struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Age      int64  `json:"age"`
	Status   int64  `json:"status"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type UserLoginResponse struct {
	*common.ApiCommon
	Data UserInfo `json:"data"`
}

// @Summary 用户注册
// @Description 用户注册
// @ID CreateUser
// @Accept json
// @Produce json
// @Param RequestBody body CreateUserRole true "新用户的详细信息"
// @Success 201 {object} common.CommonResponse "成功创建的用户信息"
// @Failure 400 {string} string "请求参数有误"
// @Router /v1/user/create [post]
// @Tags 用户管理
func CreateUser(c *gin.Context) {
	req := &CreateUserApi{
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

func (req *CreateUserApi) checkParams() {
	fmt.Println("req.Data:", req.Data)
}

func (req *CreateUserApi) createRecord() {
	now := time.Now().Unix()
	hasher := md5.New()
	hasher.Write([]byte(req.Data.Password))
	Password := hex.EncodeToString(hasher.Sum(nil))
	record := &model.User{
		Username:  req.Data.Username,
		Password:  Password,
		Name:      req.Data.Name,
		Status:    0,
		Age:       req.Data.Age,
		CreatedAt: now,
		UpdatedAt: now,
	}
	err := mysql.DB.Model(&model.User{}).Create(&record).Error
	if err != nil {
		req.Logger.Error("create record failed", zap.Any("Data", record), zap.Error(err))
		req.Reply.CreateFailed()
		return
	}
	roleRecords := make([]*model.Role, 0)
	errRole := mysql.DB.Model(&model.Role{}).Where("id in (?)", req.Data.RoleIds).Find(&roleRecords).Error
	if errRole != nil {
		req.Logger.Error("get records failed", zap.Error(errRole))
		req.Reply.ReadFailed()
		return
	}
	var userRoles []*model.UserRole
	for _, roleId := range req.Data.RoleIds {
		userRoles = append(userRoles, &model.UserRole{
			RoleId:    roleId,
			UserId:    record.Id,
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

// 在数组中通过RoleId查找对应的RoleName
func GetRoleNameByRoleId(roles []*model.Role, roleId int) string {
	for _, role := range roles {
		if role.Id == roleId {
			return role.Name
		}
	}
	return ""
}

// @Summary 用户登陆
// @Description 用户登陆
// @ID UserLogin
// @Accept json
// @Produce json
// @Param RequestBody body UserLogin true "登陆用户信息信息"
// @Success 201 {object} UserLoginResponse "用户信息"
// @Failure 400 {string} string "请求参数有误"
// @Router /v1/user/login [post]
// @Tags 用户登陆
func Login(c *gin.Context) {
	req := &UserLogin{
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
	record := req.SignIn()
	if req.Reply.IsStatusFailed() {
		req.Reply.Response(c)
		return
	}
	Token, err := GenerateToken(record.Id, record.Name)
	if err != nil {
		req.Logger.Error("create token failed", zap.Error(err))
		req.Reply.CreateFailed()
		return
	}
	userInfo := UserInfo{
		Id:       record.Id,
		Name:     record.Name,
		Age:      record.Age,
		Status:   record.Status,
		Username: record.Username,
		Token:    Token,
	}

	res := UserLoginResponse{
		Data: userInfo,
	}

	req.Reply.DataSet(res)
	req.Reply.Response(c)
}

func (req *UserLogin) checkParams() {
	fmt.Println("req.Data:", req)
}

func (req *UserLogin) SignIn() (record model.User) {
	hasher := md5.New()
	hasher.Write([]byte(req.Password))
	Password := hex.EncodeToString(hasher.Sum(nil))
	err := mysql.DB.Model(&model.User{}).Where("name = ? and password = ?", req.Name, Password).First(&record).Error
	if err != nil {
		req.Logger.Error("create record failed", zap.Any("Data", record), zap.Error(err))
		req.Reply.ReadFailed()
		return
	}
	return record
}

func GenerateToken(Id int, Name string) (string, error) {
	// 创建一个 JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	// 设置 token 的声明
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = Id
	claims["name"] = Name
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 设置 token 过期时间

	// 生成 token 字符串
	tokenString, err := token.SignedString([]byte(common.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
