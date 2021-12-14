package callGo
/*
#cgo LDFLAGS: -landroid

#include <jni.h>

// Equivalent to:
// Activity.Finish()
void show_toast(uintptr_t java_vm, uintptr_t jni_env, jobject ctx) {
	JavaVM* vm = (JavaVM*)java_vm;
	JNIEnv* env = (JNIEnv*)jni_env;
	jclass clazz = (*env)->GetObjectClass(env, ctx);
	jmethodID showToast = (*env)->GetMethodID(env, clazz, "showToast", "(Ljava/lang/String;)V");
	jstring message = (*env)->NewStringUTF(env,"hello java, I from go!");
	(*env)->CallVoidMethod(env, ctx, showToast, message);
}
 */
import "C"
import (
	"golang.org/x/mobile/app"
)

func ShowToast() {
	app.RunOnJVM(func(vm, jniEnv, ctx uintptr) error {
		C.show_toast(C.uintptr_t(vm), C.uintptr_t(jniEnv), C.jobject(ctx))
		return nil
	})
}
