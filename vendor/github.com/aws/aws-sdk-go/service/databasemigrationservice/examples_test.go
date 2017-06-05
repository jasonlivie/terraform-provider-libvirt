// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleDatabaseMigrationService_AddTagsToResource() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.AddTagsToResourceInput{
		ResourceArn: aws.String("String"), // Required
		Tags: []*databasemigrationservice.Tag{ // Required
			{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.AddTagsToResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_CreateEndpoint() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.CreateEndpointInput{
		EndpointIdentifier: aws.String("String"),                       // Required
		EndpointType:       aws.String("ReplicationEndpointTypeValue"), // Required
		EngineName:         aws.String("String"),                       // Required
		CertificateArn:     aws.String("String"),
		DatabaseName:       aws.String("String"),
		DynamoDbSettings: &databasemigrationservice.DynamoDbSettings{
			ServiceAccessRoleArn: aws.String("String"), // Required
		},
		ExtraConnectionAttributes: aws.String("String"),
		KmsKeyId:                  aws.String("String"),
		MongoDbSettings: &databasemigrationservice.MongoDbSettings{
			AuthMechanism:     aws.String("AuthMechanismValue"),
			AuthSource:        aws.String("String"),
			AuthType:          aws.String("AuthTypeValue"),
			DatabaseName:      aws.String("String"),
			DocsToInvestigate: aws.String("String"),
			ExtractDocId:      aws.String("String"),
			NestingLevel:      aws.String("NestingLevelValue"),
			Password:          aws.String("SecretString"),
			Port:              aws.Int64(1),
			ServerName:        aws.String("String"),
			Username:          aws.String("String"),
		},
		Password: aws.String("SecretString"),
		Port:     aws.Int64(1),
		S3Settings: &databasemigrationservice.S3Settings{
			BucketFolder:            aws.String("String"),
			BucketName:              aws.String("String"),
			CompressionType:         aws.String("CompressionTypeValue"),
			CsvDelimiter:            aws.String("String"),
			CsvRowDelimiter:         aws.String("String"),
			ExternalTableDefinition: aws.String("String"),
			ServiceAccessRoleArn:    aws.String("String"),
		},
		ServerName: aws.String("String"),
		SslMode:    aws.String("DmsSslModeValue"),
		Tags: []*databasemigrationservice.Tag{
			{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
		Username: aws.String("String"),
	}
	resp, err := svc.CreateEndpoint(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_CreateEventSubscription() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.CreateEventSubscriptionInput{
		SnsTopicArn:      aws.String("String"), // Required
		SubscriptionName: aws.String("String"), // Required
		Enabled:          aws.Bool(true),
		EventCategories: []*string{
			aws.String("String"), // Required
			// More values...
		},
		SourceIds: []*string{
			aws.String("String"), // Required
			// More values...
		},
		SourceType: aws.String("String"),
		Tags: []*databasemigrationservice.Tag{
			{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateEventSubscription(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_CreateReplicationInstance() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.CreateReplicationInstanceInput{
		ReplicationInstanceClass:         aws.String("String"), // Required
		ReplicationInstanceIdentifier:    aws.String("String"), // Required
		AllocatedStorage:                 aws.Int64(1),
		AutoMinorVersionUpgrade:          aws.Bool(true),
		AvailabilityZone:                 aws.String("String"),
		EngineVersion:                    aws.String("String"),
		KmsKeyId:                         aws.String("String"),
		MultiAZ:                          aws.Bool(true),
		PreferredMaintenanceWindow:       aws.String("String"),
		PubliclyAccessible:               aws.Bool(true),
		ReplicationSubnetGroupIdentifier: aws.String("String"),
		Tags: []*databasemigrationservice.Tag{
			{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
		VpcSecurityGroupIds: []*string{
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.CreateReplicationInstance(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_CreateReplicationSubnetGroup() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.CreateReplicationSubnetGroupInput{
		ReplicationSubnetGroupDescription: aws.String("String"), // Required
		ReplicationSubnetGroupIdentifier:  aws.String("String"), // Required
		SubnetIds: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
		Tags: []*databasemigrationservice.Tag{
			{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateReplicationSubnetGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_CreateReplicationTask() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.CreateReplicationTaskInput{
		MigrationType:             aws.String("MigrationTypeValue"), // Required
		ReplicationInstanceArn:    aws.String("String"),             // Required
		ReplicationTaskIdentifier: aws.String("String"),             // Required
		SourceEndpointArn:         aws.String("String"),             // Required
		TableMappings:             aws.String("String"),             // Required
		TargetEndpointArn:         aws.String("String"),             // Required
		CdcStartTime:              aws.Time(time.Now()),
		ReplicationTaskSettings:   aws.String("String"),
		Tags: []*databasemigrationservice.Tag{
			{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateReplicationTask(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DeleteCertificate() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.DeleteCertificateInput{
		CertificateArn: aws.String("String"), // Required
	}
	resp, err := svc.DeleteCertificate(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DeleteEndpoint() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.DeleteEndpointInput{
		EndpointArn: aws.String("String"), // Required
	}
	resp, err := svc.DeleteEndpoint(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DeleteEventSubscription() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.DeleteEventSubscriptionInput{
		SubscriptionName: aws.String("String"), // Required
	}
	resp, err := svc.DeleteEventSubscription(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DeleteReplicationInstance() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.DeleteReplicationInstanceInput{
		ReplicationInstanceArn: aws.String("String"), // Required
	}
	resp, err := svc.DeleteReplicationInstance(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DeleteReplicationSubnetGroup() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.DeleteReplicationSubnetGroupInput{
		ReplicationSubnetGroupIdentifier: aws.String("String"), // Required
	}
	resp, err := svc.DeleteReplicationSubnetGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DeleteReplicationTask() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.DeleteReplicationTaskInput{
		ReplicationTaskArn: aws.String("String"), // Required
	}
	resp, err := svc.DeleteReplicationTask(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DescribeAccountAttributes() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	var params *databasemigrationservice.DescribeAccountAttributesInput
	resp, err := svc.DescribeAccountAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DescribeCertificates() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.DescribeCertificatesInput{
		Filters: []*databasemigrationservice.Filter{
			{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		Marker:     aws.String("String"),
		MaxRecords: aws.Int64(1),
	}
	resp, err := svc.DescribeCertificates(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DescribeConnections() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.DescribeConnectionsInput{
		Filters: []*databasemigrationservice.Filter{
			{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		Marker:     aws.String("String"),
		MaxRecords: aws.Int64(1),
	}
	resp, err := svc.DescribeConnections(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DescribeEndpointTypes() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.DescribeEndpointTypesInput{
		Filters: []*databasemigrationservice.Filter{
			{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		Marker:     aws.String("String"),
		MaxRecords: aws.Int64(1),
	}
	resp, err := svc.DescribeEndpointTypes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DescribeEndpoints() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.DescribeEndpointsInput{
		Filters: []*databasemigrationservice.Filter{
			{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		Marker:     aws.String("String"),
		MaxRecords: aws.Int64(1),
	}
	resp, err := svc.DescribeEndpoints(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DescribeEventCategories() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.DescribeEventCategoriesInput{
		Filters: []*databasemigrationservice.Filter{
			{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		SourceType: aws.String("String"),
	}
	resp, err := svc.DescribeEventCategories(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DescribeEventSubscriptions() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.DescribeEventSubscriptionsInput{
		Filters: []*databasemigrationservice.Filter{
			{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		Marker:           aws.String("String"),
		MaxRecords:       aws.Int64(1),
		SubscriptionName: aws.String("String"),
	}
	resp, err := svc.DescribeEventSubscriptions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DescribeEvents() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.DescribeEventsInput{
		Duration: aws.Int64(1),
		EndTime:  aws.Time(time.Now()),
		EventCategories: []*string{
			aws.String("String"), // Required
			// More values...
		},
		Filters: []*databasemigrationservice.Filter{
			{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		Marker:           aws.String("String"),
		MaxRecords:       aws.Int64(1),
		SourceIdentifier: aws.String("String"),
		SourceType:       aws.String("SourceType"),
		StartTime:        aws.Time(time.Now()),
	}
	resp, err := svc.DescribeEvents(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DescribeOrderableReplicationInstances() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.DescribeOrderableReplicationInstancesInput{
		Marker:     aws.String("String"),
		MaxRecords: aws.Int64(1),
	}
	resp, err := svc.DescribeOrderableReplicationInstances(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DescribeRefreshSchemasStatus() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.DescribeRefreshSchemasStatusInput{
		EndpointArn: aws.String("String"), // Required
	}
	resp, err := svc.DescribeRefreshSchemasStatus(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DescribeReplicationInstances() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.DescribeReplicationInstancesInput{
		Filters: []*databasemigrationservice.Filter{
			{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		Marker:     aws.String("String"),
		MaxRecords: aws.Int64(1),
	}
	resp, err := svc.DescribeReplicationInstances(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DescribeReplicationSubnetGroups() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.DescribeReplicationSubnetGroupsInput{
		Filters: []*databasemigrationservice.Filter{
			{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		Marker:     aws.String("String"),
		MaxRecords: aws.Int64(1),
	}
	resp, err := svc.DescribeReplicationSubnetGroups(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DescribeReplicationTasks() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.DescribeReplicationTasksInput{
		Filters: []*databasemigrationservice.Filter{
			{ // Required
				Name: aws.String("String"), // Required
				Values: []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		Marker:     aws.String("String"),
		MaxRecords: aws.Int64(1),
	}
	resp, err := svc.DescribeReplicationTasks(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DescribeSchemas() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.DescribeSchemasInput{
		EndpointArn: aws.String("String"), // Required
		Marker:      aws.String("String"),
		MaxRecords:  aws.Int64(1),
	}
	resp, err := svc.DescribeSchemas(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_DescribeTableStatistics() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.DescribeTableStatisticsInput{
		ReplicationTaskArn: aws.String("String"), // Required
		Marker:             aws.String("String"),
		MaxRecords:         aws.Int64(1),
	}
	resp, err := svc.DescribeTableStatistics(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_ImportCertificate() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.ImportCertificateInput{
		CertificateIdentifier: aws.String("String"), // Required
		CertificatePem:        aws.String("String"),
		CertificateWallet:     []byte("PAYLOAD"),
	}
	resp, err := svc.ImportCertificate(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_ListTagsForResource() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.ListTagsForResourceInput{
		ResourceArn: aws.String("String"), // Required
	}
	resp, err := svc.ListTagsForResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_ModifyEndpoint() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.ModifyEndpointInput{
		EndpointArn:    aws.String("String"), // Required
		CertificateArn: aws.String("String"),
		DatabaseName:   aws.String("String"),
		DynamoDbSettings: &databasemigrationservice.DynamoDbSettings{
			ServiceAccessRoleArn: aws.String("String"), // Required
		},
		EndpointIdentifier:        aws.String("String"),
		EndpointType:              aws.String("ReplicationEndpointTypeValue"),
		EngineName:                aws.String("String"),
		ExtraConnectionAttributes: aws.String("String"),
		MongoDbSettings: &databasemigrationservice.MongoDbSettings{
			AuthMechanism:     aws.String("AuthMechanismValue"),
			AuthSource:        aws.String("String"),
			AuthType:          aws.String("AuthTypeValue"),
			DatabaseName:      aws.String("String"),
			DocsToInvestigate: aws.String("String"),
			ExtractDocId:      aws.String("String"),
			NestingLevel:      aws.String("NestingLevelValue"),
			Password:          aws.String("SecretString"),
			Port:              aws.Int64(1),
			ServerName:        aws.String("String"),
			Username:          aws.String("String"),
		},
		Password: aws.String("SecretString"),
		Port:     aws.Int64(1),
		S3Settings: &databasemigrationservice.S3Settings{
			BucketFolder:            aws.String("String"),
			BucketName:              aws.String("String"),
			CompressionType:         aws.String("CompressionTypeValue"),
			CsvDelimiter:            aws.String("String"),
			CsvRowDelimiter:         aws.String("String"),
			ExternalTableDefinition: aws.String("String"),
			ServiceAccessRoleArn:    aws.String("String"),
		},
		ServerName: aws.String("String"),
		SslMode:    aws.String("DmsSslModeValue"),
		Username:   aws.String("String"),
	}
	resp, err := svc.ModifyEndpoint(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_ModifyEventSubscription() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.ModifyEventSubscriptionInput{
		SubscriptionName: aws.String("String"), // Required
		Enabled:          aws.Bool(true),
		EventCategories: []*string{
			aws.String("String"), // Required
			// More values...
		},
		SnsTopicArn: aws.String("String"),
		SourceType:  aws.String("String"),
	}
	resp, err := svc.ModifyEventSubscription(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_ModifyReplicationInstance() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.ModifyReplicationInstanceInput{
		ReplicationInstanceArn:        aws.String("String"), // Required
		AllocatedStorage:              aws.Int64(1),
		AllowMajorVersionUpgrade:      aws.Bool(true),
		ApplyImmediately:              aws.Bool(true),
		AutoMinorVersionUpgrade:       aws.Bool(true),
		EngineVersion:                 aws.String("String"),
		MultiAZ:                       aws.Bool(true),
		PreferredMaintenanceWindow:    aws.String("String"),
		ReplicationInstanceClass:      aws.String("String"),
		ReplicationInstanceIdentifier: aws.String("String"),
		VpcSecurityGroupIds: []*string{
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.ModifyReplicationInstance(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_ModifyReplicationSubnetGroup() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.ModifyReplicationSubnetGroupInput{
		ReplicationSubnetGroupIdentifier: aws.String("String"), // Required
		SubnetIds: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
		ReplicationSubnetGroupDescription: aws.String("String"),
	}
	resp, err := svc.ModifyReplicationSubnetGroup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_ModifyReplicationTask() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.ModifyReplicationTaskInput{
		ReplicationTaskArn:        aws.String("String"), // Required
		CdcStartTime:              aws.Time(time.Now()),
		MigrationType:             aws.String("MigrationTypeValue"),
		ReplicationTaskIdentifier: aws.String("String"),
		ReplicationTaskSettings:   aws.String("String"),
		TableMappings:             aws.String("String"),
	}
	resp, err := svc.ModifyReplicationTask(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_RefreshSchemas() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.RefreshSchemasInput{
		EndpointArn:            aws.String("String"), // Required
		ReplicationInstanceArn: aws.String("String"), // Required
	}
	resp, err := svc.RefreshSchemas(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_ReloadTables() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.ReloadTablesInput{
		ReplicationTaskArn: aws.String("String"), // Required
		TablesToReload: []*databasemigrationservice.TableToReload{ // Required
			{ // Required
				SchemaName: aws.String("String"),
				TableName:  aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.ReloadTables(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_RemoveTagsFromResource() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.RemoveTagsFromResourceInput{
		ResourceArn: aws.String("String"), // Required
		TagKeys: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.RemoveTagsFromResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_StartReplicationTask() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.StartReplicationTaskInput{
		ReplicationTaskArn:       aws.String("String"),                        // Required
		StartReplicationTaskType: aws.String("StartReplicationTaskTypeValue"), // Required
		CdcStartTime:             aws.Time(time.Now()),
	}
	resp, err := svc.StartReplicationTask(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_StopReplicationTask() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.StopReplicationTaskInput{
		ReplicationTaskArn: aws.String("String"), // Required
	}
	resp, err := svc.StopReplicationTask(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDatabaseMigrationService_TestConnection() {
	sess := session.Must(session.NewSession())

	svc := databasemigrationservice.New(sess)

	params := &databasemigrationservice.TestConnectionInput{
		EndpointArn:            aws.String("String"), // Required
		ReplicationInstanceArn: aws.String("String"), // Required
	}
	resp, err := svc.TestConnection(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}