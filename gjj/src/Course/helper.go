package Course

import "context"

func NewCourseModel(id int32,name string) *CourseModel  {
	return &CourseModel{CourseId:id,CourseName:name}
}
type CourseServiceImpl struct{}
func(this *CourseServiceImpl) ListForTop(ctx context.Context, req *ListRequest, rsp *ListResponse) error{
   ret:=make([]*CourseModel,0)
   ret=append(ret,NewCourseModel(101,"java课程"),NewCourseModel(102,"PHP课程"))
   rsp.Result=ret
   return nil
}
func NewCourseServiceImpl() *CourseServiceImpl  {
  return &CourseServiceImpl{}
}