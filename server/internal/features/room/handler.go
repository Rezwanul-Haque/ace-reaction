package room

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type RoomHandler struct {
	service *RoomService
}

func NewRoomHandler(s *RoomService) *RoomHandler {
	return &RoomHandler{service: s}
}

func (h *RoomHandler) RegisterRoutes(g *echo.Group) {
	g.POST("/rooms", h.CreateRoom)
	g.GET("/rooms/:id", h.GetRoom)
}

type CreateRoomResponse struct {
	RoomID string `json:"room_id"`
}

func (h *RoomHandler) CreateRoom(c echo.Context) error {
	room, err := h.service.CreateRoom()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, CreateRoomResponse{RoomID: room.ID})
}

type GetRoomResponse struct {
	RoomID  string `json:"room_id"`
	Players int    `json:"players"`
	Status  string `json:"status"`
}

func (h *RoomHandler) GetRoom(c echo.Context) error {
	id := c.Param("id")
	room, err := h.service.GetRoom(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "room not found"})
	}
	return c.JSON(http.StatusOK, GetRoomResponse{
		RoomID:  room.ID,
		Players: len(room.Players),
		Status:  string(room.Status),
	})
}
