package controllers

import (
    "myapp/models"
    "myapp/services"
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetUsers(c *gin.Context) {
    users, err := services.GetAllUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error en la conexion": err.Error()})
        return
    }
    if err := services.CreateUser(user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"Error en la base de datos": err.Error()})
        return
    }
    c.JSON(http.StatusOK, user)
}
