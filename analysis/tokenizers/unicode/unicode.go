//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package unicode

import (
	"github.com/blevesearch/segment"

	"github.com/blevesearch/bleve/analysis"
	"github.com/blevesearch/bleve/registry"
)

const Name = "unicode"

type UnicodeTokenizer struct {
}

func NewUnicodeTokenizer() *UnicodeTokenizer {
	return &UnicodeTokenizer{}
}

func (rt *UnicodeTokenizer) Tokenize(input []byte) analysis.TokenStream {

	rv := make(analysis.TokenStream, 0)

	segmenter := segment.NewWordSegmenterDirect(input)
	start := 0
	pos := 1
	for segmenter.Segment() {
		segmentBytes := segmenter.Bytes()
		end := start + len(segmentBytes)
		if segmenter.Type() != segment.None {
			token := analysis.Token{
				Term:     segmentBytes,
				Start:    start,
				End:      end,
				Position: pos,
				Type:     convertType(segmenter.Type()),
			}
			rv = append(rv, &token)
			pos++
		}
		start = end

	}
	return rv
}

func UnicodeTokenizerConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.Tokenizer, error) {
	return NewUnicodeTokenizer(), nil
}

func init() {
	registry.RegisterTokenizer(Name, UnicodeTokenizerConstructor)
}

func convertType(segmentWordType int) analysis.TokenType {
	switch segmentWordType {
	case segment.Ideo:
		return analysis.Ideographic
	case segment.Kana:
		return analysis.Ideographic
	case segment.Number:
		return analysis.Numeric
	}
	return analysis.AlphaNumeric
}
