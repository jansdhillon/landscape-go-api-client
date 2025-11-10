// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"net/url"
	"reflect"
	"testing"
)

func TestNewEmailPasswordProvider(t *testing.T) {
	type args struct {
		email    string
		password string
		account  *string
	}
	tests := []struct {
		name string
		args args
		want *EmailPasswordProvider
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEmailPasswordProvider(tt.args.email, tt.args.password, tt.args.account); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmailPasswordProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmailPasswordProvider_Login(t *testing.T) {
	type args struct {
		ctx context.Context
		c   *ClientWithResponses
	}
	tests := []struct {
		name    string
		p       *EmailPasswordProvider
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.Login(tt.args.ctx, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("EmailPasswordProvider.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EmailPasswordProvider.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAccessKeyProvider(t *testing.T) {
	type args struct {
		accessKey string
		secretKey string
	}
	tests := []struct {
		name string
		args args
		want *AccessKeyProvider
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAccessKeyProvider(tt.args.accessKey, tt.args.secretKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccessKeyProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccessKeyProvider_Login(t *testing.T) {
	type args struct {
		ctx context.Context
		c   *ClientWithResponses
	}
	tests := []struct {
		name    string
		p       *AccessKeyProvider
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.Login(tt.args.ctx, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccessKeyProvider.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AccessKeyProvider.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLandscapeAPIClient(t *testing.T) {
	type args struct {
		baseURL       string
		loginProvider LoginProvider
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
			got, err := NewLandscapeAPIClient(tt.args.baseURL, tt.args.loginProvider)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLandscapeAPIClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLandscapeAPIClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLegacyActionParams(t *testing.T) {
	type args struct {
		action string
	}
	tests := []struct {
		name string
		args args
		want *InvokeLegacyActionParams
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LegacyActionParams(tt.args.action); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LegacyActionParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeQueryRequestEditor(t *testing.T) {
	type args struct {
		values url.Values
	}
	tests := []struct {
		name string
		args args
		want RequestEditorFn
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeQueryRequestEditor(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeQueryRequestEditor() = %v, want %v", got, tt.want)
			}
		})
	}
}
