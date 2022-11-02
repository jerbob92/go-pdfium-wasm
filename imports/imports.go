package imports

import (
	"context"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

// Instantiate instantiates the "env" module used by Emscripten into the
// runtime default namespace.
//
// # Notes
//
//   - Closing the wazero.Runtime has the same effect as closing the result.
//   - To add more functions to the "env" module, use FunctionExporter.
//   - To instantiate into another wazero.Namespace, use FunctionExporter.
func Instantiate(ctx context.Context, r wazero.Runtime) (api.Closer, error) {
	builder := r.NewHostModuleBuilder("env")
	NewFunctionExporter().ExportFunctions(builder)
	return builder.Instantiate(ctx, r)
}

// FunctionExporter configures the functions in the "env" module used by
// Emscripten.
type FunctionExporter interface {
	// ExportFunctions builds functions to export with a wazero.HostModuleBuilder
	// named "env".
	ExportFunctions(builder wazero.HostModuleBuilder)
}

// NewFunctionExporter returns a FunctionExporter object with trace disabled.
func NewFunctionExporter() FunctionExporter {
	return &functionExporter{}
}

type functionExporter struct{}

// ExportFunctions implements FunctionExporter.ExportFunctions
func (e *functionExporter) ExportFunctions(b wazero.HostModuleBuilder) {
	b.NewFunctionBuilder().WithFunc(emscripten_notify_memory_growth).Export("emscripten_notify_memory_growth")
	b.NewFunctionBuilder().WithFunc(setTempRet0).Export("setTempRet0")
	b.NewFunctionBuilder().WithFunc(getTempRet0).Export("getTempRet0")
	b.NewFunctionBuilder().WithFunc(emscripten_throw_longjmp).Export("_emscripten_throw_longjmp")

	b.NewFunctionBuilder().WithFunc(invoke_ii).Export("invoke_ii")
	b.NewFunctionBuilder().WithFunc(invoke_iii).Export("invoke_iii")
	b.NewFunctionBuilder().WithFunc(invoke_iiii).Export("invoke_iiii")
	b.NewFunctionBuilder().WithFunc(invoke_iiiii).Export("invoke_iiiii")

	b.NewFunctionBuilder().WithFunc(invoke_v).Export("invoke_v")
	b.NewFunctionBuilder().WithFunc(invoke_vi).Export("invoke_vi")
	b.NewFunctionBuilder().WithFunc(invoke_vii).Export("invoke_vii")
	b.NewFunctionBuilder().WithFunc(invoke_viii).Export("invoke_viii")
	b.NewFunctionBuilder().WithFunc(invoke_viiii).Export("invoke_viiii")

	b.NewFunctionBuilder().WithFunc(sys_mprotect).Export("__syscall_mprotect")
	b.NewFunctionBuilder().WithFunc(sys_mprotect).Export("__sys_mprotect")
	b.NewFunctionBuilder().WithFunc(sys_madvise1).Export("__syscall_madvise1")
	b.NewFunctionBuilder().WithFunc(sys_madvise1).Export("__sys_madvise1")
	b.NewFunctionBuilder().WithFunc(sys_fstat64).Export("__syscall_fstat64")
	b.NewFunctionBuilder().WithFunc(sys_fstat64).Export("__sys_fstat64")
	b.NewFunctionBuilder().WithFunc(sys_stat64).Export("__syscall_stat64")
	b.NewFunctionBuilder().WithFunc(sys_stat64).Export("__sys_stat64")
	b.NewFunctionBuilder().WithFunc(sys_lstat64).Export("__syscall_lstat64")
	b.NewFunctionBuilder().WithFunc(sys_ftruncate64).Export("__syscall_ftruncate64")
	b.NewFunctionBuilder().WithFunc(sys_ftruncate64).Export("__sys_ftruncate64")
	b.NewFunctionBuilder().WithFunc(sys_getpid).Export("__syscall_getpid")
	b.NewFunctionBuilder().WithFunc(sys_getpid).Export("__sys_getpid")
	b.NewFunctionBuilder().WithFunc(sys_getdents64).Export("__syscall_getdents64")
	b.NewFunctionBuilder().WithFunc(sys_getdents64).Export("__sys_getdents64")
	b.NewFunctionBuilder().WithFunc(sys_unlink).Export("__syscall_unlink")
	b.NewFunctionBuilder().WithFunc(sys_unlink).Export("__sys_unlink")
	b.NewFunctionBuilder().WithFunc(sys_unlinkat).Export("__syscall_unlinkat")
	b.NewFunctionBuilder().WithFunc(sys_rmdir).Export("__syscall_rmdir")
	b.NewFunctionBuilder().WithFunc(sys_rmdir).Export("__sys_rmdir")
	b.NewFunctionBuilder().WithFunc(syscall_newfstatat).Export("__syscall_newfstatat")
	b.NewFunctionBuilder().WithFunc(sys_open).Export("__sys_open")
	b.NewFunctionBuilder().WithFunc(sys_fcntl64).Export("__sys_fcntl64")
	b.NewFunctionBuilder().WithFunc(sys_ioctl).Export("__sys_ioctl")
	b.NewFunctionBuilder().WithFunc(sys_mmap2).Export("__sys_mmap2")
	b.NewFunctionBuilder().WithFunc(sys_munmap).Export("__sys_munmap")
	b.NewFunctionBuilder().WithFunc(abort).Export("abort")
	b.NewFunctionBuilder().WithFunc(time).Export("time")
	b.NewFunctionBuilder().WithFunc(gettimeofday).Export("gettimeofday")
	b.NewFunctionBuilder().WithFunc(gmtime_r).Export("__gmtime_r")
	b.NewFunctionBuilder().WithFunc(localtime_r).Export("__localtime_r")
	b.NewFunctionBuilder().WithFunc(strftime_l).Export("strftime_l")
	b.NewFunctionBuilder().WithFunc(emscripten_resize_heap).Export("emscripten_resize_heap")
	b.NewFunctionBuilder().WithFunc(emscripten_memcpy_big).Export("emscripten_memcpy_big")
	b.NewFunctionBuilder().WithFunc(FXSYS_timePx).Export("_Z10FXSYS_timePx")
	b.NewFunctionBuilder().WithFunc(FXSYS_localtimePKx).Export("_Z15FXSYS_localtimePKx")
}
