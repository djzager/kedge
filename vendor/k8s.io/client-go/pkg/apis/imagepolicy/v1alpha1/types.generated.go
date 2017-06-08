/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// ************************************************************
// DO NOT EDIT.
// THIS FILE IS AUTO-GENERATED BY codecgen.
// ************************************************************

package v1alpha1

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	time "time"

	codec1978 "github.com/ugorji/go/codec"
	pkg1_unversioned "k8s.io/client-go/pkg/api/unversioned"
	pkg2_v1 "k8s.io/client-go/pkg/api/v1"
	pkg3_types "k8s.io/client-go/pkg/types"
)

const (
	// ----- content types ----
	codecSelferC_UTF81234 = 1
	codecSelferC_RAW1234  = 0
	// ----- value types used ----
	codecSelferValueTypeArray1234 = 10
	codecSelferValueTypeMap1234   = 9
	// ----- containerStateValues ----
	codecSelfer_containerMapKey1234    = 2
	codecSelfer_containerMapValue1234  = 3
	codecSelfer_containerMapEnd1234    = 4
	codecSelfer_containerArrayElem1234 = 6
	codecSelfer_containerArrayEnd1234  = 7
)

var (
	codecSelferBitsize1234                         = uint8(reflect.TypeOf(uint(0)).Bits())
	codecSelferOnlyMapOrArrayEncodeToStructErr1234 = errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer1234 struct{}

func init() {
	if codec1978.GenVersion != 5 {
		_, file, _, _ := runtime.Caller(0)
		err := fmt.Errorf("codecgen version mismatch: current: %v, need %v. Re-generate file: %v",
			5, codec1978.GenVersion, file)
		panic(err)
	}
	if false { // reference the types, but skip this branch at build/run time
		var v0 pkg1_unversioned.TypeMeta
		var v1 pkg2_v1.ObjectMeta
		var v2 pkg3_types.UID
		var v3 time.Time
		_, _, _, _ = v0, v1, v2, v3
	}
}

func (x *ImageReview) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [5]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = x.Kind != ""
			yyq2[1] = x.APIVersion != ""
			yyq2[2] = true
			yyq2[4] = true
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(5)
			} else {
				yynn2 = 1
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.EncodeMapStart(yynn2)
				yynn2 = 0
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[0] {
					yym4 := z.EncBinary()
					_ = yym4
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("kind"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[1] {
					yym7 := z.EncBinary()
					_ = yym7
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("apiVersion"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym8 := z.EncBinary()
					_ = yym8
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[2] {
					yy10 := &x.ObjectMeta
					yy10.CodecEncodeSelf(e)
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("metadata"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy11 := &x.ObjectMeta
					yy11.CodecEncodeSelf(e)
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				yy13 := &x.Spec
				yy13.CodecEncodeSelf(e)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("spec"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				yy14 := &x.Spec
				yy14.CodecEncodeSelf(e)
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[4] {
					yy16 := &x.Status
					yy16.CodecEncodeSelf(e)
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[4] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("status"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy17 := &x.Status
					yy17.CodecEncodeSelf(e)
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *ImageReview) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym18 := z.DecBinary()
	_ = yym18
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct19 := r.ContainerType()
		if yyct19 == codecSelferValueTypeMap1234 {
			yyl19 := r.ReadMapStart()
			if yyl19 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl19, d)
			}
		} else if yyct19 == codecSelferValueTypeArray1234 {
			yyl19 := r.ReadArrayStart()
			if yyl19 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl19, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *ImageReview) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys20Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys20Slc
	var yyhl20 bool = l >= 0
	for yyj20 := 0; ; yyj20++ {
		if yyhl20 {
			if yyj20 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys20Slc = r.DecodeBytes(yys20Slc, true, true)
		yys20 := string(yys20Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys20 {
		case "kind":
			if r.TryDecodeAsNil() {
				x.Kind = ""
			} else {
				x.Kind = string(r.DecodeString())
			}
		case "apiVersion":
			if r.TryDecodeAsNil() {
				x.APIVersion = ""
			} else {
				x.APIVersion = string(r.DecodeString())
			}
		case "metadata":
			if r.TryDecodeAsNil() {
				x.ObjectMeta = pkg2_v1.ObjectMeta{}
			} else {
				yyv23 := &x.ObjectMeta
				yyv23.CodecDecodeSelf(d)
			}
		case "spec":
			if r.TryDecodeAsNil() {
				x.Spec = ImageReviewSpec{}
			} else {
				yyv24 := &x.Spec
				yyv24.CodecDecodeSelf(d)
			}
		case "status":
			if r.TryDecodeAsNil() {
				x.Status = ImageReviewStatus{}
			} else {
				yyv25 := &x.Status
				yyv25.CodecDecodeSelf(d)
			}
		default:
			z.DecStructFieldNotFound(-1, yys20)
		} // end switch yys20
	} // end for yyj20
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *ImageReview) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj26 int
	var yyb26 bool
	var yyhl26 bool = l >= 0
	yyj26++
	if yyhl26 {
		yyb26 = yyj26 > l
	} else {
		yyb26 = r.CheckBreak()
	}
	if yyb26 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		x.Kind = string(r.DecodeString())
	}
	yyj26++
	if yyhl26 {
		yyb26 = yyj26 > l
	} else {
		yyb26 = r.CheckBreak()
	}
	if yyb26 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		x.APIVersion = string(r.DecodeString())
	}
	yyj26++
	if yyhl26 {
		yyb26 = yyj26 > l
	} else {
		yyb26 = r.CheckBreak()
	}
	if yyb26 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.ObjectMeta = pkg2_v1.ObjectMeta{}
	} else {
		yyv29 := &x.ObjectMeta
		yyv29.CodecDecodeSelf(d)
	}
	yyj26++
	if yyhl26 {
		yyb26 = yyj26 > l
	} else {
		yyb26 = r.CheckBreak()
	}
	if yyb26 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Spec = ImageReviewSpec{}
	} else {
		yyv30 := &x.Spec
		yyv30.CodecDecodeSelf(d)
	}
	yyj26++
	if yyhl26 {
		yyb26 = yyj26 > l
	} else {
		yyb26 = r.CheckBreak()
	}
	if yyb26 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Status = ImageReviewStatus{}
	} else {
		yyv31 := &x.Status
		yyv31.CodecDecodeSelf(d)
	}
	for {
		yyj26++
		if yyhl26 {
			yyb26 = yyj26 > l
		} else {
			yyb26 = r.CheckBreak()
		}
		if yyb26 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj26-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *ImageReviewSpec) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym32 := z.EncBinary()
		_ = yym32
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep33 := !z.EncBinary()
			yy2arr33 := z.EncBasicHandle().StructToArray
			var yyq33 [3]bool
			_, _, _ = yysep33, yyq33, yy2arr33
			const yyr33 bool = false
			yyq33[0] = len(x.Containers) != 0
			yyq33[1] = len(x.Annotations) != 0
			yyq33[2] = x.Namespace != ""
			var yynn33 int
			if yyr33 || yy2arr33 {
				r.EncodeArrayStart(3)
			} else {
				yynn33 = 0
				for _, b := range yyq33 {
					if b {
						yynn33++
					}
				}
				r.EncodeMapStart(yynn33)
				yynn33 = 0
			}
			if yyr33 || yy2arr33 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq33[0] {
					if x.Containers == nil {
						r.EncodeNil()
					} else {
						yym35 := z.EncBinary()
						_ = yym35
						if false {
						} else {
							h.encSliceImageReviewContainerSpec(([]ImageReviewContainerSpec)(x.Containers), e)
						}
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq33[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("containers"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					if x.Containers == nil {
						r.EncodeNil()
					} else {
						yym36 := z.EncBinary()
						_ = yym36
						if false {
						} else {
							h.encSliceImageReviewContainerSpec(([]ImageReviewContainerSpec)(x.Containers), e)
						}
					}
				}
			}
			if yyr33 || yy2arr33 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq33[1] {
					if x.Annotations == nil {
						r.EncodeNil()
					} else {
						yym38 := z.EncBinary()
						_ = yym38
						if false {
						} else {
							z.F.EncMapStringStringV(x.Annotations, false, e)
						}
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq33[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("annotations"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					if x.Annotations == nil {
						r.EncodeNil()
					} else {
						yym39 := z.EncBinary()
						_ = yym39
						if false {
						} else {
							z.F.EncMapStringStringV(x.Annotations, false, e)
						}
					}
				}
			}
			if yyr33 || yy2arr33 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq33[2] {
					yym41 := z.EncBinary()
					_ = yym41
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Namespace))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq33[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("namespace"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym42 := z.EncBinary()
					_ = yym42
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Namespace))
					}
				}
			}
			if yyr33 || yy2arr33 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *ImageReviewSpec) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym43 := z.DecBinary()
	_ = yym43
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct44 := r.ContainerType()
		if yyct44 == codecSelferValueTypeMap1234 {
			yyl44 := r.ReadMapStart()
			if yyl44 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl44, d)
			}
		} else if yyct44 == codecSelferValueTypeArray1234 {
			yyl44 := r.ReadArrayStart()
			if yyl44 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl44, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *ImageReviewSpec) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys45Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys45Slc
	var yyhl45 bool = l >= 0
	for yyj45 := 0; ; yyj45++ {
		if yyhl45 {
			if yyj45 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys45Slc = r.DecodeBytes(yys45Slc, true, true)
		yys45 := string(yys45Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys45 {
		case "containers":
			if r.TryDecodeAsNil() {
				x.Containers = nil
			} else {
				yyv46 := &x.Containers
				yym47 := z.DecBinary()
				_ = yym47
				if false {
				} else {
					h.decSliceImageReviewContainerSpec((*[]ImageReviewContainerSpec)(yyv46), d)
				}
			}
		case "annotations":
			if r.TryDecodeAsNil() {
				x.Annotations = nil
			} else {
				yyv48 := &x.Annotations
				yym49 := z.DecBinary()
				_ = yym49
				if false {
				} else {
					z.F.DecMapStringStringX(yyv48, false, d)
				}
			}
		case "namespace":
			if r.TryDecodeAsNil() {
				x.Namespace = ""
			} else {
				x.Namespace = string(r.DecodeString())
			}
		default:
			z.DecStructFieldNotFound(-1, yys45)
		} // end switch yys45
	} // end for yyj45
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *ImageReviewSpec) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj51 int
	var yyb51 bool
	var yyhl51 bool = l >= 0
	yyj51++
	if yyhl51 {
		yyb51 = yyj51 > l
	} else {
		yyb51 = r.CheckBreak()
	}
	if yyb51 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Containers = nil
	} else {
		yyv52 := &x.Containers
		yym53 := z.DecBinary()
		_ = yym53
		if false {
		} else {
			h.decSliceImageReviewContainerSpec((*[]ImageReviewContainerSpec)(yyv52), d)
		}
	}
	yyj51++
	if yyhl51 {
		yyb51 = yyj51 > l
	} else {
		yyb51 = r.CheckBreak()
	}
	if yyb51 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Annotations = nil
	} else {
		yyv54 := &x.Annotations
		yym55 := z.DecBinary()
		_ = yym55
		if false {
		} else {
			z.F.DecMapStringStringX(yyv54, false, d)
		}
	}
	yyj51++
	if yyhl51 {
		yyb51 = yyj51 > l
	} else {
		yyb51 = r.CheckBreak()
	}
	if yyb51 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Namespace = ""
	} else {
		x.Namespace = string(r.DecodeString())
	}
	for {
		yyj51++
		if yyhl51 {
			yyb51 = yyj51 > l
		} else {
			yyb51 = r.CheckBreak()
		}
		if yyb51 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj51-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *ImageReviewContainerSpec) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym57 := z.EncBinary()
		_ = yym57
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep58 := !z.EncBinary()
			yy2arr58 := z.EncBasicHandle().StructToArray
			var yyq58 [1]bool
			_, _, _ = yysep58, yyq58, yy2arr58
			const yyr58 bool = false
			yyq58[0] = x.Image != ""
			var yynn58 int
			if yyr58 || yy2arr58 {
				r.EncodeArrayStart(1)
			} else {
				yynn58 = 0
				for _, b := range yyq58 {
					if b {
						yynn58++
					}
				}
				r.EncodeMapStart(yynn58)
				yynn58 = 0
			}
			if yyr58 || yy2arr58 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq58[0] {
					yym60 := z.EncBinary()
					_ = yym60
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Image))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq58[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("image"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym61 := z.EncBinary()
					_ = yym61
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Image))
					}
				}
			}
			if yyr58 || yy2arr58 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *ImageReviewContainerSpec) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym62 := z.DecBinary()
	_ = yym62
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct63 := r.ContainerType()
		if yyct63 == codecSelferValueTypeMap1234 {
			yyl63 := r.ReadMapStart()
			if yyl63 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl63, d)
			}
		} else if yyct63 == codecSelferValueTypeArray1234 {
			yyl63 := r.ReadArrayStart()
			if yyl63 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl63, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *ImageReviewContainerSpec) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys64Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys64Slc
	var yyhl64 bool = l >= 0
	for yyj64 := 0; ; yyj64++ {
		if yyhl64 {
			if yyj64 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys64Slc = r.DecodeBytes(yys64Slc, true, true)
		yys64 := string(yys64Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys64 {
		case "image":
			if r.TryDecodeAsNil() {
				x.Image = ""
			} else {
				x.Image = string(r.DecodeString())
			}
		default:
			z.DecStructFieldNotFound(-1, yys64)
		} // end switch yys64
	} // end for yyj64
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *ImageReviewContainerSpec) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj66 int
	var yyb66 bool
	var yyhl66 bool = l >= 0
	yyj66++
	if yyhl66 {
		yyb66 = yyj66 > l
	} else {
		yyb66 = r.CheckBreak()
	}
	if yyb66 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Image = ""
	} else {
		x.Image = string(r.DecodeString())
	}
	for {
		yyj66++
		if yyhl66 {
			yyb66 = yyj66 > l
		} else {
			yyb66 = r.CheckBreak()
		}
		if yyb66 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj66-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *ImageReviewStatus) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym68 := z.EncBinary()
		_ = yym68
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep69 := !z.EncBinary()
			yy2arr69 := z.EncBasicHandle().StructToArray
			var yyq69 [2]bool
			_, _, _ = yysep69, yyq69, yy2arr69
			const yyr69 bool = false
			yyq69[1] = x.Reason != ""
			var yynn69 int
			if yyr69 || yy2arr69 {
				r.EncodeArrayStart(2)
			} else {
				yynn69 = 1
				for _, b := range yyq69 {
					if b {
						yynn69++
					}
				}
				r.EncodeMapStart(yynn69)
				yynn69 = 0
			}
			if yyr69 || yy2arr69 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				yym71 := z.EncBinary()
				_ = yym71
				if false {
				} else {
					r.EncodeBool(bool(x.Allowed))
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("allowed"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				yym72 := z.EncBinary()
				_ = yym72
				if false {
				} else {
					r.EncodeBool(bool(x.Allowed))
				}
			}
			if yyr69 || yy2arr69 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq69[1] {
					yym74 := z.EncBinary()
					_ = yym74
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Reason))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq69[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("reason"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym75 := z.EncBinary()
					_ = yym75
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Reason))
					}
				}
			}
			if yyr69 || yy2arr69 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *ImageReviewStatus) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym76 := z.DecBinary()
	_ = yym76
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct77 := r.ContainerType()
		if yyct77 == codecSelferValueTypeMap1234 {
			yyl77 := r.ReadMapStart()
			if yyl77 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl77, d)
			}
		} else if yyct77 == codecSelferValueTypeArray1234 {
			yyl77 := r.ReadArrayStart()
			if yyl77 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl77, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *ImageReviewStatus) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys78Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys78Slc
	var yyhl78 bool = l >= 0
	for yyj78 := 0; ; yyj78++ {
		if yyhl78 {
			if yyj78 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys78Slc = r.DecodeBytes(yys78Slc, true, true)
		yys78 := string(yys78Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys78 {
		case "allowed":
			if r.TryDecodeAsNil() {
				x.Allowed = false
			} else {
				x.Allowed = bool(r.DecodeBool())
			}
		case "reason":
			if r.TryDecodeAsNil() {
				x.Reason = ""
			} else {
				x.Reason = string(r.DecodeString())
			}
		default:
			z.DecStructFieldNotFound(-1, yys78)
		} // end switch yys78
	} // end for yyj78
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *ImageReviewStatus) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj81 int
	var yyb81 bool
	var yyhl81 bool = l >= 0
	yyj81++
	if yyhl81 {
		yyb81 = yyj81 > l
	} else {
		yyb81 = r.CheckBreak()
	}
	if yyb81 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Allowed = false
	} else {
		x.Allowed = bool(r.DecodeBool())
	}
	yyj81++
	if yyhl81 {
		yyb81 = yyj81 > l
	} else {
		yyb81 = r.CheckBreak()
	}
	if yyb81 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Reason = ""
	} else {
		x.Reason = string(r.DecodeString())
	}
	for {
		yyj81++
		if yyhl81 {
			yyb81 = yyj81 > l
		} else {
			yyb81 = r.CheckBreak()
		}
		if yyb81 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj81-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x codecSelfer1234) encSliceImageReviewContainerSpec(v []ImageReviewContainerSpec, e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.EncodeArrayStart(len(v))
	for _, yyv84 := range v {
		z.EncSendContainerState(codecSelfer_containerArrayElem1234)
		yy85 := &yyv84
		yy85.CodecEncodeSelf(e)
	}
	z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x codecSelfer1234) decSliceImageReviewContainerSpec(v *[]ImageReviewContainerSpec, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv86 := *v
	yyh86, yyl86 := z.DecSliceHelperStart()
	var yyc86 bool
	if yyl86 == 0 {
		if yyv86 == nil {
			yyv86 = []ImageReviewContainerSpec{}
			yyc86 = true
		} else if len(yyv86) != 0 {
			yyv86 = yyv86[:0]
			yyc86 = true
		}
	} else if yyl86 > 0 {
		var yyrr86, yyrl86 int
		var yyrt86 bool
		if yyl86 > cap(yyv86) {

			yyrg86 := len(yyv86) > 0
			yyv286 := yyv86
			yyrl86, yyrt86 = z.DecInferLen(yyl86, z.DecBasicHandle().MaxInitLen, 16)
			if yyrt86 {
				if yyrl86 <= cap(yyv86) {
					yyv86 = yyv86[:yyrl86]
				} else {
					yyv86 = make([]ImageReviewContainerSpec, yyrl86)
				}
			} else {
				yyv86 = make([]ImageReviewContainerSpec, yyrl86)
			}
			yyc86 = true
			yyrr86 = len(yyv86)
			if yyrg86 {
				copy(yyv86, yyv286)
			}
		} else if yyl86 != len(yyv86) {
			yyv86 = yyv86[:yyl86]
			yyc86 = true
		}
		yyj86 := 0
		for ; yyj86 < yyrr86; yyj86++ {
			yyh86.ElemContainerState(yyj86)
			if r.TryDecodeAsNil() {
				yyv86[yyj86] = ImageReviewContainerSpec{}
			} else {
				yyv87 := &yyv86[yyj86]
				yyv87.CodecDecodeSelf(d)
			}

		}
		if yyrt86 {
			for ; yyj86 < yyl86; yyj86++ {
				yyv86 = append(yyv86, ImageReviewContainerSpec{})
				yyh86.ElemContainerState(yyj86)
				if r.TryDecodeAsNil() {
					yyv86[yyj86] = ImageReviewContainerSpec{}
				} else {
					yyv88 := &yyv86[yyj86]
					yyv88.CodecDecodeSelf(d)
				}

			}
		}

	} else {
		yyj86 := 0
		for ; !r.CheckBreak(); yyj86++ {

			if yyj86 >= len(yyv86) {
				yyv86 = append(yyv86, ImageReviewContainerSpec{}) // var yyz86 ImageReviewContainerSpec
				yyc86 = true
			}
			yyh86.ElemContainerState(yyj86)
			if yyj86 < len(yyv86) {
				if r.TryDecodeAsNil() {
					yyv86[yyj86] = ImageReviewContainerSpec{}
				} else {
					yyv89 := &yyv86[yyj86]
					yyv89.CodecDecodeSelf(d)
				}

			} else {
				z.DecSwallow()
			}

		}
		if yyj86 < len(yyv86) {
			yyv86 = yyv86[:yyj86]
			yyc86 = true
		} else if yyj86 == 0 && yyv86 == nil {
			yyv86 = []ImageReviewContainerSpec{}
			yyc86 = true
		}
	}
	yyh86.End()
	if yyc86 {
		*v = yyv86
	}
}
