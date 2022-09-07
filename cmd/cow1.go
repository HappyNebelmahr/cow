
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)
var cow1Cmd = &cobra.Command{
	Use:   "cow1",
	Short: "create a cow",
	Long:  `creates a Big cow`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("cow1 called")
		createcow(args)
	},
}
func init() {
	rootCmd.AddCommand(cow1Cmd)
}
func createcow(was []string) {
	var x = ""
	for i, s := range was {
		x = x + " " + s
		i++
	}

	fmt.Println(`
				 _________
				<` + x + ` >
			    ------------
					\    ^__^
					 \   (oo)\___________
					     (__)\           )\/\/
					         ||---------w|
						 ||         ||


  `)

}

//https://www.youtube.com/watch?v=SSRIn5DAmyw
//https://www.youtube.com/watch?v=yl0phDUrnwc
