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
	// 将极坐标正方向反转后逆时针旋转90度（求余用于溢出角度）
	rad := rect2Polar(x, y)
	originDegree := rad2degree(rad)
	utils.TBPrint(40, 0, 0, 0, fmt.Sprintf("%.2f", originDegree))
	offsetDegree := int(originDegree)
	reverseDegree := offsetDegree*(-1) + 360
	degree := (reverseDegree + 90) % 360
	utils.TBPrint(60, 0, 0, 0, fmt.Sprintf("%8d", degree))
	interval := 360.0 / l

	if degree/interval > l {
		errStr := fmt.Sprintf("out of range: Thumb Degree")
		panic(errors.New(errStr))
	}
	return degree / interval
}

func rect2Polar(x, y float64) (rad float64) {
	if y > 0 {
		rad = math.Acos(x / math.Sqrt(x*x+y*y))
	} else {
		rad = math.Pi*2 - math.Acos(x/math.Sqrt(x*x+y*y))
	}
	return rad
}

func rad2degree(r float64) float64 {
	return r * 180 / math.Pi
}
