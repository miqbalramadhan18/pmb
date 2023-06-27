package controller

import (
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/miqbalramadhan18/balbackend"
	"github.com/miqbalramadhan18/pmb/config"
	"github.com/whatsauth/whatsauth"
)

var DataJalurPenerimaan = "JalurPenerimaan"
var DataInformasi = "Informasi"
var DataBiaya = "Biaya"

// type HTTPRequest struct {
// 	Header string `json:"header"`
// 	Body   string `json:"body"`
// }

// func Sink(c *fiber.Ctx) error {
// 	var req HTTPRequest
// 	req.Header = string(c.Request().Header.Header())
// 	req.Body = string(c.Request().Body())
// 	return c.JSON(req)
// }

func WsWhatsAuthQR(c *websocket.Conn) {
	whatsauth.RunSocket(c, config.PublicKey, config.Usertables[:], config.Ulbimariaconn)
}

func PostWhatsAuthRequest(c *fiber.Ctx) error {
	if string(c.Request().Host()) == config.Internalhost {
		var req whatsauth.WhatsauthRequest
		err := c.BodyParser(&req)
		if err != nil {
			return err
		}
		ntfbtn := whatsauth.RunModuleLegacy(req, config.PrivateKey, config.Usertables[:], config.Ulbimariaconn)
		return c.JSON(ntfbtn)
	} else {
		var ws whatsauth.WhatsauthStatus
		ws.Status = string(c.Request().Host())
		return c.JSON(ws)
	}

}

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetJalurPenerimaan(c *fiber.Ctx) error {
	getstatus := balbackend.GetDataJalurPenerimaan("Mandiri")
	return c.JSON(getstatus)
}

func GetInformasi(c *fiber.Ctx) error {
	getstatus := balbackend.GetDataInformasi("catatan")
	return c.JSON(getstatus)
}

func GetBiaya(c *fiber.Ctx) error {
	getstatus := balbackend.GetDataBiaya("7.700.000.00")
	return c.JSON(getstatus)
}
