package models

type TaskResponse struct {
	Return struct {
		NCode   int    `json:"nCode"`
		StrText string `json:"strText"`
		StrInfo string `json:"strInfo"`
	} `json:"return"`
	Response struct {
		CdosPlanList []*CdosPlan `json:"cdosPlanList"`
	} `json:"response"`
}

type CdosPlan struct {
	LId               string `json:"lId"`
	LPlanId           string `json:"lPlanId"`
	StrPlanName       string `json:"strPlanName"`
	NPlanState        int    `json:"nPlanState"`
	StrStartDate      string `json:"strStartDate"`
	StrEndDate        string `json:"strEndDate"`
	NSumScore         int    `json:"nSumScore"`
	NVersion          int    `json:"nVersion"`
	NCourseCount      int    `json:"nCourseCount"`
	NStudyTime        int    `json:"nStudyTime"`
	StrEvaluationData string `json:"strEvaluationData"`
	StrState          string `json:"strState"`
	BScoreState       bool   `json:"bScoreState"`
	NInvalid          int    `json:"nInvalid"`
	LPrePlanId        string `json:"lPrePlanId"`
	NTotalStudyTime   int    `json:"nTotalStudyTime"`
	NTotalCourseTime  int    `json:"nTotalCourseTime"`
	NCoursewareCount  int    `json:"nCoursewareCount"`
	FirstCourseName   string `json:"firstCourseName"`
}

type CdosPlanList struct {
}

type CourseResopnse struct {
	Return struct {
		NCode   int    `json:"nCode"`
		StrText string `json:"strText"`
		StrInfo string `json:"strInfo"`
	} `json:"return"`
	Response struct {
		NCourseCount  int          `json:"nCourseCount"`
		CdoCourseList []*CdoCourse `json:"cdoCourseList"`
		StrEndDate    string       `json:"strEndDate"`
		CdoTeachPlan  struct {
			LId                string `json:"lId"`
			LSchoolId          string `json:"lSchoolId"`
			StrName            string `json:"strName"`
			NType              int    `json:"nType"`
			NVersion           int    `json:"nVersion"`
			NCourseCount       int    `json:"nCourseCount"`
			StrCourseData      string `json:"strCourseData"`
			NStudyTime         int    `json:"nStudyTime"`
			NStudyRate         int    `json:"nStudyRate"`
			NWorkRate          int    `json:"nWorkRate"`
			NExaminRate        int    `json:"nExaminRate"`
			NCareerPlanRate    int    `json:"nCareerPlanRate"`
			StrStartDate       string `json:"strStartDate"`
			StrEndDate         string `json:"strEndDate"`
			DtUpdateTime       string `json:"dtUpdateTime"`
			DtPublishTime      string `json:"dtPublishTime"`
			NCourseHour        int    `json:"nCourseHour"`
			StrStudyCredit     string `json:"strStudyCredit"`
			StrTeacherName     string `json:"strTeacherName"`
			StrDest            string `json:"strDest"`
			LTeacherId         string `json:"lTeacherId"`
			NInvalid           int    `json:"nInvalid"`
			NAutoScore         int    `json:"nAutoScore"`
			NPlanState         int    `json:"nPlanState"`
			NCreativePlanRate  int    `json:"nCreativePlanRate"`
			NEvaluationRate    int    `json:"nEvaluationRate"`
			DtCreateTime       string `json:"dtCreateTime"`
			StrEvaluationData  string `json:"strEvaluationData"`
			NEvaluationCount   int    `json:"nEvaluationCount"`
			StrIntro           string `json:"strIntro"`
			StrAttachmentUrl   string `json:"strAttachmentUrl"`
			StrAttachmentName  string `json:"strAttachmentName"`
			BScoreState        bool   `json:"bScoreState"`
			NAutoCreativeScore int    `json:"nAutoCreativeScore"`
			NCareerPlanScore   int    `json:"nCareerPlanScore"`
			NCreativePlanScore int    `json:"nCreativePlanScore"`
			NPublishScope      int    `json:"nPublishScope"`
			StrStudentIds      string `json:"strStudentIds"`
			NStudentCount      int    `json:"nStudentCount"`
			LPrePlanId         string `json:"lPrePlanId"`
			NShow              int    `json:"nShow"`
			NCourseCountV1     int    `json:"nCourseCountV1"`
			NStudyTimeV1       int    `json:"nStudyTimeV1"`
		} `json:"cdoTeachPlan"`
	} `json:"response"`
}

type CdoCourse struct {
	LId             string `json:"lId"`
	StrName         string `json:"strName"`
	NTimeLength     int    `json:"nTimeLength"`
	NCourseMin      int    `json:"nCourseMin"`
	StrPicture      string `json:"strPicture"`
	LLecturerId     string `json:"lLecturerId"`
	StrLecturerName string `json:"strLecturerName"`
	StrPhoto        string `json:"strPhoto"`
	NStudyMin       int    `json:"nStudyMin"`
}

type VideoResponse struct {
	Return struct {
		NCode   int    `json:"nCode"`
		StrText string `json:"strText"`
		StrInfo string `json:"strInfo"`
	} `json:"return"`
	Response struct {
		CdoList      []*Video `json:"cdoList"`
		NRecordCount int      `json:"nRecordCount"`
		NTotalCount  string   `json:"nTotalCount"`
	} `json:"response"`
}

type Video struct {
	LId             string `json:"lId"`
	LNewCourseId    string `json:"lNewCourseId"`
	LVideoId        string `json:"lVideoId"`
	NOrder          int    `json:"nOrder"`
	DtCreateTime    string `json:"dtCreateTime"`
	DtUpdateTime    string `json:"dtUpdateTime"`
	StrName         string `json:"strName"`
	StrVideoName    string `json:"strVideoName"`
	StrTimeLength   string `json:"strTimeLength"`
	LCourseId       string `json:"lCourseId"`
	LCoursewareId   string `json:"lCoursewareId"`
	NVideoSecond    int    `json:"nVideoSecond"` // 视频总时长
	PdfName         string `json:"pdfName"`
	PdfURL          string `json:"pdfURL"`
	NTimeLength     int    `json:"nTimeLength"`
	NViewTimeLength int    `json:"nViewTimeLength"`
	NViewSecond     int    `json:"nViewSecond"` // 已观看时长
}
