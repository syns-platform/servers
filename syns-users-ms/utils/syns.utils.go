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
	"log"
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
		AllowOrigins: 		[]string{os.Getenv("CORS_ALLOW_LOCAL_ORIGIN"), os.Getenv("CORS_ALLOW_VERCEL_ORIGIN")},
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


// @dev Report ip address to SYNS_PLATFORM_EMAIL
// 
// @param ip string
func ReportVisitor(ip string) {
	// Set up authentication information for Gmail's SMTP server
	SYNS_EMAIL := os.Getenv("SYNS_PLATFORM_EMAIL")
	auth := smtp.PlainAuth("", SYNS_EMAIL, os.Getenv("SYNS_PLATFORM_EMAIL_PASSWORD"), "smtp.gmail.com")

	// Compose the email message
	to := []string{SYNS_EMAIL}
	msg := []byte("To: " +SYNS_EMAIL+ " +\r\n" +
		"Subject: Visitor IP address\r\n" +
		"\r\n" +
		"The IP address of the client is: " + ip + "\r\n")

	// Send the email using Gmail's SMTP server
	err := smtp.SendMail("smtp.gmail.com:587", auth, SYNS_EMAIL, to, msg)
	if err != nil {
		// Handle any errors that occur while sending the email
		panic(err)
	}
}