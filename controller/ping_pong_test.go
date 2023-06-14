package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type PingPongControllerTestSuite struct {
	suite.Suite
	ctx                *gin.Context
	recorder           *httptest.ResponseRecorder
	pingPongController PingPongController
}

func TestPingPongControllerTestSuite(t *testing.T) {
	suite.Run(t, new(PingPongControllerTestSuite))
}

func (suite *PingPongControllerTestSuite) SetupTest() {
	suite.recorder = httptest.NewRecorder()
	suite.ctx, _ = gin.CreateTestContext(suite.recorder)
	suite.pingPongController = NewPingPongController()
}

func (suite *PingPongControllerTestSuite) TearDownTest() {
}

func (suite *PingPongControllerTestSuite) Test_Ping_ShouldReturnStatusOkWithMessageInResponse_WhenMessageIsPresentInRequest() {
	response := `{"message":"pong"}`
	suite.ctx.Request = httptest.NewRequest("GET", "/ping", nil)

	suite.pingPongController.Ping(suite.ctx)

	suite.Equal(http.StatusOK, suite.recorder.Code)
	suite.Equal([]byte(response), suite.recorder.Body.Bytes())
}
