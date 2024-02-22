package colors

// Для вывода цветного текста в терминале
// пример 
// fmt.Println(colors.ColorRed(),"text",colors.ColorReset())
// Первый аргумент цвет красный, третий удаление цвета, иначе остальной
// вывод в терминале будет красным


func ColorReset() string {
	return "\033[0m" // очистка
}
func ColorRed() string {
	return  "\033[31m" // красный
}
func ColorGreen() string {
	return  "\033[32m" // зеленый
}
func ColorYellow() string {
	return "\033[33m" // желтый
}
func ColorBlue() string {
	return  "\033[34m" // синий
}
func ColorPurple() string {
	return  "\033[35m" // фиолетовый
}
func ColorCyan() string {
	return "\033[36m" // голубой
}
func ColorWhite() string {
	return  "\033[37m" // белый
}


