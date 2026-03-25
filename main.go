//go:build cgo

package main

/*
#include <stdint.h>
void cpuid_native(uint32_t eaxIn, uint32_t ecxIn, uint32_t *regs);
*/
import "C"

import (
	"fmt"
)

func cpuid(eaxIn, ecxIn uint32) (eax, ebx, ecx, edx uint32) {
	var regs [4]C.uint32_t
	C.cpuid_native(C.uint32_t(eaxIn), C.uint32_t(ecxIn), &regs[0])
	return uint32(regs[0]), uint32(regs[1]), uint32(regs[2]), uint32(regs[3])
}

func main() {
	// Folha 0: obter Vendor ID e máximo nível suportado
	eax, ebx, ecx, edx := cpuid(0, 0)

	fmt.Printf("CPUID(0, 0):\n")
	fmt.Printf("EAX: %08x (Max Leaf)\n", eax)

	// Vendor ID string
	vendor := string([]byte{
		byte(ebx), byte(ebx >> 8), byte(ebx >> 16), byte(ebx >> 24),
		byte(edx), byte(edx >> 8), byte(edx >> 16), byte(edx >> 24),
		byte(ecx), byte(ecx >> 8), byte(ecx >> 16), byte(ecx >> 24),
	})
	fmt.Printf("Vendor ID: %s\n", vendor)

	// Folha 1: obter informações do processador e recursos
	eax, ebx, ecx, edx = cpuid(1, 0)
	fmt.Printf("\nCPUID(1, 0):\n")
	fmt.Printf("EAX: %08x\n", eax)
	fmt.Printf("EBX: %08x\n", ebx)
	fmt.Printf("ECX: %08x\n", ecx)
	fmt.Printf("EDX: %08x\n", edx)
}
