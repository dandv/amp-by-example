// Copyright Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hello

import (
	"fmt"
	"net/http"
	"strings"
)

const MAX_AGE_IN_SECONDS = 60 * 60 * 24 // 1 day
const OLD_ADDRESS = "amp-by-example.appspot.com"
const NEW_ADDRESS = "https://ampbyexample.com"

func init() {
	fs := http.FileServer(http.Dir("dist"))
	http.Handle("/", RedirectDomain(fs))
}

func RedirectDomain(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Host == OLD_ADDRESS ||
			(r.TLS == nil && !strings.HasPrefix(r.Host, "localhost")) {
			http.Redirect(w, r, NEW_ADDRESS+r.URL.Path, http.StatusMovedPermanently)
		} else {
			w.Header().Add("Cache-Control", fmt.Sprintf("max-age=%d, public, must-revalidate, proxy-revalidate", MAX_AGE_IN_SECONDS))
			h.ServeHTTP(w, r)
		}
	})
}
