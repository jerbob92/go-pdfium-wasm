package imports

import (
	"log"
	"os"
)

func sys_mprotect(_ int32, _ int32, _ int32) int32 {
	log.Fatal("Called into __sys_mprotect")
	return 0
}

func sys_madvise1(_ int32, _ int32, _ int32) int32 {
	log.Fatal("Called into __sys_madvise1")
	// JS implementation says: advice is welcome, but ignored
	return 0
}

func sys_fstat64(_ int32, _ int32) int32 {
	log.Fatal("Called into __sys_fstat64")
	return 0
}

func sys_stat64(_ int32, _ int32) int32 {
	log.Fatal("Called into __sys_stat64")
	return 0
}

func sys_lstat64(_ int32, _ int32) int32 {
	log.Fatal("Called into sys_lstat64")
	return 0
}

func sys_ftruncate64(_ int32, _ int64) int32 {
	log.Fatal("Called into __sys_ftruncate64")
	return 0
}

func sys_getpid() int32 {
	return int32(os.Getpid())
}

func sys_getdents64(_ int32, _ int32, _ int32) int32 {
	log.Fatal("Called into __sys_getdents64")
	return 0
}

func sys_unlink(_ int32) int32 {
	log.Fatal("Called into __sys_unlink")
	return 0
}

func sys_unlinkat(_ int32, _ int32, _ int32) int32 {
	log.Fatal("Called into __sys_unlinkat")
	return 0
}

func sys_rmdir(_ int32) int32 {
	log.Fatal("Called into __sys_rmdir")
	return 0
}

func syscall_newfstatat(int32, int32, int32, int32) int32 {
	log.Fatal("Called into __syscall_newfstatat")
	return 0
}

func abort() {
	log.Fatal("Called into abort")
}
func time(int32) int32 {
	log.Fatal("Called into time")
	return 0
}

func gettimeofday(int32, int32) int32 {
	log.Fatal("Called into gettimeofday")
	return 0
}

func sys_open(int32, int32, int32) int32 {
	log.Fatal("Called into sys_open")
	return 0
}

func sys_fcntl64(int32, int32, int32) int32 {
	log.Fatal("Called into sys_fcntl64")
	return 0
}

func sys_ioctl(int32, int32, int32) int32 {
	log.Fatal("Called into sys_ioctl")
	return 0
}

func sys_mmap2(int32, int32, int32, int32, int32, int32) int32 {
	log.Fatal("Called into sys_mmap2")
	return 0
}

func sys_munmap(int32, int32) int32 {
	log.Fatal("Called into sys_munmap")
	return 0
}

func gmtime_r(int32, int32) int32 {
	log.Fatal("Called into gmtime_r")
	return 0
}

func localtime_r(int32, int32) int32 {
	log.Fatal("Called into localtime_r")
	return 0
}

func strftime_l(int32, int32, int32, int32, int32) int32 {
	log.Fatal("Called into strftime_l")
	return 0
}

func emscripten_resize_heap(int32) int32 {
	log.Fatal("Called into emscripten_resize_heap")
	return 0
}

func emscripten_memcpy_big(int32, int32, int32) int32 {
	log.Fatal("Called into emscripten_memcpy_big")
	return 0
}

func FXSYS_timePx(int32) int64 { return 0 }

func FXSYS_localtimePKx(int32) int32 { return 0 }
