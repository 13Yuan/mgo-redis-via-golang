package handlers

import (
	db "om-api/src/db"
	m "om-api/src/models"
	"net/http"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strings"
)

func isFuzzySearch(_type string) bool {
	return _type == "all"
}

func generateKey(_type, id string) string {
	return strings.Replace(_type, " ", "", -1) + "-" + strings.Replace(id, " ", "", -1)
}

func GetIdentifiers(c *gin.Context) {
	var (
		identifers m.IdentifiersObj
		err error
		result []string
	)
	m_id := c.Param("id")
	if m_type := c.Param("type"); !isFuzzySearch(m_type) {
		result, err = db.Get(generateKey(m_type, m_id))
		identifers.Count = len(result)
	} else {
		result, err = db.Get(m_id)
		identifers.Count = len(result)
	}
	if err != nil {
		c.JSON(http.StatusNoContent, "Get data error!")
	} else {
		resultStr := "[" + strings.Join(result, ",") + "]"
		json.Unmarshal([]byte(resultStr), &identifers.Identifiers)
		c.JSON(http.StatusOK, identifers)
	}
}