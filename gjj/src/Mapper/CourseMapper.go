package Mapper

import (
	"github.com/jinzhu/gorm"
	"jtthink/src/Boot"
	"jtthink/src/Vars"
)

func GetCourseList() *gorm.DB {
	return Boot.GetDB().Table(Vars.Table_CourseMain).
		Order("course_id desc").Limit(3)
}

const course_list="select * from course_main order by course_id desc limit ?"
func GetCourseListBySql(args ...interface{})  *gorm.DB{
	return Boot.GetDB().Raw(course_list,args...)
}
func GetCourseDetail(course_id int ) *gorm.DB{
	return Boot.GetDB().Table(Vars.Table_CourseMain).Where("course_id=?",course_id)
}
//取计数表
func GetCourseCounts(course_id int ) *gorm.DB{
	return Boot.GetDB().Table(Vars.Table_CourseCounts).Where("course_id=?",course_id)
}


