package utils

import (
	"encoding/json"
	"io"
	"log"
	"strings"

	"CareerCourse/models"
)

// 任务Task -> 课程Course -> 视频Video

// GetTaskId 获取任务 id
func GetTaskId() string {
	res, err := Post("https://bistu.wnssedu.com/student/rest/v1/study/getTeachPlanList", nil)
	if err != nil {
		log.Fatal("Failed to encrypt: ", err)
		return ""
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Failed to encrypt: ", err)
		return ""
	}
	task := &models.TaskResponse{}
	err = json.Unmarshal(body, task)
	if err != nil {
		log.Fatal("Failed to encrypt: ", err)
		return ""
	}
	return task.Response.CdosPlanList[0].LPlanId
}

// GetCourses 获取课程
func GetCourses() []*models.CdoCourse {
	url := "https://bistu.wnssedu.com/student/rest/v1/study/getPlanCourseList"
	payload := strings.NewReader("lPlanId=" + GetTaskId())

	res, err := Post(url, payload)
	if err != nil {
		log.Fatal("Failed to encrypt: ", err)
		return nil
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	courseResponse := &models.CourseResopnse{}
	err = json.Unmarshal(body, courseResponse)
	if err != nil {
		log.Fatal("Failed to encrypt: ", err)
		return nil
	}
	return courseResponse.Response.CdoCourseList
}

// GetVideos 获取视频
func GetVideos(lNewCourseId string) []*models.Video {

	url := "https://bistu.wnssedu.com/course/rest/v1/newcourse/getNewCourseVideoList"
	payload := strings.NewReader("lNewCourseId=" + lNewCourseId)

	res, err := Post(url, payload)
	if err != nil {
		log.Fatal("Failed to encrypt: ", err)
		return nil
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	videoResponse := &models.VideoResponse{}
	err = json.Unmarshal(body, videoResponse)
	if err != nil {
		log.Fatal("Failed to encrypt: ", err)
		return nil
	}
	return videoResponse.Response.CdoList
}
