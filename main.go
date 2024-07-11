package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/budgets"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// The email addresses to receive notifications.
		const myEmail = "aws_admin@random.domain.com"

		// Create a new budget.
		_, err := budgets.NewBudget(ctx, "budget", &budgets.BudgetArgs{
			Name:            pulumi.String("My Pulumi Budget"),
			BudgetType:      pulumi.String("COST"),
			LimitAmount:     pulumi.String("15.00"), // Monthly budget amount in USD.
			LimitUnit:       pulumi.String("USD"),
			TimePeriodStart: pulumi.String("2024-05-01_00:00"),
			TimeUnit:        pulumi.String("MONTHLY"),
			Notifications: budgets.BudgetNotificationArray{
				&budgets.BudgetNotificationArgs{
					ComparisonOperator: pulumi.String("GREATER_THAN"),
					Threshold:          pulumi.Float64(75),
					ThresholdType:      pulumi.String("PERCENTAGE"),
					NotificationType:   pulumi.String("ACTUAL"),
					SubscriberEmailAddresses: pulumi.StringArray{
						pulumi.String(myEmail),
					},
				},
				&budgets.BudgetNotificationArgs{
					ComparisonOperator: pulumi.String("GREATER_THAN"),
					Threshold:          pulumi.Float64(100),
					ThresholdType:      pulumi.String("PERCENTAGE"),
					NotificationType:   pulumi.String("ACTUAL"),
					SubscriberEmailAddresses: pulumi.StringArray{
						pulumi.String(myEmail),
					},
				},
			},
		})
		if err != nil {
			return err
		}

		return nil
	})
}
