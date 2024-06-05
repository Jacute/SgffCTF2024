package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	content, err := os.ReadFile("templates/index.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error reading file: %v", err)
		return
	}
	c.Header("Flag", "SgffCTF{why_c4nt_w3_n0t_b3_s0b3r}")
	c.Data(http.StatusOK, "text/html; charset=utf-8", content)
}

func GetSecret1(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte("SgffCTF{1_kn0w_"))
}

func GetSecret2(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte("th3_p13c35_f1t}"))
}
