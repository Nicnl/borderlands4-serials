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
	"github.com/gin-gonic/gin/render"
)

var (
	//go:embed index.html
	indexHtmlBytes []byte

	//go:embed matt_special_place.html
	mattHtmlBytes []byte
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
		for _, stat := range strings.Split(baseType.Stats, " ") {
			if strings.Contains(stat, "=") {
				continue
			}
			additionalDataArr = append(additionalDataArr, stat)
		}

		additionalDataArr = append(additionalDataArr, "\""+baseType.Name+"\"")
	}

	if len(additionalDataArr) > 0 {
		return strings.Join(additionalDataArr, " ")
	}
	return ""
}

func CORSMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}

const README_JSON = `
Hey you, peeking at the JSON!

No need to peek anymore 😉
The complete source code and documentation is available here on Github: 
https://github.com/Nicnl/borderlands4-serials

With love,
- Nicnl & InflamedSebi
`

func main() {
	r := gin.Default()

	r.OPTIONS("/api/v1/deserialize_bulk", CORSMiddleware)
	r.OPTIONS("/api/v1/deserialize", CORSMiddleware)
	r.OPTIONS("/api/v1/reserialize", CORSMiddleware)
	r.OPTIONS("/api/v1/serialize_bulk", CORSMiddleware)

	r.POST("/api/v1/deserialize_bulk", CORSMiddleware, func(c *gin.Context) {
		var jsonReq []string
		if err := c.BindJSON(&jsonReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		results := make(map[string]gin.H, len(jsonReq))
		for _, serialB85 := range jsonReq {
			item, err := codex.Deserialize(serialB85)
			if err != nil {
				results[serialB85] = gin.H{
					"success": false,
				}
				continue
			} else {
				result := gin.H{
					"deserialized": item.Serial.String(),
					"success":      true,
				}

				if level, found := item.Level(); found {
					result["deserialized_level"] = level
				}

				if itemType, found := item.Type(); found {
					result["deserialized_manufacturer"] = itemType.Manufacturer
					result["deserialized_type"] = itemType.Type
				}

				if baseType, found := item.BaseBarrel(); found {
					result["deserialized_base_name"] = baseType.Name

					for _, stat := range strings.Split(baseType.Stats, " ") {
						switch stat {
						case "common", "uncommon", "rare", "epic", "legendary":
							result["deserialized_rarity"] = stat
						}
					}
				}

				deserializedParts := make([]string, 0, len(item.Serial))
				{
					pos := 0
					for {
						part := item.FindPartAtPos(pos, true)
						if part == nil {
							break
						}

						deserializedParts = append(deserializedParts, part.String())
						pos++
					}
				}

				result["deserialized_parts"] = deserializedParts

				results[serialB85] = result
			}
		}

		c.Render(http.StatusOK, render.IndentedJSON{Data: results})
	})

	r.POST("/api/v1/serialize_bulk", CORSMiddleware, func(c *gin.Context) {
		var jsonReq []string
		if err := c.BindJSON(&jsonReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		output := make([]string, 0, len(jsonReq))
		for i, serialB85 := range jsonReq {
			output = append(output, "")

			// Clean the reserialized data
			{
				// Add a terminator if missing
				if !strings.HasSuffix(serialB85, "|") {
					serialB85 = serialB85 + "|"
				}

				// If more than two terminators, keep only the first two
				for len(serialB85) >= 2 && serialB85[len(serialB85)-1] == '|' && serialB85[len(serialB85)-2] == '|' {
					serialB85 = serialB85[:len(serialB85)-1]
				}
			}

			//fmt.Fprintln(os.Stderr, "# Reserialize:")
			//fmt.Fprintln(os.Stderr, " - From: ", serialB85)

			s := serial.Serial{}
			err := s.FromString(serialB85)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Failed to import:", serialB85, "=>", err.Error())
				continue
			}

			data := serial.Serialize(s)
			b85Serial := b85.Encode(data)

			output[i] = b85Serial
		}

		c.Render(http.StatusOK, render.IndentedJSON{Data: output})
	})

	r.POST("/api/v1/deserialize", CORSMiddleware, func(c *gin.Context) {
		var jsonReq struct {
			SerialB85 string `json:"serial_b85"`
		}
		if err := c.BindJSON(&jsonReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		fmt.Fprintln(os.Stderr, "# Deserialize:")
		fmt.Fprintln(os.Stderr, " - From: ", jsonReq.SerialB85)

		item, err := codex.Deserialize(jsonReq.SerialB85)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to deserialize:", jsonReq.SerialB85, "=>", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid serial"})
			return
		}

		fmt.Fprintln(os.Stderr, " - To:   ", item.Serial.String())
		fmt.Fprintln(os.Stderr, " - Infos:", additionalDataFunc(item))
		c.JSON(http.StatusOK, gin.H{
			"deserialized":    item.Serial.String(),
			"additional_data": additionalDataFunc(item),
			"readme":          README_JSON,
		})
	})

	r.POST("/api/v1/reserialize", CORSMiddleware, func(c *gin.Context) {
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

		fmt.Fprintln(os.Stderr, "# Reserialize:")
		fmt.Fprintln(os.Stderr, " - From: ", jsonReq.Deserialized)

		s := serial.Serial{}
		err := s.FromString(jsonReq.Deserialized)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to import:", jsonReq.Deserialized, "=>", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid deserialized data"})
			return
		}

		data := serial.Serialize(s)
		b85Serial := b85.Encode(data)

		item, err := codex.Deserialize(b85Serial)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to re-deserialize:", b85Serial, "=>", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to re-deserialize data"})
			return
		}

		fmt.Fprintln(os.Stderr, " - To:   ", item.B85)
		fmt.Fprintln(os.Stderr, " - Infos:", additionalDataFunc(item))
		c.JSON(http.StatusOK, gin.H{
			"serial_b85":      item.B85,
			"additional_data": additionalDataFunc(item),
			"readme":          README_JSON,
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

	routeMattFromDisk := func(c *gin.Context) {
		// Load file directly from disk for debug
		mattHtmlBytes, err := os.ReadFile("C:\\Users\\Nicnl\\GolandProjects\\borderlands_4_serials\\cmd\\api\\matt_special_place.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to load index.html")
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", mattHtmlBytes)
	}

	routeIndexEmbeded := func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", indexHtmlBytes)
	}

	routeMattEmbeded := func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", mattHtmlBytes)
	}

	if routeMattFromDisk != nil || routeIndexFromDisk != nil || routeMattEmbeded != nil || routeIndexEmbeded != nil {
		// Prevent go compiler from being annoying
	}

	r.GET("/", routeIndexEmbeded)
	r.GET("/index.html", routeIndexEmbeded)
	r.GET("/matt.html", routeMattEmbeded)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
