package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDumpWithBody(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	s := "1234"
	r := fmt.Sprintf("{\"body\":\"%s\"}", s)

	req, _ := http.NewRequest("GET", "/dump", strings.NewReader(s))

	c.Request = req

	requestDump(c)

	assert.Equal(t, 200, w.Code)

	body, _ := io.ReadAll(w.Result().Body)
	assert.Equal(t, r, string(body))
}

func TestDump(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	s := ""
	r := fmt.Sprintf("{\"body\":\"%s\"}", s)

	req, _ := http.NewRequest("GET", "/dump", nil)

	c.Request = req

	requestDump(c)

	assert.Equal(t, 200, w.Code)

	body, _ := io.ReadAll(w.Result().Body)
	assert.Equal(t, r, string(body))
}
