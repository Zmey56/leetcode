package isLongPressedName

//Your friend is typing his name into a keyboard. Sometimes, when typing a character c, the key might get long pressed,
//and the character will be typed 1 or more times.
//
//You examine the typed characters of the keyboard. Return True if it is possible that it was your friends name,
//with some characters (possibly none) being long pressed.

func isLongPressedName(name string, typed string) bool {
	i, j := 0, 0
	for j < len(typed) {
		if i < len(name) && name[i] == typed[j] {
			i++
			j++
		} else if j > 0 && typed[j] == typed[j-1] {
			j++
		} else {
			return false
		}
	}
	return i == len(name)

}
