package keyboard

import "math"

// judgePosSection 计算所处第几个扇形（以圆形正上方为起点）
// x, y: 横纵坐标（代表摇杆对应的百分比）
// l: 代表该圆中存在几个扇形
func judgePosSection(x, y float64, l int) int {
	temp := math.Atanh(y / x)
	// 将极坐标正方向反转后逆时针旋转90度（求余用于溢出角度）
	degree := int(-temp+90) % 360
	interval := 360.0 / l
	return degree % interval
}
