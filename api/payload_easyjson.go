// AUTOGENERATED FILE: easyjson marshaller/unmarshallers.

package api

import (
	json "encoding/json"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ = json.RawMessage{}
	_ = jlexer.Lexer{}
	_ = jwriter.Writer{}
)

func easyjsonA8a797f8DecodeGithubComTopfreegamesKhanApi(in *jlexer.Lexer, out *ApproveOrDenyMembershipInvitationPayload) {
	if in.IsNull() {
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "playerPublicID":
			out.PlayerPublicID = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjsonA8a797f8EncodeGithubComTopfreegamesKhanApi(out *jwriter.Writer, in ApproveOrDenyMembershipInvitationPayload) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"playerPublicID\":")
	out.String(string(in.PlayerPublicID))
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ApproveOrDenyMembershipInvitationPayload) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA8a797f8EncodeGithubComTopfreegamesKhanApi(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ApproveOrDenyMembershipInvitationPayload) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA8a797f8DecodeGithubComTopfreegamesKhanApi(l, v)
}
func easyjsonA8a797f8DecodeGithubComTopfreegamesKhanApi1(in *jlexer.Lexer, out *BasePayloadWithRequestorAndPlayerPublicIDs) {
	if in.IsNull() {
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "playerPublicID":
			out.PlayerPublicID = string(in.String())
		case "requestorPublicID":
			out.RequestorPublicID = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjsonA8a797f8EncodeGithubComTopfreegamesKhanApi1(out *jwriter.Writer, in BasePayloadWithRequestorAndPlayerPublicIDs) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"playerPublicID\":")
	out.String(string(in.PlayerPublicID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"requestorPublicID\":")
	out.String(string(in.RequestorPublicID))
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BasePayloadWithRequestorAndPlayerPublicIDs) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA8a797f8EncodeGithubComTopfreegamesKhanApi1(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BasePayloadWithRequestorAndPlayerPublicIDs) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA8a797f8DecodeGithubComTopfreegamesKhanApi1(l, v)
}
func easyjsonA8a797f8DecodeGithubComTopfreegamesKhanApi2(in *jlexer.Lexer, out *InviteForMembershipPayload) {
	if in.IsNull() {
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "level":
			out.Level = string(in.String())
		case "playerPublicID":
			out.PlayerPublicID = string(in.String())
		case "requestorPublicID":
			out.RequestorPublicID = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjsonA8a797f8EncodeGithubComTopfreegamesKhanApi2(out *jwriter.Writer, in InviteForMembershipPayload) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"level\":")
	out.String(string(in.Level))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"playerPublicID\":")
	out.String(string(in.PlayerPublicID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"requestorPublicID\":")
	out.String(string(in.RequestorPublicID))
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v InviteForMembershipPayload) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA8a797f8EncodeGithubComTopfreegamesKhanApi2(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *InviteForMembershipPayload) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA8a797f8DecodeGithubComTopfreegamesKhanApi2(l, v)
}
func easyjsonA8a797f8DecodeGithubComTopfreegamesKhanApi3(in *jlexer.Lexer, out *ApplyForMembershipPayload) {
	if in.IsNull() {
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "level":
			out.Level = string(in.String())
		case "playerPublicID":
			out.PlayerPublicID = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjsonA8a797f8EncodeGithubComTopfreegamesKhanApi3(out *jwriter.Writer, in ApplyForMembershipPayload) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"level\":")
	out.String(string(in.Level))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"playerPublicID\":")
	out.String(string(in.PlayerPublicID))
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ApplyForMembershipPayload) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA8a797f8EncodeGithubComTopfreegamesKhanApi3(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ApplyForMembershipPayload) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA8a797f8DecodeGithubComTopfreegamesKhanApi3(l, v)
}
func easyjsonA8a797f8DecodeGithubComTopfreegamesKhanApi4(in *jlexer.Lexer, out *CreateGamePayload) {
	if in.IsNull() {
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "publicID":
			out.PublicID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "membershipLevels":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.MembershipLevels = make(map[string]interface{})
				} else {
					out.MembershipLevels = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 interface{}
					v1 = in.Interface()
					(out.MembershipLevels)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "metadata":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Metadata = make(map[string]interface{})
				} else {
					out.Metadata = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v2 interface{}
					v2 = in.Interface()
					(out.Metadata)[key] = v2
					in.WantComma()
				}
				in.Delim('}')
			}
		case "minLevelToAcceptApplication":
			out.MinLevelToAcceptApplication = int(in.Int())
		case "minLevelToCreateInvitation":
			out.MinLevelToCreateInvitation = int(in.Int())
		case "minLevelToRemoveMember":
			out.MinLevelToRemoveMember = int(in.Int())
		case "minLevelOffsetToRemoveMember":
			out.MinLevelOffsetToRemoveMember = int(in.Int())
		case "minLevelOffsetToPromoteMember":
			out.MinLevelOffsetToPromoteMember = int(in.Int())
		case "minLevelOffsetToDemoteMember":
			out.MinLevelOffsetToDemoteMember = int(in.Int())
		case "maxMembers":
			out.MaxMembers = int(in.Int())
		case "maxClansPerPlayer":
			out.MaxClansPerPlayer = int(in.Int())
		case "cooldownAfterDeny":
			out.CooldownAfterDeny = int(in.Int())
		case "cooldownAfterDelete":
			out.CooldownAfterDelete = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjsonA8a797f8EncodeGithubComTopfreegamesKhanApi4(out *jwriter.Writer, in CreateGamePayload) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"publicID\":")
	out.String(string(in.PublicID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"name\":")
	out.String(string(in.Name))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"membershipLevels\":")
	if in.MembershipLevels == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v3First := true
		for v3Name, v3Value := range in.MembershipLevels {
			if !v3First {
				out.RawByte(',')
			}
			v3First = false
			out.String(string(v3Name))
			out.RawByte(':')
			out.Raw(json.Marshal(v3Value))
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"metadata\":")
	if in.Metadata == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v4First := true
		for v4Name, v4Value := range in.Metadata {
			if !v4First {
				out.RawByte(',')
			}
			v4First = false
			out.String(string(v4Name))
			out.RawByte(':')
			out.Raw(json.Marshal(v4Value))
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"minLevelToAcceptApplication\":")
	out.Int(int(in.MinLevelToAcceptApplication))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"minLevelToCreateInvitation\":")
	out.Int(int(in.MinLevelToCreateInvitation))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"minLevelToRemoveMember\":")
	out.Int(int(in.MinLevelToRemoveMember))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"minLevelOffsetToRemoveMember\":")
	out.Int(int(in.MinLevelOffsetToRemoveMember))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"minLevelOffsetToPromoteMember\":")
	out.Int(int(in.MinLevelOffsetToPromoteMember))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"minLevelOffsetToDemoteMember\":")
	out.Int(int(in.MinLevelOffsetToDemoteMember))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"maxMembers\":")
	out.Int(int(in.MaxMembers))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"maxClansPerPlayer\":")
	out.Int(int(in.MaxClansPerPlayer))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"cooldownAfterDeny\":")
	out.Int(int(in.CooldownAfterDeny))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"cooldownAfterDelete\":")
	out.Int(int(in.CooldownAfterDelete))
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CreateGamePayload) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA8a797f8EncodeGithubComTopfreegamesKhanApi4(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CreateGamePayload) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA8a797f8DecodeGithubComTopfreegamesKhanApi4(l, v)
}
func easyjsonA8a797f8DecodeGithubComTopfreegamesKhanApi5(in *jlexer.Lexer, out *UpdateGamePayload) {
	if in.IsNull() {
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "membershipLevels":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.MembershipLevels = make(map[string]interface{})
				} else {
					out.MembershipLevels = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v5 interface{}
					v5 = in.Interface()
					(out.MembershipLevels)[key] = v5
					in.WantComma()
				}
				in.Delim('}')
			}
		case "metadata":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Metadata = make(map[string]interface{})
				} else {
					out.Metadata = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v6 interface{}
					v6 = in.Interface()
					(out.Metadata)[key] = v6
					in.WantComma()
				}
				in.Delim('}')
			}
		case "minLevelToAcceptApplication":
			out.MinLevelToAcceptApplication = int(in.Int())
		case "minLevelToCreateInvitation":
			out.MinLevelToCreateInvitation = int(in.Int())
		case "minLevelToRemoveMember":
			out.MinLevelToRemoveMember = int(in.Int())
		case "minLevelOffsetToRemoveMember":
			out.MinLevelOffsetToRemoveMember = int(in.Int())
		case "minLevelOffsetToPromoteMember":
			out.MinLevelOffsetToPromoteMember = int(in.Int())
		case "minLevelOffsetToDemoteMember":
			out.MinLevelOffsetToDemoteMember = int(in.Int())
		case "maxMembers":
			out.MaxMembers = int(in.Int())
		case "maxClansPerPlayer":
			out.MaxClansPerPlayer = int(in.Int())
		case "cooldownAfterDeny":
			out.CooldownAfterDeny = int(in.Int())
		case "cooldownAfterDelete":
			out.CooldownAfterDelete = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjsonA8a797f8EncodeGithubComTopfreegamesKhanApi5(out *jwriter.Writer, in UpdateGamePayload) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"name\":")
	out.String(string(in.Name))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"membershipLevels\":")
	if in.MembershipLevels == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v7First := true
		for v7Name, v7Value := range in.MembershipLevels {
			if !v7First {
				out.RawByte(',')
			}
			v7First = false
			out.String(string(v7Name))
			out.RawByte(':')
			out.Raw(json.Marshal(v7Value))
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"metadata\":")
	if in.Metadata == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v8First := true
		for v8Name, v8Value := range in.Metadata {
			if !v8First {
				out.RawByte(',')
			}
			v8First = false
			out.String(string(v8Name))
			out.RawByte(':')
			out.Raw(json.Marshal(v8Value))
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"minLevelToAcceptApplication\":")
	out.Int(int(in.MinLevelToAcceptApplication))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"minLevelToCreateInvitation\":")
	out.Int(int(in.MinLevelToCreateInvitation))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"minLevelToRemoveMember\":")
	out.Int(int(in.MinLevelToRemoveMember))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"minLevelOffsetToRemoveMember\":")
	out.Int(int(in.MinLevelOffsetToRemoveMember))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"minLevelOffsetToPromoteMember\":")
	out.Int(int(in.MinLevelOffsetToPromoteMember))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"minLevelOffsetToDemoteMember\":")
	out.Int(int(in.MinLevelOffsetToDemoteMember))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"maxMembers\":")
	out.Int(int(in.MaxMembers))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"maxClansPerPlayer\":")
	out.Int(int(in.MaxClansPerPlayer))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"cooldownAfterDeny\":")
	out.Int(int(in.CooldownAfterDeny))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"cooldownAfterDelete\":")
	out.Int(int(in.CooldownAfterDelete))
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UpdateGamePayload) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA8a797f8EncodeGithubComTopfreegamesKhanApi5(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UpdateGamePayload) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA8a797f8DecodeGithubComTopfreegamesKhanApi5(l, v)
}
func easyjsonA8a797f8DecodeGithubComTopfreegamesKhanApi6(in *jlexer.Lexer, out *UpdatePlayerPayload) {
	if in.IsNull() {
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "metadata":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Metadata = make(map[string]interface{})
				} else {
					out.Metadata = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v9 interface{}
					v9 = in.Interface()
					(out.Metadata)[key] = v9
					in.WantComma()
				}
				in.Delim('}')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjsonA8a797f8EncodeGithubComTopfreegamesKhanApi6(out *jwriter.Writer, in UpdatePlayerPayload) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"name\":")
	out.String(string(in.Name))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"metadata\":")
	if in.Metadata == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v10First := true
		for v10Name, v10Value := range in.Metadata {
			if !v10First {
				out.RawByte(',')
			}
			v10First = false
			out.String(string(v10Name))
			out.RawByte(':')
			out.Raw(json.Marshal(v10Value))
		}
		out.RawByte('}')
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UpdatePlayerPayload) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA8a797f8EncodeGithubComTopfreegamesKhanApi6(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UpdatePlayerPayload) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA8a797f8DecodeGithubComTopfreegamesKhanApi6(l, v)
}
func easyjsonA8a797f8DecodeGithubComTopfreegamesKhanApi7(in *jlexer.Lexer, out *CreatePlayerPayload) {
	if in.IsNull() {
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "publicID":
			out.PublicID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "metadata":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Metadata = make(map[string]interface{})
				} else {
					out.Metadata = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v11 interface{}
					v11 = in.Interface()
					(out.Metadata)[key] = v11
					in.WantComma()
				}
				in.Delim('}')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjsonA8a797f8EncodeGithubComTopfreegamesKhanApi7(out *jwriter.Writer, in CreatePlayerPayload) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"publicID\":")
	out.String(string(in.PublicID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"name\":")
	out.String(string(in.Name))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"metadata\":")
	if in.Metadata == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v12First := true
		for v12Name, v12Value := range in.Metadata {
			if !v12First {
				out.RawByte(',')
			}
			v12First = false
			out.String(string(v12Name))
			out.RawByte(':')
			out.Raw(json.Marshal(v12Value))
		}
		out.RawByte('}')
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CreatePlayerPayload) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA8a797f8EncodeGithubComTopfreegamesKhanApi7(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CreatePlayerPayload) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA8a797f8DecodeGithubComTopfreegamesKhanApi7(l, v)
}
func easyjsonA8a797f8DecodeGithubComTopfreegamesKhanApi8(in *jlexer.Lexer, out *TransferClanOwnershipPayload) {
	if in.IsNull() {
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "playerPublicID":
			out.PlayerPublicID = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjsonA8a797f8EncodeGithubComTopfreegamesKhanApi8(out *jwriter.Writer, in TransferClanOwnershipPayload) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"playerPublicID\":")
	out.String(string(in.PlayerPublicID))
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TransferClanOwnershipPayload) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA8a797f8EncodeGithubComTopfreegamesKhanApi8(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TransferClanOwnershipPayload) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA8a797f8DecodeGithubComTopfreegamesKhanApi8(l, v)
}
func easyjsonA8a797f8DecodeGithubComTopfreegamesKhanApi9(in *jlexer.Lexer, out *UpdateClanPayload) {
	if in.IsNull() {
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "ownerPublicID":
			out.OwnerPublicID = string(in.String())
		case "metadata":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Metadata = make(map[string]interface{})
				} else {
					out.Metadata = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v13 interface{}
					v13 = in.Interface()
					(out.Metadata)[key] = v13
					in.WantComma()
				}
				in.Delim('}')
			}
		case "allowApplication":
			out.AllowApplication = bool(in.Bool())
		case "autoJoin":
			out.AutoJoin = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjsonA8a797f8EncodeGithubComTopfreegamesKhanApi9(out *jwriter.Writer, in UpdateClanPayload) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"name\":")
	out.String(string(in.Name))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"ownerPublicID\":")
	out.String(string(in.OwnerPublicID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"metadata\":")
	if in.Metadata == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v14First := true
		for v14Name, v14Value := range in.Metadata {
			if !v14First {
				out.RawByte(',')
			}
			v14First = false
			out.String(string(v14Name))
			out.RawByte(':')
			out.Raw(json.Marshal(v14Value))
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"allowApplication\":")
	out.Bool(bool(in.AllowApplication))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"autoJoin\":")
	out.Bool(bool(in.AutoJoin))
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UpdateClanPayload) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA8a797f8EncodeGithubComTopfreegamesKhanApi9(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UpdateClanPayload) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA8a797f8DecodeGithubComTopfreegamesKhanApi9(l, v)
}
func easyjsonA8a797f8DecodeGithubComTopfreegamesKhanApi10(in *jlexer.Lexer, out *CreateClanPayload) {
	if in.IsNull() {
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "publicID":
			out.PublicID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "ownerPublicID":
			out.OwnerPublicID = string(in.String())
		case "metadata":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Metadata = make(map[string]interface{})
				} else {
					out.Metadata = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v15 interface{}
					v15 = in.Interface()
					(out.Metadata)[key] = v15
					in.WantComma()
				}
				in.Delim('}')
			}
		case "allowApplication":
			out.AllowApplication = bool(in.Bool())
		case "autoJoin":
			out.AutoJoin = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjsonA8a797f8EncodeGithubComTopfreegamesKhanApi10(out *jwriter.Writer, in CreateClanPayload) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"publicID\":")
	out.String(string(in.PublicID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"name\":")
	out.String(string(in.Name))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"ownerPublicID\":")
	out.String(string(in.OwnerPublicID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"metadata\":")
	if in.Metadata == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v16First := true
		for v16Name, v16Value := range in.Metadata {
			if !v16First {
				out.RawByte(',')
			}
			v16First = false
			out.String(string(v16Name))
			out.RawByte(':')
			out.Raw(json.Marshal(v16Value))
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"allowApplication\":")
	out.Bool(bool(in.AllowApplication))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"autoJoin\":")
	out.Bool(bool(in.AutoJoin))
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CreateClanPayload) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA8a797f8EncodeGithubComTopfreegamesKhanApi10(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CreateClanPayload) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA8a797f8DecodeGithubComTopfreegamesKhanApi10(l, v)
}
func easyjsonA8a797f8DecodeGithubComTopfreegamesKhanApi11(in *jlexer.Lexer, out *Validation) {
	if in.IsNull() {
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjsonA8a797f8EncodeGithubComTopfreegamesKhanApi11(out *jwriter.Writer, in Validation) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Validation) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA8a797f8EncodeGithubComTopfreegamesKhanApi11(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Validation) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA8a797f8DecodeGithubComTopfreegamesKhanApi11(l, v)
}
