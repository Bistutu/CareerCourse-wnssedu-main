package utils

import (
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

var client = &http.Client{}

func init() {
	jar, _ := cookiejar.New(nil)
	client.Jar = jar
}

func Post(URL string, data io.Reader) (*http.Response, error) {
	req, _ := http.NewRequest("POST", URL, data)
	req.Header.Add("DNT", "1")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("User-Agent", UserAgent)
	return client.Do(req)
}

func setCoursewareId(lCoursewareId string) {
	url, _ := url.Parse("https://bistu.wnssedu.com")
	client.Jar.SetCookies(url, []*http.Cookie{
		{Name: "Courseware_9000001132" + lCoursewareId, Value: "1c1e2698fd13a4e9828322e7ff9a2563"},
		{Name: "Courseware_12200005370" + lCoursewareId, Value: "85bc5786f0737b717842c638b508bae9"},
	})
}
