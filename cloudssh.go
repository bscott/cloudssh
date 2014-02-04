package main

import (
	"fmt"
	"github.com/hailocab/goamz/aws"
	"github.com/hailocab/goamz/ec2"
	"github.com/spf13/cobra"
	"github.com/stevedomin/termtable"
)

func main() {

	auth, err := aws.EnvAuth()
	if err != nil {
		fmt.Println(err)
		// panic(err.String())
	}

	e := ec2.New(auth, aws.USEast)

	var cloudsshCmd = &cobra.Command{
		Use:   "ec2",
		Short: "cloudssh lists cloud instances and allows you to ssh the target node",
		Long:  "cloudssh lists cloud instances and allows you to ssh the target node",
		Run: func(c *cobra.Command, arg []string) {
			filter := ec2.NewFilter()
			resp, err := e.DescribeInstances(nil, filter)
			if err != nil {
				panic(err)
			}
			t := termtable.NewTable(nil, nil)
			t.SetHeader([]string{"Instance ID", "IP Address", "DNS Name", "State", "Key Pair"})
			for _, instance := range resp.Reservations {
				for _, reservation := range instance.Instances {
					t.AddRow([]string{reservation.InstanceId, reservation.IPAddress, reservation.DNSName, reservation.State.Name, reservation.KeyName})
				}
			}
			fmt.Println(t.Render())
		},
	}

	var rootCmd = &cobra.Command{Use: "cloudssh"}
	rootCmd.AddCommand(cloudsshCmd)
	rootCmd.Execute()
}
