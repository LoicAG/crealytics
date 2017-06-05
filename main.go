package main

import (
    "fmt"
    "io/ioutil"
    "github.com/gin-gonic/gin"
    "golang.org/x/oauth2/google"
    "google.golang.org/api/compute/v1"
)

func main() {
    r := gin.Default()

    // A service account credentials file is expected
    // https://developers.google.com/identity/protocols/OAuth2ServiceAccount
    configBytes, err := ioutil.ReadFile("./credentials.json")
    if err != nil {
        panic(err)
    }

    jwtConfig, err := google.JWTConfigFromJSON(configBytes, "https://www.googleapis.com/auth/compute")
    if err != nil {
        panic(err)
    }

    httpClient := jwtConfig.Client(nil)

    service, err := compute.New(httpClient)
    if err != nil {
        panic(err)
    }

    r.GET("/healthcheck", HealthCheck)
    r.POST("/v1/instances/create", CreateInstance(service))

    //TODO make port configurable
    r.Run(":8080")
}

func HealthCheck(c *gin.Context) {
    c.String(200, "")
}

func CreateInstance(service *compute.Service) gin.HandlerFunc {
    fn := func(c *gin.Context) {

	//TODO put these parameters in a config file
        projectId := "crealytics-169710"
        prefix := "https://www.googleapis.com/compute/v1/projects/" + projectId
        imageURL := "https://www.googleapis.com/compute/v1/projects/debian-cloud/global/images/debian-7-wheezy-v20140606"
        zone := "europe-west1-c"
        machineType := "f1-micro"

        instance := &compute.Instance{
	    Name: "dummy",
            MachineType: prefix + "/zones/" + zone + "/machineTypes/" + machineType,
            NetworkInterfaces: []*compute.NetworkInterface{
                &compute.NetworkInterface{
                    AccessConfigs: []*compute.AccessConfig{
                        &compute.AccessConfig{
                            Type: "ONE_TO_ONE_NAT",
                            Name: "External NAT",
                        },
                    },
                    Network: prefix + "/global/networks/default",
                },
            },
            Disks: []*compute.AttachedDisk{
                {
                    AutoDelete: true,
                    Boot:       true,
                    Type:       "PERSISTENT",
                    InitializeParams: &compute.AttachedDiskInitializeParams{
                        DiskName:    "my-root-pd",
                        SourceImage: imageURL,
                    },
                },
            },
	}

        op, err := service.Instances.Insert(projectId, zone, instance).Do()
        fmt.Printf("Got compute.Operation, err: %#v, %v", op, err)
    }
    return gin.HandlerFunc(fn)
}
