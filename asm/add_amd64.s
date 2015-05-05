TEXT ·add(SB),0,$0-24
	ADDQ	a+0(FP), b+8(FP), AX
	MOVQ	AX, ret+16(FP)
	RET
