package backend

// Bootstrap - Parent of Element Type and Elements. This is the outer case of
// Bootstrap API
type Bootstrap struct {
	ElementTypes []ElementType `json:"element_types"`
	Events       []Event       `json:"events"`
	Players      []Elements    `json:"elements"`
	Teams        []TeamsStruct `json:"teams"`
}

// ChipPlay - Used in Event struct, it has data about which chips were played
// the most during a specific wee
type ChipPlay struct {
	ChipName  string `json:"chip_name"`
	NumPlayed int64  `json:"num_played"`
}

// Elements - Overall stats about each player
type Elements struct {
	Assists                         int    `json:"assists" validate:"gte=0"`
	Bonus                           int    `json:"bonus" validate:"gte=0"`
	BPS                             int    `json:"bps" validate:"gte=0"`
	ChanceOfPlayingNextRound        int    `json:"chance_of_playing_next_round" validate:"gte=0"`
	ChanceOfPlayingThisRound        int    `json:"chance_of_playing_this_round" validate:"gte=0"`
	CleanSheets                     int    `json:"clean_sheets" validate:"gte=0"`
	CornersAndIndrectFreekicksOrder int    `json:"corners_and_indirect_freekicks_order" validate:"gte=0"`
	Creativity                      string `json:"creativity" validate:"required"`
	CreativityRank                  int    `json:"creativity_rank" validate:"gte=0"`
	CreativityRankType              int    `json:"creativity_rank_type" validate:"gte=0"`
	DirectFreekicksOrder            int    `json:"direct_freekicks_order" validate:"gte=0"`
	DreamTeamCount                  int    `json:"dreamteam_count" validate:"gte=0"`
	ElementType                     int    `json:"element_type" validate:"gte=0"`
	EPNext                          string `json:"ep_next" validate:"required"`
	EPThis                          string `json:"ep_this" validate:"required"`
	EventPoints                     int    `json:"event_points" validate:"number"`
	FirstName                       string `json:"first_name"`
	Form                            string `json:"form"`
	GoalsConceded                   int    `json:"goals_conceded" validate:"gte=0"`
	GoalsScored                     int    `json:"goals_scored" validate:"gte=0"`
	IctIndex                        string `json:"ict_index"`
	IctIndexRank                    int    `json:"ict_index_rank" validate:"gte=0"`
	IctIndexRankType                int    `json:"ict_index_rank_type" validate:"gte=0"`
	Influence                       string `json:"influence"`
	InfluenceRank                   int    `json:"influence_rank" validate:"gte=0"`
	InfluenceRankType               int    `json:"influence_rank_type" validate:"gte=0"`
	InDreamteam                     bool   `json:"in_dreamteam"`
	ID                              int    `json:"id" gorm:"primaryKey"` // Player ID
	Minutes                         int    `json:"minutes" validate:"gte=0"`
	News                            string `json:"news"`
	NewsAdded                       string `json:"news_added"`
	OwnGoals                        int    `json:"own_goals" validate:"gte=0"`
	PenaltiesOrder                  int    `json:"penalties_order" validate:"gte=0"`
	PenaltiesSaved                  int    `json:"penalties_saved" validate:"gte=0"`
	PenaltiesMissed                 int    `json:"penalties_missed" validate:"gte=0"`
	PointsPerGame                   string `json:"points_per_game" validate:"required"`
	RedCards                        int    `json:"red_cards" validate:"gte=0"`
	Saves                           int    `json:"saves" validate:"gte=0"`
	SecondName                      string `json:"second_name" validate:"required"`
	SelectedByPercent               string `json:"selected_by_percent" validate:"required"`
	Team                            int    `json:"team" validate:"gte=0"`
	TeamCode                        int    `json:"team_code" validate:"gte=0"`
	Threat                          string `json:"threat" validate:"required"`
	ThreatRank                      int    `json:"threat_rank" validate:"gte=0"`
	ThreatRankType                  int    `json:"threat_rank_type" validate:"gte=0"`
	TransfersIn                     int    `json:"transfers_in" validate:"gte=0"`
	TransfersInEvent                int    `json:"transfers_in_event" validate:"gte=0"`
	TransfersOut                    int    `json:"transfers_outte:"gte=0"`
	TransfersOutEvent               int    `json:"transfers_out_event" "gte=0"`
	TotalPoints                     int    `json:"total_points" validate:"number"`
	ValueForm                       string `json:"value_form" validate:"required"`
	ValueSeason                     string `json:"value_season" validate:"required"`
	YellowCards                     int    `json:"yellow_cards" validate:"gte=0"`
}

// ElementType - Information about which type of player it is, like forward
// goalkeeper, midfielder, etc.
type ElementType struct {
	ElementCount      int    `json:"element_count"`
	ID                int    `json:"id"`
	PluralName        string `json:"plural_name"`
	PluralNameShort   string `json:"plural_name_short"`
	SingularName      string `json:"singular_name"`
	SingularNameShort string `json:"singular_name_short"`
	SquadSelect       int    `json:"squad_select"`
	SquadMaxPlay      int    `json:"squad_max_play"`
	SquadMinPlay      int    `json:"squad_min_play"`
}

