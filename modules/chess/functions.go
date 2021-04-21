package chess

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"net/url"
	"path"
	"time"
)

var (
	apiUrl, _ = url.Parse(config.SwayItConfig.Chess.URL)
)

func buildUrl(paths ...string) string {
	u := *apiUrl
	paths = append([]string{u.Path}, paths...)
	u.Path = path.Join(paths...)
	return u.String()
}

func getPlayer(username string) (*PlayerProfile, error) {
	var p PlayerProfile

	err := common.GetJson(common.BuildUrl(apiUrl, "player", username), &p)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func getPlayerStats(username string) (*PlayerStats, error) {
	var p PlayerStats

	err := common.GetJson(common.BuildUrl(apiUrl, "player", username, "stats"), &p)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func getPlayerGames(username string) (*PlayerGames, error) {
	var p PlayerGames

	err := common.GetJson(common.BuildUrl(apiUrl, "player", username, "games"), &p)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func getPlayerToMove(username string) (*PlayerToMove, error) {
	var p PlayerToMove

	err := common.GetJson(common.BuildUrl(apiUrl, "player", username, "games", "to-move"), &p)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func run(w *common.WaybarOutput) error {
	toMove, err := getPlayerToMove(config.SwayItConfig.Chess.Username)
	if err != nil {
		return err
	}

	if len(toMove.Games) > 0 {
		player, err := getPlayer(config.SwayItConfig.Chess.Username)

		if err != nil {
			return err
		}

		w.Text = fmt.Sprintf("\uf439 %d", len(toMove.Games))

		w.Tooltip = "<b>Source:</b> Chess.com"
		w.Tooltip += fmt.Sprintf("\n<b>Last Update:</b> %s", module.GetLastRun())

		w.Tooltip += fmt.Sprintf("\n\n<b>Name:</b> %s", player.Name)
		if player.Title != "" {
			w.Tooltip += fmt.Sprintf("\n<b>Title:</b> %s", player.Title)
		}
		w.Tooltip += fmt.Sprintf("\n<b>Followers:</b> %d", player.Followers)

		joined := time.Unix(int64(player.Joined), 0)
		w.Tooltip += fmt.Sprintf("\n\n<b>Joined: %s", joined.String())
		lastOnline := time.Unix(int64(player.LastOnline), 0)
		w.Tooltip += fmt.Sprintf("\n\n<b>Joined: %s", lastOnline.String())
	}

	return nil
}
