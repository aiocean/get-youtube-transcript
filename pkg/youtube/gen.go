package youtube

// TranscriptRequest is the request body for the transcript request.
// This struct was generated.
type TranscriptResponse struct {
	ResponseContext struct {
		VisitorData           string `json:"visitorData"`
		ServiceTrackingParams []struct {
			Service string `json:"service"`
			Params  []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"params"`
		} `json:"serviceTrackingParams"`
		MainAppWebResponseContext struct {
			LoggedOut bool `json:"loggedOut"`
		} `json:"mainAppWebResponseContext"`
		WebResponseContextExtensionData struct {
			HasDecorated bool `json:"hasDecorated"`
		} `json:"webResponseContextExtensionData"`
	} `json:"responseContext"`
	Actions []struct {
		ClickTrackingParams         string `json:"clickTrackingParams"`
		UpdateEngagementPanelAction struct {
			TargetId string `json:"targetId"`
			Content  struct {
				TranscriptRenderer struct {
					Body struct {
						TranscriptBodyRenderer struct {
							CueGroups []struct {
								TranscriptCueGroupRenderer struct {
									FormattedStartOffset struct {
										SimpleText string `json:"simpleText"`
									} `json:"formattedStartOffset"`
									Cues []struct {
										TranscriptCueRenderer struct {
											Cue struct {
												SimpleText string `json:"simpleText"`
											} `json:"cue"`
											StartOffsetMs string `json:"startOffsetMs"`
											DurationMs    string `json:"durationMs"`
										} `json:"transcriptCueRenderer"`
									} `json:"cues"`
								} `json:"transcriptCueGroupRenderer"`
							} `json:"cueGroups"`
						} `json:"transcriptBodyRenderer"`
					} `json:"body"`
					Footer struct {
						TranscriptFooterRenderer struct {
							LanguageMenu struct {
								SortFilterSubMenuRenderer struct {
									SubMenuItems []struct {
										Title        string `json:"title"`
										Selected     bool   `json:"selected"`
										Continuation struct {
											ReloadContinuationData struct {
												Continuation        string `json:"continuation"`
												ClickTrackingParams string `json:"clickTrackingParams"`
											} `json:"reloadContinuationData"`
										} `json:"continuation"`
										TrackingParams string `json:"trackingParams"`
									} `json:"subMenuItems"`
									TrackingParams string `json:"trackingParams"`
								} `json:"sortFilterSubMenuRenderer"`
							} `json:"languageMenu"`
						} `json:"transcriptFooterRenderer"`
					} `json:"footer"`
					TrackingParams string `json:"trackingParams"`
				} `json:"transcriptRenderer"`
			} `json:"content"`
		} `json:"updateEngagementPanelAction"`
	} `json:"actions"`
	TrackingParams string `json:"trackingParams"`
}
