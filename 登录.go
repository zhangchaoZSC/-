package main






import (

"fmt"

"github.com/gin-gonic/gin"

"github.com/gin-gonic/gin/binding"

"log"

"net/http"

)

func main() {

	router := gin.Default()

	router.GET("/auth/signin", func(c *gin.Context) {

		cookie := &http.Cookie{

			Name:     "session_id",

			Value:    "123",

			Path:     "/",

			HttpOnly: true,

		}

		http.SetCookie(c.Writer, cookie)

		c.String(http.StatusOK, "Login successful")

	})



	router.GET("/home", AuthMiddleWare(), func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{"data": "home"})

	})

	router.Run(":8080")

}



func AuthMiddleWare() gin.HandlerFunc {

	return func(c *gin.Context) {

		if cookie, err := c.Request.Cookie("session_id"); err == nil {

			value := cookie.Value

			fmt.Println(value)

			if value == "123" {

				c.Next()

				return

			}

		}

		c.JSON(http.StatusUnauthorized, gin.H{

			"error": "Unauthorized",

		})

		c.Abort()

		return

	}

}









