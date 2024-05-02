package main

import "github.com/gin-gonic/gin"

var ninjaWeapon = map[string]string{
	"ninjaStar": "Beginner - Damage 1",
}

func GetWeapon(c *gin.Context) {
	weaponType := c.Query("type")
	weaponName, ol := ninjaWeapon[weaponType]
	if !ol {
		c.JSON(404, gin.H{
			weaponType: "Not Found",
		})
		return
	}
	c.JSON(200, gin.H{
		weaponType: weaponName,
	})
}

func PostWeapon(c *gin.Context) {
	weaponType := c.Query("type")
	weaponName := c.Query("name")

	if len(weaponType) == 0 || len(weaponName) == 0 {
		c.JSON(404, gin.H{
			weaponType: weaponName,
		})
		return
	}

	if _, ok := ninjaWeapon[weaponType]; ok {
	    c.JSON(409, gin.H{
			"message": "Weapon already exists",
		})
		return
	}

	ninjaWeapon[weaponType] = weaponName
	c.JSON(201, gin.H{
	    weaponType: weaponName,
	})
}

func DeleteWeapon(c *gin.Context) {
	weaponType := c.Query("type")
	weaponName, ok := ninjaWeapon[weaponType]
	if !ok {
	    c.JSON(404, gin.H{
	        weaponType: "",
	    })
		return
	}
    delete(ninjaWeapon, weaponType)
	c.JSON(200, gin.H{
        weaponType: weaponName,
    })
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	rGroup := r.Group("/weapon")
	rGroup.GET("", GetWeapon)
	rGroup.POST("", PostWeapon)
	rGroup.DELETE("", DeleteWeapon)

	r.Run() // listen and serve on 0.0.0.0:8080
}
