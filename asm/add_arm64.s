TEXT ·Add64(SB), 0, $0-24
	MOVD a+0(FP), R9
	MOVD b+8(FP), R10
	ADD  R9, R10, R11
	MOVD R11, ret+16(FP)
	RET
