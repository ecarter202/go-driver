//
// DISCLAIMER
//
// Copyright 2017 ArangoDB GmbH, Cologne, Germany
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
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//
// Author Ewout Prangsma
//

package test

type UserDoc struct {
	Name string `arango:"name"`
	Age  int    `arango:"age"`
}

type UserDocWithKey struct {
	Key  string `arango:"_key"`
	Name string `arango:"name"`
	Age  int    `arango:"age"`
}

type Account struct {
	ID   string   `arango:"id"`
	User *UserDoc `arango:"user"`
}

type Book struct {
	Title string
}

type RouteEdge struct {
	From     string `arango:"_from,omitempty"`
	To       string `arango:"_to,omitempty"`
	Distance int    `arango:"distance,omitempty"`
}

type RouteEdgeWithKey struct {
	Key      string `arango:"_key"`
	From     string `arango:"_from,omitempty"`
	To       string `arango:"_to,omitempty"`
	Distance int    `arango:"distance,omitempty"`
}

type RelationEdge struct {
	From string `arango:"_from,omitempty"`
	To   string `arango:"_to,omitempty"`
	Type string `arango:"type,omitempty"`
}

type AccountEdge struct {
	From string   `arango:"_from,omitempty"`
	To   string   `arango:"_to,omitempty"`
	User *UserDoc `arango:"user"`
}
