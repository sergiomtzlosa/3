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

// CUDA handle for reducesum kernel
var reducesum_code cu.Function

// Stores the arguments for reducesum kernel invocation
type reducesum_args_t struct {
	arg_src     unsafe.Pointer
	arg_dst     unsafe.Pointer
	arg_initVal float32
	arg_n       int
	argptr      [4]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for reducesum kernel invocation
var reducesum_args reducesum_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	reducesum_args.argptr[0] = unsafe.Pointer(&reducesum_args.arg_src)
	reducesum_args.argptr[1] = unsafe.Pointer(&reducesum_args.arg_dst)
	reducesum_args.argptr[2] = unsafe.Pointer(&reducesum_args.arg_initVal)
	reducesum_args.argptr[3] = unsafe.Pointer(&reducesum_args.arg_n)
}

// Wrapper for reducesum CUDA kernel, asynchronous.
func k_reducesum_async(src unsafe.Pointer, dst unsafe.Pointer, initVal float32, n int, cfg *config) {
	if Synchronous { // debug
		Sync()
		timer.Start("reducesum")
	}

	reducesum_args.Lock()
	defer reducesum_args.Unlock()

	if reducesum_code == 0 {
		reducesum_code = fatbinLoad(reducesum_map, "reducesum")
	}

	reducesum_args.arg_src = src
	reducesum_args.arg_dst = dst
	reducesum_args.arg_initVal = initVal
	reducesum_args.arg_n = n

	args := reducesum_args.argptr[:]
	cu.LaunchKernel(reducesum_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
		timer.Stop("reducesum")
	}
}

// maps compute capability on PTX code for reducesum kernel.
var reducesum_map = map[int]string{0: "",
	20: reducesum_ptx_20,
	30: reducesum_ptx_30,
	35: reducesum_ptx_35,
	50: reducesum_ptx_50}

