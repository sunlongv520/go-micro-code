package service

import (
	"context"
	"jtthink/src/Boot"
	. "jtthink/src/Course"
)

type CourseTopicServiceImpl struct {}
func(this *CourseTopicServiceImpl) GetTopic(ctx context.Context, in *TopicRequest, out *TopicResponse) error{
		if err:=Boot.GetDB().Table("course_topic").Find(&out.Result).Error;err!=nil {
			return err
		}
		return nil
}
func NewCourseTopicServiceImpl() *CourseTopicServiceImpl  {
	return &CourseTopicServiceImpl{}
}