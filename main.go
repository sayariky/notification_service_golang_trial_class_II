package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

type MailerRequest struct {
	BuyerEmail   string `json:"buyer_email"`
	BuyerAddress string `json:"buyer_address"`
	ProductName  string `json:"product_name"`
}

func main() {
	s := gin.Default()

	s.POST("/mailer", func(ctx *gin.Context) {
		var mailerBody MailerRequest
		if err := ctx.ShouldBindJSON(&mailerBody); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		m := gomail.NewMessage()
		m.SetHeader("From", "trial-class@golang.com")
		m.SetHeader("To", mailerBody.BuyerEmail)
		// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
		m.SetHeader("Subject", "Success Create Order")
		m.SetBody("text/html", fmt.Sprintf("Order produk %s berhasil, dan akan dikirim ke %s segera", mailerBody.ProductName, mailerBody.BuyerAddress))
		// m.Attach("/home/Alex/lolcat.jpg")

		d := gomail.NewDialer(
			"smtp-relay.sendinblue.com",
			587,
			"dayatketangga95@gmail.com",
			"xsmtpsib-e93e29d754e19d4990f146b772701d93e6dc72a8cc571ed45efdef4ee1f9db95-dbSD7LkntEzcCUg0",
		)

		// Send the email to Bob, Cora and Dan.
		if err := d.DialAndSend(m); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	s.Run(":8001")

}
