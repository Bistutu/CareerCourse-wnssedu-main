package models

import (
	"fmt"
	"net/url"
	"strings"
)

type LoginPayload struct {
	LSchoolId      int    `url:"lSchoolId"`
	StrLoginId     string `url:"strLoginId"`
	StrRSAPassword string `url:"strRSAPassword"`
	StrPassword    string `url:"strPassword"`
	StrVerifyCode  string `url:"strVerifyCode"`
}

func (payload *LoginPayload) ToIOReader() *strings.Reader {
	data := url.Values{}
	data.Add("lSchoolId", fmt.Sprintf("%d", payload.LSchoolId))
	data.Add("strLoginId", payload.StrLoginId)
	data.Add("strRSAPassword", payload.StrRSAPassword)
	data.Add("strPassword", payload.StrPassword)
	data.Add("strVerifyCode", payload.StrVerifyCode)
	return strings.NewReader(data.Encode())
}
