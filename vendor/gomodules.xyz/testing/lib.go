/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package testing

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"

	"github.com/pkg/errors"
	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
	"gomodules.xyz/encoding/yaml"
)

var rootDir = func() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Dir(filepath.Dir(filepath.Dir(filepath.Dir(b))))
}()

func ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filepath.Join(rootDir, filename))
}

func LoadFile(filename string) (map[string]any, error) {
	data, err := os.ReadFile(filepath.Join(rootDir, filename))
	if err != nil {
		return nil, err
	}

	var obj map[string]any
	err = yaml.Unmarshal(data, &obj)
	return obj, err
}

func Diff(oldObj, newObj any) error {
	old, err := json.Marshal(oldObj)
	if err != nil {
		return err
	}

	nu, err := json.Marshal(newObj)
	if err != nil {
		return err
	}

	return DiffJSON(old, nu)
}

func DiffJSON(old, nu []byte) error {
	differ := diff.New()
	d, err := differ.Compare(old, nu)
	if err != nil {
		return err
	}

	if d.Modified() {
		var original map[string]any
		err := yaml.Unmarshal(old, &original)
		if err != nil {
			return err
		}

		config := formatter.AsciiFormatterConfig{
			ShowArrayIndex: true,
			Coloring:       true,
		}

		f := formatter.NewAsciiFormatter(original, config)
		result, err := f.Format(d)
		if err != nil {
			return err
		}
		return errors.New(result)
	}
	return nil
}

func RoundTripFile(filename string, v any) error {
	if reflect.TypeOf(v).Kind() != reflect.Pointer {
		return fmt.Errorf("v is expected to be a pointer, found %T", v)
	}

	data, err := ReadFile(filename)
	if err != nil {
		return err
	}
	old, err := yaml.ToJSON(data)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, v)
	if err != nil {
		return err
	}
	nu, err := json.Marshal(v)
	if err != nil {
		return err
	}

	return DiffJSON(old, nu)
}
