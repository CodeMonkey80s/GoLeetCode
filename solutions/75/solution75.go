package solution75

func sortColors(nums []int) {

	var onesCount int
	var twosCount int
	var zerosCount int

	for i := 0; i < len(nums); i++ {
		switch nums[i] {
		case 0:
			zerosCount++
		case 1:
			onesCount++
		case 2:
			twosCount++
		}
	}

	for i := 0; i < len(nums); i++ {
		switch {
		case zerosCount > 0:
			nums[i] = 0
			zerosCount--
		case onesCount > 0:
			nums[i] = 1
			onesCount--
		case twosCount > 0:
			nums[i] = 2
			twosCount--
		}
	}
}
