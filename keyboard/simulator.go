package keyboard

import (
	"github.com/SpicyChickenFLY/xinput2mouse/event"
	"github.com/SpicyChickenFLY/xinput2mouse/utils"
	"github.com/SpicyChickenFLY/xinput2mouse/xgc"
	"github.com/micmonay/keybd_event"
	"github.com/nsf/termbox-go"
)

//Simulator is the struct of keyboard simulator
type Simulator struct {
	ltPulled, rtPulled bool
	buttons            uint16
	lSec, rSec         int
	alphabetDict       *[][]string
}

// NewSimulator init new keyboard simulator
func NewSimulator() *Simulator {
	return &Simulator{false, false, 0, -1, -1, nil}
}

// Update xinput event
func (s *Simulator) Update(xg *xgc.XinputGamepad) error {
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
	return s.render()
}

func (s *Simulator) render() error {
	if err := utils.Clear(0, 0); err != nil {
		return err
	}
	utils.TBPrint(0, 0, 0, termbox.ColorRed, ">>keyboard")
	for i, sec := range *s.alphabetDict {
		if i == s.lSec {
			utils.TBPrint(i*3, 1, 0, termbox.ColorRed, sec[0])
			if s.rSec >= 0 {
				for j, str := range (*s.alphabetDict)[s.lSec] {
					if j == s.rSec {
						utils.TBPrint(j*3, 2, 0, termbox.ColorRed, str)
					} else {
						utils.TBPrint(j*3, 2, 0, 0, str)
					}
				}
			}
		} else {
			utils.TBPrint(i*3, 1, 0, 0, sec[0])
		}
	}

	utils.Flush()
	return nil
}

// HandleEvent which manager dispatched to keyborad event
func (s *Simulator) HandleEvent(eventIdx uint16) error {
	// handle thumb output
	if err := s.handleThumb(eventIdx); err != nil {
		return err
	}
	switch {
	case eventIdx&xgc.XinputGamepadDpad > 0:
		return s.handleDpad(eventIdx)
	case eventIdx&xgc.XinputGamepadMain > 0:
		return s.handleMain(eventIdx)
	case eventIdx&xgc.XinputGamepadFunc > 0:
		return s.handleFunc(eventIdx)
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
	if xg.JudgeThumbRPulled() && s.lSec != -1 {
		rx := float64(xg.ThumbRX) / xgc.ThumbMax
		ry := float64(xg.ThumbRY) / xgc.ThumbMax
		rl := len((*s.alphabetDict)[s.lSec])
		s.rSec = judgePosSection(rx, ry, rl)
	} else {
		s.rSec = -1
	}
}

func (s *Simulator) handleThumb(eventIdx uint16) error {
	kbBond, err := keybd_event.NewKeyBonding()
	if err == nil {
		return err
	}
	// if Right Trigger pulled, judge which key should be simulated
	if s.rtPulled && s.lSec > 0 && s.rSec > 0 {
		keyVal := (*s.alphabetDict)[s.lSec][s.rSec]
		kbBond.AddKey(event.KeyInfoMap[keyVal].VK)
		kbBond.HasSHIFT(
			event.KeyInfoMap[keyVal].Flag&event.KBEventHasShift > 0)
		return kbBond.Launching()
	}
	return nil
}

func (s *Simulator) handleDpad(eventIdx uint16) error {
	kbBond, err := keybd_event.NewKeyBonding()
	if err == nil {
		return err
	}
	if s.ltPulled {
		for _, vk := range dpadKeysDict[0][eventIdx] {
			kbBond.AddKey(event.KeyInfoMap[vk].VK)
		}
	} else {
		for _, vk := range dpadKeysDict[1][eventIdx] {
			kbBond.AddKey(event.KeyInfoMap[vk].VK)
		}
	}
	return kbBond.Launching()
}

func (s *Simulator) handleMain(eventIdx uint16) error {
	kbBond, err := keybd_event.NewKeyBonding()
	if err == nil {
		return err
	}
	if s.ltPulled {
		for _, vk := range mainKeysDict[0][eventIdx] {
			kbBond.AddKey(event.KeyInfoMap[vk].VK)
		}
		switch eventIdx {
		case xgc.XinputGamepadY:
			kbBond.HasSHIFT(true)
		}
	} else {
		for _, vk := range mainKeysDict[0][eventIdx] {
			kbBond.AddKey(event.KeyInfoMap[vk].VK)
		}
	}
	return kbBond.Launching()
}

func (s *Simulator) handleFunc(eventIdx uint16) error {
	kbBond, err := keybd_event.NewKeyBonding()
	if err == nil {
		return err
	}
	return kbBond.Launching()
}
