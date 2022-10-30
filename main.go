package main

import (
	"context"
	"crypto/rand"
	_ "embed"
	"log"
	"os"

	"jerbob92/go-pdfium-wasm/imports"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/experimental"
	"github.com/tetratelabs/wazero/experimental/logging"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

//go:embed pdfium.wasm
var pdfiumWasm []byte

// main shows how to instantiate the same module name multiple times in the same runtime.
//
// See README.md for a full description.
func main() {
	// Choose the context to use for function calls.
	// Set context to one that has an experimental listener
	ctx := context.WithValue(context.Background(), experimental.FunctionListenerFactoryKey{}, logging.NewLoggingListenerFactory(os.Stdout))

	//ctx := context.Background()
	// Create a new WebAssembly Runtime.
	r := wazero.NewRuntimeWithConfig(ctx, wazero.NewRuntimeConfigInterpreter())
	defer r.Close(ctx) // This closes everything this Runtime created.

	if _, err := wasi_snapshot_preview1.Instantiate(ctx, r); err != nil {
		log.Panicln(err)
	}

	// Add missing emscripten and syscalls
	if _, err := imports.Instantiate(ctx, r); err != nil {
		log.Panicln(err)
	}

	// Compile WebAssembly that requires its own "env" module.
	compiled, err := r.CompileModule(ctx, pdfiumWasm)
	if err != nil {
		log.Panicln(err)
	}

	mod, err := r.InstantiateModule(ctx, compiled, wazero.NewModuleConfig().WithStdout(os.Stdout).WithStderr(os.Stderr).WithRandSource(rand.Reader).WithFS(os.DirFS("")))
	if err != nil {
		log.Panicln(err)
	}

	// We need to call this Emscripten _initialize because PDFium has no main method.
	_, err = mod.ExportedFunction("_initialize").Call(ctx)
	if err != nil {
		log.Panicln(err)
	}

	openRet, err := mod.ExportedFunction("FPDF_InitLibrary").Call(ctx)
	if err != nil {
		log.Panicln(err)
	}

	log.Println(openRet)

	//malloc := mod.ExportedFunction("malloc")
	//free := mod.ExportedFunction("free")

	/*
		filePath := "test.pdf"
		filePathSize := uint64(len(filePath))

		results, err := malloc.Call(ctx, filePathSize)
		if err != nil {
			log.Panicln(err)
		}
		filePathPtr := results[0]
		defer free.Call(ctx, filePathPtr, filePathSize)

		// The pointer is a linear memory offset, which is where we write the name.
		if !mod.Memory().Write(ctx, uint32(filePathPtr), []byte(filePath)) {
			log.Panicf("Memory.Write(%d, %d) out of range of memory size %d",
				filePathPtr, filePathSize, mod.Memory().Size(ctx))
		}

		doc, err := mod.ExportedFunction("FPDF_LoadDocument").Call(ctx, filePathPtr, 0)
		if err != nil {
			log.Panicln(err)
		}*/

	/*
		fileData, err := ioutil.ReadFile("test.pdf")
		if err != nil {
			log.Panicln(err)
		}

		results, err := malloc.Call(ctx, uint64(len(fileData)))
		if err != nil {
			log.Panicln(err)
		}
		filePathPtr := results[0]
		defer free.Call(ctx, filePathPtr, uint64(len(fileData)))

		// The pointer is a linear memory offset, which is where we write the name.
		if !mod.Memory().Write(ctx, uint32(filePathPtr), fileData) {
			log.Panicf("Memory.Write(%d, %d) out of range of memory size %d",
				filePathPtr, uint64(len(fileData)), mod.Memory().Size(ctx))
		}

		doc, err := mod.ExportedFunction("FPDF_LoadMemDocument").Call(ctx, filePathPtr, uint64(len(fileData)), 0)
		if err != nil {
			log.Panicln(err)
		}

		log.Println("Got doc")
		log.Println(doc)

		errorCode, err := mod.ExportedFunction("FPDF_GetLastError").Call(ctx)
		if err != nil {
			log.Panicln(err)
		}

		log.Println("Errorcode")
		log.Println(errorCode)*/
}
