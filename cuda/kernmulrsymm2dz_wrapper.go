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

// CUDA handle for kernmulRSymm2Dz kernel
var kernmulRSymm2Dz_code cu.Function

// Stores the arguments for kernmulRSymm2Dz kernel invocation
type kernmulRSymm2Dz_args_t struct {
	arg_fftMz  unsafe.Pointer
	arg_fftKzz unsafe.Pointer
	arg_Nx     int
	arg_Ny     int
	argptr     [4]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for kernmulRSymm2Dz kernel invocation
var kernmulRSymm2Dz_args kernmulRSymm2Dz_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	kernmulRSymm2Dz_args.argptr[0] = unsafe.Pointer(&kernmulRSymm2Dz_args.arg_fftMz)
	kernmulRSymm2Dz_args.argptr[1] = unsafe.Pointer(&kernmulRSymm2Dz_args.arg_fftKzz)
	kernmulRSymm2Dz_args.argptr[2] = unsafe.Pointer(&kernmulRSymm2Dz_args.arg_Nx)
	kernmulRSymm2Dz_args.argptr[3] = unsafe.Pointer(&kernmulRSymm2Dz_args.arg_Ny)
}

// Wrapper for kernmulRSymm2Dz CUDA kernel, asynchronous.
func k_kernmulRSymm2Dz_async(fftMz unsafe.Pointer, fftKzz unsafe.Pointer, Nx int, Ny int, cfg *config) {
	if Synchronous { // debug
		Sync()
		timer.Start("kernmulRSymm2Dz")
	}

	kernmulRSymm2Dz_args.Lock()
	defer kernmulRSymm2Dz_args.Unlock()

	if kernmulRSymm2Dz_code == 0 {
		kernmulRSymm2Dz_code = fatbinLoad(kernmulRSymm2Dz_map, "kernmulRSymm2Dz")
	}

	kernmulRSymm2Dz_args.arg_fftMz = fftMz
	kernmulRSymm2Dz_args.arg_fftKzz = fftKzz
	kernmulRSymm2Dz_args.arg_Nx = Nx
	kernmulRSymm2Dz_args.arg_Ny = Ny

	args := kernmulRSymm2Dz_args.argptr[:]
	cu.LaunchKernel(kernmulRSymm2Dz_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
		timer.Stop("kernmulRSymm2Dz")
	}
}

// maps compute capability on PTX code for kernmulRSymm2Dz kernel.
var kernmulRSymm2Dz_map = map[int]string{0: "",
	20: kernmulRSymm2Dz_ptx_20,
	30: kernmulRSymm2Dz_ptx_30,
	35: kernmulRSymm2Dz_ptx_35,
	50: kernmulRSymm2Dz_ptx_50}

