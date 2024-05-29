#!/bin/sh

# Script to generate info.go.

outfile=${1:-/dev/stdout}

cat <<'EOF' > $outfile
// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// SetGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Also passes any
// provided key-value pairs. Intended for use by Google-written clients.
//
// Internal use only.

package generativelanguage

//go:generate ./gen_info.sh info.go

EOF

awk '/^func \(c \*[A-Z].*\) setGoogleClientInfo/ {
  printf("func (c %s SetGoogleClientInfo(keyval ...string) {\n", $3);
	printf("  c.setGoogleClientInfo(keyval...)\n");
  printf("}\n\n");
}' *_client.go >> $outfile

gofmt -w $outfile
