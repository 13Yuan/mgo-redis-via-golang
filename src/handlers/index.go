package handlers

import (
	"github.com/gin-gonic/gin"
	services "MA.Content.Services.OrgMapper/src/services"
	. "MA.Content.Services.OrgMapper/src/models"
	"fmt"
)

func loopGetIndentifiers(ch chan Identifier, m_id string) {
	types := []string{
		"mdy_id",
		"lei",
		"seldo",
	}
	for i := 0; i < len(types); i++ {
		go services.GetIdentifiers(ch, m_id, types[i])
	}
}


func GetIdentifiersViaId(c *gin.Context) {
	m_id := c.Param("id")
	ch := make(chan Identifier, 3)
	loopGetIndentifiers(ch, m_id)
	for i := range ch {
		fmt.Println(i)
	}
	close(ch)
}

func GetIdentifiers(c *gin.Context) { 
	m_id := c.Param("id")
	m_type := c.Param("type")
	ch := make(chan Identifier)
	go services.GetIdentifiers(ch, m_id, m_type)
	c.JSON(200, gin.H{
		"message": <- ch,
	})
}