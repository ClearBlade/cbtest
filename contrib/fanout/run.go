package fanout

import (
	"fmt"
	"reflect"
	"testing"
)

// runHandler is an alias for raw interface{}. Used as an alias since it looks
// better and more clear in the documentation.
type runHandler interface{}

func Run(t tWithRunParallel, name string, numParallel int, fn runHandler) bool {

	fnval, err := requireRunHandler(fn)
	if err != nil {
		t.Error(err)
		t.FailNow()
		return false
	}

	t.Run("Worker", func(t *testing.T) {
		for idx := 0; idx < numParallel; idx++ {

			currentIdx := idx
			name := fmt.Sprintf("%d", currentIdx)

			t.Run(name, func(t *testing.T) {
				t.Parallel()
				fnval.Call([]reflect.Value{reflect.ValueOf(t), reflect.ValueOf(currentIdx)})
			})

		}
	})
	return false
}

func requireRunHandler(fn runHandler) (reflect.Value, error) {

	var err error
	reterr := fmt.Errorf("runHandler must be a function that takes a testing type and index; and returns any type")

	value := reflect.ValueOf(fn)

	if value.Kind() != reflect.Func {
		return reflect.ValueOf(nil), reterr
	}

	fntype := value.Type()

	if fntype.NumIn() != 2 {
		return reflect.ValueOf(nil), reterr
	}

	if fntype.NumOut() != 0 && fntype.NumOut() != 1 {
		return reflect.ValueOf(nil), reterr
	}

	err = requireTestingType(fntype.In(0))
	if err != nil {
		return reflect.ValueOf(nil), reterr
	}

	return value, nil
}

func requireTestingType(t reflect.Type) error {
	return nil
}
