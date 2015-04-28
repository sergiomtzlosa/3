package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/mumax/3/cuda/cu"
	"github.com/mumax/3/timer"
	"sync"
	"unsafe"
)

// CUDA handle for shiftbytes kernel
var shiftbytes_code cu.Function

// Stores the arguments for shiftbytes kernel invocation
type shiftbytes_args_t struct {
	arg_dst   unsafe.Pointer
	arg_src   unsafe.Pointer
	arg_Nx    int
	arg_Ny    int
	arg_Nz    int
	arg_shx   int
	arg_clamp byte
	argptr    [7]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for shiftbytes kernel invocation
var shiftbytes_args shiftbytes_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	shiftbytes_args.argptr[0] = unsafe.Pointer(&shiftbytes_args.arg_dst)
	shiftbytes_args.argptr[1] = unsafe.Pointer(&shiftbytes_args.arg_src)
	shiftbytes_args.argptr[2] = unsafe.Pointer(&shiftbytes_args.arg_Nx)
	shiftbytes_args.argptr[3] = unsafe.Pointer(&shiftbytes_args.arg_Ny)
	shiftbytes_args.argptr[4] = unsafe.Pointer(&shiftbytes_args.arg_Nz)
	shiftbytes_args.argptr[5] = unsafe.Pointer(&shiftbytes_args.arg_shx)
	shiftbytes_args.argptr[6] = unsafe.Pointer(&shiftbytes_args.arg_clamp)
}

// Wrapper for shiftbytes CUDA kernel, asynchronous.
func k_shiftbytes_async(dst unsafe.Pointer, src unsafe.Pointer, Nx int, Ny int, Nz int, shx int, clamp byte, cfg *config) {
	if Synchronous { // debug
		Sync()
		timer.Start("shiftbytes")
	}

	shiftbytes_args.Lock()
	defer shiftbytes_args.Unlock()

	if shiftbytes_code == 0 {
		shiftbytes_code = fatbinLoad(shiftbytes_map, "shiftbytes")
	}

	shiftbytes_args.arg_dst = dst
	shiftbytes_args.arg_src = src
	shiftbytes_args.arg_Nx = Nx
	shiftbytes_args.arg_Ny = Ny
	shiftbytes_args.arg_Nz = Nz
	shiftbytes_args.arg_shx = shx
	shiftbytes_args.arg_clamp = clamp

	args := shiftbytes_args.argptr[:]
	cu.LaunchKernel(shiftbytes_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
		timer.Stop("shiftbytes")
	}
}

// maps compute capability on PTX code for shiftbytes kernel.
var shiftbytes_map = map[int]string{0: "",
	20: shiftbytes_ptx_20,
	30: shiftbytes_ptx_30,
	35: shiftbytes_ptx_35,
	50: shiftbytes_ptx_50}

