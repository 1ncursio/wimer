package main

import (
	"context"
	"fmt"
	"log"
	"syscall"
	"unsafe"

	"github.com/lextoumbourou/idle"
	"github.com/mitchellh/go-ps"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) PS() {

	ps, _ := ps.Processes()
	fmt.Println(ps[0].Executable())

	for pp := range ps {
		fmt.Printf("%d %s\n", ps[pp].Pid(), ps[pp].Executable())
	}
}

type HANDLE uintptr
type HWND HANDLE
type Response struct {
	WindowText string `json:"windowText"`
	ClassName  string `json:"className"`
	Hwnd       HWND   `json:"hwnd"`
}

var user32 = syscall.NewLazyDLL("user32.dll")

func (a *App) Test() Response {
	hwnd := GetForegroundWindow()
	className := GetClassName(hwnd)
	windowText := GetWindowText(hwnd)

	reponse := Response{
		WindowText: windowText,
		ClassName:  className,
		Hwnd:       hwnd,
	}

	return reponse
}

func GetForegroundWindow() HWND {
	hwnd, _, _ := user32.NewProc("GetForegroundWindow").Call()

	return HWND(hwnd)
}

func GetWindowText(hwnd HWND) string {
	length := GetWindowTextLength(hwnd) + 1
	buffer := make([]uint16, length)
	user32.NewProc("GetWindowTextW").Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(length),
	)

	return syscall.UTF16ToString(buffer)
}

func GetWindowTextLength(hwnd HWND) int {
	ret, _, _ := user32.NewProc("GetWindowTextLengthW").Call(uintptr(hwnd))

	return int(ret)
}

func GetClassName(hwnd HWND) string {
	// Get the window class name
	buf := make([]uint16, 256)
	user32.NewProc("GetClassNameW").Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)),
	)

	return syscall.UTF16ToString(buf)
}

func (a *App) GetIdleTime() int {
	idleTime, err := idle.Get()

	if err != nil {
		log.Fatal(err)
	}

	return int(idleTime.Seconds())
}
