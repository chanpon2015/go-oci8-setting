package main

import (
	"fmt"
	"reflect"
	"testing"
)

// ポインタpの指す先にvの値を入れる関数。
// 不適切な値があったらエラーを返す。
// 変更するのはSetの内部だけ。
func Set(p, v interface{}) error {
	pv := reflect.ValueOf(p)
	// TODO: pがポインタかどうかKindをチェック
	test := pv.Elem()
	fmt.Println(test.Kind())
	if pv.Kind() != reflect.Ptr {
		return fmt.Errorf("is not pointer")
	}

	vv := reflect.ValueOf(v)
	// TODO： pに値がセットできるかつ、pの先にvが代入できるかチェック
	// ヒント1： https://golang.org/pkg/reflect/#Type のAssignableTo
	// ヒント2: https://golang.org/pkg/reflect/#Value.CanSet
	// fmt.Println(pv.Elem().Type(), vv.Type())
	fmt.Println("aaa", pv.Elem().Type().Name(), vv.Type().Name())
	if !vv.Type().AssignableTo(pv.Elem().Type()) {
		return fmt.Errorf("err")
	}
	if pv.CanSet() {
		return fmt.Errorf("is set found")
	}

	// Elemでポインタの指してる先を取得して代入
	pv.Elem().Set(vv)
	return nil

}

type myint int

func (n myint) String() string {
	return fmt.Sprint(int(n))
}

func TestSet(t *testing.T) {
	var (
		v1 int
		v2 fmt.Stringer
	)

	cases := map[string]struct {
		p, v      interface{}
		expectErr bool
	}{
		"ok":       {&v1, 100, false},
		"stringer": {&v2, myint(100), false},
		// "myint":    {&v1, myint(100), true},
		// "nil":      {nil, 100, true},
		// "string":   {&v1, "hoge", true},
		// "float64":  {&v1, 1.11, true},
		// "func":     {&v1, func() {}, true},
	}

	for tn, tc := range cases {
		tc := tc
		t.Run(tn, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Fatal("panicが起きた:", r)
				}
			}()

			err := Set(tc.p, tc.v)
			switch {
			case tc.expectErr && err == nil:
				t.Fatal("期待したエラーが発生しませんでした")
			case !tc.expectErr && err != nil:
				t.Fatal("予想外のエラー:", err)
			case err == nil:
				v := reflect.ValueOf(tc.p).Elem().Interface()
				if !reflect.DeepEqual(v, tc.v) {
					t.Errorf("want %v but got %v", tc.v, v)
				}
			}
		})
	}
}
