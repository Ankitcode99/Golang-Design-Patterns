package main

import (
	"fmt"
	"observer-pattern/NewsFeed/models"
	"observer-pattern/NewsFeed/observables"
	"observer-pattern/NewsFeed/observers"
)

func main() {
	// Create a new category
	politicsCategory := observables.NewCategory("POLITICS")
	nseCategory := observables.NewCategory("NSE")
	technologyCategory := observables.NewCategory("Technology")

	addNewNewsFeed := func(news models.NewsFeed) {
		switch news.Category {
		case "POLITICS":
			politicsCategory.AddNewNewsFeed(news)

		case "NSE":
			nseCategory.AddNewNewsFeed(news)

		case "Technology":
			technologyCategory.AddNewNewsFeed(news)

		default:
			fmt.Printf("~~~~~~   A NEW CATEGORY \"%s\" HAS BEEN MENTIONED, PLEASE ADD IT IN LIST OF CATEGORIES  ~~~~~~\n\n", news.Category)
		}
	}

	// Create some observers (e.g., email notifier)
	emailNotifier := observers.EmailNotifier{EmailAddress: "user@example.com"}

	// Subscribe the observer to the category
	politicsCategory.Subscribe(&emailNotifier)
	nseCategory.Subscribe(&emailNotifier)

	// Create some news feeds
	newsFeed1 := models.NewsFeed{
		Headline:    "New Political Leader Emerges",
		Description: "A new political leader has gained significant momentum in recent polls.",
		PublishedOn: "2024-03-24",
		Category:    "POLITICS",
	}

	newsFeed2 := models.NewsFeed{
		Headline:    "Economic Policy Update",
		Description: "The government has announced changes to its economic policies.",
		PublishedOn: "2024-03-24",
		Category:    "JOB",
	}

	// Add news feeds to the category
	// politicsCategory.AddNewNewsFeed(newsFeed1)
	// politicsCategory.AddNewNewsFeed(newsFeed2)

	addNewNewsFeed(newsFeed1)
	addNewNewsFeed(newsFeed2)

	// (Optional) Unsubscribe the observer
	// politicsCategory.unsubscribe(&emailNotifier)
}
