package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetStudentsHandler(t *testing.T) {
	assert := assert.New(t)
	r := gin.Default()
	SetupHandlers(r)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/students", nil)
	var list []Student
	r.ServeHTTP(res, req)

	err := json.NewDecoder(res.Body).Decode(&list) // 결과 반환
	assert.Nil(err)
	assert.Len(list, 2)
	assert.Equal("aaa", list[0].Name)
}
