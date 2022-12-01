package imports

import (
	"context"
	"github.com/tetratelabs/wazero/api"
)

type FPDF_FILEACCESS_CB struct {
}

func (cb FPDF_FILEACCESS_CB) Call(ctx context.Context, mod api.Module, stack []uint64) {
	stack[0] = uint64(0)
	return
}

type FPDF_FILEWRITE_CB struct {
}

func (cb FPDF_FILEWRITE_CB) Call(ctx context.Context, mod api.Module, stack []uint64) {
	stack[0] = uint64(0)
	return
}
