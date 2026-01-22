package controllers

import (
	"golang_mvc_starterpack/dto"
	"golang_mvc_starterpack/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PeopleHandler struct {
	service services.PeopleService
}

func NewPeopleHandler(s services.PeopleService) *PeopleHandler {
	return &PeopleHandler{s}
}

func (h *PeopleHandler) IndexPeople(c *gin.Context) {
	people, err := h.service.Index()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, people)
}

func (h *PeopleHandler) ShowPeople(c *gin.Context) {
	id := c.Param("id")
	people, err := h.service.Show(id)
	if err != nil {
		log.Println(err)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, people)
}

func (h *PeopleHandler) StorePeople(c *gin.Context) {
	var people dto.PeopleCreate

	if err := c.ShouldBindJSON(&people); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// call layer service people
	if err := h.service.Store(&people); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"status":  "success",
		"message": "Data berhasil disimpan",
		"data":    people,
	})

}

func (h *PeopleHandler) UpdatePeople(c *gin.Context) {
	idParam := c.Param("id")
	idUint, _ := strconv.ParseUint(idParam, 10, 32)

	var input dto.PeoplePatch
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Update(uint(idUint), input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil diupdate"})
}

func (h *PeopleHandler) DeletePeople(c *gin.Context) {
	idParam := c.Param("id")
	if err := h.service.Delete(idParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "data berhasil dihapus"})
}
