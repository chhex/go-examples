package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"log"
	"syscall"
	"unsafe"
)

// TH32CS_SNAPPROCESS is described in https://msdn.microsoft.com/de-de/library/windows/desktop/ms682489(v=vs.85).aspx
const TH32CS_SNAPPROCESS = 0x00000002

func main() {
	err := processes()
	if err != nil {
		log.Fatal(err)
	}
}


func processes() error {
	handle, err := windows.CreateToolhelp32Snapshot(TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return err
	}
	defer windows.CloseHandle(handle)

	var entry windows.ProcessEntry32
	entry.Size = uint32(unsafe.Sizeof(entry))
	// get the first process
	err = windows.Process32First(handle, &entry)
	if err != nil {
		return err
	}

	for {
		printProcess(&entry)

		err = windows.Process32Next(handle, &entry)
		if err != nil {
			// windows sends ERROR_NO_MORE_FILES on last process
			if err == syscall.ERROR_NO_MORE_FILES {
				return nil
			}
			return err
		}
	}
}

func printProcess(e *windows.ProcessEntry32)  {
	// Find when the string ends for decoding
	end := 0
	for {
		if e.ExeFile[end] == 0 {
			break
		}
		end++
	}
	fmt.Printf("ProcessId: %d, ParentProcessId: %d,  ExeFile %s\n", int(e.ProcessID), int(e.ParentProcessID),syscall.UTF16ToString(e.ExeFile[:end]))
}