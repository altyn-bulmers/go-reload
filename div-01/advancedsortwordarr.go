package piscine

func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	arrLen := 0
	for r := range array {
		arrLen = r + 1
	}
	temp := 0
	for i:=0; i < arrLen-1; i++{
		temp = i
		for j:= i + 1; j < arrLen; j++{
			if f (array[j], array[temp]) < 0 {
				temp = j
			}
		}
		array[i], array[temp] = array[temp], array[i]
	}
}


