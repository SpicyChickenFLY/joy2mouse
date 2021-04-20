package keyboard

import (
	"errors"
	"fmt"
	"math"

	"github.com/SpicyChickenFLY/xinput2mouse/utils"
)

// judgePosSection 计算所处第几个扇形（以圆形正上方为起点）
// x, y: 横纵坐标（代表摇杆对应的百分比）
// l: 代表该圆中存在几个扇形
func judgePosSection(x, y float64, l int) int {
	// FIXME: How to translate (x,y) to degree properly
	temp := math.Atanh(y/x) * 180 / math.Pi
	// 将极坐标正方向反转后逆时针旋转90度（求余用于溢出角度）
	degree := int(-temp+90) % 360
	interval := 360.0 / l
	utils.TBPrint(20, 0, 0, 0, fmt.Sprintf("%.2f/%.2f %.2f %3d/%3d", x, y, temp, degree, interval))
	if degree/interval > l {
		errStr := fmt.Sprintf("t:%f,l:%d,d:%d,i:%d-%d", temp, l, degree, interval, degree/interval)
		panic(errors.New(errStr))
	}
	return degree / interval
}
