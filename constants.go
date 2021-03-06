/*
 * Ported from three.js by @rydrman
 */

package three

const (
	REVISION                         string = "85"
	MouseLeft                        int    = 0
	MouseMiddle                             = 1
	MouseRight                              = 2
	CullFaceNone                            = 0
	CullFaceBack                            = 1
	CullFaceFront                           = 2
	CullFaceFrontBack                       = 3
	FrontFaceDirectionCW                    = 0
	FrontFaceDirectionCCW                   = 1
	BasicShadowMap                          = 0
	PCFShadowMap                            = 1
	PCFSoftShadowMap                        = 2
	FrontSide                               = 0
	BackSide                                = 1
	DoubleSide                              = 2
	FlatShading                             = 1
	SmoothShading                           = 2
	NoColors                                = 0
	FaceColors                              = 1
	VertexColors                            = 2
	NoBlending                              = 0
	NormalBlending                          = 1
	AdditiveBlending                        = 2
	SubtractiveBlending                     = 3
	MultiplyBlending                        = 4
	CustomBlending                          = 5
	AddEquation                             = 100
	SubtractEquation                        = 101
	ReverseSubtractEquation                 = 102
	MinEquation                             = 103
	MaxEquation                             = 104
	ZeroFactor                              = 200
	OneFactor                               = 201
	SrcColorFactor                          = 202
	OneMinusSrcColorFactor                  = 203
	SrcAlphaFactor                          = 204
	OneMinusSrcAlphaFactor                  = 205
	DstAlphaFactor                          = 206
	OneMinusDstAlphaFactor                  = 207
	DstColorFactor                          = 208
	OneMinusDstColorFactor                  = 209
	SrcAlphaSaturateFactor                  = 210
	NeverDepth                              = 0
	AlwaysDepth                             = 1
	LessDepth                               = 2
	LessEqualDepth                          = 3
	EqualDepth                              = 4
	GreaterEqualDepth                       = 5
	GreaterDepth                            = 6
	NotEqualDepth                           = 7
	MultiplyOperation                       = 0
	MixOperation                            = 1
	AddOperation                            = 2
	NoToneMapping                           = 0
	LinearToneMapping                       = 1
	ReinhardToneMapping                     = 2
	Uncharted2ToneMapping                   = 3
	CineonToneMapping                       = 4
	UVMapping                               = 300
	CubeReflectionMapping                   = 301
	CubeRefractionMapping                   = 302
	EquirectangularReflectionMapping        = 303
	EquirectangularRefractionMapping        = 304
	SphericalReflectionMapping              = 305
	CubeUVReflectionMapping                 = 306
	CubeUVRefractionMapping                 = 307
	RepeatWrapping                          = 1000
	ClampToEdgeWrapping                     = 1001
	MirroredRepeatWrapping                  = 1002
	NearestFilter                           = 1003
	NearestMipMapNearestFilter              = 1004
	NearestMipMapLinearFilter               = 1005
	LinearFilter                            = 1006
	LinearMipMapNearestFilter               = 1007
	LinearMipMapLinearFilter                = 1008
	UnsignedByteType                        = 1009
	ByteType                                = 1010
	ShortType                               = 1011
	UnsignedShortType                       = 1012
	IntType                                 = 1013
	UnsignedIntType                         = 1014
	FloatType                               = 1015
	HalfFloatType                           = 1016
	UnsignedShort4444Type                   = 1017
	UnsignedShort5551Type                   = 1018
	UnsignedShort565Type                    = 1019
	UnsignedInt248Type                      = 1020
	AlphaFormat                             = 1021
	RGBFormat                               = 1022
	RGBAFormat                              = 1023
	LuminanceFormat                         = 1024
	LuminanceAlphaFormat                    = 1025
	RGBEFormat                              = RGBAFormat
	DepthFormat                             = 1026
	DepthStencilFormat                      = 1027
	RGBS3TCDXT1Format                       = 2001
	RGBAS3TCDXT1Format                      = 2002
	RGBAS3TCDXT3Format                      = 2003
	RGBAS3TCDXT5Format                      = 2004
	RGBPVRTC4BPPV1Format                    = 2100
	RGBPVRTC2BPPV1Format                    = 2101
	RGBAPVRTC4BPPV1Format                   = 2102
	RGBAPVRTC2BPPV1Format                   = 2103
	RGBETC1Format                           = 2151
	LoopOnce                                = 2200
	LoopRepeat                              = 2201
	LoopPingPong                            = 2202
	InterpolateDiscrete                     = 2300
	InterpolateLinear                       = 2301
	InterpolateSmooth                       = 2302
	ZeroCurvatureEnding                     = 2400
	ZeroSlopeEnding                         = 2401
	WrapAroundEnding                        = 2402
	TrianglesDrawMode                       = 0
	TriangleStripDrawMode                   = 1
	TriangleFanDrawMode                     = 2
	LinearEncoding                          = 3000
	sRGBEncoding                            = 3001
	GammaEncoding                           = 3007
	RGBEEncoding                            = 3002
	LogLuvEncoding                          = 3003
	RGBM7Encoding                           = 3004
	RGBM16Encoding                          = 3005
	RGBDEncoding                            = 3006
	BasicDepthPacking                       = 3200
	RGBADepthPacking                        = 3201
)
