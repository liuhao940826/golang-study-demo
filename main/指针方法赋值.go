package main

import "fmt"

//*操作符作为右值时，意义是取指针的值；作为左值时，也就是放在赋值操作符的左边时，表示 a 指向的变量。
// 其实归纳起来，*操作符的根本意义就是操作指针指向的变量。当操作在右值时，就是取指向变量的值；当操作在左值时，就是将值设置给指向的变量。

func checkAuth(currentUserId *int64, orgId *int64) {

	//传递  把 1 2 赋值给指定地址
	*currentUserId = 1

	*orgId = 2

}

//数组长度在定义后无法再次修改 数组是值类型 每次传递产生一个副本  用slice来弥补不足
func main() {
	//var currentUserId ,orgId int64
	//fmt.Println("currentUserId= ,orgId",currentUserId ,orgId)
	//
	//checkAuth(&currentUserId,&orgId)
	//
	//fmt.Println("后面的currentUserId= ,orgId",currentUserId ,orgId)

	var zero int64 = 0

	s := "asd"
	var n int64 = 1

	var input = CreateProjectReq{
		Code:          &s,
		PreCode:       &s,
		PriorityID:    &n,
		ProjectTypeID: &n,
	}
	fmt.Printf("地址 %T ", &input)
	fmt.Println("前一次 ", input, *input.Code, *input.PreCode, *input.PriorityID, *input.ProjectTypeID)

	blankString := ""

	assgin(zero, &input, blankString)

	fmt.Println("后一次 ", *input.Code, *input.PreCode, *input.PriorityID, *input.ProjectTypeID)

}

//后去地址的值  值  值=>值
func assgin(zero int64, input *CreateProjectReq, blankString string) {
	*input.Code = blankString
	*input.PreCode = blankString
	*input.PriorityID = zero
	*input.ProjectTypeID = zero
}

type CreateProjectReq struct {
	// 编号
	Code *string `json:"code"`
	// 名称
	Name string `json:"name"`
	// 前缀编号
	PreCode *string `json:"preCode"`
	// 负责人id
	Owner int64 `json:"owner"`
	// 项目类型
	ProjectTypeID *int64 `json:"projectTypeId"`
	// 优先级
	PriorityID *int64 `json:"priorityId"`
	// 计划开始时间
	// 计划结束时间
	// 项目公开性,1公开,2私有
	PublicStatus int `json:"publicStatus"`
	// 资源id
	ResourceID *int64 `json:"resourceId"`
	// 是否归档,1归档,2未归档
	IsFiling *int `json:"isFiling"`
	// 描述
	Remark *string `json:"remark"`
	// 项目状态
	Status *int64 `json:"status"`
	// 更新人
	Updator *int64 `json:"updator"`
	// 资源路径
	ResourcePath string `json:"resourcePath"`
	// 资源类型1本地2oss3钉盘
	ResourceType int `json:"resourceType"`
	// 用户成员id
	MemberIds []int64 `json:"memberIds"`
}
