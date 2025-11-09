package client

import (
	"context"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestLoginResponse_Get(t *testing.T) {
	type args struct {
		fieldName string
	}
	tests := []struct {
		name      string
		a         LoginResponse
		args      args
		wantValue interface{}
		wantFound bool
	}{
		{
			name: "get existing field",
			a: LoginResponse{
				AdditionalProperties: map[string]interface{}{
					"custom_field": "custom_value",
				},
			},
			args:      args{fieldName: "custom_field"},
			wantValue: "custom_value",
			wantFound: true,
		},
		{
			name: "get non-existing field",
			a: LoginResponse{
				AdditionalProperties: map[string]interface{}{},
			},
			args:      args{fieldName: "missing"},
			wantValue: nil,
			wantFound: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotFound := tt.a.Get(tt.args.fieldName)
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("LoginResponse.Get() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotFound != tt.wantFound {
				t.Errorf("LoginResponse.Get() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}

func TestLoginResponse_Set(t *testing.T) {
	type args struct {
		fieldName string
		value     interface{}
	}
	tests := []struct {
		name string
		a    *LoginResponse
		args args
	}{
		{
			name: "set new field",
			a:    &LoginResponse{},
			args: args{fieldName: "test", value: "value"},
		},
		{
			name: "overwrite existing",
			a: &LoginResponse{
				AdditionalProperties: map[string]interface{}{"test": "old"},
			},
			args: args{fieldName: "test", value: "new"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.Set(tt.args.fieldName, tt.args.value)
			val, found := tt.a.Get(tt.args.fieldName)
			if !found {
				t.Errorf("LoginResponse.Set() failed to set field")
			}
			if !reflect.DeepEqual(val, tt.args.value) {
				t.Errorf("LoginResponse.Set() value = %v, want %v", val, tt.args.value)
			}
		})
	}
}

func TestLoginResponse_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		a       *LoginResponse
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("LoginResponse.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoginResponse_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		a       LoginResponse
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginResponse.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoginResponse.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScriptResult_AsV1Script(t *testing.T) {
	tests := []struct {
		name    string
		tr      ScriptResult
		want    V1Script
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tr.AsV1Script()
			if (err != nil) != tt.wantErr {
				t.Errorf("ScriptResult.AsV1Script() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScriptResult.AsV1Script() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScriptResult_FromV1Script(t *testing.T) {
	type args struct {
		v V1Script
	}
	tests := []struct {
		name    string
		tr      *ScriptResult
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.FromV1Script(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("ScriptResult.FromV1Script() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestScriptResult_MergeV1Script(t *testing.T) {
	type args struct {
		v V1Script
	}
	tests := []struct {
		name    string
		tr      *ScriptResult
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.MergeV1Script(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("ScriptResult.MergeV1Script() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestScriptResult_AsV2Script(t *testing.T) {
	tests := []struct {
		name    string
		tr      ScriptResult
		want    V2Script
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tr.AsV2Script()
			if (err != nil) != tt.wantErr {
				t.Errorf("ScriptResult.AsV2Script() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScriptResult.AsV2Script() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScriptResult_FromV2Script(t *testing.T) {
	type args struct {
		v V2Script
	}
	tests := []struct {
		name    string
		tr      *ScriptResult
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.FromV2Script(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("ScriptResult.FromV2Script() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestScriptResult_MergeV2Script(t *testing.T) {
	type args struct {
		v V2Script
	}
	tests := []struct {
		name    string
		tr      *ScriptResult
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.MergeV2Script(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("ScriptResult.MergeV2Script() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestScriptResult_Discriminator(t *testing.T) {
	tests := []struct {
		name    string
		tr      ScriptResult
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tr.Discriminator()
			if (err != nil) != tt.wantErr {
				t.Errorf("ScriptResult.Discriminator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ScriptResult.Discriminator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScriptResult_ValueByDiscriminator(t *testing.T) {
	tests := []struct {
		name    string
		tr      ScriptResult
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tr.ValueByDiscriminator()
			if (err != nil) != tt.wantErr {
				t.Errorf("ScriptResult.ValueByDiscriminator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScriptResult.ValueByDiscriminator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScriptResult_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		tr      ScriptResult
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tr.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("ScriptResult.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScriptResult.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScriptResult_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		tr      *ScriptResult
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("ScriptResult.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLegacyActionResponse_AsV1Script(t *testing.T) {
	tests := []struct {
		name    string
		tr      LegacyActionResponse
		want    V1Script
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tr.AsV1Script()
			if (err != nil) != tt.wantErr {
				t.Errorf("LegacyActionResponse.AsV1Script() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LegacyActionResponse.AsV1Script() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLegacyActionResponse_FromV1Script(t *testing.T) {
	type args struct {
		v V1Script
	}
	tests := []struct {
		name    string
		tr      *LegacyActionResponse
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.FromV1Script(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("LegacyActionResponse.FromV1Script() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLegacyActionResponse_MergeV1Script(t *testing.T) {
	type args struct {
		v V1Script
	}
	tests := []struct {
		name    string
		tr      *LegacyActionResponse
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.MergeV1Script(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("LegacyActionResponse.MergeV1Script() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLegacyActionResponse_AsV2Script(t *testing.T) {
	tests := []struct {
		name    string
		tr      LegacyActionResponse
		want    V2Script
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tr.AsV2Script()
			if (err != nil) != tt.wantErr {
				t.Errorf("LegacyActionResponse.AsV2Script() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LegacyActionResponse.AsV2Script() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLegacyActionResponse_FromV2Script(t *testing.T) {
	type args struct {
		v V2Script
	}
	tests := []struct {
		name    string
		tr      *LegacyActionResponse
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.FromV2Script(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("LegacyActionResponse.FromV2Script() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLegacyActionResponse_MergeV2Script(t *testing.T) {
	type args struct {
		v V2Script
	}
	tests := []struct {
		name    string
		tr      *LegacyActionResponse
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.MergeV2Script(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("LegacyActionResponse.MergeV2Script() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLegacyActionResponse_AsLegacyScriptAttachment(t *testing.T) {
	tests := []struct {
		name    string
		tr      LegacyActionResponse
		want    LegacyScriptAttachment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tr.AsLegacyScriptAttachment()
			if (err != nil) != tt.wantErr {
				t.Errorf("LegacyActionResponse.AsLegacyScriptAttachment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LegacyActionResponse.AsLegacyScriptAttachment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLegacyActionResponse_FromLegacyScriptAttachment(t *testing.T) {
	type args struct {
		v LegacyScriptAttachment
	}
	tests := []struct {
		name    string
		tr      *LegacyActionResponse
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.FromLegacyScriptAttachment(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("LegacyActionResponse.FromLegacyScriptAttachment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLegacyActionResponse_MergeLegacyScriptAttachment(t *testing.T) {
	type args struct {
		v LegacyScriptAttachment
	}
	tests := []struct {
		name    string
		tr      *LegacyActionResponse
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.MergeLegacyScriptAttachment(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("LegacyActionResponse.MergeLegacyScriptAttachment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLegacyActionResponse_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		tr      LegacyActionResponse
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tr.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("LegacyActionResponse.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LegacyActionResponse.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLegacyActionResponse_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		tr      *LegacyActionResponse
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("LegacyActionResponse.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		server string
		opts   []ClientOption
	}
	tests := []struct {
		name    string
		args    args
		want    *Client
		wantErr bool
	}{
		{
			name:    "valid server",
			args:    args{server: "https://test.com"},
			wantErr: false,
		},
		{
			name:    "valid server with options",
			args:    args{server: "https://test.com", opts: []ClientOption{WithHTTPClient(&http.Client{})}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.args.server, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("NewClient() returned nil")
			}
		})
	}
}

func TestWithHTTPClient(t *testing.T) {
	type args struct {
		doer HttpRequestDoer
	}
	tests := []struct {
		name string
		args args
		want ClientOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithHTTPClient(tt.args.doer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithHTTPClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithRequestEditorFn(t *testing.T) {
	type args struct {
		fn RequestEditorFn
	}
	tests := []struct {
		name string
		args args
		want ClientOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithRequestEditorFn(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithRequestEditorFn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_InvokeLegacyAction(t *testing.T) {
	type args struct {
		ctx        context.Context
		params     *InvokeLegacyActionParams
		reqEditors []RequestEditorFn
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.InvokeLegacyAction(tt.args.ctx, tt.args.params, tt.args.reqEditors...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.InvokeLegacyAction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.InvokeLegacyAction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_LoginWithPasswordWithBody(t *testing.T) {
	type args struct {
		ctx         context.Context
		contentType string
		body        io.Reader
		reqEditors  []RequestEditorFn
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.LoginWithPasswordWithBody(tt.args.ctx, tt.args.contentType, tt.args.body, tt.args.reqEditors...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.LoginWithPasswordWithBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.LoginWithPasswordWithBody() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_LoginWithPassword(t *testing.T) {
	type args struct {
		ctx        context.Context
		body       LoginWithPasswordJSONRequestBody
		reqEditors []RequestEditorFn
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.LoginWithPassword(tt.args.ctx, tt.args.body, tt.args.reqEditors...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.LoginWithPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.LoginWithPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_LoginWithAccessKeyWithBody(t *testing.T) {
	type args struct {
		ctx         context.Context
		contentType string
		body        io.Reader
		reqEditors  []RequestEditorFn
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.LoginWithAccessKeyWithBody(tt.args.ctx, tt.args.contentType, tt.args.body, tt.args.reqEditors...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.LoginWithAccessKeyWithBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.LoginWithAccessKeyWithBody() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_LoginWithAccessKey(t *testing.T) {
	type args struct {
		ctx        context.Context
		body       LoginWithAccessKeyJSONRequestBody
		reqEditors []RequestEditorFn
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.LoginWithAccessKey(tt.args.ctx, tt.args.body, tt.args.reqEditors...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.LoginWithAccessKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.LoginWithAccessKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetScript(t *testing.T) {
	type args struct {
		ctx        context.Context
		scriptId   ScriptIdPathParam
		reqEditors []RequestEditorFn
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetScript(tt.args.ctx, tt.args.scriptId, tt.args.reqEditors...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetScript() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetScript() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_ArchiveScript(t *testing.T) {
	type args struct {
		ctx        context.Context
		scriptId   ScriptIdPathParam
		reqEditors []RequestEditorFn
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.ArchiveScript(tt.args.ctx, tt.args.scriptId, tt.args.reqEditors...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ArchiveScript() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.ArchiveScript() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_RedactScript(t *testing.T) {
	type args struct {
		ctx        context.Context
		scriptId   ScriptIdPathParam
		reqEditors []RequestEditorFn
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.RedactScript(tt.args.ctx, tt.args.scriptId, tt.args.reqEditors...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.RedactScript() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.RedactScript() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInvokeLegacyActionRequest(t *testing.T) {
	type args struct {
		server string
		params *InvokeLegacyActionParams
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewInvokeLegacyActionRequest(tt.args.server, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewInvokeLegacyActionRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInvokeLegacyActionRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLoginWithPasswordRequest(t *testing.T) {
	type args struct {
		server string
		body   LoginWithPasswordJSONRequestBody
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLoginWithPasswordRequest(tt.args.server, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLoginWithPasswordRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoginWithPasswordRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLoginWithPasswordRequestWithBody(t *testing.T) {
	type args struct {
		server      string
		contentType string
		body        io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLoginWithPasswordRequestWithBody(tt.args.server, tt.args.contentType, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLoginWithPasswordRequestWithBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoginWithPasswordRequestWithBody() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLoginWithAccessKeyRequest(t *testing.T) {
	type args struct {
		server string
		body   LoginWithAccessKeyJSONRequestBody
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLoginWithAccessKeyRequest(tt.args.server, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLoginWithAccessKeyRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoginWithAccessKeyRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLoginWithAccessKeyRequestWithBody(t *testing.T) {
	type args struct {
		server      string
		contentType string
		body        io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLoginWithAccessKeyRequestWithBody(tt.args.server, tt.args.contentType, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLoginWithAccessKeyRequestWithBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoginWithAccessKeyRequestWithBody() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewGetScriptRequest(t *testing.T) {
	type args struct {
		server   string
		scriptId ScriptIdPathParam
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGetScriptRequest(tt.args.server, tt.args.scriptId)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGetScriptRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGetScriptRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewArchiveScriptRequest(t *testing.T) {
	type args struct {
		server   string
		scriptId ScriptIdPathParam
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewArchiveScriptRequest(tt.args.server, tt.args.scriptId)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewArchiveScriptRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArchiveScriptRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRedactScriptRequest(t *testing.T) {
	type args struct {
		server   string
		scriptId ScriptIdPathParam
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRedactScriptRequest(tt.args.server, tt.args.scriptId)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRedactScriptRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRedactScriptRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_applyEditors(t *testing.T) {
	type args struct {
		ctx               context.Context
		req               *http.Request
		additionalEditors []RequestEditorFn
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.applyEditors(tt.args.ctx, tt.args.req, tt.args.additionalEditors); (err != nil) != tt.wantErr {
				t.Errorf("Client.applyEditors() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewClientWithResponses(t *testing.T) {
	type args struct {
		server string
		opts   []ClientOption
	}
	tests := []struct {
		name    string
		args    args
		want    *ClientWithResponses
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClientWithResponses(tt.args.server, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClientWithResponses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClientWithResponses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithBaseURL(t *testing.T) {
	type args struct {
		baseURL string
	}
	tests := []struct {
		name string
		args args
		want ClientOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithBaseURL(tt.args.baseURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithBaseURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInvokeLegacyActionResponse_Status(t *testing.T) {
	tests := []struct {
		name string
		r    InvokeLegacyActionResponse
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Status(); got != tt.want {
				t.Errorf("InvokeLegacyActionResponse.Status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInvokeLegacyActionResponse_StatusCode(t *testing.T) {
	tests := []struct {
		name string
		r    InvokeLegacyActionResponse
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.StatusCode(); got != tt.want {
				t.Errorf("InvokeLegacyActionResponse.StatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoginWithPasswordResponse_Status(t *testing.T) {
	tests := []struct {
		name string
		r    LoginWithPasswordResponse
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Status(); got != tt.want {
				t.Errorf("LoginWithPasswordResponse.Status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoginWithPasswordResponse_StatusCode(t *testing.T) {
	tests := []struct {
		name string
		r    LoginWithPasswordResponse
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.StatusCode(); got != tt.want {
				t.Errorf("LoginWithPasswordResponse.StatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoginWithAccessKeyResponse_Status(t *testing.T) {
	tests := []struct {
		name string
		r    LoginWithAccessKeyResponse
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Status(); got != tt.want {
				t.Errorf("LoginWithAccessKeyResponse.Status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoginWithAccessKeyResponse_StatusCode(t *testing.T) {
	tests := []struct {
		name string
		r    LoginWithAccessKeyResponse
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.StatusCode(); got != tt.want {
				t.Errorf("LoginWithAccessKeyResponse.StatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetScriptResponse_Status(t *testing.T) {
	tests := []struct {
		name string
		r    GetScriptResponse
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Status(); got != tt.want {
				t.Errorf("GetScriptResponse.Status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetScriptResponse_StatusCode(t *testing.T) {
	tests := []struct {
		name string
		r    GetScriptResponse
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.StatusCode(); got != tt.want {
				t.Errorf("GetScriptResponse.StatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArchiveScriptResponse_Status(t *testing.T) {
	tests := []struct {
		name string
		r    ArchiveScriptResponse
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Status(); got != tt.want {
				t.Errorf("ArchiveScriptResponse.Status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArchiveScriptResponse_StatusCode(t *testing.T) {
	tests := []struct {
		name string
		r    ArchiveScriptResponse
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.StatusCode(); got != tt.want {
				t.Errorf("ArchiveScriptResponse.StatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedactScriptResponse_Status(t *testing.T) {
	tests := []struct {
		name string
		r    RedactScriptResponse
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Status(); got != tt.want {
				t.Errorf("RedactScriptResponse.Status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedactScriptResponse_StatusCode(t *testing.T) {
	tests := []struct {
		name string
		r    RedactScriptResponse
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.StatusCode(); got != tt.want {
				t.Errorf("RedactScriptResponse.StatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientWithResponses_InvokeLegacyActionWithResponse(t *testing.T) {
	type args struct {
		ctx        context.Context
		params     *InvokeLegacyActionParams
		reqEditors []RequestEditorFn
	}
	tests := []struct {
		name    string
		c       *ClientWithResponses
		args    args
		want    *InvokeLegacyActionResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.InvokeLegacyActionWithResponse(tt.args.ctx, tt.args.params, tt.args.reqEditors...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.InvokeLegacyActionWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.InvokeLegacyActionWithResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientWithResponses_LoginWithPasswordWithBodyWithResponse(t *testing.T) {
	type args struct {
		ctx         context.Context
		contentType string
		body        io.Reader
		reqEditors  []RequestEditorFn
	}
	tests := []struct {
		name    string
		c       *ClientWithResponses
		args    args
		want    *LoginWithPasswordResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.LoginWithPasswordWithBodyWithResponse(tt.args.ctx, tt.args.contentType, tt.args.body, tt.args.reqEditors...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.LoginWithPasswordWithBodyWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.LoginWithPasswordWithBodyWithResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientWithResponses_LoginWithPasswordWithResponse(t *testing.T) {
	type args struct {
		ctx        context.Context
		body       LoginWithPasswordJSONRequestBody
		reqEditors []RequestEditorFn
	}
	tests := []struct {
		name    string
		c       *ClientWithResponses
		args    args
		want    *LoginWithPasswordResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.LoginWithPasswordWithResponse(tt.args.ctx, tt.args.body, tt.args.reqEditors...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.LoginWithPasswordWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.LoginWithPasswordWithResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientWithResponses_LoginWithAccessKeyWithBodyWithResponse(t *testing.T) {
	type args struct {
		ctx         context.Context
		contentType string
		body        io.Reader
		reqEditors  []RequestEditorFn
	}
	tests := []struct {
		name    string
		c       *ClientWithResponses
		args    args
		want    *LoginWithAccessKeyResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.LoginWithAccessKeyWithBodyWithResponse(tt.args.ctx, tt.args.contentType, tt.args.body, tt.args.reqEditors...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.LoginWithAccessKeyWithBodyWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.LoginWithAccessKeyWithBodyWithResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientWithResponses_LoginWithAccessKeyWithResponse(t *testing.T) {
	type args struct {
		ctx        context.Context
		body       LoginWithAccessKeyJSONRequestBody
		reqEditors []RequestEditorFn
	}
	tests := []struct {
		name    string
		c       *ClientWithResponses
		args    args
		want    *LoginWithAccessKeyResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.LoginWithAccessKeyWithResponse(tt.args.ctx, tt.args.body, tt.args.reqEditors...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.LoginWithAccessKeyWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.LoginWithAccessKeyWithResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientWithResponses_GetScriptWithResponse(t *testing.T) {
	type args struct {
		ctx        context.Context
		scriptId   ScriptIdPathParam
		reqEditors []RequestEditorFn
	}
	tests := []struct {
		name    string
		c       *ClientWithResponses
		args    args
		want    *GetScriptResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetScriptWithResponse(tt.args.ctx, tt.args.scriptId, tt.args.reqEditors...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.GetScriptWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.GetScriptWithResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientWithResponses_ArchiveScriptWithResponse(t *testing.T) {
	type args struct {
		ctx        context.Context
		scriptId   ScriptIdPathParam
		reqEditors []RequestEditorFn
	}
	tests := []struct {
		name    string
		c       *ClientWithResponses
		args    args
		want    *ArchiveScriptResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.ArchiveScriptWithResponse(tt.args.ctx, tt.args.scriptId, tt.args.reqEditors...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.ArchiveScriptWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.ArchiveScriptWithResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientWithResponses_RedactScriptWithResponse(t *testing.T) {
	type args struct {
		ctx        context.Context
		scriptId   ScriptIdPathParam
		reqEditors []RequestEditorFn
	}
	tests := []struct {
		name    string
		c       *ClientWithResponses
		args    args
		want    *RedactScriptResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.RedactScriptWithResponse(tt.args.ctx, tt.args.scriptId, tt.args.reqEditors...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.RedactScriptWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.RedactScriptWithResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseInvokeLegacyActionResponse(t *testing.T) {
	type args struct {
		rsp *http.Response
	}
	tests := []struct {
		name    string
		args    args
		want    *InvokeLegacyActionResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInvokeLegacyActionResponse(tt.args.rsp)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInvokeLegacyActionResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseInvokeLegacyActionResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseLoginWithPasswordResponse(t *testing.T) {
	type args struct {
		rsp *http.Response
	}
	tests := []struct {
		name    string
		args    args
		want    *LoginWithPasswordResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseLoginWithPasswordResponse(tt.args.rsp)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseLoginWithPasswordResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseLoginWithPasswordResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseLoginWithAccessKeyResponse(t *testing.T) {
	type args struct {
		rsp *http.Response
	}
	tests := []struct {
		name    string
		args    args
		want    *LoginWithAccessKeyResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseLoginWithAccessKeyResponse(tt.args.rsp)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseLoginWithAccessKeyResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseLoginWithAccessKeyResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseGetScriptResponse(t *testing.T) {
	type args struct {
		rsp *http.Response
	}
	tests := []struct {
		name    string
		args    args
		want    *GetScriptResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseGetScriptResponse(tt.args.rsp)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseGetScriptResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseGetScriptResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseArchiveScriptResponse(t *testing.T) {
	type args struct {
		rsp *http.Response
	}
	tests := []struct {
		name    string
		args    args
		want    *ArchiveScriptResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseArchiveScriptResponse(tt.args.rsp)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseArchiveScriptResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseArchiveScriptResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseRedactScriptResponse(t *testing.T) {
	type args struct {
		rsp *http.Response
	}
	tests := []struct {
		name    string
		args    args
		want    *RedactScriptResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseRedactScriptResponse(tt.args.rsp)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseRedactScriptResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseRedactScriptResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