// reducesum PTX code for various compute capabilities.
const (
	reducesum_ptx_20 = `
.version 4.0
.target sm_20
.address_size 64


.visible .entry reducesum(
	.param .u64 reducesum_param_0,
	.param .u64 reducesum_param_1,
	.param .f32 reducesum_param_2,
	.param .u32 reducesum_param_3
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<15>;
	.reg .f32 	%f<30>;
	.reg .s64 	%rd<13>;
	// demoted variable
	.shared .align 4 .b8 reducesum$__cuda_local_var_34156_32_non_const_sdata[2048];

	ld.param.u64 	%rd4, [reducesum_param_0];
	ld.param.u64 	%rd3, [reducesum_param_1];
	ld.param.f32 	%f29, [reducesum_param_2];
	ld.param.u32 	%r9, [reducesum_param_3];
	cvta.to.global.u64 	%rd1, %rd4;
	mov.u32 	%r14, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r13, %r14, %r10, %r2;
	mov.u32 	%r11, %nctaid.x;
	mul.lo.s32 	%r4, %r11, %r14;
	setp.ge.s32	%p1, %r13, %r9;
	@%p1 bra 	BB0_2;

BB0_1:
	mul.wide.s32 	%rd5, %r13, 4;
	add.s64 	%rd6, %rd1, %rd5;
	ld.global.f32 	%f5, [%rd6];
	add.f32 	%f29, %f29, %f5;
	add.s32 	%r13, %r13, %r4;
	setp.lt.s32	%p2, %r13, %r9;
	@%p2 bra 	BB0_1;

BB0_2:
	mul.wide.s32 	%rd7, %r2, 4;
	mov.u64 	%rd8, reducesum$__cuda_local_var_34156_32_non_const_sdata;
	add.s64 	%rd2, %rd8, %rd7;
	st.shared.f32 	[%rd2], %f29;
	bar.sync 	0;
	setp.lt.u32	%p3, %r14, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	mov.u32 	%r7, %r14;
	shr.u32 	%r14, %r7, 1;
	setp.ge.u32	%p4, %r2, %r14;
	@%p4 bra 	BB0_5;

	ld.shared.f32 	%f6, [%rd2];
	add.s32 	%r12, %r14, %r2;
	mul.wide.u32 	%rd9, %r12, 4;
	add.s64 	%rd11, %rd8, %rd9;
	ld.shared.f32 	%f7, [%rd11];
	add.f32 	%f8, %f6, %f7;
	st.shared.f32 	[%rd2], %f8;

BB0_5:
	bar.sync 	0;
	setp.gt.u32	%p5, %r7, 131;
	@%p5 bra 	BB0_3;

BB0_6:
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	ld.volatile.shared.f32 	%f9, [%rd2];
	ld.volatile.shared.f32 	%f10, [%rd2+128];
	add.f32 	%f11, %f9, %f10;
	st.volatile.shared.f32 	[%rd2], %f11;
	ld.volatile.shared.f32 	%f12, [%rd2+64];
	ld.volatile.shared.f32 	%f13, [%rd2];
	add.f32 	%f14, %f13, %f12;
	st.volatile.shared.f32 	[%rd2], %f14;
	ld.volatile.shared.f32 	%f15, [%rd2+32];
	ld.volatile.shared.f32 	%f16, [%rd2];
	add.f32 	%f17, %f16, %f15;
	st.volatile.shared.f32 	[%rd2], %f17;
	ld.volatile.shared.f32 	%f18, [%rd2+16];
	ld.volatile.shared.f32 	%f19, [%rd2];
	add.f32 	%f20, %f19, %f18;
	st.volatile.shared.f32 	[%rd2], %f20;
	ld.volatile.shared.f32 	%f21, [%rd2+8];
	ld.volatile.shared.f32 	%f22, [%rd2];
	add.f32 	%f23, %f22, %f21;
	st.volatile.shared.f32 	[%rd2], %f23;
	ld.volatile.shared.f32 	%f24, [%rd2+4];
	ld.volatile.shared.f32 	%f25, [%rd2];
	add.f32 	%f26, %f25, %f24;
	st.volatile.shared.f32 	[%rd2], %f26;

BB0_8:
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	cvta.to.global.u64 	%rd12, %rd3;
	ld.shared.f32 	%f27, [reducesum$__cuda_local_var_34156_32_non_const_sdata];
	atom.global.add.f32 	%f28, [%rd12], %f27;

BB0_10:
	ret;
}


`
	reducesum_ptx_30 = `
.version 4.0
.target sm_30
.address_size 64


.visible .entry reducesum(
	.param .u64 reducesum_param_0,
	.param .u64 reducesum_param_1,
	.param .f32 reducesum_param_2,
	.param .u32 reducesum_param_3
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<15>;
	.reg .f32 	%f<30>;
	.reg .s64 	%rd<13>;
	// demoted variable
	.shared .align 4 .b8 reducesum$__cuda_local_var_34229_32_non_const_sdata[2048];

	ld.param.u64 	%rd4, [reducesum_param_0];
	ld.param.u64 	%rd3, [reducesum_param_1];
	ld.param.f32 	%f29, [reducesum_param_2];
	ld.param.u32 	%r9, [reducesum_param_3];
	cvta.to.global.u64 	%rd1, %rd4;
	mov.u32 	%r14, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r13, %r14, %r10, %r2;
	mov.u32 	%r11, %nctaid.x;
	mul.lo.s32 	%r4, %r11, %r14;
	setp.ge.s32	%p1, %r13, %r9;
	@%p1 bra 	BB0_2;

BB0_1:
	mul.wide.s32 	%rd5, %r13, 4;
	add.s64 	%rd6, %rd1, %rd5;
	ld.global.f32 	%f5, [%rd6];
	add.f32 	%f29, %f29, %f5;
	add.s32 	%r13, %r13, %r4;
	setp.lt.s32	%p2, %r13, %r9;
	@%p2 bra 	BB0_1;

BB0_2:
	mul.wide.s32 	%rd7, %r2, 4;
	mov.u64 	%rd8, reducesum$__cuda_local_var_34229_32_non_const_sdata;
	add.s64 	%rd2, %rd8, %rd7;
	st.shared.f32 	[%rd2], %f29;
	bar.sync 	0;
	setp.lt.u32	%p3, %r14, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	mov.u32 	%r7, %r14;
	shr.u32 	%r14, %r7, 1;
	setp.ge.u32	%p4, %r2, %r14;
	@%p4 bra 	BB0_5;

	ld.shared.f32 	%f6, [%rd2];
	add.s32 	%r12, %r14, %r2;
	mul.wide.u32 	%rd9, %r12, 4;
	add.s64 	%rd11, %rd8, %rd9;
	ld.shared.f32 	%f7, [%rd11];
	add.f32 	%f8, %f6, %f7;
	st.shared.f32 	[%rd2], %f8;

BB0_5:
	bar.sync 	0;
	setp.gt.u32	%p5, %r7, 131;
	@%p5 bra 	BB0_3;

BB0_6:
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	ld.volatile.shared.f32 	%f9, [%rd2];
	ld.volatile.shared.f32 	%f10, [%rd2+128];
	add.f32 	%f11, %f9, %f10;
	st.volatile.shared.f32 	[%rd2], %f11;
	ld.volatile.shared.f32 	%f12, [%rd2+64];
	ld.volatile.shared.f32 	%f13, [%rd2];
	add.f32 	%f14, %f13, %f12;
	st.volatile.shared.f32 	[%rd2], %f14;
	ld.volatile.shared.f32 	%f15, [%rd2+32];
	ld.volatile.shared.f32 	%f16, [%rd2];
	add.f32 	%f17, %f16, %f15;
	st.volatile.shared.f32 	[%rd2], %f17;
	ld.volatile.shared.f32 	%f18, [%rd2+16];
	ld.volatile.shared.f32 	%f19, [%rd2];
	add.f32 	%f20, %f19, %f18;
	st.volatile.shared.f32 	[%rd2], %f20;
	ld.volatile.shared.f32 	%f21, [%rd2+8];
	ld.volatile.shared.f32 	%f22, [%rd2];
	add.f32 	%f23, %f22, %f21;
	st.volatile.shared.f32 	[%rd2], %f23;
	ld.volatile.shared.f32 	%f24, [%rd2+4];
	ld.volatile.shared.f32 	%f25, [%rd2];
	add.f32 	%f26, %f25, %f24;
	st.volatile.shared.f32 	[%rd2], %f26;

BB0_8:
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	cvta.to.global.u64 	%rd12, %rd3;
	ld.shared.f32 	%f27, [reducesum$__cuda_local_var_34229_32_non_const_sdata];
	atom.global.add.f32 	%f28, [%rd12], %f27;

BB0_10:
	ret;
}


`
	reducesum_ptx_35 = `
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

.visible .entry reducesum(
	.param .u64 reducesum_param_0,
	.param .u64 reducesum_param_1,
	.param .f32 reducesum_param_2,
	.param .u32 reducesum_param_3
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<15>;
	.reg .f32 	%f<30>;
	.reg .s64 	%rd<13>;
	// demoted variable
	.shared .align 4 .b8 reducesum$__cuda_local_var_34394_32_non_const_sdata[2048];

	ld.param.u64 	%rd4, [reducesum_param_0];
	ld.param.u64 	%rd3, [reducesum_param_1];
	ld.param.f32 	%f29, [reducesum_param_2];
	ld.param.u32 	%r9, [reducesum_param_3];
	cvta.to.global.u64 	%rd1, %rd4;
	mov.u32 	%r14, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r13, %r14, %r10, %r2;
	mov.u32 	%r11, %nctaid.x;
	mul.lo.s32 	%r4, %r11, %r14;
	setp.ge.s32	%p1, %r13, %r9;
	@%p1 bra 	BB2_2;

BB2_1:
	mul.wide.s32 	%rd5, %r13, 4;
	add.s64 	%rd6, %rd1, %rd5;
	ld.global.nc.f32 	%f5, [%rd6];
	add.f32 	%f29, %f29, %f5;
	add.s32 	%r13, %r13, %r4;
	setp.lt.s32	%p2, %r13, %r9;
	@%p2 bra 	BB2_1;

BB2_2:
	mul.wide.s32 	%rd7, %r2, 4;
	mov.u64 	%rd8, reducesum$__cuda_local_var_34394_32_non_const_sdata;
	add.s64 	%rd2, %rd8, %rd7;
	st.shared.f32 	[%rd2], %f29;
	bar.sync 	0;
	setp.lt.u32	%p3, %r14, 66;
	@%p3 bra 	BB2_6;

BB2_3:
	mov.u32 	%r7, %r14;
	shr.u32 	%r14, %r7, 1;
	setp.ge.u32	%p4, %r2, %r14;
	@%p4 bra 	BB2_5;

	ld.shared.f32 	%f6, [%rd2];
	add.s32 	%r12, %r14, %r2;
	mul.wide.u32 	%rd9, %r12, 4;
	add.s64 	%rd11, %rd8, %rd9;
	ld.shared.f32 	%f7, [%rd11];
	add.f32 	%f8, %f6, %f7;
	st.shared.f32 	[%rd2], %f8;

BB2_5:
	bar.sync 	0;
	setp.gt.u32	%p5, %r7, 131;
	@%p5 bra 	BB2_3;

BB2_6:
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB2_8;

	ld.volatile.shared.f32 	%f9, [%rd2];
	ld.volatile.shared.f32 	%f10, [%rd2+128];
	add.f32 	%f11, %f9, %f10;
	st.volatile.shared.f32 	[%rd2], %f11;
	ld.volatile.shared.f32 	%f12, [%rd2+64];
	ld.volatile.shared.f32 	%f13, [%rd2];
	add.f32 	%f14, %f13, %f12;
	st.volatile.shared.f32 	[%rd2], %f14;
	ld.volatile.shared.f32 	%f15, [%rd2+32];
	ld.volatile.shared.f32 	%f16, [%rd2];
	add.f32 	%f17, %f16, %f15;
	st.volatile.shared.f32 	[%rd2], %f17;
	ld.volatile.shared.f32 	%f18, [%rd2+16];
	ld.volatile.shared.f32 	%f19, [%rd2];
	add.f32 	%f20, %f19, %f18;
	st.volatile.shared.f32 	[%rd2], %f20;
	ld.volatile.shared.f32 	%f21, [%rd2+8];
	ld.volatile.shared.f32 	%f22, [%rd2];
	add.f32 	%f23, %f22, %f21;
	st.volatile.shared.f32 	[%rd2], %f23;
	ld.volatile.shared.f32 	%f24, [%rd2+4];
	ld.volatile.shared.f32 	%f25, [%rd2];
	add.f32 	%f26, %f25, %f24;
	st.volatile.shared.f32 	[%rd2], %f26;

BB2_8:
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB2_10;

	cvta.to.global.u64 	%rd12, %rd3;
	ld.shared.f32 	%f27, [reducesum$__cuda_local_var_34394_32_non_const_sdata];
	atom.global.add.f32 	%f28, [%rd12], %f27;

BB2_10:
	ret;
}


`
	reducesum_ptx_50 = `
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

.visible .entry reducesum(
	.param .u64 reducesum_param_0,
	.param .u64 reducesum_param_1,
	.param .f32 reducesum_param_2,
	.param .u32 reducesum_param_3
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<15>;
	.reg .f32 	%f<30>;
	.reg .s64 	%rd<13>;
	// demoted variable
	.shared .align 4 .b8 reducesum$__cuda_local_var_34394_32_non_const_sdata[2048];

	ld.param.u64 	%rd4, [reducesum_param_0];
	ld.param.u64 	%rd3, [reducesum_param_1];
	ld.param.f32 	%f29, [reducesum_param_2];
	ld.param.u32 	%r9, [reducesum_param_3];
	cvta.to.global.u64 	%rd1, %rd4;
	mov.u32 	%r14, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r13, %r14, %r10, %r2;
	mov.u32 	%r11, %nctaid.x;
	mul.lo.s32 	%r4, %r11, %r14;
	setp.ge.s32	%p1, %r13, %r9;
	@%p1 bra 	BB2_2;

BB2_1:
	mul.wide.s32 	%rd5, %r13, 4;
	add.s64 	%rd6, %rd1, %rd5;
	ld.global.nc.f32 	%f5, [%rd6];
	add.f32 	%f29, %f29, %f5;
	add.s32 	%r13, %r13, %r4;
	setp.lt.s32	%p2, %r13, %r9;
	@%p2 bra 	BB2_1;

BB2_2:
	mul.wide.s32 	%rd7, %r2, 4;
	mov.u64 	%rd8, reducesum$__cuda_local_var_34394_32_non_const_sdata;
	add.s64 	%rd2, %rd8, %rd7;
	st.shared.f32 	[%rd2], %f29;
	bar.sync 	0;
	setp.lt.u32	%p3, %r14, 66;
	@%p3 bra 	BB2_6;

BB2_3:
	mov.u32 	%r7, %r14;
	shr.u32 	%r14, %r7, 1;
	setp.ge.u32	%p4, %r2, %r14;
	@%p4 bra 	BB2_5;

	ld.shared.f32 	%f6, [%rd2];
	add.s32 	%r12, %r14, %r2;
	mul.wide.u32 	%rd9, %r12, 4;
	add.s64 	%rd11, %rd8, %rd9;
	ld.shared.f32 	%f7, [%rd11];
	add.f32 	%f8, %f6, %f7;
	st.shared.f32 	[%rd2], %f8;

BB2_5:
	bar.sync 	0;
	setp.gt.u32	%p5, %r7, 131;
	@%p5 bra 	BB2_3;

BB2_6:
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB2_8;

	ld.volatile.shared.f32 	%f9, [%rd2];
	ld.volatile.shared.f32 	%f10, [%rd2+128];
	add.f32 	%f11, %f9, %f10;
	st.volatile.shared.f32 	[%rd2], %f11;
	ld.volatile.shared.f32 	%f12, [%rd2+64];
	ld.volatile.shared.f32 	%f13, [%rd2];
	add.f32 	%f14, %f13, %f12;
	st.volatile.shared.f32 	[%rd2], %f14;
	ld.volatile.shared.f32 	%f15, [%rd2+32];
	ld.volatile.shared.f32 	%f16, [%rd2];
	add.f32 	%f17, %f16, %f15;
	st.volatile.shared.f32 	[%rd2], %f17;
	ld.volatile.shared.f32 	%f18, [%rd2+16];
	ld.volatile.shared.f32 	%f19, [%rd2];
	add.f32 	%f20, %f19, %f18;
	st.volatile.shared.f32 	[%rd2], %f20;
	ld.volatile.shared.f32 	%f21, [%rd2+8];
	ld.volatile.shared.f32 	%f22, [%rd2];
	add.f32 	%f23, %f22, %f21;
	st.volatile.shared.f32 	[%rd2], %f23;
	ld.volatile.shared.f32 	%f24, [%rd2+4];
	ld.volatile.shared.f32 	%f25, [%rd2];
	add.f32 	%f26, %f25, %f24;
	st.volatile.shared.f32 	[%rd2], %f26;

BB2_8:
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB2_10;

	cvta.to.global.u64 	%rd12, %rd3;
	ld.shared.f32 	%f27, [reducesum$__cuda_local_var_34394_32_non_const_sdata];
	atom.global.add.f32 	%f28, [%rd12], %f27;

BB2_10:
	ret;
}


`
)
