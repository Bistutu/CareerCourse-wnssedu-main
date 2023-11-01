package main

import (
	"fmt"
	"os"
	"sync"

	"CareerCourse/utils"
)

const poolSize = 3 // 限制协程并发数
var wg sync.WaitGroup

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide username and password!")
		return
	}

	username := os.Args[1] // 账号
	password := os.Args[2] // 密码
	Run(username, password)
}

func Run(username, password string) {
	//ch := make(chan struct{}, 3)
	utils.Login(username, password) // 登录
	courses := utils.GetCourses()   // 获取全部课程
	for _, v := range courses {     // 遍历课程
		courseId := v.LId
		videos := utils.GetVideos(courseId) // 获取课程下的视频集合
		for _, v := range videos {
			for i := 0; i < 30; i++ { // 30 是个阈值（可以动态调整），每次请求都会为视频增加一些观看时间
				wg.Add(1)
				//ch <- struct{}{}
				func() {
					defer func() {
						//<-ch
						wg.Done()
					}()
					utils.Watch(courseId, v.LCoursewareId) // 观看视频
				}()
			}
		}
	}
	wg.Wait()
}
