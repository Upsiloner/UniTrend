package user_ctl

import (
	"net/http"

	"github.com/Upsiloner/UniTrend/domain"
	"github.com/Upsiloner/UniTrend/domain/user_domain"

	"github.com/gin-gonic/gin"
)

type GetUserController struct {
	GetUserUsecase user_domain.GetUserUsecase
}

func (lc *GetUserController) GetUser(c *gin.Context) {
	var request user_domain.GetUserRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.NewDefaultResponse(400, "参数错误"))
		return
	}

	user, err := lc.GetUserUsecase.GetUserByUnionID(c, request.Union_ID)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.NewDefaultResponse(400, "未查询到当前union_id对应的用户"))
		return
	}

	GetUserResponse := user_domain.GetUserResponse{
		DefaultResponse: domain.NewDefaultResponse(200),
		User:            user,
	}

	c.JSON(http.StatusOK, GetUserResponse)
}
