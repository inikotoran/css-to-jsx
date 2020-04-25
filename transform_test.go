package main

import "testing"

func Test_transform(t *testing.T) {
	type args struct {
		css string
	}
	tests := []struct {
		name    string
		args    args
		wantJSX string
	}{
		{
			name: "convert single line css to jsx success",
			args: args{
				css: "margin-top: 666px;",
			},
			wantJSX: "marginTop: '666px',\n",
		},
		{
			name: "convert multi line css to jsx success",
			args: args{
				css: `background-color: #282c34;
color: white;
padding: 40px;
font-family: Arial;
text-align: center;`,
			},
			wantJSX: `backgroundColor: '#282c34',
color: 'white',
padding: '40px',
fontFamily: 'Arial',
textAlign: 'center',
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotJSX, err := transform(tt.args.css)
			if err != nil {
				t.Errorf("transform() return error: %v", err.Error())
			}
			if gotJSX != tt.wantJSX {
				t.Errorf("transform() = %v, want %v", gotJSX, tt.wantJSX)
			}
		})
	}
}
