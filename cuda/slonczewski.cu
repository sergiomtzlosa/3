// Original implementation by Mykola Dvornik for mumax2
// Modified for mumax3 by Arne Vansteenkiste, 2013

#include "float3.h"
#include "constants.h"

extern "C" __global__ void
addslonczewskitorque(float* __restrict__ tx, float* __restrict__ ty, float* __restrict__ tz,
                     float* __restrict__ mx, float* __restrict__ my, float* __restrict__ mz,
                     float* __restrict__ jx, float* __restrict__ jy, float* __restrict__ jz,
                     float* __restrict__ pxLUT, float* __restrict__ pyLUT, float* __restrict__ pzLUT,
                     float* __restrict__ msatLUT,
                     float* __restrict__ alphaLUT,
                     float* __restrict__ fltLUT,
                     float* __restrict__ polLUT,
                     float* __restrict__ lambdaLUT,
                     float* __restrict__ epsilonPrimeLUT,
                     int8_t* __restrict__ regions,
                     int N) {

    int I =  ( blockIdx.y*gridDim.x + blockIdx.x ) * blockDim.x + threadIdx.x;
    if (I < N) {

        float3 m = make_float3(mx[I], my[I], mz[I]);
        float3 J = make_float3(jx[I], jy[I], jz[I]);

        // read parameters
        int8_t region       = regions[I];
        float3 p            = normalized(make_float3(pxLUT[region], pyLUT[region], pzLUT[region]));
        float  Ms           = msatLUT[region];
        float  alpha        = alphaLUT[region];
        float  flt          = fltLUT[region]; // free-layer thickness
        float  pol          = polLUT[region];
        float  lambda       = lambdaLUT[region];
        float  epsilonPrime = epsilonPrimeLUT[region];

        if (len(J) == 0.0f || Ms == 0.0f || flt == 0.0f) {
            return;
        }

        float3 pxm      = cross(p, m);
        float3 mxpxm    = cross(m, pxm);
        float  pdotm    = dot(p, m);

        // derived parameters
        float alphaFac   = 1.0f / (1.0f + alpha * alpha);
        float beta       = HBAR * GAMMA0 / (MU0 * QE * Ms);    // njn is missing ??
        float beta_prime = pol * beta;                         // epsilon is missing??
        float pre_fld    = beta * epsilonPrime;
        float lambda2    = lambda * lambda;
        float epsilon    = lambda2 / ((lambda2 + 1.0f) + (lambda2 - 1.0f) * pdotm);

        float Jdir  = dot(make_float3(1.0f, 1.0f, 1.0f), normalized(J));
        float Jsign = Jdir / fabsf(Jdir);
        float nJn   = len(J) * Jsign;

        float preX = beta_prime *nJn*flt*epsilon*pol / Ms;
        float preY = pre_fld *nJn*flt*epsilonPrime / Ms;

        float  mxpxmFac = alphaFac*(preX - alpha * preY);
        float  pxmFac   = alphaFac*(preY - alpha * preX);

        tx[I] += mxpxmFac * mxpxm.x + pxmFac * pxm.x;
        ty[I] += mxpxmFac * mxpxm.y + pxmFac * pxm.y;
        tz[I] += mxpxmFac * mxpxm.z + pxmFac * pxm.z;
    }
}