// kernmulRSymm2Dz PTX code for various compute capabilities.
const (
	kernmulRSymm2Dz_ptx_20 = `
.version 4.0
.target sm_20
.address_size 64


.visible .entry kernmulRSymm2Dz(
	.param .u64 kernmulRSymm2Dz_param_0,
	.param .u64 kernmulRSymm2Dz_param_1,
	.param .u32 kernmulRSymm2Dz_param_2,
	.param .u32 kernmulRSymm2Dz_param_3
)
{
	.reg .pred 	%p<5>;
	.reg .s32 	%r<19>;
	.reg .f32 	%f<6>;
	.reg .s64 	%rd<9>;


	ld.param.u64 	%rd1, [kernmulRSymm2Dz_param_0];
	ld.param.u64 	%rd2, [kernmulRSymm2Dz_param_1];
	ld.param.u32 	%r3, [kernmulRSymm2Dz_param_2];
	ld.param.u32 	%r4, [kernmulRSymm2Dz_param_3];
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r1, %r5, %r6, %r7;
	mov.u32 	%r8, %ntid.y;
	mov.u32 	%r9, %ctaid.y;
	mov.u32 	%r10, %tid.y;
	mad.lo.s32 	%r2, %r8, %r9, %r10;
	setp.ge.s32	%p1, %r2, %r4;
	setp.ge.s32	%p2, %r1, %r3;
	or.pred  	%p3, %p2, %p1;
	@%p3 bra 	BB0_2;

	cvta.to.global.u64 	%rd3, %rd2;
	cvta.to.global.u64 	%rd4, %rd1;
	mad.lo.s32 	%r11, %r2, %r3, %r1;
	shl.b32 	%r12, %r11, 1;
	mul.wide.s32 	%rd5, %r12, 4;
	add.s64 	%rd6, %rd4, %rd5;
	ld.global.f32 	%f1, [%rd6+4];
	shr.u32 	%r13, %r4, 31;
	add.s32 	%r14, %r4, %r13;
	shr.s32 	%r15, %r14, 1;
	setp.gt.s32	%p4, %r2, %r15;
	sub.s32 	%r16, %r4, %r2;
	selp.b32	%r17, %r16, %r2, %p4;
	mad.lo.s32 	%r18, %r17, %r3, %r1;
	mul.wide.s32 	%rd7, %r18, 4;
	add.s64 	%rd8, %rd3, %rd7;
	ld.global.f32 	%f2, [%rd8];
	ld.global.f32 	%f3, [%rd6];
	mul.f32 	%f4, %f3, %f2;
	st.global.f32 	[%rd6], %f4;
	mul.f32 	%f5, %f1, %f2;
	st.global.f32 	[%rd6+4], %f5;

BB0_2:
	ret;
}


`
	kernmulRSymm2Dz_ptx_30 = `
.version 4.0
.target sm_30
.address_size 64


.visible .entry kernmulRSymm2Dz(
	.param .u64 kernmulRSymm2Dz_param_0,
	.param .u64 kernmulRSymm2Dz_param_1,
	.param .u32 kernmulRSymm2Dz_param_2,
	.param .u32 kernmulRSymm2Dz_param_3
)
{
	.reg .pred 	%p<5>;
	.reg .s32 	%r<19>;
	.reg .f32 	%f<6>;
	.reg .s64 	%rd<9>;


	ld.param.u64 	%rd1, [kernmulRSymm2Dz_param_0];
	ld.param.u64 	%rd2, [kernmulRSymm2Dz_param_1];
	ld.param.u32 	%r3, [kernmulRSymm2Dz_param_2];
	ld.param.u32 	%r4, [kernmulRSymm2Dz_param_3];
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r1, %r5, %r6, %r7;
	mov.u32 	%r8, %ntid.y;
	mov.u32 	%r9, %ctaid.y;
	mov.u32 	%r10, %tid.y;
	mad.lo.s32 	%r2, %r8, %r9, %r10;
	setp.ge.s32	%p1, %r2, %r4;
	setp.ge.s32	%p2, %r1, %r3;
	or.pred  	%p3, %p2, %p1;
	@%p3 bra 	BB0_2;

	cvta.to.global.u64 	%rd3, %rd2;
	cvta.to.global.u64 	%rd4, %rd1;
	mad.lo.s32 	%r11, %r2, %r3, %r1;
	shl.b32 	%r12, %r11, 1;
	mul.wide.s32 	%rd5, %r12, 4;
	add.s64 	%rd6, %rd4, %rd5;
	ld.global.f32 	%f1, [%rd6+4];
	shr.u32 	%r13, %r4, 31;
	add.s32 	%r14, %r4, %r13;
	shr.s32 	%r15, %r14, 1;
	setp.gt.s32	%p4, %r2, %r15;
	sub.s32 	%r16, %r4, %r2;
	selp.b32	%r17, %r16, %r2, %p4;
	mad.lo.s32 	%r18, %r17, %r3, %r1;
	mul.wide.s32 	%rd7, %r18, 4;
	add.s64 	%rd8, %rd3, %rd7;
	ld.global.f32 	%f2, [%rd8];
	ld.global.f32 	%f3, [%rd6];
	mul.f32 	%f4, %f3, %f2;
	st.global.f32 	[%rd6], %f4;
	mul.f32 	%f5, %f1, %f2;
	st.global.f32 	[%rd6+4], %f5;

BB0_2:
	ret;
}


`
	kernmulRSymm2Dz_ptx_35 = `
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

.visible .entry kernmulRSymm2Dz(
	.param .u64 kernmulRSymm2Dz_param_0,
	.param .u64 kernmulRSymm2Dz_param_1,
	.param .u32 kernmulRSymm2Dz_param_2,
	.param .u32 kernmulRSymm2Dz_param_3
)
{
	.reg .pred 	%p<5>;
	.reg .s32 	%r<19>;
	.reg .f32 	%f<6>;
	.reg .s64 	%rd<9>;


	ld.param.u64 	%rd1, [kernmulRSymm2Dz_param_0];
	ld.param.u64 	%rd2, [kernmulRSymm2Dz_param_1];
	ld.param.u32 	%r3, [kernmulRSymm2Dz_param_2];
	ld.param.u32 	%r4, [kernmulRSymm2Dz_param_3];
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r1, %r5, %r6, %r7;
	mov.u32 	%r8, %ntid.y;
	mov.u32 	%r9, %ctaid.y;
	mov.u32 	%r10, %tid.y;
	mad.lo.s32 	%r2, %r8, %r9, %r10;
	setp.ge.s32	%p1, %r2, %r4;
	setp.ge.s32	%p2, %r1, %r3;
	or.pred  	%p3, %p2, %p1;
	@%p3 bra 	BB2_2;

	cvta.to.global.u64 	%rd3, %rd2;
	cvta.to.global.u64 	%rd4, %rd1;
	mad.lo.s32 	%r11, %r2, %r3, %r1;
	shl.b32 	%r12, %r11, 1;
	mul.wide.s32 	%rd5, %r12, 4;
	add.s64 	%rd6, %rd4, %rd5;
	ld.global.f32 	%f1, [%rd6+4];
	shr.u32 	%r13, %r4, 31;
	add.s32 	%r14, %r4, %r13;
	shr.s32 	%r15, %r14, 1;
	setp.gt.s32	%p4, %r2, %r15;
	sub.s32 	%r16, %r4, %r2;
	selp.b32	%r17, %r16, %r2, %p4;
	mad.lo.s32 	%r18, %r17, %r3, %r1;
	mul.wide.s32 	%rd7, %r18, 4;
	add.s64 	%rd8, %rd3, %rd7;
	ld.global.nc.f32 	%f2, [%rd8];
	ld.global.f32 	%f3, [%rd6];
	mul.f32 	%f4, %f3, %f2;
	st.global.f32 	[%rd6], %f4;
	mul.f32 	%f5, %f1, %f2;
	st.global.f32 	[%rd6+4], %f5;

BB2_2:
	ret;
}


`
	kernmulRSymm2Dz_ptx_50 = `
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

.visible .entry kernmulRSymm2Dz(
	.param .u64 kernmulRSymm2Dz_param_0,
	.param .u64 kernmulRSymm2Dz_param_1,
	.param .u32 kernmulRSymm2Dz_param_2,
	.param .u32 kernmulRSymm2Dz_param_3
)
{
	.reg .pred 	%p<5>;
	.reg .s32 	%r<19>;
	.reg .f32 	%f<6>;
	.reg .s64 	%rd<9>;


	ld.param.u64 	%rd1, [kernmulRSymm2Dz_param_0];
	ld.param.u64 	%rd2, [kernmulRSymm2Dz_param_1];
	ld.param.u32 	%r3, [kernmulRSymm2Dz_param_2];
	ld.param.u32 	%r4, [kernmulRSymm2Dz_param_3];
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r1, %r5, %r6, %r7;
	mov.u32 	%r8, %ntid.y;
	mov.u32 	%r9, %ctaid.y;
	mov.u32 	%r10, %tid.y;
	mad.lo.s32 	%r2, %r8, %r9, %r10;
	setp.ge.s32	%p1, %r2, %r4;
	setp.ge.s32	%p2, %r1, %r3;
	or.pred  	%p3, %p2, %p1;
	@%p3 bra 	BB2_2;

	cvta.to.global.u64 	%rd3, %rd2;
	cvta.to.global.u64 	%rd4, %rd1;
	mad.lo.s32 	%r11, %r2, %r3, %r1;
	shl.b32 	%r12, %r11, 1;
	mul.wide.s32 	%rd5, %r12, 4;
	add.s64 	%rd6, %rd4, %rd5;
	ld.global.f32 	%f1, [%rd6+4];
	shr.u32 	%r13, %r4, 31;
	add.s32 	%r14, %r4, %r13;
	shr.s32 	%r15, %r14, 1;
	setp.gt.s32	%p4, %r2, %r15;
	sub.s32 	%r16, %r4, %r2;
	selp.b32	%r17, %r16, %r2, %p4;
	mad.lo.s32 	%r18, %r17, %r3, %r1;
	mul.wide.s32 	%rd7, %r18, 4;
	add.s64 	%rd8, %rd3, %rd7;
	ld.global.nc.f32 	%f2, [%rd8];
	ld.global.f32 	%f3, [%rd6];
	mul.f32 	%f4, %f3, %f2;
	st.global.f32 	[%rd6], %f4;
	mul.f32 	%f5, %f1, %f2;
	st.global.f32 	[%rd6+4], %f5;

BB2_2:
	ret;
}


`
)
