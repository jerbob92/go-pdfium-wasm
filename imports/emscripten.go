package imports

import "log"

func emscripten_notify_memory_growth(i int32) {
	log.Printf("Called into emscripten_notify_memory_growth with argument %d", i)
}

var tempRet0 = int32(0)

func setTempRet0(input int32) {
	tempRet0 = input
}

func getTempRet0() int32 {
	return tempRet0
}

func emscripten_throw_longjmp() {
	log.Fatal("Called into emscripten_throw_longjmp")
}

func invoke_v(_ int32) {
	log.Fatal("Called into invoke_v")
}

func invoke_vi(_ int32, _ int32) {
	log.Fatal("Called into invoke_vi")
}

func invoke_vii(_ int32, _ int32, _ int32) {
	log.Fatal("Called into invoke_vii")
}

func invoke_viii(_ int32, _ int32, _ int32, _ int32) {
	log.Fatal("Called into invoke_viii")
}

func invoke_viiii(_ int32, _ int32, _ int32, _ int32, _ int32) {
	log.Fatal("Called into invoke_viiii")
}

func invoke_ii(_ int32, _ int32) int32 {
	log.Fatal("Called into invoke_ii")
	return 0
}

func invoke_iii(_ int32, _ int32, _ int32) int32 {
	log.Fatal("Called into invoke_iii")
	return 0
}

func invoke_iiii(_ int32, _ int32, _ int32, _ int32) int32 {
	log.Fatal("Called into invoke_iiii")
	return 0
}

func invoke_iiiii(_ int32, _ int32, _ int32, _ int32, _ int32) int32 {
	log.Fatal("Called into invoke_iiiii")
	return 0
}
