package main

import (
    "errors"
    "flag"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
    "syscall"
    "time"
    "unsafe"
)

var (
    resetFrame  = []byte{0x7E, 0xFF, 0x06, 0x0C, 0x00, 0x00, 0x00, 0xFE, 0xEF, 0xEF}
    volumeFrame = []byte{0x7E, 0xFF, 0x06, 0x06, 0x00, 0x00, 0x14, 0xFE, 0xE1, 0xEF}
    stopFrame   = []byte{0x7E, 0xFF, 0x06, 0x16, 0x00, 0x00, 0x00, 0xFE, 0xE5, 0xEF}
)

func main() {
    device := flag.String("device", "", "serial device path; empty means auto-detect")
    baud := flag.Int("baud", 9600, "serial baud rate")
    playDuration := flag.Duration("play-duration", 5*time.Second, "how long to play each track before sending stop")
    cycles := flag.Int("cycles", 1, "how many times to loop tracks 1..3")
    flag.Parse()

    resolved, err := resolveDevice(*device)
    if err != nil {
        log.Fatal(err)
    }

    serial, err := openSerial(resolved, *baud)
    if err != nil {
        log.Fatalf("open serial %s: %v", resolved, err)
    }
    defer serial.Close()

    log.Printf("using serial device %s at %d 8N1", resolved, *baud)

    send(serial, "reset", resetFrame)
    time.Sleep(200 * time.Millisecond)

    send(serial, "volume=20", volumeFrame)
    time.Sleep(100 * time.Millisecond)

    for cycle := 1; cycle <= *cycles; cycle++ {
        for track := 1; track <= 3; track++ {
            frame := trackFrame(track)
            send(serial, fmt.Sprintf("play track %04d", track), frame)
            log.Printf("waiting %s for track %04d", playDuration.String(), track)
            time.Sleep(*playDuration)
            send(serial, fmt.Sprintf("stop track %04d", track), stopFrame)
            time.Sleep(300 * time.Millisecond)
        }
        log.Printf("completed cycle %d/%d", cycle, *cycles)
    }
}

func send(serial *os.File, label string, frame []byte) {
    n, err := serial.Write(frame)
    if err != nil {
        log.Fatalf("send %s failed: %v", label, err)
    }
    if n != len(frame) {
        log.Fatalf("send %s short write: %d/%d", label, n, len(frame))
    }
    log.Printf("sent %s: % X", label, frame)
}

func trackFrame(track int) []byte {
    if track < 1 || track > 0xFF {
        log.Fatalf("unsupported track number %d", track)
    }
    cmd := byte(0x12)
    feedback := byte(0x00)
    high := byte(0x00)
    low := byte(track)
    checksum := checksum(cmd, feedback, high, low)
    return []byte{0x7E, 0xFF, 0x06, cmd, feedback, high, low, byte(checksum >> 8), byte(checksum), 0xEF}
}

func checksum(cmd, feedback, high, low byte) uint16 {
    sum := uint16(0xFF) + uint16(0x06) + uint16(cmd) + uint16(feedback) + uint16(high) + uint16(low)
    return 0 - sum
}

func resolveDevice(explicit string) (string, error) {
    if strings.TrimSpace(explicit) != "" {
        return explicit, nil
    }

    candidates := []string{
        "/dev/serial/by-id/usb-1a86_USB_Serial-if00-port0",
        "/dev/ttyUSB0",
        "/dev/ttyUSB1",
        "/dev/ttyACM0",
    }
    for _, candidate := range candidates {
        if _, err := os.Stat(candidate); err == nil {
            return candidate, nil
        }
    }

    matches, _ := filepath.Glob("/dev/ttyUSB*")
    if len(matches) > 0 {
        return matches[0], nil
    }
    matches, _ = filepath.Glob("/dev/ttyACM*")
    if len(matches) > 0 {
        return matches[0], nil
    }
    return "", errors.New("no serial device found for DFPlayer")
}

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
