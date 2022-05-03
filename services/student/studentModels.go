package student_services

import (
	student_repositories "RU-Smart-Workspace/ru-smart-api/repositories/student"

	"github.com/go-redis/redis/v8"
)

type (

	studentServices struct {
		studentRepo student_repositories.StudentRepoInterface
		redis_cache *redis.Client
	}

	AuthenPlayload struct {
		Std_code string `json:"std_code"`
		Refresh_token string `json:"refresh_token"`
	}

	TokenResponse struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
		IsAuth       bool   `json:"isAuth"`
		Message      string `json:"message"`
		StatusCode   int    `json:"status_code"`
	}

	// claims คือข้อมูลที่อยู่ในส่วน Payload ของ Token
	// -iss (issuer) : เว็บหรือบริษัทเจ้าของ token
	// -sub (subject) : subject ของ token
	// -aud (audience) : ผู้รับ token
	// -exp (expiration time) : เวลาหมดอายุของ token
	// -nbf (not before) : เป็นเวลาที่บอกว่า token จะเริ่มใช้งานได้เมื่อไหร่
	// -iat (issued at) : ใช้เก็บเวลาที่ token นี้เกิดปัญหา
	// -jti (JWT id) : เอาไว้เก็บไอดีของ JWT แต่ละตัวนะครับ
	// -name (Full name) : เอาไว้เก็บชื่อ
	ClaimsToken struct {
		Issuer              string `json:"issuer"`
		Subject             string `json:"subject"`
		Role                string `json:"role"`
		ExpiresAccessToken  string `json:"expires_access_token"`
		ExpiresRefreshToken string `json:"expiration_refresh_token"`
	}

	StudentProfileService struct {
		STD_CODE             string `json:"STD_CODE"`
		NAME_THAI            string `json:"NAME_THAI"`
		NAME_ENG             string `json:"NAME_ENG"`
		BIRTH_DATE           string `json:"BIRTH_DATE"`
		STD_STATUS_DESC_THAI string `json:"STD_STATUS_DESC_THAI"`
		CITIZEN_ID           string `json:"CITIZEN_ID"`
		REGIONAL_NAME_THAI   string `json:"REGIONAL_NAME_THAI"`
		STD_TYPE_DESC_THAI   string `json:"STD_TYPE_DESC_THAI"`
		FACULTY_NAME_THAI    string `json:"FACULTY_NAME_THAI"`
		MAJOR_NAME_THAI      string `json:"MAJOR_NAME_THAI"`
		WAIVED_NO            string `json:"WAIVED_NO"`
		WAIVED_PAID          string `json:"WAIVED_PAID"`
		WAIVED_TOTAL_CREDIT  int    `json:"WAIVED_TOTAL_CREDIT"`
		CHK_CERT_NAME_THAI   string `json:"CHK_CERT_NAME_THAI"`
		PENAL_NAME_THAI      string `json:"PENAL_NAME_THAI"`
		MOBILE_TELEPHONE     string `json:"MOBILE_TELEPHONE"`
		EMAIL_ADDRESS        string `json:"EMAIL_ADDRESS"`
	}

	StudentServicesInterface interface {
		Authentication(stdCode string) (*TokenResponse, error)
		RefreshAuthentication(refreshToken, stdCode string) (*TokenResponse, error)
		GetStudentProfile(stdCode string) (*StudentProfileService, error)
	}
)

func NewStudentServices(studentRepo student_repositories.StudentRepoInterface, redis_cache *redis.Client) StudentServicesInterface {
	return &studentServices{
		studentRepo: studentRepo,
		redis_cache: redis_cache,
	}
}

