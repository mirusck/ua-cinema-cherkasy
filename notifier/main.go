package notifier

import (
	"github.com/aws/aws-lambda-go/lambda"
)


/*
 * TODO send day-to-day notifications about new movies that appeared in DB (mongo or redis?)
 */
func Handler() error {
	// TODO integrate with FB messenger

	// TODO fetch all movies that appeared in DB no longer than 2 days
	// and their notifications still haven't been sent

	// TODO schedule broadcast message for all subscribed users with movies info

	return nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(Handler)
}

