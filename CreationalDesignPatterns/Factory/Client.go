package factory

// The key aim of this pattern is to hide details needed to create concrete instance
func Main() {
	dog := NewGog("papa")
	cat := NewCat("kiddy")
	dog.saying()
	cat.saying()
}
