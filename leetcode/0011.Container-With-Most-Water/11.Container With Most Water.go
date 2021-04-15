package leetcode

func maxArea(height []int) int {
	var i, j int = 0, len(height) - 1
	var finalArea int
	for i < j {
		width := j - i
		var h int

		if height[i] < height[j] {
			h = height[i]
			i++
		} else {
			h = height[j]
			j--
		}

		currArea := width * h
		if currArea > finalArea {
			finalArea = currArea
		}
	}
	return finalArea
}