// shiftbytes PTX code for various compute capabilities.
const (
	shiftbytes_ptx_20 = `
.version 4.0
.target sm_20
.address_size 64


.visible .entry shiftbytes(
	.param .u64 shiftbytes_param_0,
	.param .u64 shiftbytes_param_1,
	.param .u32 shiftbytes_param_2,
	.param .u32 shiftbytes_param_3,
	.param .u32 shiftbytes_param_4,
	.param .u32 shiftbytes_param_5,
	.param .u8 shiftbytes_param_6
)
{
	.reg .pred 	%p<9>;
	.reg .s16 	%rs<5>;
	.reg .s32 	%r<21>;
	.reg .s64 	%rd<9>;


	ld.param.u64 	%rd1, [shiftbytes_param_0];
	ld.param.u64 	%rd2, [shiftbytes_param_1];
	ld.param.u32 	%r6, [shiftbytes_param_2];
	ld.param.u32 	%r7, [shiftbytes_param_3];
	ld.param.u32 	%r9, [shiftbytes_param_4];
	ld.param.u32 	%r8, [shiftbytes_param_5];
	ld.param.u8 	%rs4, [shiftbytes_param_6];
	mov.u32 	%r10, %ntid.x;
	mov.u32 	%r11, %ctaid.x;
	mov.u32 	%r12, %tid.x;
	mad.lo.s32 	%r1, %r10, %r11, %r12;
	mov.u32 	%r13, %ntid.y;
	mov.u32 	%r14, %ctaid.y;
	mov.u32 	%r15, %tid.y;
	mad.lo.s32 	%r2, %r13, %r14, %r15;
	mov.u32 	%r16, %ntid.z;
	mov.u32 	%r17, %ctaid.z;
	mov.u32 	%r18, %tid.z;
	mad.lo.s32 	%r3, %r16, %r17, %r18;
	setp.lt.s32	%p1, %r1, %r6;
	setp.lt.s32	%p2, %r2, %r7;
	and.pred  	%p3, %p1, %p2;
	setp.lt.s32	%p4, %r3, %r9;
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_4;
	bra.uni 	BB0_1;

BB0_1:
	sub.s32 	%r4, %r1, %r8;
	setp.lt.s32	%p6, %r4, 0;
	setp.ge.s32	%p7, %r4, %r6;
	or.pred  	%p8, %p6, %p7;
	mad.lo.s32 	%r5, %r3, %r7, %r2;
	@%p8 bra 	BB0_3;

	cvta.to.global.u64 	%rd3, %rd2;
	mad.lo.s32 	%r19, %r5, %r6, %r4;
	cvt.s64.s32	%rd4, %r19;
	add.s64 	%rd5, %rd3, %rd4;
	ld.global.u8 	%rs4, [%rd5];

BB0_3:
	cvta.to.global.u64 	%rd6, %rd1;
	mad.lo.s32 	%r20, %r5, %r6, %r1;
	cvt.s64.s32	%rd7, %r20;
	add.s64 	%rd8, %rd6, %rd7;
	st.global.u8 	[%rd8], %rs4;

BB0_4:
	ret;
}


`
	shiftbytes_ptx_30 = `
.version 4.0
.target sm_30
.address_size 64


.visible .entry shiftbytes(
	.param .u64 shiftbytes_param_0,
	.param .u64 shiftbytes_param_1,
	.param .u32 shiftbytes_param_2,
	.param .u32 shiftbytes_param_3,
	.param .u32 shiftbytes_param_4,
	.param .u32 shiftbytes_param_5,
	.param .u8 shiftbytes_param_6
)
{
	.reg .pred 	%p<9>;
	.reg .s16 	%rs<5>;
	.reg .s32 	%r<21>;
	.reg .s64 	%rd<9>;


	ld.param.u64 	%rd1, [shiftbytes_param_0];
	ld.param.u64 	%rd2, [shiftbytes_param_1];
	ld.param.u32 	%r6, [shiftbytes_param_2];
	ld.param.u32 	%r7, [shiftbytes_param_3];
	ld.param.u32 	%r9, [shiftbytes_param_4];
	ld.param.u32 	%r8, [shiftbytes_param_5];
	ld.param.u8 	%rs4, [shiftbytes_param_6];
	mov.u32 	%r10, %ntid.x;
	mov.u32 	%r11, %ctaid.x;
	mov.u32 	%r12, %tid.x;
	mad.lo.s32 	%r1, %r10, %r11, %r12;
	mov.u32 	%r13, %ntid.y;
	mov.u32 	%r14, %ctaid.y;
	mov.u32 	%r15, %tid.y;
	mad.lo.s32 	%r2, %r13, %r14, %r15;
	mov.u32 	%r16, %ntid.z;
	mov.u32 	%r17, %ctaid.z;
	mov.u32 	%r18, %tid.z;
	mad.lo.s32 	%r3, %r16, %r17, %r18;
	setp.lt.s32	%p1, %r1, %r6;
	setp.lt.s32	%p2, %r2, %r7;
	and.pred  	%p3, %p1, %p2;
	setp.lt.s32	%p4, %r3, %r9;
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_4;
	bra.uni 	BB0_1;

BB0_1:
	sub.s32 	%r4, %r1, %r8;
	setp.lt.s32	%p6, %r4, 0;
	setp.ge.s32	%p7, %r4, %r6;
	or.pred  	%p8, %p6, %p7;
	mad.lo.s32 	%r5, %r3, %r7, %r2;
	@%p8 bra 	BB0_3;

	cvta.to.global.u64 	%rd3, %rd2;
	mad.lo.s32 	%r19, %r5, %r6, %r4;
	cvt.s64.s32	%rd4, %r19;
	add.s64 	%rd5, %rd3, %rd4;
	ld.global.u8 	%rs4, [%rd5];

BB0_3:
	cvta.to.global.u64 	%rd6, %rd1;
	mad.lo.s32 	%r20, %r5, %r6, %r1;
	cvt.s64.s32	%rd7, %r20;
	add.s64 	%rd8, %rd6, %rd7;
	st.global.u8 	[%rd8], %rs4;

BB0_4:
	ret;
}


`
	shiftbytes_ptx_35 = `
.version 4.0
.target sm_35
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.visible .entry shiftbytes(
	.param .u64 shiftbytes_param_0,
	.param .u64 shiftbytes_param_1,
	.param .u32 shiftbytes_param_2,
	.param .u32 shiftbytes_param_3,
	.param .u32 shiftbytes_param_4,
	.param .u32 shiftbytes_param_5,
	.param .u8 shiftbytes_param_6
)
{
	.reg .pred 	%p<9>;
	.reg .s16 	%rs<5>;
	.reg .s32 	%r<21>;
	.reg .s64 	%rd<9>;


	ld.param.u64 	%rd1, [shiftbytes_param_0];
	ld.param.u64 	%rd2, [shiftbytes_param_1];
	ld.param.u32 	%r6, [shiftbytes_param_2];
	ld.param.u32 	%r7, [shiftbytes_param_3];
	ld.param.u32 	%r9, [shiftbytes_param_4];
	ld.param.u32 	%r8, [shiftbytes_param_5];
	ld.param.u8 	%rs4, [shiftbytes_param_6];
	mov.u32 	%r10, %ntid.x;
	mov.u32 	%r11, %ctaid.x;
	mov.u32 	%r12, %tid.x;
	mad.lo.s32 	%r1, %r10, %r11, %r12;
	mov.u32 	%r13, %ntid.y;
	mov.u32 	%r14, %ctaid.y;
	mov.u32 	%r15, %tid.y;
	mad.lo.s32 	%r2, %r13, %r14, %r15;
	mov.u32 	%r16, %ntid.z;
	mov.u32 	%r17, %ctaid.z;
	mov.u32 	%r18, %tid.z;
	mad.lo.s32 	%r3, %r16, %r17, %r18;
	setp.lt.s32	%p1, %r1, %r6;
	setp.lt.s32	%p2, %r2, %r7;
	and.pred  	%p3, %p1, %p2;
	setp.lt.s32	%p4, %r3, %r9;
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB2_4;
	bra.uni 	BB2_1;

BB2_1:
	sub.s32 	%r4, %r1, %r8;
	setp.lt.s32	%p6, %r4, 0;
	setp.ge.s32	%p7, %r4, %r6;
	or.pred  	%p8, %p6, %p7;
	mad.lo.s32 	%r5, %r3, %r7, %r2;
	@%p8 bra 	BB2_3;

	cvta.to.global.u64 	%rd3, %rd2;
	mad.lo.s32 	%r19, %r5, %r6, %r4;
	cvt.s64.s32	%rd4, %r19;
	add.s64 	%rd5, %rd3, %rd4;
	ld.global.nc.u8 	%rs4, [%rd5];

BB2_3:
	cvta.to.global.u64 	%rd6, %rd1;
	mad.lo.s32 	%r20, %r5, %r6, %r1;
	cvt.s64.s32	%rd7, %r20;
	add.s64 	%rd8, %rd6, %rd7;
	st.global.u8 	[%rd8], %rs4;

BB2_4:
	ret;
}


`
	shiftbytes_ptx_50 = `
.version 4.0
.target sm_50
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.visible .entry shiftbytes(
	.param .u64 shiftbytes_param_0,
	.param .u64 shiftbytes_param_1,
	.param .u32 shiftbytes_param_2,
	.param .u32 shiftbytes_param_3,
	.param .u32 shiftbytes_param_4,
	.param .u32 shiftbytes_param_5,
	.param .u8 shiftbytes_param_6
)
{
	.reg .pred 	%p<9>;
	.reg .s16 	%rs<5>;
	.reg .s32 	%r<21>;
	.reg .s64 	%rd<9>;


	ld.param.u64 	%rd1, [shiftbytes_param_0];
	ld.param.u64 	%rd2, [shiftbytes_param_1];
	ld.param.u32 	%r6, [shiftbytes_param_2];
	ld.param.u32 	%r7, [shiftbytes_param_3];
	ld.param.u32 	%r9, [shiftbytes_param_4];
	ld.param.u32 	%r8, [shiftbytes_param_5];
	ld.param.u8 	%rs4, [shiftbytes_param_6];
	mov.u32 	%r10, %ntid.x;
	mov.u32 	%r11, %ctaid.x;
	mov.u32 	%r12, %tid.x;
	mad.lo.s32 	%r1, %r10, %r11, %r12;
	mov.u32 	%r13, %ntid.y;
	mov.u32 	%r14, %ctaid.y;
	mov.u32 	%r15, %tid.y;
	mad.lo.s32 	%r2, %r13, %r14, %r15;
	mov.u32 	%r16, %ntid.z;
	mov.u32 	%r17, %ctaid.z;
	mov.u32 	%r18, %tid.z;
	mad.lo.s32 	%r3, %r16, %r17, %r18;
	setp.lt.s32	%p1, %r1, %r6;
	setp.lt.s32	%p2, %r2, %r7;
	and.pred  	%p3, %p1, %p2;
	setp.lt.s32	%p4, %r3, %r9;
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB2_4;
	bra.uni 	BB2_1;

BB2_1:
	sub.s32 	%r4, %r1, %r8;
	setp.lt.s32	%p6, %r4, 0;
	setp.ge.s32	%p7, %r4, %r6;
	or.pred  	%p8, %p6, %p7;
	mad.lo.s32 	%r5, %r3, %r7, %r2;
	@%p8 bra 	BB2_3;

	cvta.to.global.u64 	%rd3, %rd2;
	mad.lo.s32 	%r19, %r5, %r6, %r4;
	cvt.s64.s32	%rd4, %r19;
	add.s64 	%rd5, %rd3, %rd4;
	ld.global.nc.u8 	%rs4, [%rd5];

BB2_3:
	cvta.to.global.u64 	%rd6, %rd1;
	mad.lo.s32 	%r20, %r5, %r6, %r1;
	cvt.s64.s32	%rd7, %r20;
	add.s64 	%rd8, %rd6, %rd7;
	st.global.u8 	[%rd8], %rs4;

BB2_4:
	ret;
}


`
)
