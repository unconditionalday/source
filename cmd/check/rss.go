package check

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/unconditionalday/source-checker/internal/checker"
	"github.com/unconditionalday/source-checker/internal/source"
	iox "github.com/unconditionalday/source-checker/internal/x/io"

	"github.com/pterm/pterm"
)

var (
	SuccessStyle = pterm.NewStyle(pterm.FgGreen)
	FailureStyle = pterm.NewStyle(pterm.FgRed)
	ActualStyle  = pterm.NewStyle(pterm.FgLightGreen)

	AvailableStr   = SuccessStyle.Sprint("Available")
	UnavailableStr = FailureStyle.Sprint("Unavailable")
	ErrorStr       = FailureStyle.Sprint("Error")
)

func NewRSSCheckCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rss [source.json]",
		Short: "Check rss sources ",
		Long:  `Check rss source`,
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := iox.ReadJSON(args[0], source.Source{})
			if err != nil {
				return err
			}

			c := checker.NewRSSChecker()

			x := pterm.TableData{
				{"Name", "URL", "Availability", "Latency"},
			}

			pb, _ := pterm.DefaultProgressbar.WithTotal(len(s)).WithTitle("Installing stuff").Start()
			var errors []error
			for i := 0; i < pb.Total; i++ {
				pb.UpdateTitle("Checking " + ActualStyle.Sprint(s[i].Name))
				aRes, lRes, err := ProcessRSSFeed(s[i].URL, c)
				if err != nil {
					errors = append(errors, err)
				}

				x = append(x, []string{s[i].Name, s[i].URL, aRes, lRes})

				pb.Increment()
			}

			pterm.DefaultTable.WithHasHeader().WithData(x).Render()
			pb.Stop()

			if len(errors) > 0 {
				return fmt.Errorf("%v\n", errors)
			}

			return nil
		},
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("unconditional")

	cmd.Flags().StringP("availability", "a", "", "Availability")
	cmd.Flags().StringP("latency", "l", "", "Latency")

	return cmd
}

func ProcessRSSFeed(rssURL string, c checker.RSSChecker) (string, string, error) {
	aRes := UnavailableStr
	lRes := ErrorStr

	if c.Availability(rssURL) {
		aRes = AvailableStr
	} else {
		return aRes, lRes, fmt.Errorf("%s is unavailable", rssURL)
	}

	if latency, err := c.Latency(rssURL); err == nil {
		if latency > 1000 {
			lRes = FailureStyle.Sprintf("%dms", latency)
			return aRes, lRes, fmt.Errorf("%s is too slow", rssURL)
		}

		lRes = SuccessStyle.Sprintf("%dms", latency)
	}

	return aRes, lRes, nil
}
