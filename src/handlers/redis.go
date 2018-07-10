package handlers

import (
	"fmt"
	db "MA.Content.Services.OrgMapper/src/db"
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
	var result string
	m_id := c.Param("id")
	if m_type := c.Param("type"); !isFuzzySearch(m_type) {
		results, err := db.Get(generateKey(m_type, m_id))
		if err != nil {
			result = "Get data error!"
			fmt.Println(result)
		} else {
			result = strings.Join(results, ",")
		}
	} else {
		results, err := db.Get(m_id)
		if err != nil {
			result = "Get data error!"
			fmt.Println(result)
		} else {
			result = strings.Join(results, ",")
		}
	}
	c.JSON(200, result)
}