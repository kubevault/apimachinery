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
	"os"
	"path/filepath"
	"runtime"

	"github.com/pkg/errors"
	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
	"sigs.k8s.io/yaml"
)

var rootDir = func() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Dir(filepath.Dir(b))
}()

func ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filepath.Join(rootDir, filename))
}

func CheckDiff(oldObj, newObj interface{}) error {
	old, err := json.Marshal(oldObj)
	if err != nil {
		return err
	}

	new, err := json.Marshal(newObj)
	if err != nil {
		return err
	}

	differ := diff.New()
	d, err := differ.Compare(old, new)
	if err != nil {
		return err
	}

	if d.Modified() {
		var original map[string]interface{}
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
