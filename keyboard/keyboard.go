package keyboard

import (
	"math"

	"github.com/SpicyChickenFLY/xinput2mouse/xgc"
	"github.com/micmonay/keybd_event"
)

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

//Simulator is the struct of keyboard simulator
type Simulator struct {
	kbBond             keybd_event.KeyBonding
	ltPulled, rtPulled bool
	buttons            uint16
	lSec, rSec         int
	alphabetDict       *[][]string
}

// NewSimulator init new keyboard simulator
func NewSimulator() *Simulator {
	if kbBond, err := keybd_event.NewKeyBonding(); err == nil {
		return &Simulator{kbBond, false, false, 0, -1, -1, nil}
	}
	return nil
}

// Handle xinput event
func (s *Simulator) Handle(xg *xgc.XinputGamepad) error {
	s.ltPulled = xg.JudgeLTPulled()
	s.rtPulled = xg.JudgeRTPulled()
	if s.ltPulled {
		s.alphabetDict = &alphabetHigherCaseDict
	} else {
		s.alphabetDict = &alphabetLoweraseDict
	}
	s.buttons = xg.Buttons
	if err := s.render(); err != nil {
		return err
	}
	return s.handle()
}

func (s *Simulator) render() error {
	return nil
}

func (s *Simulator) handle() error {
	// if Right Trigger pulled, judge which key should be simulated
	if s.rtPulled && s.lSec > 0 && s.rSec > 0 {
		keyVal := (*s.alphabetDict)[s.lSec][s.rSec]
		s.kbBond.AddKey(keyInfoMap[keyVal].VK)
		s.kbBond.HasSHIFT(keyInfoMap[keyVal].Flag&kbEventHasShift > 0)
		err := s.kbBond.Launching()
		if err != nil {
			return err
		}
	}
	// judge rest xinput by order
	if s.buttons&xgc.XinputGamepadDpad > 0 {
		if s.ltPulled {
			switch {
			case s.buttons&xgc.XinputGamepadDpadUp > 0:
				s.kbBond.AddKey(keyInfoMap["PAGEUP"].VK)
			case s.buttons&xgc.XinputGamepadDpadDown > 0:
				s.kbBond.AddKey(keyInfoMap["PAGEDOWN"].VK)
			case s.buttons&xgc.XinputGamepadDpadLeft > 0:
				s.kbBond.AddKey(keyInfoMap["HOME"].VK)
			case s.buttons&xgc.XinputGamepadDpadRight > 0:
				s.kbBond.AddKey(keyInfoMap["HOME"].VK)
			}
		} else {
			switch {
			case s.buttons&xgc.XinputGamepadDpadUp > 0:
				s.kbBond.AddKey(keyInfoMap["UP"].VK)
			case s.buttons&xgc.XinputGamepadDpadDown > 0:
				s.kbBond.AddKey(keyInfoMap["DOWN"].VK)
			case s.buttons&xgc.XinputGamepadDpadLeft > 0:
				s.kbBond.AddKey(keyInfoMap["LEFT"].VK)
			case s.buttons&xgc.XinputGamepadDpadRight > 0:
				s.kbBond.AddKey(keyInfoMap["RIGHT"].VK)
			}
		}
		err := s.kbBond.Launching()
		if err != nil {
			return err
		}
		return nil
	}

	if s.buttons&xgc.XinputGamepadMain > 0 {
		if s.ltPulled {
			switch {
			case s.buttons&xgc.XinputGamepadY > 0:
				s.kbBond.AddKey(keyInfoMap["TAB"].VK)
				s.kbBond.HasSHIFT(true)
			}
		} else {
			switch {
			case s.buttons&xgc.XinputGamepadA > 0:
				s.kbBond.AddKey(keyInfoMap["SPACE"].VK)
			case s.buttons&xgc.XinputGamepadB > 0:
				s.kbBond.AddKey(keyInfoMap["ENTER"].VK)
			case s.buttons&xgc.XinputGamepadX > 0:
				s.kbBond.AddKey(keyInfoMap["BACKSPACE"].VK)
			case s.buttons&xgc.XinputGamepadY > 0:
				s.kbBond.AddKey(keyInfoMap["TAB"].VK)
			}
		}
		s.kbBond.HasSHIFT(false)
		err := s.kbBond.Launching()
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}

func (s *Simulator) judgeLPosSec(xg *xgc.XinputGamepad) {
	if xg.JudgeThumbLPulled() {
		lx := float64(xg.ThumbLX) / xgc.ThumbMax
		ly := float64(xg.ThumbLY) / xgc.ThumbMax
		ll := len(*s.alphabetDict)
		s.lSec = judgePosSection(lx, ly, ll)
	} else {
		s.lSec = -1
	}

}

func (s *Simulator) judgeRPosSec(xg *xgc.XinputGamepad) {
	if xg.JudgeThumbRPulled() {
		rx := float64(xg.ThumbRX) / xgc.ThumbMax
		ry := float64(xg.ThumbRY) / xgc.ThumbMax
		rl := len(*s.alphabetDict)
		s.rSec = judgePosSection(rx, ry, rl)
	} else {
		s.rSec = -1
	}
}

// func (s *Simulator)
