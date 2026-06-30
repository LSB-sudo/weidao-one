//go:build linux

package gps

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

func openSerial(device string, baud int) (*os.File, error) {
	speed, err := baudAttr(baud)
	if err != nil {
		return nil, err
	}

	serial, err := os.OpenFile(device, os.O_RDWR|syscall.O_NOCTTY|syscall.O_NONBLOCK, 0)
	if err != nil {
		return nil, err
	}
	if err := configureSerial(serial.Fd(), speed); err != nil {
		_ = serial.Close()
		return nil, err
	}
	if err := syscall.SetNonblock(int(serial.Fd()), false); err != nil {
		_ = serial.Close()
		return nil, err
	}
	return serial, nil
}

func configureSerial(fd uintptr, speed uint32) error {
	termios := &syscall.Termios{}
	if _, _, errno := syscall.Syscall6(syscall.SYS_IOCTL, fd, uintptr(syscall.TCGETS), uintptr(unsafe.Pointer(termios)), 0, 0, 0); errno != 0 {
		return errno
	}

	termios.Iflag = 0
	termios.Oflag = 0
	termios.Lflag = 0
	termios.Cflag &^= syscall.CSIZE | syscall.PARENB | syscall.CSTOPB
	termios.Cflag |= syscall.CREAD | syscall.CLOCAL | syscall.CS8
	termios.Ispeed = speed
	termios.Ospeed = speed
	termios.Cc[syscall.VMIN] = 1
	termios.Cc[syscall.VTIME] = 0

	if _, _, errno := syscall.Syscall6(syscall.SYS_IOCTL, fd, uintptr(syscall.TCSETS), uintptr(unsafe.Pointer(termios)), 0, 0, 0); errno != 0 {
		return errno
	}
	return nil
}

func baudAttr(baud int) (uint32, error) {
	switch baud {
	case 4800:
		return syscall.B4800, nil
	case 9600:
		return syscall.B9600, nil
	case 19200:
		return syscall.B19200, nil
	case 38400:
		return syscall.B38400, nil
	case 57600:
		return syscall.B57600, nil
	case 115200:
		return syscall.B115200, nil
	default:
		return 0, fmt.Errorf("unsupported baud rate %d", baud)
	}
}
