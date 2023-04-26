/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package utils

// @import
import (
	"Syns/servers/syns-users-ms/models"
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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

// @dev Handle testing `signature` to match ETH signed signature
// 
// @param signature *string
// 
// @return bool
// 
// @return error
func TestSignature(signature *string) (bool, error) {
	pattern := os.Getenv("SIGNATURE_REGEX")
	return regexp.MatchString(pattern, *signature)
}

// @dev Handle striping off all special characters in `username` passed in from http request in replace them with "-"
// 
// @param username *string
// 
// @return string
// 
// @return error
func SanitizeUsername(username *string) (*string, error) {
	pattern := os.Getenv("USERNAME_REGEX")
	reg, err := regexp.Compile(pattern)
	validUsername :=strings.ToLower(strings.Trim(reg.ReplaceAllString(*username, "-"), "-"))
	return &validUsername, err
}

// @dev Calculate random avatar for users
func RandomizeAvatar() string {
	// Seed the random number generator with the current time
    rand.Seed(time.Now().UnixNano())

    // Generate a random number between 1 and 7
    randomNum := rand.Intn(7) + 1

	// return `avatar` string
	return fmt.Sprintf("demo-avatar-%d.png", randomNum)
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

	switch opt := args.(type) {
	case models.Feedback:
		if opt.Email == "" {
			description = "From an anonymous: " +opt.Feedback
		} else {
			description = "From " +opt.Email+ ": " +opt.Feedback
		}
	default:
		return
	}

	if mode == "FEEDBACK" {
		subject = "Subject: New Feedback Alert"
	}

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