package gdk
/*
#include <gdk/gdk.h>
#include <gdk/gdkx.h>
// #cgo pkg-config: gdk-3.0
*/
import "C"

func (v *GdkWindow) GetNativeWindowID() int32 {
	return int32(C.gdk_x11_window_get_xid(v.Window))
}

