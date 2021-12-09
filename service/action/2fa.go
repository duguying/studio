package action

import (
	"duguying/studio/g"
	"duguying/studio/modules/db"
	"encoding/base32"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/dgryski/dgoogauth"
	"github.com/gin-gonic/gin"
	"github.com/gogather/json"
	"rsc.io/qr"
)

func QrGoogleAuth(c *gin.Context) {
	uidStr := c.DefaultQuery("uid", "")

	uid, err := strconv.ParseUint(uidStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	user, err := db.GetUserById(g.Db, uint(uid))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	secretBase32 := base32.StdEncoding.EncodeToString([]byte(user.TfaSecret))
	account := fmt.Sprintf("%s@duguying.net", user.Username)
	issuer := "duguying.net"

	URL, err := url.Parse("otpauth://totp")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	URL.Path += "/" + url.PathEscape(issuer) + ":" + url.PathEscape(account)

	params := url.Values{}
	params.Add("secret", secretBase32)
	params.Add("issuer", issuer)

	URL.RawQuery = params.Encode()
	fmt.Printf("URL is %s\n", URL.String())

	code, err := qr.Encode(URL.String(), qr.Q)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	c.Data(http.StatusOK, "image/png", code.PNG())
}

type TfaAuthRequest struct {
	Uid   uint   `json:"uid"`
	Token string `json:"token"`
}

func (tar *TfaAuthRequest) String() string {
	c, _ := json.Marshal(tar)
	return string(c)
}

func TfaAuth(c *gin.Context) {
	tar := TfaAuthRequest{}
	err := c.BindJSON(&tar)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	user, err := db.GetUserById(g.Db, tar.Uid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	secretBase32 := base32.StdEncoding.EncodeToString([]byte(user.TfaSecret))
	otpc := &dgoogauth.OTPConfig{
		Secret:      secretBase32,
		WindowSize:  1,
		HotpCounter: 0,
		UTC:         true,
	}

	val, err := otpc.Authenticate(tar.Token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	if !val {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": "Sorry, Not Authenticated",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"ok":  true,
			"err": "Authenticated!",
		})
		return
	}

}
