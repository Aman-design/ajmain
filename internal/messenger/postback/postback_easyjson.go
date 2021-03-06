// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package postback

import (
	json "encoding/json"
	models "github.com/knadh/listmonk/models"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonDf11841fDecodeGithubComKnadhListmonkInternalMessengerPostback(in *jlexer.Lexer, out *postback) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "subject":
			out.Subject = string(in.String())
		case "content_type":
			out.ContentType = string(in.String())
		case "body":
			out.Body = string(in.String())
		case "recipients":
			if in.IsNull() {
				in.Skip()
				out.Recipients = nil
			} else {
				in.Delim('[')
				if out.Recipients == nil {
					if !in.IsDelim(']') {
						out.Recipients = make([]recipient, 0, 0)
					} else {
						out.Recipients = []recipient{}
					}
				} else {
					out.Recipients = (out.Recipients)[:0]
				}
				for !in.IsDelim(']') {
					var v1 recipient
					easyjsonDf11841fDecodeGithubComKnadhListmonkInternalMessengerPostback1(in, &v1)
					out.Recipients = append(out.Recipients, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "campaign":
			if in.IsNull() {
				in.Skip()
				out.Campaign = nil
			} else {
				if out.Campaign == nil {
					out.Campaign = new(campaign)
				}
				easyjsonDf11841fDecodeGithubComKnadhListmonkInternalMessengerPostback2(in, out.Campaign)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonDf11841fEncodeGithubComKnadhListmonkInternalMessengerPostback(out *jwriter.Writer, in postback) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"subject\":"
		out.RawString(prefix[1:])
		out.String(string(in.Subject))
	}
	{
		const prefix string = ",\"content_type\":"
		out.RawString(prefix)
		out.String(string(in.ContentType))
	}
	{
		const prefix string = ",\"body\":"
		out.RawString(prefix)
		out.String(string(in.Body))
	}
	{
		const prefix string = ",\"recipients\":"
		out.RawString(prefix)
		if in.Recipients == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Recipients {
				if v2 > 0 {
					out.RawByte(',')
				}
				easyjsonDf11841fEncodeGithubComKnadhListmonkInternalMessengerPostback1(out, v3)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"campaign\":"
		out.RawString(prefix)
		if in.Campaign == nil {
			out.RawString("null")
		} else {
			easyjsonDf11841fEncodeGithubComKnadhListmonkInternalMessengerPostback2(out, *in.Campaign)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v postback) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDf11841fEncodeGithubComKnadhListmonkInternalMessengerPostback(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v postback) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDf11841fEncodeGithubComKnadhListmonkInternalMessengerPostback(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *postback) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDf11841fDecodeGithubComKnadhListmonkInternalMessengerPostback(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *postback) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDf11841fDecodeGithubComKnadhListmonkInternalMessengerPostback(l, v)
}
func easyjsonDf11841fDecodeGithubComKnadhListmonkInternalMessengerPostback2(in *jlexer.Lexer, out *campaign) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "uuid":
			out.UUID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "headers":
			if in.IsNull() {
				in.Skip()
				out.Headers = nil
			} else {
				in.Delim('[')
				if out.Headers == nil {
					if !in.IsDelim(']') {
						out.Headers = make(models.Headers, 0, 8)
					} else {
						out.Headers = models.Headers{}
					}
				} else {
					out.Headers = (out.Headers)[:0]
				}
				for !in.IsDelim(']') {
					var v4 map[string]string
					if in.IsNull() {
						in.Skip()
					} else {
						in.Delim('{')
						v4 = make(map[string]string)
						for !in.IsDelim('}') {
							key := string(in.String())
							in.WantColon()
							var v5 string
							v5 = string(in.String())
							(v4)[key] = v5
							in.WantComma()
						}
						in.Delim('}')
					}
					out.Headers = append(out.Headers, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make([]string, 0, 4)
					} else {
						out.Tags = []string{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v6 string
					v6 = string(in.String())
					out.Tags = append(out.Tags, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonDf11841fEncodeGithubComKnadhListmonkInternalMessengerPostback2(out *jwriter.Writer, in campaign) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"uuid\":"
		out.RawString(prefix[1:])
		out.String(string(in.UUID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"headers\":"
		out.RawString(prefix)
		if in.Headers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v7, v8 := range in.Headers {
				if v7 > 0 {
					out.RawByte(',')
				}
				if v8 == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
					out.RawString(`null`)
				} else {
					out.RawByte('{')
					v9First := true
					for v9Name, v9Value := range v8 {
						if v9First {
							v9First = false
						} else {
							out.RawByte(',')
						}
						out.String(string(v9Name))
						out.RawByte(':')
						out.String(string(v9Value))
					}
					out.RawByte('}')
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v10, v11 := range in.Tags {
				if v10 > 0 {
					out.RawByte(',')
				}
				out.String(string(v11))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjsonDf11841fDecodeGithubComKnadhListmonkInternalMessengerPostback1(in *jlexer.Lexer, out *recipient) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "uuid":
			out.UUID = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "attribs":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Attribs = make(models.SubscriberAttribs)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v12 interface{}
					if m, ok := v12.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v12.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v12 = in.Interface()
					}
					(out.Attribs)[key] = v12
					in.WantComma()
				}
				in.Delim('}')
			}
		case "status":
			out.Status = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonDf11841fEncodeGithubComKnadhListmonkInternalMessengerPostback1(out *jwriter.Writer, in recipient) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"uuid\":"
		out.RawString(prefix[1:])
		out.String(string(in.UUID))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"attribs\":"
		out.RawString(prefix)
		if in.Attribs == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v13First := true
			for v13Name, v13Value := range in.Attribs {
				if v13First {
					v13First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v13Name))
				out.RawByte(':')
				if m, ok := v13Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v13Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v13Value))
				}
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.Status))
	}
	out.RawByte('}')
}
