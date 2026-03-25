.text
.globl cpuid_native
.type cpuid_native, @function

cpuid_native:
    pushq   %rbx

    mov     %rdx, %r8

    mov     %edi, %eax
    mov     %esi, %ecx
    cpuid

    mov     %eax, (%r8)
    mov     %ebx, 4(%r8)
    mov     %ecx, 8(%r8)
    mov     %edx, 12(%r8)

    popq    %rbx
    ret

.size cpuid_native, .-cpuid_native
