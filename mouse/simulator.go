package mouse

import (
	"github.com/SpicyChickenFLY/xinput2mouse/xgc"
	"github.com/micmonay/keybd_event"
)

//Simulator is the struct of keyboard simulator
type Simulator struct {
	kbBond             keybd_event.KeyBonding
	ltPulled, rtPulled bool
	buttons            uint16
}

// NewSimulator init new keyboard simulator
func NewSimulator() *Simulator {
	if kbBond, err := keybd_event.NewKeyBonding(); err == nil {
		return &Simulator{kbBond, false, false, 0}
	}
	return nil
}

// Handle xinput event
func (s *Simulator) Handle(xg *xgc.XinputGamepad) error {
	return nil
}

func (s *Simulator) render() error {
	return nil
}

func (s *Simulator) handle() error {
	return nil
}
