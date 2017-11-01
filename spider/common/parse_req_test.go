package common

import (
	"Vua-Crawler/spider/model"
	"reflect"
	"testing"
)

func TestParseReq(t *testing.T) {
	type args struct {
		r     []*model.Request
		query map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want []*model.Request
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseReq(tt.args.r, tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRuleRequests(t *testing.T) {
	type args struct {
		r     *model.Request
		query map[string]interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  []*model.Request
		want1 bool
	}{
		// TODO: Add test cases.
		{
			name: "easy test",
			args: args{
				r: &model.Request{
					Url: "http://{1-10, 1}",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getRuleRequests(tt.args.r, tt.args.query)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRuleRequests() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getRuleRequests() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getRuleAndMatch(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name      string
		args      args
		wantRule  string
		wantMatch bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRule, gotMatch := getRuleAndMatch(tt.args.url)
			if gotRule != tt.wantRule {
				t.Errorf("getRuleAndMatch() gotRule = %v, want %v", gotRule, tt.wantRule)
			}
			if gotMatch != tt.wantMatch {
				t.Errorf("getRuleAndMatch() gotMatch = %v, want %v", gotMatch, tt.wantMatch)
			}
		})
	}
}

func TestParseOffset(t *testing.T) {
	type args struct {
		r    *model.Request
		rule string
	}
	tests := []struct {
		name  string
		args  args
		want  []*model.Request
		want1 bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ParseOffset(tt.args.r, tt.args.rule)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseOffset() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ParseOffset() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
