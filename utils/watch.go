package utils

import (
	"io"
	"log"
	"strconv"
	"time"
)

// Watch 观看视频
func Watch(lNewCourseId, lCoursewareId string) bool {
	setCoursewareId(lCoursewareId)
	stampTime := time.Now().UnixMilli()
	url := "https://bistu.wnssedu.com/course/Servlet/recordStudy.svl?lCoursewareId=" + lCoursewareId +
		"&strStartTime=" + strconv.FormatInt(stampTime, 10) +
		"&nCurSeconds=181&lNewCourseId=" + lNewCourseId
	res, err := Post(url, nil)
	defer res.Body.Close()
	if err != nil {
		log.Fatal("Failed to encrypt: ", err)
		return false
	}
	bytes, _ := io.ReadAll(res.Body)
	log.Println("观看视频：", lCoursewareId, "状态码：", res.StatusCode, "返回信息：", string(bytes))
	return true
}
