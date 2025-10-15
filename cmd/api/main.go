package main

import (
	"borderlands_4_serials/b4s/b85"
	"borderlands_4_serials/b4s/codex"
	"borderlands_4_serials/b4s/serial"
	_ "embed"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	//go:embed index.html
	indexHtmlBytes []byte
)

func additionalDataFunc(item *codex.Item) string {
	additionalDataArr := make([]string, 0, 5)

	if level, found := item.Level(); found {
		additionalDataArr = append(additionalDataArr, "L"+fmt.Sprint(level))
	}

	if itemType, found := item.Type(); found {
		additionalDataArr = append(additionalDataArr, itemType.Type+":"+itemType.Manufacturer)
	} else {
		additionalDataArr = append(additionalDataArr, "?")
	}

	if baseType, found := item.BaseBarrel(); found {
		additionalDataArr = append(additionalDataArr, baseType.Name)
	}

	if len(additionalDataArr) > 0 {
		return strings.Join(additionalDataArr, " ")
	}
	return ""
}

func main() {
	r := gin.Default()

	r.POST("/api/v1/deserialize", func(c *gin.Context) {
		var jsonReq struct {
			SerialB85 string `json:"serial_b85"`
		}
		if err := c.BindJSON(&jsonReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		item, err := codex.Deserialize(jsonReq.SerialB85)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid serial"})
			return
		}

		fmt.Println("# Deserialize:")
		fmt.Println(" - From: ", jsonReq.SerialB85)
		fmt.Println(" - To:   ", item.Serial.String())
		fmt.Println(" - Infos:", additionalDataFunc(item))
		c.JSON(http.StatusOK, gin.H{
			"deserialized":    item.Serial.String(),
			"additional_data": additionalDataFunc(item),
		})
	})

	r.POST("/api/v1/reserialize", func(c *gin.Context) {
		var jsonReq struct {
			Deserialized string `json:"deserialized"`
		}
		if err := c.BindJSON(&jsonReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		// Clean the reserialized data
		{
			// Add a terminator if missing
			if !strings.HasSuffix(jsonReq.Deserialized, "|") {
				jsonReq.Deserialized = jsonReq.Deserialized + "|"
			}

			// If more than two terminators, keep only the first two
			for len(jsonReq.Deserialized) >= 2 && jsonReq.Deserialized[len(jsonReq.Deserialized)-1] == '|' && jsonReq.Deserialized[len(jsonReq.Deserialized)-2] == '|' {
				jsonReq.Deserialized = jsonReq.Deserialized[:len(jsonReq.Deserialized)-1]
			}
		}

		s := serial.Serial{}
		err := s.FromString(jsonReq.Deserialized)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid deserialized data"})
			return
		}

		data := serial.Serialize(s)
		b85Serial := b85.Encode(data)

		item, err := codex.Deserialize(b85Serial)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to re-deserialize data"})
			return
		}

		fmt.Println("# Reserialize:")
		fmt.Println(" - From: ", jsonReq.Deserialized)
		fmt.Println(" - To:   ", item.B85)
		fmt.Println(" - Infos:", additionalDataFunc(item))
		c.JSON(http.StatusOK, gin.H{
			"serial_b85":      item.B85,
			"additional_data": additionalDataFunc(item),
		})
	})

	routeIndexFromDisk := func(c *gin.Context) {
		// Load file directly from disk for debug
		indexHtmlBytes, err := os.ReadFile("C:\\Users\\Nicnl\\GolandProjects\\borderlands_4_serials\\cmd\\api\\index.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to load index.html")
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", indexHtmlBytes)
	}
	if routeIndexFromDisk != nil {
		// Prevent go compiler from being annoying
	}

	routeIndexEmbeded := func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", indexHtmlBytes)
	}

	r.GET("/", routeIndexEmbeded)
	r.GET("/index.html", routeIndexEmbeded)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
