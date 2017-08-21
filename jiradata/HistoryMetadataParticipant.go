package jiradata

/////////////////////////////////////////////////////////////////////////
// This Code is Generated by SlipScheme Project:
// https://github.com/coryb/slipscheme
//
// Generated with command:
// slipscheme -pkg jiradata -dir jiradata -overwrite schemas/IssueUpdate.json
/////////////////////////////////////////////////////////////////////////
//                            DO NOT EDIT                              //
/////////////////////////////////////////////////////////////////////////

// HistoryMetadataParticipant defined from schema:
// {
//   "title": "History Metadata Participant",
//   "type": "object",
//   "properties": {
//     "avatarUrl": {
//       "type": "string"
//     },
//     "displayName": {
//       "type": "string"
//     },
//     "displayNameKey": {
//       "type": "string"
//     },
//     "id": {
//       "type": "string"
//     },
//     "type": {
//       "type": "string"
//     },
//     "url": {
//       "type": "string"
//     }
//   }
// }
type HistoryMetadataParticipant struct {
	AvatarURL      string `json:"avatarUrl,omitempty" yaml:"avatarUrl,omitempty"`
	DisplayName    string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	DisplayNameKey string `json:"displayNameKey,omitempty" yaml:"displayNameKey,omitempty"`
	ID             string `json:"id,omitempty" yaml:"id,omitempty"`
	Type           string `json:"type,omitempty" yaml:"type,omitempty"`
	URL            string `json:"url,omitempty" yaml:"url,omitempty"`
}