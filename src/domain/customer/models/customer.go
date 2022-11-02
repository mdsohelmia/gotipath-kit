package models

import (
	"database/sql"
	"time"
)

// Customer Model
type Customer struct {
	ID   uint   `json:"id"`
	UUID string `json:"uuid"`
	// Customer Account Type [PERSONAL|BUSINESS]
	AccountType string `json:"account_type"`
	// Customer Billing Type [POST_PAID | PRE_PAID]
	BillingType string `json:"billing_type"`
	//Personal Inoformation
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Email              string `json:"email"`
	Phone              string `json:"phone"`
	Password           string `json:"password"`
	CountryCallingCode string `json:"country_calling_code"`
	Avatar             string `json:"avatar"`

	// Bussiness Information
	CompanyName            string `json:"company_name"`
	Nid                    string `json:"nid"`
	TradeLicenseNumberFile string `json:"trade_license_image_path"`
	TradeLicenseNumber     string `json:"trade_license_number"`
	VatNumber              string `json:"vat_number"`
	VatFile                string `json:"vat_file"`
	SelfPayVat             bool   `json:"self_pay_vat"`
	// Customer Address
	StAddress string `json:"st_address"`
	City      string `json:"city"`
	State     string `json:"state"`
	CountryId int    `json:"country_id"`

	// Customer Verification
	EmailVerifyAt sql.NullTime `json:"email_verify_at"`
	PhoneVerifyAt sql.NullTime `json:"phone_verify_at"`
	//Cusotmer accrute Charege
	Balance            float64 `json:"balance"`
	AccruedCharge      float64 `json:"accrued_charge"`       //main accrued charge
	CdnAccruedCharge   float64 `json:"cdn_accrued_charge"`   //cdn accrued charge
	CloudAccruedCharge float64 `json:"cloud_accrued_charge"` //cloud accrued charge
	IsCloudEnable      bool    `json:"is_cloud_enable"`
	CloudId            string  `json:"cloud_id"`
	// Securty Fileds
	Otp          int
	OtpExpiredAt sql.NullTime
	Is2faActive  bool `json:"is_2fa_active"`
	// Thirdparty SSO
	OauthId    string `json:"oauth_id"`
	OauthType  string `json:"oauth_type"`
	OauthToken string `json:"oauth_token"`
	// Api Token
	ApiToken  sql.NullString `json:"api_token"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func (c Customer) TableName() string {
	return "custoemrs"
}
