package handler

import (
	"net/http"
	"strconv"

	"github.com/Doer-org/hack-camp_vol9_2022-1/presentation/json"
	"github.com/Doer-org/hack-camp_vol9_2022-1/usecase"
	"github.com/gin-gonic/gin"
)

type MemberHandler struct {
	uc usecase.IMemberUsecase
}

func NewMemberHandler(uc usecase.IMemberUsecase) *MemberHandler {
	return &MemberHandler{
		uc: uc,
	}
}

func (rh *MemberHandler) CreateMember(ctx *gin.Context) {
	var memberjson json.MemberJson
	if err := ctx.BindJSON(&memberjson); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	member := json.MemberJsonToEntity(&memberjson)
	member, err := rh.uc.CreateMember(member.Name, member.RoomId)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	memberJson := json.MemberEntityToJson(member)
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": memberJson},
	)

}

func (rh *MemberHandler) GetAllMembersOfRoomID(ctx *gin.Context) {
	roomId := ctx.Param("roomId")
	member, err := rh.uc.GetAllMembersOfRoomID(roomId)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	memberJson := json.MembersEntityToJson(member)
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": memberJson},
	)

}

func (rh *MemberHandler) DeleteAllMembersOfRoomID(ctx *gin.Context) {
	roomId := ctx.Param("roomId")
	err := rh.uc.DeleteAllMembersOfRoomID(roomId)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{"data": "success"},
	)

}

func (rh *MemberHandler) GetMemberOfId(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	member, err := rh.uc.GetMemberOfId(id)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	memberJson := json.MemberEntityToJson(member)
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": memberJson},
	)
}
