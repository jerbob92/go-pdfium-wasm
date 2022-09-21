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
