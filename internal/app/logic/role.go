package logic

import (
	common "github.com/CodeLine-95/go-cloud-native/common/models"
	"github.com/CodeLine-95/go-cloud-native/initial/store/db"
	"github.com/CodeLine-95/go-cloud-native/internal/app/constant"
	"github.com/CodeLine-95/go-cloud-native/internal/app/models"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/base"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/response"
	"github.com/CodeLine-95/go-cloud-native/tools/structs"
	"github.com/gin-gonic/gin"
	"time"
)

func RoleResp(c *gin.Context) {
	var params common.SearchRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}

	var role models.CloudRole
	selectFields := structs.ToTags(role, "json")

	var roleResp []*models.CloudRole
	err := db.D().Select(selectFields).
		Where("position(concat(?) in concat(role_key,role_name)) > 0", params.SearchKey).
		Scopes(base.Paginate(params.Page, params.PageSize)).
		Find(&roleResp).Error
	if err != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}

	response.PageOK(c, roleResp, len(roleResp), params.Page, params.PageSize, constant.ErrorMsg[constant.Success])
}

// RoleAdd 添加角色
func RoleAdd(c *gin.Context) {
	var params common.RoleRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}

	// 验证roleKey标识，唯一
	roleResp := models.CloudRole{}
	err := db.D().Where("role_key = ?", params.Key).Find(&roleResp).Error
	if err != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}
	if roleResp.RoleKey == params.Key {
		response.Error(c, constant.ErrorDBRecordExist, nil, constant.ErrorMsg[constant.ErrorDBRecordExist])
		return
	}

	auth, err := base.GetAuth(c)
	if err != nil {
		response.Error(c, constant.ErrorNotLogin, err, constant.ErrorMsg[constant.ErrorNotLogin])
		return
	}

	cloudRole := &models.CloudRole{
		RoleName:   params.Name,
		RoleRemark: params.Remark,
		RoleKey:    params.Key,
		RoleSort:   params.Sort,
		Status:     params.Status,
		ControlBy: common.ControlBy{
			CreateBy: uint32(auth.UID),
		},
		ModelTime: common.ModelTime{
			CreateTime: uint32(time.Now().Unix()),
		},
	}
	res := db.D().Create(cloudRole)
	if res.RowsAffected == 0 || res.Error != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}

	response.OK(c, nil, constant.ErrorMsg[constant.Success])
}

// RoleEdit 编辑角色
func RoleEdit(c *gin.Context) {
	var params common.RoleRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}

	auth, err := base.GetAuth(c)
	if err != nil {
		response.Error(c, constant.ErrorNotLogin, err, constant.ErrorMsg[constant.ErrorNotLogin])
		return
	}

	cloudRole := &models.CloudRole{
		RoleId:     uint32(params.Id),
		RoleName:   params.Name,
		RoleRemark: params.Remark,
		RoleKey:    params.Key,
		RoleSort:   params.Sort,
		Status:     params.Status,
		ControlBy: common.ControlBy{
			UpdateBy: uint32(auth.UID),
		},
		ModelTime: common.ModelTime{
			UpdateTime: uint32(time.Now().Unix()),
		},
	}
	res := db.D().Updates(cloudRole)
	if res.RowsAffected == 0 || res.Error != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}

	response.OK(c, nil, constant.ErrorMsg[constant.Success])
}

// RoleDel 删除角色
func RoleDel(c *gin.Context) {
	var params common.RoleRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}

	err := db.D().Delete(params).Error
	if err != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}

	response.OK(c, nil, constant.ErrorMsg[constant.Success])
}