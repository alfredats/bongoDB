package main

/*
#include <stdlib.h>
#include <string.h>
#include <dlfcn.h>

static inline int bongo_stub() { return 42; }
int call_init(void* fptr_init) {
	return ((int (*)())fptr_init)();
};

int call_create(void* fptr_create, const char* k, const char* v) {
	return ((int (*)(const char*, const char*))fptr_create)(k, v);
};

int call_read(void* fptr_read, const char* k, char* v, int v_size) {
	return ((int (*)(const char*, char*, int))fptr_read)(k, v, v_size);
};

int call_update(void* fptr_update, const char* k, const char* v) {
	return ((int (*)(const char*, const char*))fptr_update)(k, v);
};

int call_delete(void* fptr_delete, const char* k) {
	return ((int (*)(const char*))fptr_delete)(k);
};
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func BongoStub() int {
	return (int)(C.bongo_stub())
}

type BongoAPIs struct {
	Init   unsafe.Pointer
	Create unsafe.Pointer
	Read   unsafe.Pointer
	Update unsafe.Pointer
	Delete unsafe.Pointer
}

type BongoHandle struct {
	handle unsafe.Pointer
	apis   BongoAPIs
}

func (bh *BongoHandle) loadLibrary(libPath string) error {
	if bh.handle != nil {
		return nil // Library already loaded
	}
	cLibPath := C.CString(libPath)
	defer C.free(unsafe.Pointer(cLibPath))

	bh.handle = C.dlopen(cLibPath, C.RTLD_LAZY)
	if bh.handle == nil {
		return fmt.Errorf("failed to load library: %s", C.GoString(C.dlerror()))
	}
	return nil
}

func (bh *BongoHandle) getSymbol(symbolName string) (unsafe.Pointer, error) {
	cSymbolName := C.CString(symbolName)
	defer C.free(unsafe.Pointer(cSymbolName))

	symbol := C.dlsym(bh.handle, cSymbolName)
	if symbol == nil {
		return nil, fmt.Errorf("failed to get symbol: %s", C.GoString(C.dlerror()))
	}
	return symbol, nil
}

func (bh *BongoHandle) Init(libPath string) error {
	err := bh.loadLibrary(libPath)
	if err != nil {
		return err
	}

	bh.apis = BongoAPIs{}
	bh.apis.Init, err = bh.getSymbol("bongoDB_init")
	if err != nil {
		return err
	}
	bh.apis.Create, err = bh.getSymbol("bongoDB_create")
	if err != nil {
		return err
	}
	bh.apis.Read, err = bh.getSymbol("bongoDB_read")
	if err != nil {
		return err
	}
	bh.apis.Update, err = bh.getSymbol("bongoDB_update")
	if err != nil {
		return err
	}
	bh.apis.Delete, err = bh.getSymbol("bongoDB_delete")
	if err != nil {
		return err
	}

	C.call_init(bh.apis.Init)

	return nil
}

func (bh *BongoHandle) Create(key, value string) int {
	cKey := C.CString(key)
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cKey))
	defer C.free(unsafe.Pointer(cValue))

	return int(C.call_create(bh.apis.Create, cKey, cValue))
}

func (bh *BongoHandle) Read(key string) (string, int) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	// Assuming a maximum value size of 1024 for demonstration purposes
	valueSize := 1024
	cValue := (*C.char)(C.malloc(C.size_t(valueSize)))
	defer C.free(unsafe.Pointer(cValue))

	result := int(C.call_read(bh.apis.Read, cKey, cValue, C.int(valueSize)))
	if result != 0 {
		return "", result
	}

	value := C.GoString(cValue)
	return value, 0
}

func (bh *BongoHandle) Update(key, value string) int {
	cKey := C.CString(key)
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cKey))
	defer C.free(unsafe.Pointer(cValue))

	return int(C.call_update(bh.apis.Update, cKey, cValue))
}

func (bh *BongoHandle) Delete(key string) int {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	return int(C.call_delete(bh.apis.Delete, cKey))
}

func (bh *BongoHandle) Close() error {
	if bh.handle != nil {
		if C.dlclose(bh.handle) != 0 {
			return fmt.Errorf("failed to close library: %s", C.GoString(C.dlerror()))
		}
		bh.handle = nil
	}
	return nil
}
