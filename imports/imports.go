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

	b.NewFunctionBuilder().WithFunc(sys_ftruncate64).Export("__syscall_ftruncate64")
	b.NewFunctionBuilder().WithFunc(sys_ftruncate64).Export("__sys_ftruncate64")
	b.NewFunctionBuilder().WithFunc(sys_unlink).Export("__syscall_unlink")
	b.NewFunctionBuilder().WithFunc(sys_unlink).Export("__sys_unlink")
	b.NewFunctionBuilder().WithFunc(sys_unlinkat).Export("__syscall_unlinkat")
	b.NewFunctionBuilder().WithFunc(sys_rmdir).Export("__syscall_rmdir")
	b.NewFunctionBuilder().WithFunc(sys_rmdir).Export("__sys_rmdir")
}
