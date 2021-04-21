package chess

type PlayerProfile struct {
	Avatar     string `json:"avatar"`
	PlayerID   int    `json:"player_id"`
	ID         string `json:"@id"`
	URL        string `json:"url"`
	Name       string `json:"name"`
	Title      string `json:"title"`
	Username   string `json:"username"`
	Followers  int    `json:"followers"`
	Country    string `json:"country"`
	LastOnline int    `json:"last_online"`
	Joined     int    `json:"joined"`
	Status     string `json:"status"`
	IsStreamer bool   `json:"is_streamer"`
}

type PlayerStats struct {
	ChessDaily *struct {
		Last struct {
			Rating int `json:"rating"`
			Date   int `json:"date"`
			Rd     int `json:"rd"`
		} `json:"last"`
		Best struct {
			Rating int    `json:"rating"`
			Date   int    `json:"date"`
			Game   string `json:"game"`
		} `json:"best"`
		Record struct {
			Win            int `json:"win"`
			Loss           int `json:"loss"`
			Draw           int `json:"draw"`
			TimePerMove    int `json:"time_per_move"`
			TimeoutPercent int `json:"timeout_percent"`
		} `json:"record"`
		Tournament struct {
			Points        int `json:"points"`
			Withdraw      int `json:"withdraw"`
			Count         int `json:"count"`
			HighestFinish int `json:"highest_finish"`
		} `json:"tournament"`
	} `json:"chess_daily"`
	Chess960Daily *struct {
		Last struct {
			Rating int `json:"rating"`
			Date   int `json:"date"`
			Rd     int `json:"rd"`
		} `json:"last"`
		Best struct {
			Rating int    `json:"rating"`
			Date   int    `json:"date"`
			Game   string `json:"game"`
		} `json:"best"`
		Record struct {
			Win            int `json:"win"`
			Loss           int `json:"loss"`
			Draw           int `json:"draw"`
			TimePerMove    int `json:"time_per_move"`
			TimeoutPercent int `json:"timeout_percent"`
		} `json:"record"`
		Tournament struct {
			Points        int `json:"points"`
			Withdraw      int `json:"withdraw"`
			Count         int `json:"count"`
			HighestFinish int `json:"highest_finish"`
		} `json:"tournament"`
	} `json:"chess960_daily"`
	ChessRapid *struct {
		Last struct {
			Rating int `json:"rating"`
			Date   int `json:"date"`
			Rd     int `json:"rd"`
		} `json:"last"`
		Best struct {
			Rating int    `json:"rating"`
			Date   int    `json:"date"`
			Game   string `json:"game"`
		} `json:"best"`
		Record struct {
			Win  int `json:"win"`
			Loss int `json:"loss"`
			Draw int `json:"draw"`
		} `json:"record"`
	} `json:"chess_rapid"`
	ChessBullet *struct {
		Last struct {
			Rating int `json:"rating"`
			Date   int `json:"date"`
			Rd     int `json:"rd"`
		} `json:"last"`
		Best struct {
			Rating int    `json:"rating"`
			Date   int    `json:"date"`
			Game   string `json:"game"`
		} `json:"best"`
		Record struct {
			Win  int `json:"win"`
			Loss int `json:"loss"`
			Draw int `json:"draw"`
		} `json:"record"`
	} `json:"chess_bullet"`
	ChessBlitz *struct {
		Last struct {
			Rating int `json:"rating"`
			Date   int `json:"date"`
			Rd     int `json:"rd"`
		} `json:"last"`
		Best struct {
			Rating int    `json:"rating"`
			Date   int    `json:"date"`
			Game   string `json:"game"`
		} `json:"best"`
		Record struct {
			Win  int `json:"win"`
			Loss int `json:"loss"`
			Draw int `json:"draw"`
		} `json:"record"`
	} `json:"chess_blitz"`
	Tactics *struct {
		Highest struct {
			Rating int `json:"rating"`
			Date   int `json:"date"`
		} `json:"highest"`
		Lowest struct {
			Rating int `json:"rating"`
			Date   int `json:"date"`
		} `json:"lowest"`
	} `json:"tactics"`
	Lessons *struct {
		Highest struct {
			Rating int `json:"rating"`
			Date   int `json:"date"`
		} `json:"highest"`
		Lowest struct {
			Rating int `json:"rating"`
			Date   int `json:"date"`
		} `json:"lowest"`
	} `json:"lessons"`
	PuzzleRush *struct {
		Best struct {
			TotalAttempts int `json:"total_attempts"`
			Score         int `json:"score"`
		} `json:"best"`
	} `json:"puzzle_rush"`
}

type PlayerGames struct {
	Games []struct {
		URL          string `json:"url"`
		MoveBy       int    `json:"move_by"`
		Pgn          string `json:"pgn"`
		TimeControl  string `json:"time_control"`
		LastActivity int    `json:"last_activity"`
		Rated        bool   `json:"rated"`
		Turn         string `json:"turn"`
		Fen          string `json:"fen"`
		StartTime    int    `json:"start_time"`
		TimeClass    string `json:"time_class"`
		Rules        string `json:"rules"`
		White        string `json:"white"`
		Black        string `json:"black"`
		Tournament   string `json:"tournament,omitempty"`
	} `json:"games"`
}

type PlayerToMove struct {
	Games []struct {
		URL          string `json:"url"`
		MoveBy       int    `json:"move_by"`
		LastActivity int    `json:"last_activity"`
	} `json:"games"`
}
