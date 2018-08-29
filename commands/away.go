package commands

import (
	"fmt"
	"time"
)

func Away(dcs DiscordSession, args []string) {
	var sDate = args[1]
	if validateTime(sDate) == false {
		_, _ = dcs.Session.ChannelMessageSend(dcs.Msg.ChannelID, "Invalid date given")
		return
	}

	if args[0] == "add" {

		layout := "2006-01-02"
		t, _ := time.Parse(layout, args[1])

		fmt.Println(t) // 2017-08-31 00:00:00 +0000 UTC
		_, _ = dcs.Session.ChannelMessageSend(dcs.Msg.ChannelID, "Your entry as been added")
	}
	if args[0] == "remove" {

		layout := "2006-01-02"
		t, _ := time.Parse(layout, args[1])
		fmt.Println(t) // 2017-08-31 00:00:00 +0000 UTC
		_, _ = dcs.Session.ChannelMessageSend(dcs.Msg.ChannelID, "Your entry as been removed")
	}

}

func parseTimeStrict(layout, value string) (time.Time, error) {
	t, err := time.Parse(layout, value)
	if err != nil {
		return t, err
	}
	if t.Format(layout) != value {
		return t, fmt.Errorf("invalid time: %q", value)
	}
	return t, nil
}

func validateTime(s string) bool {
	t, err := parseTimeStrict("1/2/06", s)
	if err != nil {
		fmt.Printf("error parsing %q: %v\n", s, err)
		return false
	}
	fmt.Printf("parsed %q: %v\n", s, t)
	return true
}
