package controllers

import (
	"awesomeProject/database"
	"awesomeProject/middlewares"
	"awesomeProject/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func AllRoles(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "roles"); err != nil {
		return err
	}

	var roles []models.Role

	database.DB.Find(&roles)

	return c.JSON(roles)
}

//type RoleCreateDTO struct {
//	name        string
//	permissions []string
//}

func CreateRole(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "roles"); err != nil {
		return err
	}

	//var role models.Role
	//var roleDTO RoleCreateDTO
	var roleDTO fiber.Map

	if err := c.BodyParser(&roleDTO); err != nil {
		return err
	}

	//var permissions []models.Permission
	list := roleDTO["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))

		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	role := models.Role{
		Name:        roleDTO["name"].(string),
		Permissions: permissions,
	}
	database.DB.Create(&role)

	return c.JSON(role)
}

func GetRole(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "roles"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Preload("Permissions").Find(&role)

	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "roles"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))
	//role := models.Role{
	//	Id: uint(id),
	//}
	var roleDTO fiber.Map
	if err := c.BodyParser(&roleDTO); err != nil {
		return err
	}

	list := roleDTO["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))

		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}
	var result interface{}

	database.DB.Table("role_permissions").Where("role_id= ?", id).Delete(result)
	//if err := database.DB.Table("role_permissions").Where(deletionConditions).Delete(&models.Permission{}).Error; err != nil {
	//	return err
	//}

	role := models.Role{
		Id:          uint(id),
		Name:        roleDTO["name"].(string),
		Permissions: permissions,
	}

	database.DB.Model(&role).Updates(role)

	return c.JSON(role)
}

func DeleteRole(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "roles"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))
	role := models.Role{
		Id: uint(id),
	}
	database.DB.Delete(&role)

	return nil
}
