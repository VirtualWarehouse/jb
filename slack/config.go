package slack

import (
	"fmt"

	"github.com/spf13/viper"
)

func checkConfig() {
	v := []string{
		"d",
		"token",
		"touchchannel",
		"apptoken",
		"userid",
		"slackdomain",
	}
	for _, s := range v {
		check(s)
	}
}

func check(s string) {
	target := viper.GetString("slackdomain")
	if len(target) == 0 {
		panic(fmt.Sprintf("%s is empty", s))
	}
}
