package xgc

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	MaxControllerNum = 4 // support 4 controller at one time
)

var (
	LTDeadZone = 50
	LSDeadZone = 50
)

var (
	LoadError                 error
	xinputDLL                 *windows.DLL
	procXInputEnable          *windows.Proc
	procXInputGetCapabilities *windows.Proc
	procXInputGetKeystroke    *windows.Proc
	procXInputGetState        *windows.Proc
	procXInputSetState        *windows.Proc
)

func init() {
	if xinputDLL, LoadError = windows.LoadDLL("Xinput1_4.dll"); LoadError != nil {
		fmt.Println("failed to load dll: XInput")
		return
	}
	if procXInputEnable, LoadError = xinputDLL.FindProc("XInputEnable"); LoadError != nil {
		fmt.Println("failed to load proc func: XInputEnable")
		return
	}
	if procXInputGetCapabilities, LoadError = xinputDLL.FindProc("XInputGetCapabilities"); LoadError != nil {
		fmt.Println("failed to load proc func: XInputGetCapabilities")
		return
	}
	if procXInputGetKeystroke, LoadError = xinputDLL.FindProc("XInputGetKeystroke"); LoadError != nil {
		fmt.Println("failed to load proc func: XInputGetKeystroke")
		return
	}
	if procXInputGetState, LoadError = xinputDLL.FindProc("XInputGetState"); LoadError != nil {
		fmt.Println("failed to load proc func: XInputGetState")
		return
	}
	if procXInputSetState, LoadError = xinputDLL.FindProc("XInputSetState"); LoadError != nil {
		fmt.Println("failed to load proc func: XInputSetState")
		return
	}
}

type Joystick interface {
	Name() string
	GetState() (interface{}, error)
	SetState(LeftMotorSpeed, RightMotorSpeed uint16) error
	Close()
}

type xgcImpl struct {
	id        int
	name      string
	state     XINPUT_STATE
	vibration XINPUT_VIBRATION
}

func OpenXGC(id int) (Joystick, error) {
	if LoadError != nil {
		return nil, LoadError
	}
	xgc := &xgcImpl{}
	xgc.id = id
	r, _, _ := procXInputEnable.Call(
		uintptr(1))
	if r == 0 {
		return xgc, nil
	}
	return nil, syscall.Errno(r)
}

func (xgc *xgcImpl) GetCap() error {
	r, _, _ := procXInputGetCapabilities.Call(
		uintptr(xgc.id),
		uintptr(xgc.id),
		uintptr(unsafe.Pointer(&xgc.state)))
	if r == 0 {
		return nil
	}
	return syscall.Errno(r)
}

func (xgc *xgcImpl) Name() string {
	return xgc.name
}

func (xgc *xgcImpl) GetState() (interface{}, error) {
	r, _, _ := procXInputGetState.Call(
		uintptr(xgc.id),
		uintptr(unsafe.Pointer(&xgc.state)))
	if r == 0 {
		return xgc.state, nil
	}
	return xgc.state, syscall.Errno(r)
}

func (xgc *xgcImpl) SetState(LeftMotorSpeed, RightMotorSpeed uint16) error {
	xgc.vibration.wLeftMotorSpeed = WORD(LeftMotorSpeed)
	xgc.vibration.wRightMotorSpeed = WORD(RightMotorSpeed)
	fmt.Println(xgc)
	r, _, _ := procXInputSetState.Call(
		uintptr(xgc.id),
		uintptr(unsafe.Pointer(&xgc.vibration)))
	if r == 0 {
		return nil
	}
	return syscall.Errno(r)
}

func (xgc *xgcImpl) Close() {
	// no impl under windows
}
