// import "github.com/nekonbu72/sjson/sjson"

package sjson

import (
	"encoding/json"
	"os"
)

// OpenDecode ...
// 使用側のパッケージでは、構造体をフィールド含め大文字で定義すること
// .
func OpenDecode(path string, v interface{}) error {

	r, err := os.Open(path)
	defer r.Close()
	if err != nil {
		return err
	}

	if err := json.NewDecoder(r).Decode(v); err != nil {
		return err
	}

	return nil
}
