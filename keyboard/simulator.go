package keyboard

import (
	"fmt"

	"github.com/SpicyChickenFLY/xinput2mouse/event"
	"github.com/SpicyChickenFLY/xinput2mouse/xgc"
	"github.com/micmonay/keybd_event"
)

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
	s.judgeLPosSec(xg)
	s.judgeRPosSec(xg)
	return s.handle()
}

func (s *Simulator) render() error {
	fmt.Println(*s.alphabetDict)
	if s.lSec > 0 {
		fmt.Println((*s.alphabetDict)[s.lSec])
		if s.rSec > 0 {
			fmt.Println((*s.alphabetDict)[s.lSec][s.rSec])
		}
	}

	return nil
}

func (s *Simulator) handle() error {
	// handle thumb output
	if err := s.handleThumb(); err != nil {
		return err
	}
	// judge rest xinput by order
	if err := s.handleDpad(); err != nil {
		return err
	}
	if err := s.handleMain(); err != nil {
		return err
	}
	return s.handleFunc()
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

func (s *Simulator) handleThumb() error {
	// if Right Trigger pulled, judge which key should be simulated
	if s.rtPulled && s.lSec > 0 && s.rSec > 0 {
		keyVal := (*s.alphabetDict)[s.lSec][s.rSec]
		s.kbBond.AddKey(event.KeyInfoMap[keyVal].VK)
		s.kbBond.HasSHIFT(event.KeyInfoMap[keyVal].Flag&event.KBEventHasShift > 0)
		err := s.kbBond.Launching()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Simulator) handleDpad() error {
	if s.buttons&xgc.XinputGamepadDpad > 0 {
		if s.ltPulled {
			switch {
			case s.buttons&xgc.XinputGamepadDpadUp > 0:
				s.kbBond.AddKey(event.KeyInfoMap["PAGEUP"].VK)
			case s.buttons&xgc.XinputGamepadDpadDown > 0:
				s.kbBond.AddKey(event.KeyInfoMap["PAGEDOWN"].VK)
			case s.buttons&xgc.XinputGamepadDpadLeft > 0:
				s.kbBond.AddKey(event.KeyInfoMap["HOME"].VK)
			case s.buttons&xgc.XinputGamepadDpadRight > 0:
				s.kbBond.AddKey(event.KeyInfoMap["HOME"].VK)
			}
		} else {
			switch {
			case s.buttons&xgc.XinputGamepadDpadUp > 0:
				s.kbBond.AddKey(event.KeyInfoMap["UP"].VK)
			case s.buttons&xgc.XinputGamepadDpadDown > 0:
				s.kbBond.AddKey(event.KeyInfoMap["DOWN"].VK)
			case s.buttons&xgc.XinputGamepadDpadLeft > 0:
				s.kbBond.AddKey(event.KeyInfoMap["LEFT"].VK)
			case s.buttons&xgc.XinputGamepadDpadRight > 0:
				s.kbBond.AddKey(event.KeyInfoMap["RIGHT"].VK)
			}
		}
		err := s.kbBond.Launching()
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (s *Simulator) handleMain() error {
	if s.buttons&xgc.XinputGamepadMain > 0 {
		if s.ltPulled {
			switch {
			case s.buttons&xgc.XinputGamepadY > 0:
				s.kbBond.AddKey(event.KeyInfoMap["TAB"].VK)
				s.kbBond.HasSHIFT(true)
			}
		} else {
			switch {
			case s.buttons&xgc.XinputGamepadA > 0:
				s.kbBond.AddKey(event.KeyInfoMap["SPACE"].VK)
			case s.buttons&xgc.XinputGamepadB > 0:
				s.kbBond.AddKey(event.KeyInfoMap["ENTER"].VK)
			case s.buttons&xgc.XinputGamepadX > 0:
				s.kbBond.AddKey(event.KeyInfoMap["BACKSPACE"].VK)
			case s.buttons&xgc.XinputGamepadY > 0:
				s.kbBond.AddKey(event.KeyInfoMap["TAB"].VK)
			}
		}
		s.kbBond.HasSHIFT(false)
		err := s.kbBond.Launching()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Simulator) handleFunc() error {
	return nil
}
