// +build fedora

package gtkspell

/*
#include <gtk/gtk.h>
#include <gtkspell/gtkspell.h>
#include <stdlib.h>

static GtkTextView* to_GtkTextView(void* w) { return GTK_TEXT_VIEW(w); }
static inline gchar* to_gcharptr(const char* s) { return (gchar*)s; }
*/
// #cgo pkg-config: gtkspell3-3.0 gtk+-3.0
import "C"
import "unsafe"

import "github.com/agl/go-gtk/glib"
import "github.com/agl/go-gtk/gtk"

//-----------------------------------------------------------------------
// GtkSpellChecker
//-----------------------------------------------------------------------
type GtkSpellChecker struct {
	Spell *C.GtkSpellChecker
}

func New(textview *gtk.GtkTextView, language string) (*GtkSpellChecker, *glib.Error) {
	var lang *C.char
	if len(language) > 0 {
		lang = C.CString(language)
		defer C.free(unsafe.Pointer(lang))
	}

	var gerror *C.GError
	v := C.gtk_spell_checker_new()
	if C.gtk_spell_checker_set_language(v, (*C.gchar)(unsafe.Pointer(lang)), &gerror) == 0 {
		return nil, glib.ErrorFromNative(unsafe.Pointer(gerror))
	}
	C.gtk_spell_checker_attach(v, C.to_GtkTextView(unsafe.Pointer(textview.Widget)))
	return &GtkSpellChecker{v}, nil
}

func (spell *GtkSpellChecker) SetLanguage(language string) *glib.Error {
	lang := C.CString(language)
	defer C.free(unsafe.Pointer(lang))

	var gerror *C.GError
	if C.gtk_spell_checker_set_language(spell.Spell, C.to_gcharptr(lang), &gerror) == 0 {
		return glib.ErrorFromNative(unsafe.Pointer(gerror))
	}
	return nil
}

func (spell *GtkSpellChecker) Recheck() {
	C.gtk_spell_checker_recheck_all(spell.Spell)
}
