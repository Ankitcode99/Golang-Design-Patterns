package observers

import "observer-pattern/NewsFeed/models"

// Observer defines the interface for receiving notifications about new news feeds.
type Observer interface {
	Notify(models.NewsFeed)
}
