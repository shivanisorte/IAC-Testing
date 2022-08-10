package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/assert"

	"github.com/gruntwork-io/terratest/modules/terraform"

	"encoding/json"
	"io/ioutil"
	"log"
)

type Data struct {
	name          string
	instance_type string
	ami           string
	region        string
}

func configurableTerraformOptions(t *testing.T) *terraform.Options {

	content, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var statvars Data
	err = json.Unmarshal(content, &statvars)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	name := fmt.Sprintf("%s-%s", statvars.name, random.UniqueId())

	region := statvars.region

	amiID := statvars.ami

	instance_type := statvars.instance_type

	terraformOptions := &terraform.Options{

		TerraformDir: "../ec2-instance",

		Vars: map[string]interface{}{
			"name":          name,
			"instance_type": instance_type,
			"ami":           amiID,
			"region":        region,
		},
	}

	return terraformOptions

}

func TestEc2Instance(t *testing.T) {
	t.Parallel()
	terraformOptions := configurableTerraformOptions(t)

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	instanceId := terraform.Output(t, terraformOptions, "instance_id")
	region := "us-east-2"
	publicInstanceIP := aws.GetPublicIpOfEc2Instance(t, instanceId, region)

	output := terraform.Output(t, terraformOptions, "ubuntu_instance_public_ip")

	fmt.Println("Expected op: ")
	fmt.Println(publicInstanceIP)
	fmt.Println(output)

	assert.Equal(t, publicInstanceIP, output)

}
