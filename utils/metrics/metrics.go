package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
    UsersRegistered = promauto.NewCounter(prometheus.CounterOpts{
        Name: "users_registered",
        Help: "Total amount of users resgitered",
    })
)