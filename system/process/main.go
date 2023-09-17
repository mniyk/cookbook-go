package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

func main() {
	handle, err := syscall.CreateToolhelp32Snapshot(syscall.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}
	defer syscall.CloseHandle(handle)

	var pe32 syscall.ProcessEntry32
	pe32.Size = uint32(unsafe.Sizeof(pe32))
	if err := syscall.Process32First(handle, &pe32); err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}

	for {
		fmt.Printf(
			"%d: %s\n",
			pe32.ProcessID,
			syscall.UTF16ToString(pe32.ExeFile[:]))
		if err := syscall.Process32Next(handle, &pe32); err != nil {
			break
		}
	}
}
