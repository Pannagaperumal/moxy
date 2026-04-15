package vm

import "pebble/package/object"

func (vm *VM) push(o object.Object) error {
	if vm.sp >= StackSize {
		return ErrStackOverflow
	}

	vm.stack[vm.sp] = o
	vm.sp++

	return nil
}

func (vm *VM) pop() object.Object {
	o := vm.stack[vm.sp-1]
	vm.sp--
	return o
}

func (vm *VM) LastPoppedStackElem() object.Object {
	return vm.stack[vm.sp]
}
