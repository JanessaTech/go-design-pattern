package flyweight

import "fmt"

type Student struct {
	name string
}

func NewStudent(name string) *Student {
	return &Student{name: name}
}

type ClassRoom struct {
	students map[int]*Student
}

func (ClassRoom *ClassRoom) getStudent(i int, name string) *Student {
	if ClassRoom.students[i] == nil {
		newStu := NewStudent(name)
		fmt.Println("Student", name, "is created at", i)
		ClassRoom.students[i] = newStu
		return newStu
	}
	return ClassRoom.students[i]
}

func Main() {
	classRoom := ClassRoom{students: map[int]*Student{}}
	stu1 := classRoom.getStudent(0, "stu1")
	stu2 := classRoom.getStudent(1, "stu2")
	stu3 := classRoom.getStudent(2, "stu3")

	fmt.Println(stu1)
	fmt.Println(stu2)
	fmt.Println(stu3)
}
