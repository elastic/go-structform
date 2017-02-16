// This file has been generated from 'unfold_sel.yml', do not edit
package gotype

import "unsafe"

func (u *unfoldCtx) trySetGotypeTarget(to interface{}) bool {
	switch ptr := to.(type) {

	// primitives

	case *bool:
		u.ptr.push(unsafe.Pointer(ptr))
		u.unfolder.push(newUnfolderBool())

	case *string:
		u.ptr.push(unsafe.Pointer(ptr))
		u.unfolder.push(newUnfolderString())

	case *uint:
		u.ptr.push(unsafe.Pointer(ptr))
		u.unfolder.push(newUnfolderUint())

	case *uint8:
		u.ptr.push(unsafe.Pointer(ptr))
		u.unfolder.push(newUnfolderUint8())

	case *uint16:
		u.ptr.push(unsafe.Pointer(ptr))
		u.unfolder.push(newUnfolderUint16())

	case *uint32:
		u.ptr.push(unsafe.Pointer(ptr))
		u.unfolder.push(newUnfolderUint32())

	case *uint64:
		u.ptr.push(unsafe.Pointer(ptr))
		u.unfolder.push(newUnfolderUint64())

	case *int:
		u.ptr.push(unsafe.Pointer(ptr))
		u.unfolder.push(newUnfolderInt())

	case *int8:
		u.ptr.push(unsafe.Pointer(ptr))
		u.unfolder.push(newUnfolderInt8())

	case *int16:
		u.ptr.push(unsafe.Pointer(ptr))
		u.unfolder.push(newUnfolderInt16())

	case *int32:
		u.ptr.push(unsafe.Pointer(ptr))
		u.unfolder.push(newUnfolderInt32())

	case *int64:
		u.ptr.push(unsafe.Pointer(ptr))
		u.unfolder.push(newUnfolderInt64())

	case *float32:
		u.ptr.push(unsafe.Pointer(ptr))
		u.unfolder.push(newUnfolderFloat32())

	case *float64:
		u.ptr.push(unsafe.Pointer(ptr))
		u.unfolder.push(newUnfolderFloat64())

	// arrays

	case *[]bool:
		u.ptr.push(unsafe.Pointer(ptr))
		u.detail.push(unfoldWaitStart)
		u.unfolder.push(newUnfolderArrBool())

	case *[]string:
		u.ptr.push(unsafe.Pointer(ptr))
		u.detail.push(unfoldWaitStart)
		u.unfolder.push(newUnfolderArrString())

	case *[]uint:
		u.ptr.push(unsafe.Pointer(ptr))
		u.detail.push(unfoldWaitStart)
		u.unfolder.push(newUnfolderArrUint())

	case *[]uint8:
		u.ptr.push(unsafe.Pointer(ptr))
		u.detail.push(unfoldWaitStart)
		u.unfolder.push(newUnfolderArrUint8())

	case *[]uint16:
		u.ptr.push(unsafe.Pointer(ptr))
		u.detail.push(unfoldWaitStart)
		u.unfolder.push(newUnfolderArrUint16())

	case *[]uint32:
		u.ptr.push(unsafe.Pointer(ptr))
		u.detail.push(unfoldWaitStart)
		u.unfolder.push(newUnfolderArrUint32())

	case *[]uint64:
		u.ptr.push(unsafe.Pointer(ptr))
		u.detail.push(unfoldWaitStart)
		u.unfolder.push(newUnfolderArrUint64())

	case *[]int:
		u.ptr.push(unsafe.Pointer(ptr))
		u.detail.push(unfoldWaitStart)
		u.unfolder.push(newUnfolderArrInt())

	case *[]int8:
		u.ptr.push(unsafe.Pointer(ptr))
		u.detail.push(unfoldWaitStart)
		u.unfolder.push(newUnfolderArrInt8())

	case *[]int16:
		u.ptr.push(unsafe.Pointer(ptr))
		u.detail.push(unfoldWaitStart)
		u.unfolder.push(newUnfolderArrInt16())

	case *[]int32:
		u.ptr.push(unsafe.Pointer(ptr))
		u.detail.push(unfoldWaitStart)
		u.unfolder.push(newUnfolderArrInt32())

	case *[]int64:
		u.ptr.push(unsafe.Pointer(ptr))
		u.detail.push(unfoldWaitStart)
		u.unfolder.push(newUnfolderArrInt64())

	case *[]float32:
		u.ptr.push(unsafe.Pointer(ptr))
		u.detail.push(unfoldWaitStart)
		u.unfolder.push(newUnfolderArrFloat32())

	case *[]float64:
		u.ptr.push(unsafe.Pointer(ptr))
		u.detail.push(unfoldWaitStart)
		u.unfolder.push(newUnfolderArrFloat64())

	// maps

	case *map[string]bool:
		u.ptr.push(unsafe.Pointer(ptr))
		u.detail.push(unfoldWaitStart)
		u.unfolder.push(newUnfolderMapBool())

	case *map[string]string:
		u.ptr.push(unsafe.Pointer(ptr))
		u.detail.push(unfoldWaitStart)
		u.unfolder.push(newUnfolderMapString())

	case *map[string]uint:
		u.ptr.push(unsafe.Pointer(ptr))
		u.detail.push(unfoldWaitStart)
		u.unfolder.push(newUnfolderMapUint())

	case *map[string]uint8:
		u.ptr.push(unsafe.Pointer(ptr))
		u.detail.push(unfoldWaitStart)
		u.unfolder.push(newUnfolderMapUint8())

	case *map[string]uint16:
		u.ptr.push(unsafe.Pointer(ptr))
		u.detail.push(unfoldWaitStart)
		u.unfolder.push(newUnfolderMapUint16())

	case *map[string]uint32:
		u.ptr.push(unsafe.Pointer(ptr))
		u.detail.push(unfoldWaitStart)
		u.unfolder.push(newUnfolderMapUint32())

	case *map[string]uint64:
		u.ptr.push(unsafe.Pointer(ptr))
		u.detail.push(unfoldWaitStart)
		u.unfolder.push(newUnfolderMapUint64())

	case *map[string]int:
		u.ptr.push(unsafe.Pointer(ptr))
		u.detail.push(unfoldWaitStart)
		u.unfolder.push(newUnfolderMapInt())

	case *map[string]int8:
		u.ptr.push(unsafe.Pointer(ptr))
		u.detail.push(unfoldWaitStart)
		u.unfolder.push(newUnfolderMapInt8())

	case *map[string]int16:
		u.ptr.push(unsafe.Pointer(ptr))
		u.detail.push(unfoldWaitStart)
		u.unfolder.push(newUnfolderMapInt16())

	case *map[string]int32:
		u.ptr.push(unsafe.Pointer(ptr))
		u.detail.push(unfoldWaitStart)
		u.unfolder.push(newUnfolderMapInt32())

	case *map[string]int64:
		u.ptr.push(unsafe.Pointer(ptr))
		u.detail.push(unfoldWaitStart)
		u.unfolder.push(newUnfolderMapInt64())

	case *map[string]float32:
		u.ptr.push(unsafe.Pointer(ptr))
		u.detail.push(unfoldWaitStart)
		u.unfolder.push(newUnfolderMapFloat32())

	case *map[string]float64:
		u.ptr.push(unsafe.Pointer(ptr))
		u.detail.push(unfoldWaitStart)
		u.unfolder.push(newUnfolderMapFloat64())

	default:
		return false
	}
	return true
}
