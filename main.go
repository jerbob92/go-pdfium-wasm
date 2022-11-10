package main

import (
	"context"
	"crypto/rand"
	_ "embed"
	"image"
	"image/jpeg"
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

	ctx = context.Background()
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

	compiled, err := r.CompileModule(ctx, pdfiumWasm)
	if err != nil {
		log.Panicln(err)
	}

	mod, err := r.InstantiateModule(ctx, compiled, wazero.NewModuleConfig().WithStartFunctions("_initialize").WithStdout(os.Stdout).WithStderr(os.Stderr).WithRandSource(rand.Reader).WithFS(os.DirFS("")))
	if err != nil {
		log.Panicln(err)
	}

	openRet, err := mod.ExportedFunction("FPDF_InitLibrary").Call(ctx)
	if err != nil {
		log.Panicln(err)
	}

	log.Println(openRet)

	malloc := mod.ExportedFunction("malloc")
	free := mod.ExportedFunction("free")

	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("could not get cwd: %v", err)
	}

	filePath := path + "/pdf-test.pdf"
	filePathSize := uint64(len(filePath)) + 1

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
	}

	log.Println("Got doc")
	log.Println(doc)

	errorCode, err := mod.ExportedFunction("FPDF_GetLastError").Call(ctx)
	if err != nil {
		log.Panicln(err)
	}

	log.Println("Errorcode")
	log.Println(errorCode)

	/*
		fileData, err := ioutil.ReadFile(filePath)
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
		log.Println(errorCode)
	*/

	rect := image.Rect(0, 0, 2000, 2000)

	// We do not allocate the memory yet since we take it directly from memory in WASM.
	img := &image.RGBA{
		Pix:    nil,
		Stride: 4 * rect.Dx(),
		Rect:   rect,
	}

	// RGBA = 4 bytes per pixel
	bufSize := 4 * rect.Dx() * rect.Dy()

	results, err = malloc.Call(ctx, uint64(bufSize))
	if err != nil {
		log.Panicln(err)
	}

	bitmap, err := mod.ExportedFunction("FPDFBitmap_CreateEx").Call(ctx, 2000, 2000, 4, results[0], uint64(img.Stride))
	if err != nil {
		log.Panicln(err)
	}

	log.Println("bitmap")
	log.Println(bitmap)

	fill, err := mod.ExportedFunction("FPDFBitmap_FillRect").Call(ctx, bitmap[0], 0, 0, 2000, 2000, 0xFFFFFFFF)
	if err != nil {
		log.Panicln(err)
	}

	log.Println("fill")
	log.Println(fill)

	//log.Println(mod.ExportedMemory("asm"))
	//log.Println(r.Module("asm").ExportedMemory("__indirect_function_table"))

	page, err := mod.ExportedFunction("FPDF_LoadPage").Call(ctx, doc[0], 0)
	if err != nil {
		log.Panicln(err)
	}

	log.Println("page")
	log.Println(page)

	render, err := mod.ExportedFunction("FPDF_RenderPageBitmap").Call(ctx, bitmap[0], page[0], 0, 0, 2000, 2000, 0, 0x10)
	if err != nil {
		log.Panicln(err)
	}

	log.Println("render")
	log.Println(render)

	b, _ := mod.Memory().Read(ctx, uint32(results[0]), uint32(bufSize))
	if err != nil {
		log.Panicln(err)
	}

	img.Pix = b

	f, err := os.Create("img.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err = jpeg.Encode(f, img, nil); err != nil {
		log.Printf("failed to encode: %v", err)
	}
}
