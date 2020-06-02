package service

import (
	"context"
	"jtthink/src/Boot"
	. "jtthink/src/Course"

	"jtthink/src/Vars"
)


func NewCourseModel(id int32,name string) *CourseModel  {
	return &CourseModel{CourseId:id,CourseName:name}
}
type CourseServiceImpl struct{}
func(this *CourseServiceImpl) ListForTop(ctx context.Context, req *ListRequest, rsp *ListResponse) error{
	course:=make([]*CourseModel,0)
	err:=Boot.GetDB().Table(Vars.Table_CourseMain).Order("course_id desc").Find(&course).Error
	if err!=nil{
		return err
	}
	rsp.Result=course
	return nil

	//ret:=make([]*CourseModel,0)
	//ret=append(ret,NewCourseModel(101,"java课程"),NewCourseModel(102,"PHP课程"))
	//rsp.Result=ret
	return nil
}
func NewCourseServiceImpl() *CourseServiceImpl  {
	return &CourseServiceImpl{}
}