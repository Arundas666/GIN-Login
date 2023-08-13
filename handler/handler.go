package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userData = make(map[string]User)
type PageData struct {
	EmailInvalid string
	PassInvalid  string
}
type User struct {
	Name     string
	Email    string
	Password string
}
func IndexPage(c *gin.Context) {
	c.HTML(http.StatusFound, "signup.html", nil)
}
func Signup(c *gin.Context) {
	c.HTML(http.StatusFound, "signup.html", nil)
}
func SignupPost(c *gin.Context) {
	name := c.Request.FormValue("name")
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")

	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")

	if email == "" {
		c.HTML(http.StatusBadGateway, "signup.html", "EmailInvalid")

		return
	}
	if password == "" {
		c.HTML(http.StatusBadGateway, "signup.html", "Password Invalid")

		return
	}
	userData[email] = User{Name: name,
		Password: password,
		Email:    email,
	}

	c.Redirect(http.StatusSeeOther, "/login")

	fmt.Printf("%+v", userData)

}
func Login(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	cookie, err := c.Cookie("logincookie")
	if err == nil && cookie != "" {
		c.Redirect(http.StatusSeeOther, "/home")

		return
	}
	c.HTML(200, "login.html", nil)

}
func LoginPost(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")

	cookie, err := c.Cookie("logincookie")
	if err != nil {
		fmt.Println(err)

	} else if cookie != "" {

		c.Redirect(303, "/loginpost")
		return
	}

	email := c.Request.FormValue("emailName")
	password := c.Request.FormValue("passwordName")

	user, ok := userData[email]

	if email == "" {
		var n = PageData{EmailInvalid: "Email is Invalid"}
		c.HTML(200, "login.html", n)

		fmt.Println("EmailEmpty")
		return
	} else if password == "" {
		var n = PageData{PassInvalid: "Password is Invalid"}
		c.HTML(200, "login.html", n)

		fmt.Println("PasswordEmpty")
		return
	}
	if ok && password == user.Password {

		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("logincookie", "123", 300, "/", "", false, true)

		cookie, _ := c.Cookie("logincookie")
		fmt.Println(cookie)
		c.HTML(http.StatusSeeOther, "index.html", cookie)
		c.HTML(200, "index.html", "Hey Yaaa")

	} else {
		// c.Redirect(303, "/login")
		c.Redirect(303,"/home")

		return
	}

}
func HomeMethod(c *gin.Context) {
	cookie, err := c.Cookie("logincookie")
	if err != nil || cookie == "" {
		c.Redirect(303, "/login")

		return
	}

	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	c.HTML(200, "index.html", nil)
}
func Logout(c *gin.Context) {
	c.SetCookie("logincookie", "", -1, "", "", true, true)
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	c.Redirect(303, "/login")
}
