/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package utils

import (
	"Syns/servers/syns-tokens-ms/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @import

// @dev Loads environment variables
func LoadEnvVars() {
	err := godotenv.Load();
	HandleException(err);
}


// @dev Handdle error exception
//
// @param e error - the passed in error
func HandleException(e error) {if (e != nil) {log.Panic(e)}}

// @dev Sets up config for cors
// 
// @return gin.HandlerFunc
func SetupCorsConfig() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: 		[]string{os.Getenv("CORS_ALLOW_LOCAL_ORIGIN"), os.Getenv("CORS_ALLOW_PRODUCTION_CLIENT_ORIGIN"), os.Getenv("CORS_ALLOW_STAGING_CLIENT_ORIGIN")},
		AllowMethods:		[]string{"POST", "PATCH", "PUT", "DELETE", "GET"},
		AllowHeaders: 		[]string{"Origin", "Authorization", "Access-Control-Allow-Origin"},	
		AllowCredentials: 	true,
		MaxAge: 			12*time.Hour,
	})
}

// @dev Handle testing `wallet_address` to match ETH crypto wallet address
// 
// @param wallet_address *string
// 
// @return bool
// 
// @return error
func TestEthAddress(wallet_address *string) (bool, error) {
	pattern := os.Getenv("ETH_ADDRESS_REGEX")
	return regexp.MatchString(pattern, *wallet_address)
}


// @dev Report ip address to SYNS_PLATFORM_EMAIL
//
// @param ip string
func EmailNotification(mode string, args interface{}) {
	// prepare states
	synsFeedbackEmail := os.Getenv("SYNS_PLATFORM_ALERT_EMAIL")
	synsFeedbackEmailPassword := os.Getenv("SYNS_PLATFORM_ALERT_EMAIL_PASSWORD")
	subject := ""
	description := ""
	smtpHost := "smtp.titan.email" // Titan Email SMTP server
	smtpPort := "587"              // Titan Email SMTP port
	isDevPurpose := false;

	// prepare subject
	if mode == "NEW NFT" {
		subject = "Subject: New NFT Minted Alert"
	}

	// prepare description
	switch obj := args.(type) {
		case *models.SynsNFT:
			if os.Getenv("GIN_MODE") != "release" {
				isDevPurpose = true
			}
			description = "Token ID: " +strconv.Itoa(obj.TokenID)+ ".\n Minter address: " +obj.TokenOwner+"."
	default:
		return
	}

	// only send email if this request is not for dev purpose
	if !isDevPurpose {
		// Set up authentication information for Gmail's SMTP server
		auth := smtp.PlainAuth("", synsFeedbackEmail, synsFeedbackEmailPassword, smtpHost)

		// Compose the email message
		to := []string{synsFeedbackEmail}
		msg := []byte("To: " +synsFeedbackEmail+ " +\r\n" +subject + "\r\n" + "\r\n" +description)

		// Send the email using Gmail's SMTP server
		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, synsFeedbackEmail, to, msg)
		if err != nil {
			// Handle any errors that occur while sending the email
			panic(err)
		}
	}
}

// @dev do http request
// 
// @param url string
// 
// @return body []byte
func DoHttp(url string, apiKeyHeader string, apiKey string, resObject *map[string]interface{}) (map[string]interface{}) {
	// prepare request
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	if strings.Compare(apiKeyHeader, "") != 0 || strings.Compare(apiKey, "") != 0 {
		req.Header.Add(apiKeyHeader, apiKey)
	}

	// ship request
	res, _ := http.DefaultClient.Do(req)

	// close request and prepare body
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// parse json from []byte to JSON
	json.Unmarshal(body, resObject)

	return *resObject
}