package listmonk

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_request_setParam(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name  string
		query url.Values
		args  args
		want  url.Values
	}{
		{
			name:  "set string",
			query: url.Values{},
			args: args{
				key:   "foo",
				value: "bar",
			},
			want: url.Values{
				"foo": []string{"bar"},
			},
		},
		{
			name:  "set slice",
			query: url.Values{},
			args: args{
				key:   "foo",
				value: []string{"bar1", "bar2"},
			},
			want: url.Values{
				"foo": []string{"bar1", "bar2"},
			},
		},
		{
			name: "replace string",
			query: url.Values{
				"foo": []string{"old"},
			},
			args: args{
				key:   "foo",
				value: "bar",
			},
			want: url.Values{
				"foo": []string{"bar"},
			},
		},
		{
			name: "replace slice",
			query: url.Values{
				"foo": []string{"old1", "old2"},
			},
			args: args{
				key:   "foo",
				value: []string{"bar1", "bar2"},
			},
			want: url.Values{
				"foo": []string{"bar1", "bar2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &request{
				query: tt.query,
			}
			r = r.setParam(tt.args.key, tt.args.value)
			assert.Equal(t, tt.want, r.query)
		})
	}
}
