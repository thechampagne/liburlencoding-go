package urlencoding

/*
#cgo LDFLAGS: -lurlencoding

#include <stdlib.h>
#include <stddef.h>

extern char* url_encoding_encode(const char* data);

extern char* url_encoding_encode_binary(const char* data, size_t length);

extern char* url_encoding_decode(const char* data);

extern char* url_encoding_decode_binary(const char* data, size_t length);

extern void url_encoding_free(char* ptr);
*/
import "C"
import "unsafe"

// Percent-encodes every byte except alphanumerics and -, _, ., ~. Assumes UTF-8 encoding.
// 
// Example:
// * *
// res := urlencoding.Encode("This is string will be encoded.")
// fmt.Println(res)
// * *
// 
// @param data
// @return encoded string
func Encode(data string) string {
	cstr := C.CString(data)
	defer C.free(unsafe.Pointer(cstr))
	res := C.url_encoding_encode(cstr)
	if res == nil {
		return ""
	}
	str := C.GoString(res)
	C.url_encoding_free(res);
	return str
}

// Percent-encodes every byte except alphanumerics and -, _, ., ~.
// 
// Example:
// * *
// res := urlencoding.EncodeBinary("This is string will be encoded.")
// fmt.Println(res)
// * *
// 
// @param data
// @return encoded string
func EncodeBinary(data string) string {
	cstr := C.CString(data)
	defer C.free(unsafe.Pointer(cstr))
	res := C.url_encoding_encode_binary(cstr, C.size_t(len(data)))
	if res == nil {
		return ""
	}
	str := C.GoString(res)
	C.url_encoding_free(res);
	return str
}

// Decode percent-encoded string assuming UTF-8 encoding.
// 
// Example:
// * *
// res := urlencoding.Decode("%F0%9F%91%BE%20Exterminate%21")
// fmt.Println(res)
// * *
// 
// @param data
// @return decoded string
func Decode(data string) string {
	cstr := C.CString(data)
	defer C.free(unsafe.Pointer(cstr))
	res := C.url_encoding_decode(cstr)
	if res == nil {
		return ""
	}
	str := C.GoString(res)
	C.url_encoding_free(res);
	return str
}

// Decode percent-encoded string as binary data, in any encoding.
// 
// Example:
// * *
// res := urlencoding.DecodeBinary("%F1%F2%F3%C0%C1%C2")
// fmt.Println(res)
// * *
// 
// @param data
// @return decoded string
func DecodeBinary(data string) string {
	cstr := C.CString(data)
	defer C.free(unsafe.Pointer(cstr))
	res := C.url_encoding_decode_binary(cstr, C.size_t(len(data)))
	if res == nil {
		return ""
	}
	str := C.GoString(res)
	C.url_encoding_free(res);
	return str
}