// Event - Gameweek info. Mostly to be used for when a week starts and ends
type Event struct {
	AverageEntryScore   int                  `json:"average_entry_score"`
	ChipPlays           []ChipPlay           `json:"chip_plays"`
	DataChecked         bool                 `json:"data_checked"`
	DeadlineTime        string               `json:"deadline_time"`
	DeadlineTimeEpoch   int64                `json:"deadline_time_epoch"`
	Finished            bool                 `json:"finished"`
	HighestScore        int                  `json:"highest_score"`
	HighestScoringEntry int                  `json:"highest_scoring_entry"`
	IsCurrent           bool                 `json:"is_current"`
	IsNext              bool                 `json:"is_next"`
	IsPrevious          bool                 `json:"is_previous"`
	ID                  int                  `json:"id"`
	MostCaptained       int                  `json:"most_captained"`
	MostPlayed          int                  `json:"most_played"`
	MostTransferredIn   int                  `json:"most_transferred_in"`
	MostViceCaptained   int                  `json:"most_vice_captained"`
	Name                string               `json:"name"`
	TopElement          int                  `json:"top_element"`
	TopElementInfo      TopElementInfoStruct `json:"top_element_info"`
	TransfersMade       int64                `json:"transfers_made"`
}

// TeamsStruct - Data of each team
type TeamsStruct struct {
	Code                int    `json:"code"`
	Draw                int    `json:"draw"`
	Form                int    `json:"form"`
	ID                  int    `json:"id"`
	Loss                int    `json:"loss"`
	Name                string `json:"name"`
	Played              int    `json:"played"`
	Points              int    `json:"points"`
	PulseID             int    `json:"pulse_id"`
	Position            int    `json:"position"`
	ShortName           string `json:"short_name"`
	Strength            int    `json:"strength"`
	StrengthAttackAway  int    `json:"strength_attack_away"`
	StrengthAttackHome  int    `json:"strength_attack_home"`
	StrengthDefenceAway int    `json:"strength_defence_away"`
	StrengthDefenceHome int    `json:"strength_defence_home"`
	StrengthOverallHome int    `json:"strength_overall_home"`
	StrengthOverallAway int    `json:"strength_overall_away"`
	Unavailable         bool   `json:"unavailable"`
	Win                 int    `json:"win"`
}

// TopElementInfoStruct - Part of Event struct. Do not know the use case yet
type TopElementInfoStruct struct {
	ID     int `json:"id"`
	Points int `json:"points"`
}

// Gameweek struct. The structs have information about each gameweek. These
// stats are combined with the element stats
type Gameweek struct {
	ID    int           `json:"id"`
	Stats GameweekStats `json:"stats"`
}

type GameweekStats struct {
	Assists         int    `json:"assists"`
	Bonus           int    `json:"bonus"`
	BPS             int    `json:"bps"`
	CleanSheets     int    `json:"clean_sheets"`
	Creativity      string `json:"creativity"`
	ElementID       int
	GameweekNumber  int
	GoalsConceded   int    `json:"goals_conceded"`
	GoalsScored     int    `json:"goals_scored"`
	IctIndex        string `json:"ict_index"`
	Influence       string `json:"influence"`
	ID              string `gorm:"primaryKey"`
	InDreamteam     bool   `json:"in_dreamteam"`
	Minutes         int    `json:"minutes"`
	OwnGoals        int    `json:"own_goals"`
	PenaltiesSaved  int    `json:"penalties_saved"`
	PenaltiesMissed int    `json:"penalties_missed"`
	RedCards        int    `json:"red_cards"`
	Saves           int    `json:"saves"`
	Threat          string `json:"threat"`
	TotalPoints     int    `json:"total_points"`
	YellowCards     int    `json:"yellow_cards"`
}

type Gameweeks struct {
	GameweekArray []Gameweek `json:"elements"`
}

// Fixtures struct. Information about each element for each fixture
type Fixtures struct {
	Fixture []FixtureStruct `json:"fixtures"`
	History []HistoryStruct `json:"history"`
}

type FixtureStruct struct {
	Code                 int    `json:"code"`
	Difficult            int    `json:"difficulty"`
	Event                int    `json:"event"`
	EventName            string `json:"event_name"`
	Finished             bool   `json:"finished"`
	IsHome               bool   `json:"is_home"`
	ID                   int    `json:"id"`
	KickOffTime          string `json:"kickoff_time"`
	Minutes              int    `json:"minutes"`
	ProvisionalStartTime bool   `json:"provisional_start_time"`
	TeamAway             int    `json:"team_a"`
	TeamAwayScore        int    `json:"team_a_score"`
	TeamHome             int    `json:"team_h"`
	TeamHomeScore        int    `json:"team_h_score"`
}

type HistoryStruct struct {
	Assists          int    `json:"assists"`
	Bonus            int    `json:"bonus"`
	BPS              int    `json:"bps"`
	CleanSheets      int    `json:"clean_sheets"`
	Creativity       string `json:"creativity"`
	Element          int    `json:"element"`
	Fixture          int    `json:"fixture"`
	KickOffTime      string `json:"kickoff_time"`
	GoalsConceded    int    `json:"goals_conceded"`
	GoalsScored      int    `json:"goals_scored"`
	IctIndex         string `json:"ict_index"`
	Influence        string `json:"influence"`
	Minutes          int    `json:"minutes"`
	OpponentTeam     int    `json:"opponent_team"`
	OwnGoals         int    `json:"own_goals"`
	PenaltiesSaved   int    `json:"penalties_saved"`
	PenaltiesMissed  int    `json:"penalties_missed"`
	RedCards         int    `json:"red_cards"`
	Round            int    `json:"round"`
	Saves            int    `json:"saves"`
	Selected         int    `json:"selected"`
	TeamAwayScore    int    `json:"team_a_score"`
	TeamHomeScore    int    `json:"team_h_score"`
	Threat           string `json:"threat"`
	TotalPoints      int    `json:"total_points"`
	TransfersBalance int    `json:"transfers_balance"`
	TransfersIn      int    `json:"transfers_in"`
	TransfersOut     int    `json:"transfers_out"`
	Value            string `json:"value"`
	WasHome          bool   `json:"was_home"`
	YellowCards      int    `json:"yellow_cards"`
}
