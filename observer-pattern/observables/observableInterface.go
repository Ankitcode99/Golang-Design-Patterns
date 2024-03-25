package observables

import (
	"observer-pattern/NewsFeed/models"
	"observer-pattern/NewsFeed/observers"
)

type CategoryObservable interface {
	Subscribe(observers.Observer)

	Unsubscribe(observers.Observer)

	notifySubscribers(models.NewsFeed)

	AddNewNewsFeed(models.NewsFeed)
}
