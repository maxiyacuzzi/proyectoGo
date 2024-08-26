package services

import (
    "myapp/config"
    "myapp/models"
    "log"
)

func GetAllUsers() ([]models.User, error) {
    rows, err := config.DB.Query("SELECT id, name, email FROM webapp.users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        err := rows.Scan(&user.ID, &user.Name, &user.Email)
        if err != nil {
            log.Println(err)
            continue
        }
        users = append(users, user)
    }
    return users, nil
}

func CreateUser(user models.User) error {
    _, err := config.DB.Exec("INSERT INTO webapp.users (name, email) VALUES ($1, $2)", user.Name, user.Email)
    return err
}
