package asset

import domainuser "github.com/mlevshin/my-finance-go-clean/internal/domain/user"

func SetName(name string) func(a *Asset) error {
	return func(a *Asset) error {
		a.Name = name
		return nil
	}
}

func SetType(name string) func(a *Asset) error {
	return func(a *Asset) error {
		a.Name = name
		return nil
	}
}

func SetOwner(u *domainuser.User) func(a *Asset) error {
	return func(a *Asset) error {
		a.Owner = u
		return nil
	}
}

func SetLimit(name string) func(a *Asset) error {
	return func(a *Asset) error {
		a.Name = name
		return nil
	}
}

func SetCurrency(c Currency) func(a *Asset) error {
	return func(a *Asset) error {
		a.Currency = c
		return nil
	}
}
