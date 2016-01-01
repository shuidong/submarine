package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

// Server represents a battle server.
type Server struct {
	*gin.Engine
	logWriter   *io.PipeWriter
	roomManager *RoomManager
}

// NewServer creates a Server.
func NewServer() *Server {
	server := &Server{gin.New(), Log.Writer(), newRoomManager()}
	server.Use(gin.Recovery(), gin.LoggerWithWriter(server.logWriter))

	server.GET("/rooms/:id", func(c *gin.Context) {
		roomID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			Log.Error(err)
			c.String(http.StatusBadRequest, "Invalid room id.")
			return
		}

		session := newSession(1, roomID)
		if err := session.Connect(c.Writer, c.Request); err != nil {
			Log.Error(err)
			c.String(http.StatusBadRequest, "Failed to upgrade the request to web socket protocol.")
			return
		}

		server.roomManager.JoinToRoom <- session
		Log.Infof("Session(%v) is created", session.id)
	})

	return server
}

// Close the Server.
func (s *Server) Close() {
	s.logWriter.Close()
	s.roomManager.Close <- struct{}{}
}
