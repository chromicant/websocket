package websocket

import "syscall/js"

// Callback is for backward compatibility. Use Func instead.
type Callback = js.Func

// EventCallbackFlag is for backward compatibility.
type EventCallbackFlag int

const (
	PreventDefault EventCallbackFlag = 1 << iota
	StopPropagation
	StopImmediatePropagation
)

// NewCallback is for backward compatibility. Use FuncOf instead.
func NewCallback(fn func([]js.Value)) Callback {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go func() {
			fn(args)
		}()
		return nil
	})
}

// NewEventCallback is for backward compatibility. Use FuncOf instead.
func NewEventCallback(flags EventCallbackFlag, fn func(event js.Value)) Callback {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		if flags&PreventDefault != 0 {
			e.Call("preventDefault")
		}
		if flags&StopPropagation != 0 {
			e.Call("stopPropagation")
		}
		if flags&StopImmediatePropagation != 0 {
			e.Call("stopImmediatePropagation")
		}
		go func() {
			fn(e)
		}()
		return nil
	})
}


