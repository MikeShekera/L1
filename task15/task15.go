package task15

var justString string

func main() {
	someFunc()
}

//Для создания justString необходимо явно скопировать срез, чтобы не делать ссылку на строку длиной 1024 элемента

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string(v[:100])
}
