package factory

func Main() {
	dog := NewGog("papa")
	cat := NewCat("kiddy")
	dog.Saying()
	cat.Saying()
}
