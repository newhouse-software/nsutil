package nsutil

import (
	"fmt"

	"golang.org/x/sys/windows"
)

func ThereCanBeOnlyOne(uniqueID string) (func(), error) {
	if uniqueID == "" {
		return nil, fmt.Errorf("uniqueID is empty")
	}

	name := `Local\` + uniqueID

	namePtr, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return nil, err
	}

	handle, err := windows.CreateMutex(nil, false, namePtr)
	if err != nil {
		return nil, fmt.Errorf("%s already running.", uniqueID)
	}

	if windows.GetLastError() == windows.ERROR_ALREADY_EXISTS {
		windows.CloseHandle(handle)
		return nil, fmt.Errorf("%s already running", uniqueID)
	}

	release := func() {
		windows.ReleaseMutex(handle)
		windows.CloseHandle(handle)
	}

	return release, nil
}
