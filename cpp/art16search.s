	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 10, 14
	.intel_syntax noprefix
	.globl	__Z20IndexOfByteIn16BytesPiPhS0_Pc ## -- Begin function _Z20IndexOfByteIn16BytesPiPhS0_Pc
	.p2align	4, 0x90
__Z20IndexOfByteIn16BytesPiPhS0_Pc:     ## @_Z20IndexOfByteIn16BytesPiPhS0_Pc
## %bb.0:
	push	rbp
	mov	rbp, rsp
	and	rsp, -8
	mov	rax, rcx
	movzx	ecx, byte ptr [rdx]
	vmovd	xmm0, ecx
	vpxor	xmm1, xmm1, xmm1
	vpshufb	xmm0, xmm0, xmm1
	vpcmpeqb	xmm0, xmm0, xmmword ptr [rsi]
	mov	cl, byte ptr [rdi]
	mov	edx, 1
	shl	edx, cl
	add	edx, 65535
	vpmovmskb	ecx, xmm0
	and	ecx, edx
	je	LBB0_2
## %bb.1:
	bsf	ecx, ecx
	mov	byte ptr [rax], cl
LBB0_2:
	mov	rsp, rbp
	pop	rbp
	ret
                                        ## -- End function

.subsections_via_symbols
