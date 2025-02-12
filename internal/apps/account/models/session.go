package models

import "time"

type Session struct {
	ID                        string    `json:"$id"`
	CreatedAt                 time.Time `json:"$createdAt"`
	UpdatedAt                 time.Time `json:"$updatedAt"`
	UserID                    string    `json:"userId"`
	Expire                    time.Time `json:"expire"`
	Provider                  string    `json:"provider"`
	ProviderUid               string    `json:"providerUid"`
	ProviderAccessToken       string    `json:"providerAccessToken"`
	ProviderAccessTokenExpiry time.Time `json:"providerAccessTokenExpiry"`
	ProviderRefreshToken      string    `json:"providerRefreshToken"`
	IP                        string    `json:"ip"`
	OSCode                    string    `json:"osCode"`
	OSName                    string    `json:"osName"`
	OSVersion                 string    `json:"osVersion"`
	ClientType                string    `json:"clientType"`
	ClientCode                string    `json:"clientCode"`
	ClientName                string    `json:"clientName"`
	ClientVersion             string    `json:"clientVersion"`
	ClientEngine              string    `json:"clientEngine"`
	ClientEngineVersion       string    `json:"clientEngineVersion"`
	DeviceName                string    `json:"deviceName"`
	DeviceBrand               string    `json:"deviceBrand"`
	DeviceModel               string    `json:"deviceModel"`
	CountryCode               string    `json:"countryCode"`
	CountryName               string    `json:"countryName"`
	Current                   bool      `json:"current"`
	Factors                   []string  `json:"factors"`
	Secret                    string    `json:"secret"`
	MFAUpdatedAt              time.Time `json:"mfaUpdatedAt"`
}
