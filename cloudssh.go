// CloudSSH - SSH for the rest of us!
// Lists instances in from the cloud provider of your choosing
// Author:: Brian L. Scott

package main

import (
	"fmt"
	"github.com/hailocab/goamz/aws"
	"github.com/hailocab/goamz/ec2"
	"github.com/spf13/cobra"
	"github.com/stevedomin/termtable"
	"strconv"
)

// Returns incremented counter
func increaseInt() string {
	return "true"
}

func main() {

	auth, err := aws.EnvAuth()
	if err != nil {
		fmt.Println(err)
		// panic(err.String())
	}
	e := ec2.New(auth, aws.USWest2)

	var cloudsshCmd = &cobra.Command{
		Use:   "ec2",
		Short: "cloudssh lists cloud instances and allows you to ssh the target node",
		Long:  "cloudssh lists cloud instances and allows you to ssh the target node",
		Run: func(c *cobra.Command, arg []string) {
			filter := ec2.NewFilter()
			resp, err := e.DescribeInstances(nil, filter)
			// Check if DescribeInstances returned an error.
			if err != nil {
				panic(err)
			}
			t := termtable.NewTable(nil, nil)
			t.SetHeader([]string{"Index", "Instance ID", "IP Address", "DNS Name", "State", "Key Pair"})
			for _, instance := range resp.Reservations {
				total := strconv.Itoa(0)
				for _, reservation := range instance.Instances {
					newval, _ := strconv.Atoi(total)
					// Check if int conversion failed
					if err != nil {
						panic(err)
					}
					newval++
					newtotal := strconv.Itoa(newval)
					t.AddRow([]string{newtotal, reservation.InstanceId, reservation.IPAddress, reservation.DNSName, reservation.State.Name, reservation.KeyName})
					nt, _ := strconv.Atoi(newtotal)
					total = strconv.Itoa(nt)
				}
			}
			fmt.Println(t.Render())
		},
	}

	var rootCmd = &cobra.Command{Use: "cloudssh"}
	rootCmd.AddCommand(cloudsshCmd)
	rootCmd.Execute()
}
