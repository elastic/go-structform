// This file has been generated from 'unfold_sel.yml', do not edit
package gotype

func (u *unfoldCtx) trySetGotypeTarget(to interface{}) bool {
	switch ptr := to.(type) {

	// primitives

	case *bool:
		u.unfolder.push(newUnfolderBool(u, ptr))

	case *string:
		u.unfolder.push(newUnfolderString(u, ptr))

	case *uint:
		u.unfolder.push(newUnfolderUint(u, ptr))

	case *uint8:
		u.unfolder.push(newUnfolderUint8(u, ptr))

	case *uint16:
		u.unfolder.push(newUnfolderUint16(u, ptr))

	case *uint32:
		u.unfolder.push(newUnfolderUint32(u, ptr))

	case *uint64:
		u.unfolder.push(newUnfolderUint64(u, ptr))

	case *int:
		u.unfolder.push(newUnfolderInt(u, ptr))

	case *int8:
		u.unfolder.push(newUnfolderInt8(u, ptr))

	case *int16:
		u.unfolder.push(newUnfolderInt16(u, ptr))

	case *int32:
		u.unfolder.push(newUnfolderInt32(u, ptr))

	case *int64:
		u.unfolder.push(newUnfolderInt64(u, ptr))

	case *float32:
		u.unfolder.push(newUnfolderFloat32(u, ptr))

	case *float64:
		u.unfolder.push(newUnfolderFloat64(u, ptr))

	// arrays

	case *[]bool:
		u.unfolder.push(newUnfolderArrBool(u, ptr))
		u.detail.push(unfoldWaitStart)

	case *[]string:
		u.unfolder.push(newUnfolderArrString(u, ptr))
		u.detail.push(unfoldWaitStart)

	case *[]uint:
		u.unfolder.push(newUnfolderArrUint(u, ptr))
		u.detail.push(unfoldWaitStart)

	case *[]uint8:
		u.unfolder.push(newUnfolderArrUint8(u, ptr))
		u.detail.push(unfoldWaitStart)

	case *[]uint16:
		u.unfolder.push(newUnfolderArrUint16(u, ptr))
		u.detail.push(unfoldWaitStart)

	case *[]uint32:
		u.unfolder.push(newUnfolderArrUint32(u, ptr))
		u.detail.push(unfoldWaitStart)

	case *[]uint64:
		u.unfolder.push(newUnfolderArrUint64(u, ptr))
		u.detail.push(unfoldWaitStart)

	case *[]int:
		u.unfolder.push(newUnfolderArrInt(u, ptr))
		u.detail.push(unfoldWaitStart)

	case *[]int8:
		u.unfolder.push(newUnfolderArrInt8(u, ptr))
		u.detail.push(unfoldWaitStart)

	case *[]int16:
		u.unfolder.push(newUnfolderArrInt16(u, ptr))
		u.detail.push(unfoldWaitStart)

	case *[]int32:
		u.unfolder.push(newUnfolderArrInt32(u, ptr))
		u.detail.push(unfoldWaitStart)

	case *[]int64:
		u.unfolder.push(newUnfolderArrInt64(u, ptr))
		u.detail.push(unfoldWaitStart)

	case *[]float32:
		u.unfolder.push(newUnfolderArrFloat32(u, ptr))
		u.detail.push(unfoldWaitStart)

	case *[]float64:
		u.unfolder.push(newUnfolderArrFloat64(u, ptr))
		u.detail.push(unfoldWaitStart)

	// maps

	case *map[string]bool:
		u.unfolder.push(newUnfolderMapBool(u, ptr))
		u.detail.push(unfoldWaitStart)

	case *map[string]string:
		u.unfolder.push(newUnfolderMapString(u, ptr))
		u.detail.push(unfoldWaitStart)

	case *map[string]uint:
		u.unfolder.push(newUnfolderMapUint(u, ptr))
		u.detail.push(unfoldWaitStart)

	case *map[string]uint8:
		u.unfolder.push(newUnfolderMapUint8(u, ptr))
		u.detail.push(unfoldWaitStart)

	case *map[string]uint16:
		u.unfolder.push(newUnfolderMapUint16(u, ptr))
		u.detail.push(unfoldWaitStart)

	case *map[string]uint32:
		u.unfolder.push(newUnfolderMapUint32(u, ptr))
		u.detail.push(unfoldWaitStart)

	case *map[string]uint64:
		u.unfolder.push(newUnfolderMapUint64(u, ptr))
		u.detail.push(unfoldWaitStart)

	case *map[string]int:
		u.unfolder.push(newUnfolderMapInt(u, ptr))
		u.detail.push(unfoldWaitStart)

	case *map[string]int8:
		u.unfolder.push(newUnfolderMapInt8(u, ptr))
		u.detail.push(unfoldWaitStart)

	case *map[string]int16:
		u.unfolder.push(newUnfolderMapInt16(u, ptr))
		u.detail.push(unfoldWaitStart)

	case *map[string]int32:
		u.unfolder.push(newUnfolderMapInt32(u, ptr))
		u.detail.push(unfoldWaitStart)

	case *map[string]int64:
		u.unfolder.push(newUnfolderMapInt64(u, ptr))
		u.detail.push(unfoldWaitStart)

	case *map[string]float32:
		u.unfolder.push(newUnfolderMapFloat32(u, ptr))
		u.detail.push(unfoldWaitStart)

	case *map[string]float64:
		u.unfolder.push(newUnfolderMapFloat64(u, ptr))
		u.detail.push(unfoldWaitStart)

	default:
		return false
	}
	return true
}
