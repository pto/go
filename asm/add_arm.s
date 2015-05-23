TEXT ·Add64(SB),0,$0-24
	MOVW	a_lo+0(FP), R2
	MOVW	a_hi+4(FP), R3
	MOVW	b_lo+8(FP), R4
	MOVW	b_hi+12(FP), R5
	ADD.S	R2, R4, R0
	ADC.S	R3, R5, R1
	MOVW	R0, ret_lo+16(FP)
	MOVW	R1, ret_hi+20(FP)
	RET
