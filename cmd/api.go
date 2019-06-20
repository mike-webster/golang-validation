package wyzvalidator

import "github.com/mike-webster/golang-validation/controllers"

func main() {
	r := controllers.GetRouter()
	r.Run("0.0.0.0:3001")
}
