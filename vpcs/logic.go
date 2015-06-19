package vpcs

import (
	"../config"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/olekukonko/tablewriter"
)

func (c *Command) showList(writer io.Writer) {
	conf := config.GetConfig()

	credentials := credentials.NewStaticCredentials(
		conf.Aws.AccessKeyId,
		conf.Aws.SecretAccessKey,
		""
	)

	svc := ec2.New(
		&aws.Config{
			Region: conf.Aws.Region,
			Credentials: credentials,
		},
	)
}
