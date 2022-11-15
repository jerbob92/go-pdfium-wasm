package main

// #cgo pkg-config: pdfium
// #include "fpdfview.h"
// #include "fpdf_edit.h"
// #include <stdlib.h>
import "C"
import (
	"image"
	"io/ioutil"
	"log"
	"os"
	"time"
	"unsafe"
)

func main() {
	C.FPDF_InitLibrary()

	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("could not get cwd: %v", err)
	}

	filePath := path + "/pdf-test.pdf"

	for i := 1; i < 10; i++ {
		start := time.Now()

		func() {
			var doc C.FPDF_DOCUMENT
			fromFile := false

			var cPassword *C.char

			if fromFile {
				filePath := C.CString(filePath)
				defer C.free(unsafe.Pointer(filePath))
				doc = C.FPDF_LoadDocument(
					filePath,
					cPassword)
			} else {
				fileData, err := ioutil.ReadFile(filePath)
				if err != nil {
					log.Panicln(err)
				}

				doc = C.FPDF_LoadMemDocument(
					unsafe.Pointer(&(fileData[0])),
					C.int(len(fileData)),
					cPassword)
			}

			if doc == nil {
				log.Fatal("Could not load document")
			}

			errorCode := C.FPDF_GetLastError()

			if errorCode != C.FPDF_ERR_SUCCESS {
				log.Fatalf("Could not load document: %d", errorCode)
			}

			width := 2000
			height := 2000

			rect := image.Rect(0, 0, width, height)

			// We do not allocate the memory yet since we take it directly from memory in WASM.
			img := image.NewRGBA(rect)

			bitmap := C.FPDFBitmap_CreateEx(C.int(width), C.int(height), C.int(4), unsafe.Pointer(&img.Pix[0]), C.int(img.Stride))

			if bitmap == nil {
				log.Fatal("Bitmap could not be created")
			}

			C.FPDFBitmap_FillRect(bitmap, C.int(0), C.int(0), C.int(width), C.int(height), C.ulong(0xFFFFFFFF))
			if err != nil {
				log.Panicln(err)
			}

			page := C.FPDF_LoadPage(doc, C.int(0))
			if err != nil {
				log.Panicln(err)
			}

			if page == nil {
				log.Fatal("Page could not be loaded")
			}

			C.FPDF_RenderPageBitmap(bitmap, page, C.int(0), C.int(0), C.int(width), C.int(height), C.int(0), C.int(0x10))
			if err != nil {
				log.Panicln(err)
			}

			elapsed := time.Since(start)
			log.Printf("Rendering from file took %s", elapsed)

			C.FPDF_ClosePage(page)
			C.FPDF_CloseDocument(doc)

			/*
				f, err := os.Create("img2.jpg")
				if err != nil {
					panic(err)
				}
				defer f.Close()
				if err = jpeg.Encode(f, img, nil); err != nil {
					log.Printf("failed to encode: %v", err)
				}*/
		}()
	}
}
