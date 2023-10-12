`https://github.com/gofiber/fiber`

`go get github.com/gofiber/fiber/v2`

`export GO111MODULE="on"`

`go mod init go/admin`

`go mod tidy`

1. GORM
   `go get -u gorm.io/gorm`
   `go get -u gorm.io/driver/sqlite`
   `go get -u gorm.io/driver/postgres`

   GOPATH='/home/alivad/go'

2. Register
3. Add github.com/oxequa/realize/realize

`go get -u .`

4. Add migrations
5. Hash password
6. Roles
7. Permissions

   Add to db.permissions

| id | name          |
|----|---------------|
| 1  | view_users    |
| 2  | edit_users    |
| 3  | view_roles    |
| 4  | edit_roles    |
| 5  | view_products |
| 6  | edit_products |
| 7  | view_orders   |
| 8  | edit_orders   |

many to many to role

with gorm add table to do it `role_permissions`
