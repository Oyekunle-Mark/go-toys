package equal

func Compare(x, y map[string]int) bool {
    for key, value := range x {
        val, ok := y[key]
			if ok {
                if val == value {
                    continue
                } else {
                    return false
                }
			} else {
				return false
			}
    }  
    return true
}