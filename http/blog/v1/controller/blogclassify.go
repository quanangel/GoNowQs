package controller

import (
	"nowqs/frame/http/blog/models"

	"github.com/gin-gonic/gin"
)

type BlogClassify struct{}

func NewBlogClassify() BlogClassify {
	return BlogClassify{}
}

// blogClassifyValidate is get validate struct
type blogClassifyValidate struct {
	// Type: my/list/only
	Type string `form:"type" json:"type" xml:"type" binding:"required,oneof=my list only"`
	// Search: type is only the search is id, type is list the search is id/name/url
	Search string `form:"search" json:"search" xml:"search" binding:"required_if=Type only"`
	// Classify
	Classify int64 `form"classify" json:"classify" xml:"classify" binding:"-"`
	// Page
	Page int `form:"page" json:"page" xml:"page" binding:"-"`
	// Limit
	Limit int `form:"limit" json:"limit" xml:"limit" binding:"-"`
	// Order
	Order string `form:"order" json:"order" xml:"order" binding:"-"`
}

func (a *BlogClassify) Get(c *gin.Context) {
	returnData := gin.H{
		"code": -1,
	}
	authToken := c.GetHeader("Auth-Token")
	userID := userAuth(authToken)

	var validate blogClassifyValidate
	if err := c.Bind(&validate); err != nil {
		returnData["code"] = 10000
		returnData["msg"] = err.Error()
		jsonHandle(c, returnData)
		return
	}
	modelBlog := models.NewBlog()
	search := make(map[string]interface{})
	if validate.Page == 0 {
		validate.Page = 1
	}
	if validate.Limit == 0 {
		validate.Limit = 20
	}
	if validate.Classify != 0 {
		search["classify_id"] = validate.Classify
	}
	switch validate.Type {
	case "my":
		if userID == 0 {
			returnData["code"] = 2
			jsonHandle(c, returnData)
			return
		}
		if validate.Search != "" {
			search["id"] = validate.Search
			search["name"] = validate.Search
		}
		search["user_id"] = 
		total, result := modelBlog.GetList(search, validate.Page, validate.Limit, "")

	case "list":

	default:
		returnData["code"] = 1
		jsonHandle(c, returnData)
		return
	}
	jsonHandle(c, returnData)
}
