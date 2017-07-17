package vpcs

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"log"
	"awsplus/config"
)

var region string
var sess *session.Session
var sessionErr error

func init() {

	cfg := config.Config
	region = cfg.Region

	sess, sessionErr = session.NewSession(&aws.Config{Region: aws.String(region)})

	if sessionErr != nil {
		log.Fatalln(sessionErr.Error())
	}

}

func List() []*ec2.Vpc {

	svc := ec2.New(sess)
	resp, err := svc.DescribeVpcs(nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return resp.Vpcs
}


func GetName(id string) string {

	svc := ec2.New(sess)
	params := &ec2.DescribeVpcsInput{
		VpcIds: []*string{
			aws.String(id),
		},
	}
	resp, err := svc.DescribeVpcs(params)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, tag := range resp.Vpcs[0].Tags {
		if *tag.Key == "Name" {
			return *tag.Value
		}
	}

	return ""
}