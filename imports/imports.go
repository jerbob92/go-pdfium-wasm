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
	builder := r.NewModuleBuilder("env")
	NewFunctionExporter().ExportFunctions(builder)
	return builder.Instantiate(ctx, r)
}

// FunctionExporter configures the functions in the "env" module used by
// Emscripten.
type FunctionExporter interface {
	// ExportFunctions builds functions to export with a wazero.ModuleBuilder
	// named "env".
	ExportFunctions(builder wazero.ModuleBuilder)
}

// NewFunctionExporter returns a FunctionExporter object with trace disabled.
func NewFunctionExporter() FunctionExporter {
	return &functionExporter{}
}

type functionExporter struct{}

// ExportFunctions implements FunctionExporter.ExportFunctions
func (e *functionExporter) ExportFunctions(builder wazero.ModuleBuilder) {
	builder.ExportFunction("emscripten_notify_memory_growth", emscripten_notify_memory_growth)
	builder.ExportFunction("setTempRet0", setTempRet0)
	builder.ExportFunction("getTempRet0", getTempRet0)
	builder.ExportFunction("_emscripten_throw_longjmp", emscripten_throw_longjmp)

	builder.ExportFunction("invoke_ii", invoke_ii)
	builder.ExportFunction("invoke_iii", invoke_iii)
	builder.ExportFunction("invoke_iiii", invoke_iiii)
	builder.ExportFunction("invoke_iiiii", invoke_iiiii)

	builder.ExportFunction("invoke_v", invoke_v)
	builder.ExportFunction("invoke_vi", invoke_vi)
	builder.ExportFunction("invoke_vii", invoke_vii)
	builder.ExportFunction("invoke_viii", invoke_viii)
	builder.ExportFunction("invoke_viiii", invoke_viiii)

	builder.ExportFunction("__syscall_mprotect", sys_mprotect)
	builder.ExportFunction("__syscall_madvise1", sys_madvise1)
	builder.ExportFunction("__syscall_fstat64", sys_fstat64)
	builder.ExportFunction("__syscall_stat64", sys_stat64)
	builder.ExportFunction("__syscall_lstat64", sys_lstat64)
	builder.ExportFunction("__syscall_ftruncate64", sys_ftruncate64)
	builder.ExportFunction("__syscall_getpid", sys_getpid)
	builder.ExportFunction("__syscall_getdents64", sys_getdents64)
	builder.ExportFunction("__syscall_unlink", sys_unlink)
	builder.ExportFunction("__syscall_unlinkat", sys_unlinkat)
	builder.ExportFunction("__syscall_rmdir", sys_rmdir)
	builder.ExportFunction("__syscall_newfstatat", syscall_newfstatat)
}
