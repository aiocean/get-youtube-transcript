package youtube

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	youtubeAPIURL = "https://www.youtube.com/youtubei/v1/get_transcript?key=AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"
	clientName    = "WEB"
	clientVersion = "2.9999099"
)

type ClientInfo struct {
	ClientName    string `json:"clientName"`
	ClientVersion string `json:"clientVersion"`
}

type RequestContext struct {
	Client ClientInfo `json:"client"`
}

type TranscriptRequest struct {
	Context RequestContext `json:"context"`
	Params  string         `json:"params"`
}

func (t *Transcript) String() string {
	result := ""
	for _, segment := range t.Segments {
		result += fmt.Sprintf("[%s] %s ", segment.Time, segment.Text)
	}
	return result
}

type Transcript struct {
	VideoId  string    `json:"videoID"`
	Segments []Segment `json:"segments"`
}

type Segment struct {
	Time string `json:"time"`
	Text string `json:"text"`
}

func GetTranscript(videoID string) (*Transcript, error) {
	paramStr := fmt.Sprintf("\n\x0b%s", videoID)
	encodedParams := base64.StdEncoding.EncodeToString([]byte(paramStr))
	reqBody := TranscriptRequest{
		Context: RequestContext{
			Client: ClientInfo{
				ClientName:    clientName,
				ClientVersion: clientVersion,
			},
		},
		Params: encodedParams,
	}

	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(youtubeAPIURL, "application/json", bytes.NewBuffer(jsonReqBody))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var transResponse TranscriptResponse
	err = json.Unmarshal(body, &transResponse)
	if err != nil {
		return nil, err
	}

	transcript := &Transcript{
		VideoId:  videoID,
		Segments: []Segment{},
	}

	for _, action := range transResponse.Actions {
		segment := action.UpdateEngagementPanelAction.Content.TranscriptRenderer.Body.TranscriptBodyRenderer.CueGroups
		for _, seg := range segment {
			text := ""
			for _, run := range seg.TranscriptCueGroupRenderer.Cues {

				text += " " + strings.TrimSpace(run.TranscriptCueRenderer.Cue.SimpleText)
			}

			// make text into a single line
			text = strings.ReplaceAll(text, "\n", " ")

			transcript.Segments = append(transcript.Segments, Segment{
				Time: seg.TranscriptCueGroupRenderer.FormattedStartOffset.SimpleText,
				Text: strings.TrimSpace(text),
			})
		}
	}

	return transcript, nil
}
